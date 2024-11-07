package domain

type Game struct {
	Id   GameId   `json:"id" xml:"iD"`
	Name GameName `json:"name" xml:"name"`
}

func NewGame(name string) *Game {
	return &Game{
		Id:   NewRandomGameId(),
		Name: NewGameName(name),
	}
}
