package models

import (
	"errors"
	"strings"
	"time"
)

// Maintenance simple blue-print
// that represents an maintenance
// of a property
type Maintenance struct {
	ID           int64     `json:"id" gorm:"AUTO_INCREMENT"`
	LastName     string    `json:"last_name"`
	FirstName    string    `json:"first_name"`
	EmailAddress string    `json:"email_address"`
	PhoneNumber  string    `json:"phone_number"`
	Birthdate    string    `json:"birthdate"`
	Password     string    `json:"-"`
	Age          int       `json:"age"`
	Status       bool      `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// TableName name of our table
func (Maintenance) TableName() string {
	return "maintenance"
}

// NewMaintenance constructor
func NewMaintenance() *Maintenance {
	return &Maintenance{}
}

// Create creates a new owner
func (m *Maintenance) Create() error {
	m.EmailAddress = strings.ToLower(strings.Trim(m.EmailAddress, " \r\n\t "))

	m.Password = strings.Trim(m.Password, " \r\n\t ")

	if m.EmailAddress == "" || !strings.Contains(m.EmailAddress, "@") {
		return errors.New("Email is required")
	}
	if m.Password == "" || len(m.Password) < 6 {
		return errors.New("Password must have more than 6 characters")
	}
	aa, err := GetAccountByEmailAddress(m.EmailAddress)
	if aa.ID > 0 || err == nil {
		return errors.New("Email is already taken")
	}

	origPassword := m.Password
	m.Password = hashedPassword(origPassword)
	err = db.Debug().Model(&m).Create(&m).Error
	if err != nil {
		m.Password = origPassword
	}
	return err
}

// LoginMaintenance logins a maintenance personnel
func LoginMaintenance(email, password string) (*Maintenance, error) {
	var maintenance Maintenance

	err := db.Debug().Where("email_address = ?", email).Limit(1).First(&maintenance).Error
	if err != nil {
		return nil, errors.New("Email not found. Please register")
	}

	err = db.Debug().Where("email_address = ?", email).Where("password = ?", hashedPassword(password)).Limit(1).First(&maintenance).Error
	if err != nil {
		return nil, errors.New("Email or Password is incorrect, please try again")
	}

	if maintenance.ID == 0 {
		return nil, errors.New("User not found")
	}

	return &maintenance, err
}

// GetMaintenanceByID returns an owner specified by the ID
func GetMaintenanceByID(id int64) (*Maintenance, error) {
	var maintenance Maintenance
	if id < 0 {
		return nil, errors.New("ID is required")
	}
	err := db.Debug().Model(&maintenance).Where("id=?", id).Scan(&maintenance).Error
	return &maintenance, err
}
