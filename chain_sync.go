package ogmiosgo

const (
	FindIntersection = "findIntersection"
	NextBlock        = "nextBlock"
)

type Point struct {
	Slot uint64 `json:"slot"`
	Id   string `json:"id"`
}

type FindIntersectionParams struct {
	Points []interface{} `json:"points"`
}

func (c *Client) FindIntersection(params FindIntersectionParams) error {
	p := RPCParams{
		JsonRPC: "2.0",
		Method:  FindIntersection,
		Params:  params,
	}

	return c.conn.WriteJSON(p)
}

func (c *Client) NextBlock() error {
	p := RPCParams{
		JsonRPC: "2.0",
		Method:  NextBlock,
	}

	return c.conn.WriteJSON(p)
}
