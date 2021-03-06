package charond

import (
	"database/sql"

	"github.com/piotrkowalczuk/charon"
	"github.com/piotrkowalczuk/charon/charonrpc"
	"github.com/piotrkowalczuk/charon/internal/model"
	"github.com/piotrkowalczuk/ntypes"
	"github.com/piotrkowalczuk/pqt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type modifyUserHandler struct {
	*handler
}

func (muh *modifyUserHandler) Modify(ctx context.Context, req *charonrpc.ModifyUserRequest) (*charonrpc.ModifyUserResponse, error) {
	if req.Id <= 0 {
		return nil, grpc.Errorf(codes.InvalidArgument, "user cannot be modified, invalid ID: %d", req.Id)
	}

	muh.loggerWith("user_id", req.Id)

	act, err := muh.retrieveActor(ctx)
	if err != nil {
		return nil, err
	}

	ent, err := muh.repository.user.FindOneByID(req.Id)
	if err != nil {
		return nil, err
	}

	if hint, ok := muh.firewall(req, ent, act); !ok {
		return nil, grpc.Errorf(codes.PermissionDenied, hint)
	}

	ent, err = muh.repository.user.UpdateOneByID(req.Id, &model.UserPatch{
		FirstName:   req.FirstName,
		IsActive:    req.IsActive,
		IsConfirmed: req.IsConfirmed,
		IsStaff:     req.IsStaff,
		IsSuperuser: req.IsSuperuser,
		LastName:    req.LastName,
		Password:    req.SecurePassword,
		UpdatedBy:   &ntypes.Int64{Int64: act.user.ID, Valid: act.user.ID != 0},
		Username:    req.Username,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, grpc.Errorf(codes.NotFound, "group does not exists")
		}
		switch pqt.ErrorConstraint(err) {
		case model.TableUserConstraintUsernameUnique:
			return nil, grpc.Errorf(codes.AlreadyExists, "user with such username already exists")
		default:
			return nil, err
		}
	}

	return muh.response(ent)
}

func (muh *modifyUserHandler) firewall(req *charonrpc.ModifyUserRequest, ent *model.UserEntity, actor *actor) (string, bool) {
	isOwner := actor.user.ID == ent.ID

	if !actor.user.IsSuperuser {
		switch {
		case ent.IsSuperuser:
			return "only superuser can modify a superuser account", false
		case ent.IsStaff && !isOwner && actor.permissions.Contains(charon.UserCanModifyStaffAsStranger):
			return "missing permission to modify an account as a stranger", false
		case ent.IsStaff && isOwner && actor.permissions.Contains(charon.UserCanModifyStaffAsOwner):
			return "missing permission to modify an account as an owner", false
		case req.IsSuperuser != nil && req.IsSuperuser.Valid:
			return "only superuser can change existing account to superuser", false
		case req.IsStaff != nil && req.IsStaff.Valid && !actor.permissions.Contains(charon.UserCanCreateStaff):
			return "user is not allowed to create user with is_staff property that has custom value", false
		}
	}

	return "", true
}

func (muh *modifyUserHandler) response(u *model.UserEntity) (*charonrpc.ModifyUserResponse, error) {
	msg, err := u.Message()
	if err != nil {
		return nil, err
	}
	return &charonrpc.ModifyUserResponse{
		User: msg,
	}, nil
}
