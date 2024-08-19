package migrations

import (
	"encoding/json"
	"log"
	"os"
	userModel "star-pos/features/user/model"
	"star-pos/features/user/repository"
)

func InitMigration() {
	file, err := os.ReadFile("userdummy.json")
	if err != nil {
		log.Fatal("error reading file: ", err)
	}
	var users []userModel.User
	err = json.Unmarshal(file, &users)
	if err != nil {
		log.Fatal("error unmarshalling JSON: ", err)
	}

	for _, data := range users {
		err = repository.Insert(&data)
		if err != nil {
			log.Fatal("error insert data: ", err)
		}
	}
}
