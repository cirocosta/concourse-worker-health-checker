package healthcheck

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

const containerPayloadFormat = `{"handle":"%s", "rootfs":"raw:///tmp/"}`

type Garden struct {
	Url string
}

func (g *Garden) Check(ctx context.Context) error {
	handle, err := createHandle()
	if err != nil {
		return errors.Wrapf(err,
			"failed to create handle")
	}

	err = g.createContainer(ctx, handle)
	if err != nil {
		return errors.Wrapf(err,
			"failed to create container %s", handle)
	}

	err = g.destroyContainer(ctx, handle)
	if err != nil {
		return errors.Wrapf(err,
			"failed to destroy container %s", handle)
	}

	return nil
}

func (g *Garden) destroyContainer(ctx context.Context, handle string) error {
	var (
		url    = g.Url + "/containers/" + handle
		method = http.MethodDelete
	)

	err := doRequest(ctx, method, url, nil)
	if err != nil {
		return errors.Wrapf(err,
			"request failed")
	}

	return nil
}

func (g *Garden) createContainer(ctx context.Context, handle string) error {
	var (
		url    = g.Url + "/containers"
		method = http.MethodPost
		body   = bytes.NewBufferString(fmt.Sprintf(containerPayloadFormat, handle))
	)

	err := doRequest(ctx, method, url, body)
	if err != nil {
		return errors.Wrapf(err,
			"failed to created container")
	}

	return nil
}
