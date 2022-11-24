package seeds

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/akmalcm/team-soccer-go/models"
)

type PlayerData struct {
	PlayerData []models.Player `json:"PlayerData"`
}

func (s Seed) PlayersSeed() {
	file, err := ioutil.ReadFile("SoccerWiki_2022-11-24 - Player Data_1669265732.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	data := PlayerData{}
	_ = json.Unmarshal([]byte(file), &data)

	s.db.CreateInBatches(&data.PlayerData, 1000)
}
