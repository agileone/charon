package model

import (
	"database/sql"
	"errors"
	"strings"
	"sync"

	"github.com/piotrkowalczuk/charon"
)

// Permission returns charon.Permission value that is concatenated
// using entity properties like subsystem, module and action.
func (pe *PermissionEntity) Permission() charon.Permission {
	return charon.Permission(pe.Subsystem + ":" + pe.Module + ":" + pe.Action)
}

// PermissionProvider ...
type PermissionProvider interface {
	Find(criteria *PermissionCriteria) ([]*PermissionEntity, error)
	FindOneByID(id int64) (entity *PermissionEntity, err error)
	// FindByUserID retrieves all permissions for user represented by given id.
	FindByUserID(userID int64) (entities []*PermissionEntity, err error)
	// FindByGroupID retrieves all permissions for group represented by given id.
	FindByGroupID(groupID int64) (entities []*PermissionEntity, err error)
	Register(permissions charon.Permissions) (created, untouched, removed int64, err error)
	Insert(entity *PermissionEntity) (*PermissionEntity, error)
}

// PermissionRepository extends PermissionRepositoryBase
type PermissionRepository struct {
	PermissionRepositoryBase
}

// NewPermissionRepository ...
func NewPermissionRepository(dbPool *sql.DB) *PermissionRepository {
	return &PermissionRepository{
		PermissionRepositoryBase: PermissionRepositoryBase{
			db:      dbPool,
			table:   TablePermission,
			columns: TablePermissionColumns,
		},
	}
}

// FindByUserID implements PermissionProvider interface.
func (pr *PermissionRepository) FindByUserID(userID int64) ([]*PermissionEntity, error) {
	// TODO: does it work?
	return pr.FindBy(`
		SELECT DISTINCT ON (p.ID)
			`+columns(TablePermissionColumns, "p")+`
		FROM `+pr.table+` AS p
		LEFT JOIN `+TableUserPermissions+` AS up
			ON up.`+TableUserPermissionsColumnPermissionSubsystem+` = p.`+TablePermissionColumnSubsystem+`
			AND up.`+TableUserPermissionsColumnPermissionModule+` = p.`+TablePermissionColumnModule+`
			AND up.`+TableUserPermissionsColumnPermissionAction+` = p.`+TablePermissionColumnAction+`
		LEFT JOIN `+TableUserGroups+` AS ug ON ug.`+TableUserGroupsColumnUserID+` = $1
		LEFT JOIN `+TableGroupPermissions+` AS gp
			ON gp.`+TableGroupPermissionsColumnPermissionSubsystem+` = p.`+TablePermissionColumnSubsystem+`
			AND gp.`+TableGroupPermissionsColumnPermissionModule+` = p.`+TablePermissionColumnModule+`
			AND gp.`+TableGroupPermissionsColumnPermissionAction+` = p.`+TablePermissionColumnAction+`
			AND gp.`+TableGroupPermissionsColumnGroupID+` = ug.`+TableUserGroupsColumnGroupID+`
		WHERE up.`+TableUserPermissionsColumnUserID+` = $1 OR ug.`+TableUserGroupsColumnUserID+` = $1
	`, userID)
}

// FindByGroupID implements PermissionProvider interface.
func (pr *PermissionRepository) FindByGroupID(userID int64) ([]*PermissionEntity, error) {
	// TODO: does it work?
	return pr.FindBy(`
		SELECT DISTINCT ON (p.ID)
			`+columns(TablePermissionColumns, "p")+`
		FROM `+pr.table+` AS p
		LEFT JOIN `+TableGroupPermissions+` AS gp ON gp.permission_id = p.ID AND gp.group_id = ug.group_id
		WHERE up.user_id = $1 OR ug.user_id = $1
	`, userID)
}

