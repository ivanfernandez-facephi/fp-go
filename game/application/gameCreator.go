package application

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"fp-go/game/domain"
	"os"
)

type GameCreator struct {
	repository domain.GameRepository
}

func NewGameCreator(repository domain.GameRepository) *GameCreator {
	return &GameCreator{repository}
}

func (gc GameCreator) Execute(name string) error {
	game := domain.NewGame(name)
	fmt.Println(game.Id)

	err := gc.repository.Save(*game)
	if err != nil {
		return err
	}

	fileName := "games"
	saveExt := [...]string{"xml", "json"}
	for i := 0; i < len(saveExt); i++ {
		switch saveExt[i] {
		case "xml":
			xmlGame, err := xml.Marshal(game)
			if err != nil {
				return err
			}

			file, err := os.OpenFile(fileName+".xml", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
			defer file.Close()
			if err != nil {
				return err
			}

			_, err = file.Write(xmlGame)
			if err != nil {
				return err
			}
		case "json":
			jsonGame, err := json.Marshal(game)
			if err != nil {
				return err
			}

			file, err := os.OpenFile(fileName+".json", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
			defer file.Close()
			if err != nil {
				return err
			}

			_, err = file.Write(jsonGame)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
