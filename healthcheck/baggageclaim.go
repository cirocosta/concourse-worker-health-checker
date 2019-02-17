package healthcheck

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type Baggageclaim struct {
	Url string
}

const emptyStrategyPayloadFormat = `{"handle":"%s", "strategy":{"type":"empty"}}`

func (b *Baggageclaim) Check(ctx context.Context) error {
	handle, err := createHandle()
	if err != nil {
		return errors.Wrapf(err,
			"failed to create handle")
	}

	err = b.createVolume(ctx, handle)
	if err != nil {
		return errors.Wrapf(err,
			"failed to create volume %s", handle)
	}

	err = b.destroyVolume(ctx, handle)
	if err != nil {
		return errors.Wrapf(err,
			"failed to delete volume %s", handle)
	}

	return nil
}

func (b *Baggageclaim) destroyVolume(ctx context.Context, handle string) error {
	var (
		url    = b.Url + "/volumes/" + handle
		method = http.MethodDelete
	)

	err := doRequest(ctx, method, url, nil)
	if err != nil {
		return errors.Wrapf(err,
			"request failed")
	}

	return nil
}

func (b *Baggageclaim) createVolume(ctx context.Context, handle string) error {
	var (
		url    = b.Url + "/volumes"
		method = http.MethodPost
		body   = bytes.NewBufferString(fmt.Sprintf(emptyStrategyPayloadFormat, handle))
	)

	err := doRequest(ctx, method, url, body)
	if err != nil {
		return errors.Wrapf(err,
			"request failed")
	}

	return nil
}
