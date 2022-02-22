package main

import (
	"fmt"

	"github.com/neotoolkit/tonweb"
)

func main() {
	c := tonweb.NewClient()

	req, err := c.Blocks().GetBlockHeader(tonweb.GetBlockHeaderQuery{
		Workchain: 1,
		Shard:     1,
		Seqno:     1,
	})

	fmt.Println(req, err)
}
