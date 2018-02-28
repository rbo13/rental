package models

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"strings"
	"time"
)

// Account object
type Account struct {
	ID           int64     `json:"id" gorm:"AUTO_INCREMENT"`
	EmailAddress string    `json:"email_address"`
	Password     string    `json:"-"`
	IsAdmin      bool      `json:"is_admin"`
	Utime        int64     `json:"utime"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// GetAccountByID gets account by ID
func GetAccountByID(id int64) (*Account, error) {
	var account Account
	err := db.Debug().Model(&account).Where("id=?", id).Scan(&account).Error
	return &account, err
}

// GetAccountByEmailAddress gets account by ID
func GetAccountByEmailAddress(emailAddress string) (*Account, error) {
	var account Account
	err := db.Debug().Model(&account).Where("email_address=?", emailAddress).Scan(&account).Error
	return &account, err
}

// GetAccounts ...
func GetAccounts(limit int64, page int64) ([]Account, error) {
	var accounts []Account
	var err error
	st := page * limit
	err = db.Debug().Model(&Account{}).Order("id desc").Limit(int(limit)).Offset(int(st)).Scan(&accounts).Error
	return accounts, err
}

// NewAccount contructor
func NewAccount() *Account {
	return &Account{
		Utime:   time.Now().Unix(),
		IsAdmin: false,
	}
}

// LoginUser logins a user
func LoginUser(email, password string) (*Account, error) {
	var account Account

	err := db.Debug().Where("email_address = ?", email).Limit(1).First(&account).Error
	if err != nil {
		return &account, errors.New("Email not found. Please register")
	}

	err = db.Debug().Where("email_address = ?", email).Where("password = ?", hashedPassword(password)).Limit(1).First(&account).Error
	if err != nil {
		return &account, errors.New("Email or Password is incorrect, please try again")
	}

	if account.ID == 0 {
		return &account, errors.New("User not found")
	}

	return &account, err
}

// Create creates new record of tenants
func (a *Account) Create() error {
	a.EmailAddress = strings.ToLower(strings.Trim(a.EmailAddress, " \r\n\t "))

	a.Password = strings.Trim(a.Password, " \r\n\t ")

	if a.EmailAddress == "" || !strings.Contains(a.EmailAddress, "@") {
		return errors.New("Email is required")
	}
	if a.Password == "" || len(a.Password) < 6 {
		return errors.New("Password must have more than 6 characters")
	}
	aa, err := GetAccountByEmailAddress(a.EmailAddress)
	if aa.ID > 0 || err == nil {
		return errors.New("Email is already taken")
	}

	origPassword := a.Password
	a.Password = hashedPassword(origPassword)
	err = db.Debug().Model(&a).Create(&a).Error
	if err != nil {
		a.Password = origPassword
	}
	return err
}

// Update updates the account
func (a *Account) Update() error {
	return db.Debug().Model(&a).Where("id=?", a.ID).Update(&a).Error
}

func hashedPassword(rawPassword string) string {
	s := sha256.New()
	s.Write([]byte(rawPassword))
	return base64.URLEncoding.EncodeToString(s.Sum(nil))
}
