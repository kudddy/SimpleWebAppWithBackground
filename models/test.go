package models

import "github.com/jinzhu/gorm"

//структура для учётной записи пользователя
type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

func GetUser(u uint) *Account {

	acc := &Account{}
	GetDB().Table("accounts").Where("id = ?", u).First(acc)
	if acc.Email == "" { //Пользователь не найден!
		return nil
	}

	acc.Password = ""
	return acc
}

func AddUser() {

	//GetDB().Model(&acc).Update("Token", 200)
	db.Create(&Account{Email: "jfkddf", Password: "fsdfdsfdsf", Token: "sdfdfdsfddfs"})

}