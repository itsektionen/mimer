package templates

import (
	"context"

	"github.com/itsektionen/mimer/app/ctxs"
	"github.com/itsektionen/mimer/model"
)

func GetUser(ctx context.Context) *model.UserClaims {
	user, ok := ctxs.UserFromContext(ctx)
	if !ok {
		return nil
	}

	return user
}
