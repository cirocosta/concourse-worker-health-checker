package healthcheck

import (
	"context"
)

type ContainerProvider interface {
	Create(ctx context.Context, handle, rootfs string) (err error)
	Destroy(ctx context.Context, handle string) (err error)
}
