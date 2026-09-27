package requests

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

func DoGetRequest(ctx context.Context, URL string, out any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL, nil)
	if err != nil {
		return err
	}

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		if err := json.NewDecoder(res.Body).Decode(out); err != nil {
			return err
		}
		return nil
	}

	return errors.New("Failed to retrieve the response")
}
