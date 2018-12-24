package checkers

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type Baggageclaim struct {
	Address string
}

const emptyStrategyPayloadFormat = `{"handle":"%s", "strategy":{"type":"empty"}}`

func (b *Baggageclaim) Check(ctx context.Context) (err error) {
	handle := mustCreatedHandle()

	err = b.createVolume(ctx, handle)
	if err != nil {
		err = errors.Wrapf(err,
			"failed to create volume %s", handle)
		return
	}

	err = b.destroyVolume(ctx, handle)
	if err != nil {
		err = errors.Wrapf(err,
			"failed to delete volume %s", handle)
		return
	}

	return
}

func (b *Baggageclaim) destroyVolume(ctx context.Context, handle string) (err error) {
	var (
		url    = b.Address + "/volumes/" + handle
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

func (b *Baggageclaim) createVolume(ctx context.Context, handle string) (err error) {
	var (
		url    = b.Address + "/volumes"
		method = http.MethodPost
		body   = bytes.NewBufferString(fmt.Sprintf(emptyStrategyPayloadFormat, handle))
	)

	err = doRequest(ctx, method, url, body)
	if err != nil {
		err = errors.Wrapf(err,
			"request failed")
		return
	}

	return
}
