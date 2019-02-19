package healthcheck

import (
	"context"

	"github.com/hashicorp/go-multierror"
)

type Worker struct {
	ContainerProvider ContainerProvider
	VolumeProvider    VolumeProvider
}

func (w *Worker) Check(ctx context.Context) error {
	handle, err := createHandle()
	if err != nil {
		return err
	}

	rootfs, err := w.VolumeProvider.Create(ctx, handle)
	if err != nil {
		return err
	}

	err = w.ContainerProvider.Create(ctx, handle, rootfs.Path)
	if err != nil {
		volDestructionErr := w.VolumeProvider.Destroy(ctx, handle)
		if volDestructionErr != nil {
			return multierror.Append(err, volDestructionErr)
		}

		return err
	}

	err = w.ContainerProvider.Destroy(ctx, handle)
	if err != nil {
		return err
	}

	err = w.VolumeProvider.Destroy(ctx, handle)
	if err != nil {
		return err
	}

	return nil
}
