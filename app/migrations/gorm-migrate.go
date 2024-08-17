package migrations

import (
	"encoding/json"
	"log"
	"os"
	"star-pos/features/user"
	userData "star-pos/features/user/repository"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	file, err := os.ReadFile("userdummy.json")
	if err != nil {
		log.Fatal("error reading file: ", err)
	}
	var users []user.UserCore
	err = json.Unmarshal(file, &users)
	if err != nil {
		log.Fatal("error unmarshalling JSON: ", err)
	}

	repositoryInterface := userData.New(db)

	for _, data := range users {
		err = repositoryInterface.Insert(data)
		if err != nil {
			log.Fatal("error insert data: ", err)
		}
	}
}
