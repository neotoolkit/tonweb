package tonweb

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Run struct {
	*Client
}

func (r Run) RunGetMethod(b RunGetMethodBody) (*Response, error) {
	url := r.url()
	url.Path = r.BasePath + "/runGetMethod"

	body, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res Response

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

type RunGetMethodBody struct {
	Address string     `json:"address"`
	Method  string     `json:"method"`
	Stack   [][]string `json:"stack"`
}
