package migrations

import (
	"encoding/json"
	"log"
	"os"
	userData "star-pos/features/user/repository"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	file, err := os.ReadFile("userdummy.json")
	if err != nil {
		log.Fatal("error reading file: ", err)
	}
	var users []userData.User
	err = json.Unmarshal(file, &users)
	if err != nil {
		log.Fatal("error unmarshalling JSON: ", err)
	}

	for _, data := range users {
		err = userData.Insert(&data)
		if err != nil {
			log.Fatal("error insert data: ", err)
		}
	}
}
