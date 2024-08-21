package migrations

import (
	"encoding/json"
	"log"
	"os"
	"star-pos/app/databases"
	userModel "star-pos/features/user/model"
	"star-pos/features/user/repository"

	"gorm.io/gorm"
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
		var existingUser userModel.User
		result := databases.DB.Where("phone_number = ?", data.PhoneNumber).First(&existingUser)
		if result.Error != nil {
			log.Fatal("error check data: ", result.Error)
		}

		if result.Error == gorm.ErrRecordNotFound {
			err = repository.Insert(&data)
			if err != nil {
				log.Fatal("error insert data: ", err)
			}
		} else {
			log.Println("table already exist")
		}
	}
}
