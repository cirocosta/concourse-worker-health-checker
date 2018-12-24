package checkers

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func doRequest(ctx context.Context, method, url string, body io.Reader) (err error) {
	var (
		req  *http.Request
		resp *http.Response
	)

	req, err = http.NewRequest(method, url, body)
	if err != nil {
		return
	}

	req = req.WithContext(ctx)
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		responseMessage, _ := ioutil.ReadAll(resp.Body)
		err = errors.Errorf("non-success status code %d - %s",
			resp.StatusCode, string(responseMessage))
		return
	}

	return
}
