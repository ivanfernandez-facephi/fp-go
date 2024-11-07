package main

import (
	"fp-go/game/application"
	game_http "fp-go/game/infrastructure/http"
	game_persistence "fp-go/game/infrastructure/persistence"
	"net/http"
)

func main() {
	repo := game_persistence.NewInMemoryGameRepository()
	finder := application.NewGameFinder(repo)
	creator := application.NewGameCreator(repo)
	getGameController := game_http.NewGameGetController(finder)
	putGameController := game_http.NewGamePutController(creator)

	http.Handle("/game/find", getGameController)
	http.Handle("/game/create", putGameController)

	http.ListenAndServe(":8282", nil)
}
