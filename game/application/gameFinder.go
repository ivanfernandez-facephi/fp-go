package application

import "fp-go/game/domain"

type GameFinder struct {
	repository domain.GameRepository
}

func NewGameFinder(repository domain.GameRepository) *GameFinder {
	return &GameFinder{repository}
}

func (gf GameFinder) Execute(id string) (*domain.Game, error) {
	gameId, err := domain.NewGameId(id)
	if err != nil {
		return nil, err
	}

	game, err := gf.repository.FindById(gameId)
	if err != nil {
		return nil, err
	}

	return game, nil
}
