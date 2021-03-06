package charond

import (
	"testing"

	"github.com/piotrkowalczuk/charon"
	"github.com/piotrkowalczuk/charon/charonrpc"
	"github.com/piotrkowalczuk/charon/internal/model"
)

func TestListGroupsHandler_firewall_success(t *testing.T) {
	data := []struct {
		req charonrpc.ListGroupsRequest
		act actor
	}{
		{
			req: charonrpc.ListGroupsRequest{},
			act: actor{
				user: &model.UserEntity{ID: 1},
				permissions: charon.Permissions{
					charon.GroupCanRetrieve,
				},
			},
		},
		{
			req: charonrpc.ListGroupsRequest{},
			act: actor{
				user: &model.UserEntity{ID: 2, IsSuperuser: true},
			},
		},
	}

	h := &listGroupsHandler{}
	for _, d := range data {
		if err := h.firewall(&d.req, &d.act); err != nil {
			t.Errorf("unexpected error: %s", err.Error())
		}
	}
}

func TestListGroupsHandler_firewall_failure(t *testing.T) {
	data := []struct {
		req charonrpc.ListGroupsRequest
		act actor
	}{
		{
			req: charonrpc.ListGroupsRequest{},
			act: actor{
				user: &model.UserEntity{ID: 1},
			},
		},
		{
			req: charonrpc.ListGroupsRequest{},
			act: actor{
				user: &model.UserEntity{
					ID:      2,
					IsStaff: true,
				},
			},
		},
	}

	h := &listGroupsHandler{}
	for _, d := range data {
		if err := h.firewall(&d.req, &d.act); err == nil {
			t.Error("expected error, got nil")
		}
	}
}
