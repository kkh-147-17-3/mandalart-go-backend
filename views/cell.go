package views

import (
	"encoding/json"
	"fmt"
	"mandalart.com/services"
	"net/http"
)

type CellView struct {
	c *services.CellService
	t *services.TodoService
}

type CreateTodos struct {
	Goal        string `json:"goal"`
	Color       string `json:"color"`
	IsCompleted bool   `json:"isCompleted"`
	Todo		[]services.Todo `json:"todos"`
}

func NewCellView(c *services.CellService, t *services.TodoService) *CellView {
	return &CellView{c, t}
}

func (v *CellView) UpdateCell(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value("userID").(int)

	if !ok {
		Respond(w, r, http.StatusInternalServerError, fmt.Errorf("user id not found"))
		return
	}

	cellID, ok := ctx.Value("cellID").(int)
	if !ok {
		Respond(w, r, http.StatusBadRequest, fmt.Errorf("cell id not found"))
	}

	var reqBody CreateTodos
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		Respond(w, r, http.StatusUnprocessableEntity, err)
	}

	args := make([]services.InsertTodosInput, len(reqBody.Todo))
	for i, el := range reqBody.Todo {
		args[i] = services.InsertTodosInput{OwnerID: &userID, CellID: &cellID, Content: el.Content}
	}

	result, err := v.t.InsertTodos(r.Context(), userID, cellID, args)

	if err != nil {
		Respond(w, r, http.StatusInternalServerError, err)
	}

	Respond(w, r, http.StatusOK, result)
}