// FindBy ...
func (pr *PermissionRepository) FindBy(query string, args ...interface{}) ([]*PermissionEntity, error) {
	rows, err := pr.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	permissions := []*PermissionEntity{}
	for rows.Next() {
		var p PermissionEntity
		err = rows.Scan(
			&p.Action,
			&p.CreatedAt,
			&p.ID,
			&p.Module,
			&p.Subsystem,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		permissions = append(permissions, &p)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return permissions, nil
}

func (pr *PermissionRepository) findOneStmt() (*sql.Stmt, error) {
	return pr.db.Prepare(
		"SELECT " + strings.Join(TablePermissionColumns, ",") + " " +
			"FROM " + pr.table + " AS p " +
			"WHERE p.subsystem = $1 AND p.module = $2 AND p.action = $3",
	)
}

// Register ...
func (pr *PermissionRepository) Register(permissions charon.Permissions) (created, unt, removed int64, err error) {
	var (
		tx             *sql.Tx
		insert, delete *sql.Stmt
		rows           *sql.Rows
		res            sql.Result
		subsystem      string
		entities       []*PermissionEntity
		affected       int64
	)
	if len(permissions) == 0 {
		return 0, 0, 0, errors.New("empty slice, permissions cannot be registered")
	}

	subsystem = permissions[0].Subsystem()
	if subsystem == "" {
		return 0, 0, 0, errors.New("subsystem name is empty string, permissions cannot be registered")
	}

	for _, p := range permissions {
		if p.Subsystem() != subsystem {
			return 0, 0, 0, errors.New("provided permissions do not belong to one subsystem, permissions cannot be registered")
		}
	}

	tx, err = pr.db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
			unt = untouched(int64(len(permissions)), created, removed)
		}
	}()

	rows, err = tx.Query("SELECT "+strings.Join(TablePermissionColumns, ",")+" FROM "+pr.table+" AS p WHERE p.subsystem = $1", subsystem)
	if err != nil {
		return
	}
	defer rows.Close()

	entities = []*PermissionEntity{}
	for rows.Next() {
		var entity PermissionEntity
		err = rows.Scan(
			&entity.Action,
			&entity.CreatedAt,
			&entity.ID,
			&entity.Module,
			&entity.Subsystem,
			&entity.UpdatedAt,
		)
		if err != nil {
			return
		}
		entities = append(entities, &entity)
	}
	if rows.Err() != nil {
		return 0, 0, 0, rows.Err()
	}

	insert, err = tx.Prepare("INSERT INTO " + pr.table + " (subsystem, module, action) VALUES ($1, $2, $3)")
	if err != nil {
		return
	}

MissingPermissionsLoop:
	for _, p := range permissions {
		for _, e := range entities {
			if p == e.Permission() {
				continue MissingPermissionsLoop
			}
		}

		if res, err = insert.Exec(p.Split()); err != nil {
			return
		}
		if affected, err = res.RowsAffected(); err != nil {
			return
		}
		created += affected
	}

	delete, err = tx.Prepare("DELETE FROM " + pr.table + " AS p WHERE p.ID = $1")
	if err != nil {
		return
	}

RedundantPermissionsLoop:
	for _, e := range entities {
		for _, p := range permissions {
			if e.Permission() == p {
				continue RedundantPermissionsLoop
			}
		}

		if res, err = delete.Exec(e.ID); err != nil {
			return
		}
		if affected, err = res.RowsAffected(); err != nil {
			return
		}

		removed += affected
	}

	return
}

// PermissionRegistry is an interface that describes in memory storage that holds information
// about permissions that was registered by 3rd party services.
// Should be only used as a proxy for registration process to avoid multiple sql hits.
type PermissionRegistry interface {
	// Exists returns true if given charon.Permission was already registered.
	Exists(permission charon.Permission) (exists bool)
	// Register checks if given collection is valid and
	// calls PermissionProvider to store provided permissions
	// in persistent way.
	Register(permissions charon.Permissions) (created, untouched, removed int64, err error)
}

// PermissionReg ...
type PermissionReg struct {
	sync.RWMutex
	repository  PermissionProvider
	permissions map[charon.Permission]struct{}
}

// NewPermissionRegistry ...
func NewPermissionRegistry(r PermissionProvider) *PermissionReg {
	return &PermissionReg{
		repository:  r,
		permissions: make(map[charon.Permission]struct{}),
	}
}

// Exists ...
func (pr *PermissionReg) Exists(permission charon.Permission) (ok bool) {
	pr.RLock()
	pr.RUnlock()

	_, ok = pr.permissions[permission]
	return
}

// Register ...
func (pr *PermissionReg) Register(permissions charon.Permissions) (created, untouched, removed int64, err error) {
	pr.Lock()
	defer pr.Unlock()

	nb := 0
	for _, p := range permissions {
		if _, ok := pr.permissions[p]; !ok {
			pr.permissions[p] = struct{}{}
			nb++
		}
	}

	if nb > 0 {
		return pr.repository.Register(permissions)
	}

	return 0, 0, 0, nil
}

// FindByTag ...
func (pr *PermissionRepository) FindByTag(userID int64) ([]*PermissionEntity, error) {
	query := `
		SELECT DISTINCT ON (p.ID)
			` + columns(TablePermissionColumns, "p") + `
		FROM ` + pr.table + ` AS p
		LEFT JOIN ` + TableUserPermissions + ` AS up ON up.permission_id = p.ID AND up.user_id = $1
		LEFT JOIN ` + TableUserGroups + ` AS ug ON ug.user_id = $1
		LEFT JOIN ` + TableGroupPermissions + ` AS gp ON gp.permission_id = p.ID AND gp.group_id = ug.group_id
		WHERE up.user_id = $1 OR ug.user_id = $1
	`

	rows, err := pr.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	permissions := []*PermissionEntity{}
	for rows.Next() {
		var p PermissionEntity
		err = rows.Scan(
			&p.Action,
			&p.CreatedAt,
			&p.ID,
			&p.Module,
			&p.Subsystem,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		permissions = append(permissions, &p)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return permissions, nil
}
