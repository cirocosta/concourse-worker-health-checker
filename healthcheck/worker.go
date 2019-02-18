package healthcheck

import (
	"context"
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
		return err
	}

	err = w.ContainerProvider.Destroy(ctx, handle)
	if err != nil {
		return err
	}

	err = w.VolumeProvider.Destroy(ctx, rootfs.Handle)
	if err != nil {
		return err
	}

	return nil
}
