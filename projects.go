package gonduit

import (
	"github.com/danieldanciu/gonduit/entities"
	"github.com/danieldanciu/gonduit/requests"
	"github.com/danieldanciu/gonduit/responses"
)

// ProjectQuery performs a call to project.query.
func (c *Conn) ProjectQuery(
	req requests.ProjectQueryRequest,
) (*responses.ProjectQueryResponse, error) {
	var res responses.ProjectQueryResponse

	if err := c.Call("project.query", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ProjectEdit performs a call to project.edit.
func (c *Conn) ProjectEdit(
	req requests.EditEndpointRequest,
) (*entities.Project, error) {
	var res entities.Project

	if err := c.Call("project.edit", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

