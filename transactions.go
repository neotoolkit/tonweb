package tonweb

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Transactions struct {
	*Client
}

func (t Transactions) tmp() (*Response, error) {
	return nil, nil
}

func (t Transactions) GetTransactions(q GetTransactionsQuery) (*Response, error) {
	url := t.url()
	url.Path = t.BasePath + "/getTransactions"
	url.RawQuery = fmt.Sprintf("address=%s&limit=%v&lt=%v&hash=%v&to_lt=%v&archival=%v", q.Address, q.Limit, q.LogicalTime, q.Hash, q.ToLogicalTime, q.Archival)

	var req http.Request

	req.Method = http.MethodGet
	req.URL = url

	resp, err := t.client.Do(&req)
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

func (t Transactions) GetBlockTransactions(q GetBlockTransactionsQuery) (*Response, error) {
	url := t.url()
	url.Path = t.BasePath + "/getBlockTransactions"

	if q.Count == 0 {
		q.Count = 40
	}

	url.RawQuery = fmt.Sprintf("workchain=%v&shard=%v&seqno=%v&root_hash=%s&file_hash=%s&after_lt=%v&after_hash=%s&count=%v", q.Workchain, q.Shard, q.Seqno, q.RootHash, q.FileHash, q.AfterLogicalTime, q.AfterHash, q.Count)

	var req http.Request

	req.Method = http.MethodGet
	req.URL = url

	resp, err := t.client.Do(&req)
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

func (t Transactions) TryLocateTx(q TryLocateTxQuery) (*Response, error) {
	url := t.url()
	url.Path = t.BasePath + "/tryLocateTx"
	url.RawQuery = fmt.Sprintf("source=%s&destination=%s&created_lt=%v", q.Source, q.Destination, q.CreatedLogicalTime)

	var req http.Request

	req.Method = http.MethodGet
	req.URL = url

	resp, err := t.client.Do(&req)
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

func (t Transactions) TryLocateResultTx(q TryLocateResultTxQuery) (*Response, error) {
	url := t.url()
	url.Path = t.BasePath + "/tryLocateResultTx"
	url.RawQuery = fmt.Sprintf("source=%s&destination=%s&created_lt=%v", q.Source, q.Destination, q.CreatedLogicalTime)

	var req http.Request

	req.Method = http.MethodGet
	req.URL = url

	resp, err := t.client.Do(&req)
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

func (t Transactions) TryLocateSourceTx(q TryLocateSourceTxQuery) (*Response, error) {
	url := t.url()
	url.Path = t.BasePath + "/tryLocateSourceTx"
	url.RawQuery = fmt.Sprintf("source=%s&destination=%s&created_lt=%v", q.Source, q.Destination, q.CreatedLogicalTime)

	var req http.Request

	req.Method = http.MethodGet
	req.URL = url

	resp, err := t.client.Do(&req)
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

type TryLocateTxQuery struct {
	Source             string
	Destination        string
	CreatedLogicalTime int
}

type TryLocateResultTxQuery struct {
	Source             string
	Destination        string
	CreatedLogicalTime int
}

type TryLocateSourceTxQuery struct {
	Source             string
	Destination        string
	CreatedLogicalTime int
}
