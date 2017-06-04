package responses


// Represents the Object within the Conduit transaction response
type ObjectResponse struct {
	// Conduit inconsistently returns string vs int values for the Id field, depending on whether it's a create
	// or edit operation
	// Id   string `json:"id"`
	PHID string `json:"phid"`
}

// Probably contains the list of unsuccessful transactions
type Transaction struct {
	TransactionType string `json:"type"`
	Value           interface{} `json:"value"`
}

// ManiphestEditRequest represents a request to maniphest.edit.
type EditEndpointResponse struct {
	Object       ObjectResponse        `json:"object"`
	Transactions []Transaction `json:"transactions"`
}