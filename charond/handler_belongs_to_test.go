package charond

import (
	"testing"

	"github.com/piotrkowalczuk/charon"
	"github.com/piotrkowalczuk/charon/charonrpc"
	"github.com/piotrkowalczuk/charon/internal/model"
)

func TestBelongsToHandler_firewall_success(t *testing.T) {
	data := []struct {
		req charonrpc.BelongsToRequest
		act actor
	}{
		{
			req: charonrpc.BelongsToRequest{UserId: 1},
			act: actor{
				user: &model.UserEntity{ID: 2},
				permissions: charon.Permissions{
					charon.UserGroupCanCheckBelongingAsStranger,
				},
			},
		},

		{
			req: charonrpc.BelongsToRequest{UserId: 1},
			act: actor{
				user: &model.UserEntity{ID: 2, IsSuperuser: true},
			},
		},

		{
			req: charonrpc.BelongsToRequest{UserId: 1},
			act: actor{
				user: &model.UserEntity{ID: 1},
			},
		},
	}

	h := &belongsToHandler{}
	for _, d := range data {
		if err := h.firewall(&d.req, &d.act); err != nil {
			t.Errorf("unexpected error: %s", err.Error())
		}
	}
}

func TestBelongsToHandler_firewall_failure(t *testing.T) {
	data := []struct {
		req charonrpc.BelongsToRequest
		act actor
	}{
		{
			req: charonrpc.BelongsToRequest{UserId: 1},
			act: actor{
				user: &model.UserEntity{ID: 2},
			},
		},
		{
			req: charonrpc.BelongsToRequest{UserId: 1},
			act: actor{
				user: &model.UserEntity{
					ID:      2,
					IsStaff: true,
				},
			},
		},
	}

	h := &belongsToHandler{}
	for _, d := range data {
		if err := h.firewall(&d.req, &d.act); err == nil {
			t.Error("expected error, got nil")
		}
	}
}
