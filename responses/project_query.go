package responses

import (
	"github.com/danieldanciu/gonduit/entities"
	"encoding/json"
	"fmt"
)

// ProjectQueryResponse represents a response from calling project.query.
type ProjectQueryResponse struct {
	Data    map[string]entities.Project `json:"data"`
	SlugMap json.RawMessage           `json:"slugMap"`
	Cursor  entities.Cursor             `json:"cursor"`
}

type projectQueryResponse ProjectQueryResponse

func (pqr *ProjectQueryResponse) UnmarshalJSON(data []byte) error {
	res := projectQueryResponse{}
	if err := json.Unmarshal(data, &res); err == nil {
		*pqr = ProjectQueryResponse(res)
		return nil
	} else {

		fmt.Printf("Error is %v", err)
	}
	// check to see if we maybe have an empty response
	resEmpty := struct {
		Data json.RawMessage `json:"data"`
	}{}
	if err := json.Unmarshal(data, &resEmpty); err == nil {
		fmt.Printf("Raw message is %v", resEmpty.Data)
		*pqr = ProjectQueryResponse{}
		return nil
	} else {
		return err
	}

}
