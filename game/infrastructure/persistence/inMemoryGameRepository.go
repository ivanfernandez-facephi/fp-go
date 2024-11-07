package persistence

import "fp-go/game/domain"

type InMemoryGameRepository struct {
	data map[string]domain.Game
}

func NewInMemoryGameRepository() *InMemoryGameRepository {
	return &InMemoryGameRepository{
		data: make(map[string]domain.Game),
	}
}

func (r *InMemoryGameRepository) Save(value domain.Game) error {
	r.data[string(value.Id)] = value
	return nil
}

func (r *InMemoryGameRepository) FindById(id domain.GameId) (*domain.Game, error) {
	game, exists := r.data[string(id)]
	if !exists {
		return nil, nil
	}

	return &game, nil
}
