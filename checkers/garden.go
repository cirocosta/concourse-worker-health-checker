package checkers

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

const containerPayloadFormat = `{"handle":"%s", "rootfs":"raw:///tmp/"}`

type Garden struct {
	Address string
}

func (g *Garden) Check(ctx context.Context) (err error) {
	handle := mustCreatedHandle()

	err = g.createContainer(ctx, handle)
	if err != nil {
		err = errors.Wrapf(err,
			"failed to create container %s", handle)
		return
	}

	err = g.destroyContainer(ctx, handle)
	if err != nil {
		err = errors.Wrapf(err,
			"failed to destroy container %s", handle)
		return
	}

	return
}

func (g *Garden) destroyContainer(ctx context.Context, handle string) (err error) {
	var (
		url    = g.Address + "/containers/" + handle
		method = http.MethodDelete
	)

	err = doRequest(ctx, method, url, nil)
	if err != nil {
		err = errors.Wrapf(err,
			"request failed")
		return
	}

	return
}

func (g *Garden) createContainer(ctx context.Context, handle string) (err error) {
	var (
		url    = g.Address + "/containers"
		method = http.MethodPost
		body   = bytes.NewBufferString(fmt.Sprintf(containerPayloadFormat, handle))
	)

	err = doRequest(ctx, method, url, body)
	if err != nil {
		err = errors.Wrapf(err,
			"failed to created container")
		return
	}

	return
}
