package models

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// User is our simple struct
type User struct {
	ID            int64  `json:"id" gorm:"AUTO_INCREMENT"`
	Email         string `json:"email" gorm:"unique_index"`
	Firstname     string `json:"first_name"`
	MiddleInitial string `json:"middle_initial"`
	Lastname      string `json:"last_name"`
	Fullname      string `json:"full_name"`
	Address       string `json:"address"`
	Password      string `json:"-" `
	VerifyHash    string `json:"-"` // A field that checks if the password is correctly hashed
	PhoneNumber   string `json:"phone_number"`
	Gender        string `json:"gender"`
	Utime         int64  `json:"utime"`
}

// TableName ...
func (User) TableName() string {
	return "users"
}

// GetUserByID ....
func GetUserByID(id int64) (*User, error) {
	var user User
	err := db.Debug().Model(&user).Where("id=?", id).Scan(&user).Error
	return &user, err
}

// GetUserByEmail ....
func GetUserByEmail(email string) (*User, error) {
	var user User
	email = strings.Trim(strings.ToLower(email), "\r\n\t ")
	err := db.Debug().Model(&user).Where("email=?", email).Scan(&user).Limit(1).Error
	if err != nil {
		log.Print(err)
	}
	return &user, err
}

// NewUser ...
func NewUser() User {
	return User{
		Utime: time.Now().Unix(),
	}
}

// Create creates a new user
func (u *User) Create() error {
	u.Email = strings.ToLower(strings.Trim(u.Email, " \r\n\t "))

	u.Password = strings.Trim(u.Password, " \r\n\t ")

	if u.Email == "" || !strings.Contains(u.Email, "@") {
		return fmt.Errorf("Email is required")
	}
	if u.Password == "" || len(u.Password) < 6 {
		return fmt.Errorf("Password must have more than 6 characters")
	}
	uu, err := GetUserByEmail(u.Email)
	if uu.ID > 0 || err == nil {
		return fmt.Errorf("Email is already taken")
	}

	origPassword := u.Password
	u.Password = hashedPassword(origPassword)
	u.VerifyHash = hashedPassword(u.Password + u.Email + string(u.Utime) + "-")
	err = db.Debug().Model(&u).Create(&u).Error
	if err != nil {
		u.Password = origPassword
	}
	return err
}

// LoginUser ...
// func LoginUser(email, password string) (*User, error) {
// 	var user User

// 	err := db.Debug().Where("email = ?", email).Limit(1).First(&user).Error
// 	if err != nil {
// 		return &user, errors.New("Email not found. Please register")
// 	}

// 	err = db.Debug().Where("email = ?", email).Where("password = ?", hashedPassword(password)).Limit(1).First(&user).Error
// 	if err != nil {
// 		return &user, errors.New("Email or Password is incorrect, please try again")
// 	}

// 	if user.ID == 0 {
// 		return &user, errors.New("User not found")
// 	}

// 	return &user, err
// }

// Update ...
func (u *User) Update() error {
	return db.Debug().Model(&u).Update(&u).Error
}

// ChangeEmail handles changing of user email
func (u *User) ChangeEmail(newEmail string) error {
	u.Email = strings.ToLower(strings.Trim(newEmail, " \n\r"))
	return db.Debug().Model(&u).Exec("update users set verified = 0, email=? where id = ? limit 1 ", u.Email, u.ID).Error
}

// ChangePassword handles changing of user password
func (u *User) ChangePassword(newPassword string) error {
	newPassword = strings.Trim(newPassword, " \r\n\t ")

	if newPassword == "" || len(newPassword) < 6 {
		return fmt.Errorf("Password must have more than 6 characters")
	}

	origPassword := u.Password
	u.Password = hashedPassword(newPassword)

	err := db.Debug().Model(&u).Exec("update users set password = ? where id = ? limit 1 ", u.Password, u.ID).Error
	if err != nil {
		u.Password = origPassword
	}
	return err
}

// GetUsers ...
func GetUsers(limit int64, page int64) ([]User, error) {
	var users []User
	var err error
	st := page * limit
	err = db.Debug().Model(&User{}).Order("id desc").Limit(int(limit)).Offset(int(st)).Scan(&users).Error
	return users, err
}

// func hashedPassword(rawPassword string) string {
// 	s := sha256.New()
// 	s.Write([]byte(rawPassword))
// 	return base64.URLEncoding.EncodeToString(s.Sum(nil))
// }
