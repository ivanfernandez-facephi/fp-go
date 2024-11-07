package domain

type GameRepository interface {
	Save(value Game) error
	FindById(id GameId) (*Game, error)
}
