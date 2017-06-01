package requests

type Transaction struct {
	TransactionType string `json:"type"`
	Value           interface{} `json:"value"`

}

// ManiphestEditRequest represents a request to maniphest.edit.
type ManiphestEditTaskRequest struct {
	ObjectIdentifier string        `json:"objectIdentifier"`
	Transactions     []Transaction `json:"transactions"`
	Request
}
