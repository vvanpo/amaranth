package cmf

import (
	"context"
)

type Resource interface {
	Router
	Serve(context.Context, *Exchange) error
}
