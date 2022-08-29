package models

type User struct {
	Id       uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null;column:id"`
	UserName string `json:"userName" gorm:"column:user_name"`
	Password string `json:"password" gorm:"column:password"`
}

func (User) TableName() string {
	return "user"
}

func GetUserById(id string) (user User, err error) {
	db.
		Model(&User{}).
		Where("id = %s", id).First(&user)

	return
}
