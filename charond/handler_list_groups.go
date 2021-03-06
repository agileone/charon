package charond

import (
	"github.com/piotrkowalczuk/charon"
	"github.com/piotrkowalczuk/charon/charonrpc"
	"github.com/piotrkowalczuk/charon/internal/model"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type listGroupsHandler struct {
	*handler
}

func (lgh *listGroupsHandler) List(ctx context.Context, req *charonrpc.ListGroupsRequest) (*charonrpc.ListGroupsResponse, error) {
	act, err := lgh.retrieveActor(ctx)
	if err != nil {
		return nil, err
	}
	if err = lgh.firewall(req, act); err != nil {
		return nil, err
	}

	ents, err := lgh.repository.group.Find(&model.GroupCriteria{
		Offset: req.Offset.Int64Or(0),
		Limit:  req.Limit.Int64Or(10),
	})
	if err != nil {
		return nil, err
	}

	return lgh.response(ents)
}

func (lgh *listGroupsHandler) firewall(req *charonrpc.ListGroupsRequest, act *actor) error {
	if act.user.IsSuperuser {
		return nil
	}
	if act.permissions.Contains(charon.GroupCanRetrieve) {
		return nil
	}

	return grpc.Errorf(codes.PermissionDenied, "list of groups cannot be retrieved, missing permission")
}

func (lgh *listGroupsHandler) response(ents []*model.GroupEntity) (*charonrpc.ListGroupsResponse, error) {
	resp := &charonrpc.ListGroupsResponse{
		Groups: make([]*charonrpc.Group, 0, len(ents)),
	}
	var (
		err error
		msg *charonrpc.Group
	)
	for _, e := range ents {
		if msg, err = e.Message(); err != nil {
			return nil, err
		}
		resp.Groups = append(resp.Groups, msg)
	}

	return resp, nil
}
