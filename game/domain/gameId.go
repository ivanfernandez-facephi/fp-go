package domain

import "github.com/google/uuid"

type GameId string

func NewGameId(value string) (GameId, error) {
	res, err := uuid.Parse(value)
	if err != nil {
		return "", err
	}

	return GameId(res.String()), nil
}

func NewRandomGameId() GameId {
	value := uuid.New()

	return GameId(value.String())
}
