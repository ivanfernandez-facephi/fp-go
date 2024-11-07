package xhttp

import (
	"encoding/json"
	"fmt"
	"fp-go/game/application"
	"net/http"
)

type GameGetController struct {
	finder *application.GameFinder
}

func NewGameGetController(finder *application.GameFinder) *GameGetController {
	return &GameGetController{finder}
}

func (ggc GameGetController) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")

	game, err := ggc.finder.Execute(id)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(game)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)

		return
	}

	rw.Write(res)
}
