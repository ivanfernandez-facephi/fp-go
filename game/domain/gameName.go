package domain

import "strings"

type GameName string

func NewGameName(value string) GameName {
	return GameName(value)
}

func (g GameName) String() string {
	return strings.ToUpper(string(g))
}
