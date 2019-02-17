package healthcheck

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func doRequest(ctx context.Context, method, url string, body io.Reader) error {
	var (
		req  *http.Request
		resp *http.Response
		err  error
	)

	req, err = http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		responseMessage, _ := ioutil.ReadAll(resp.Body)

		return fmt.Errorf("non-success status code %d - %s",
			resp.StatusCode, string(responseMessage))
	}

	return nil
}