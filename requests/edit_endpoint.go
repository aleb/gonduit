package requests

type Transaction struct {
	TransactionType string `json:"type"`
	Value           interface{} `json:"value"`
}

// EditEndpointRequest represents a Conduit transaction request, such as maniphest.edit
type EditEndpointRequest struct {
	ObjectIdentifier string        `json:"objectIdentifier"`
	Transactions     []Transaction `json:"transactions"`
	Request
}

