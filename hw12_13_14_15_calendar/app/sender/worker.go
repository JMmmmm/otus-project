package sender

import "context"

type Worker interface {
	Execute(ctx context.Context, threads int) error
}
