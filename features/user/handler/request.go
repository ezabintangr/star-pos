package handler

type RequestAccount struct {
	PhoneNumber     string `gorm:"unique" json:"phone_number" form:"phone_number"`
	Password        string `json:"password" form:"password"`
	PasswordConfirm string `json:"password_confirm" form:"password_confirm"`
}
