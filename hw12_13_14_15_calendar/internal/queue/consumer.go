package queue

import (
	"context"
)

type Consumer interface {
	Handle(ctx context.Context, fn interface{}, threads int) error
}
