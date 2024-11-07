package xhttp

import (
	"encoding/json"
	"fmt"
	"fp-go/game/application"
	"io"
	"net/http"
)

type GamePutController struct {
	creator *application.GameCreator
}

type CreateGameDto struct {
	Name string `json:"name"`
}

func NewGamePutController(creator *application.GameCreator) *GamePutController {
	return &GamePutController{creator}
}

func (c GamePutController) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	reqBytes, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	var dto CreateGameDto
	err = json.Unmarshal(reqBytes, &dto)
	if err != nil {
		fmt.Println("error unmarshaling to dto", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.creator.Execute(dto.Name)
	if err != nil {
		fmt.Println("error creating game", err)
		rw.WriteHeader(http.StatusInternalServerError)
	}
}
