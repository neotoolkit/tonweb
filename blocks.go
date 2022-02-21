package tonweb

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Blocks struct {
	*Client
}

func (b Blocks) GetMasterchainInfo() (*Response, error) {
	url := b.url()
	url.Path = b.BasePath + "/getMasterchainInfo"

	resp, err := b.client.Get(url.String())
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

func (b Blocks) GetConsensusBlock() (*Response, error) {
	url := b.url()
	url.Path = b.BasePath + "/getConsensusBlock"

	resp, err := b.client.Get(url.String())
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

func (b Blocks) LookupBlock(q LookupBlockQuery) (*Response, error) {
	url := b.url()
	url.Path = b.BasePath + "/lookupBlock"
	url.RawQuery = fmt.Sprintf("workchain=%v&shard=%v&seqno=%v&lt=%v&unixtime=%v", q.Workchain, q.Shard, q.Seqno, q.LogicalTime, q.Unixtime)

	var req http.Request

	req.Method = http.MethodGet
	req.URL = url

	resp, err := b.client.Do(&req)
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

func (b Blocks) Shards(seqno int) (*Response, error) {
	url := b.url()
	url.Path = b.BasePath + "/shards"
	url.RawQuery = fmt.Sprintf("seqno=%v", seqno)

	var req http.Request

	req.Method = http.MethodGet
	req.URL = url

	resp, err := b.client.Do(&req)
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

func (b Blocks) GetBlockTransactions(q GetBlockTransactionsQuery) (*Response, error) {
	url := b.url()
	url.Path = b.BasePath + "/getBlockTransactions"

	if q.Count == 0 {
		q.Count = 40
	}

	url.RawQuery = fmt.Sprintf("workchain=%v&shard=%v&seqno=%v&root_hash=%s&file_hash=%s&after_lt=%v&after_hash=%s&count=%v", q.Workchain, q.Shard, q.Seqno, q.RootHash, q.FileHash, q.AfterLogicalTime, q.AfterHash, q.Count)

	var req http.Request

	req.Method = http.MethodGet
	req.URL = url

	resp, err := b.client.Do(&req)
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

func (b Blocks) GetBlockHeader(q GetBlockHeaderQuery) (*Response, error) {
	url := b.url()
	url.Path = b.BasePath + "/getBlockHeader"
	url.RawQuery = fmt.Sprintf("workchain=%v&shard=%v&seqno=%v&root_hash=%s&file_hash=%s", q.Workchain, q.Shard, q.Seqno, q.RootHash, q.FileHash)

	var req http.Request

	req.Method = http.MethodGet
	req.URL = url

	resp, err := b.client.Do(&req)
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

type LookupBlockQuery struct {
	Workchain   int
	Shard       int
	Seqno       int
	LogicalTime int
	Unixtime    int
}

type GetBlockTransactionsQuery struct {
	Workchain        int
	Shard            int
	Seqno            int
	RootHash         string
	FileHash         string
	AfterLogicalTime int
	AfterHash        string
	Count            int
}

type GetBlockHeaderQuery struct {
	Workchain int
	Shard     int
	Seqno     int
	RootHash  string
	FileHash  string
}
