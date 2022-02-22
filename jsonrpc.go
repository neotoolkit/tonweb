package tonweb

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type JsonRPC struct {
	*Client
}

func (jsonRPC JsonRPC) JsonRPC(b JsonRPCBody) (*JsonRPCResponse, error) {
	url := jsonRPC.url()
	url.Path = jsonRPC.BasePath + "/jsonRPC"

	body, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := jsonRPC.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res JsonRPCResponse

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

type JsonRPCBody struct {
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      string      `json:"id"`
	JsonRPC string      `json:"jsonrpc"`
}

type JsonRPCResponse struct {
	Ok      bool   `json:"ok"`
	Result  string `json:"result"`
	Error   string `json:"error"`
	Code    int    `json:"code"`
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
}
