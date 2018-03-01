package models

import (
	"errors"
	"time"

	"github.com/manveru/faker"
)

// Owner simple blue-print
// that represents an owner
// of a property
type Owner struct {
	ID                  int64     `json:"id" gorm:"AUTO_INCREMENT"`
	LastName            string    `json:"last_name"`
	FirstName           string    `json:"first_name"`
	EmailAddress        string    `json:"email_address"`
	PhoneNumber         string    `json:"phone_number"`
	Birthdate           string    `json:"birthdate"`
	TypeOfPropertyOwned string    `json:"type_of_property_owned"`
	Age                 int       `json:"age"`
	NoOfPropertyOwned   int64     `json:"no_of_property_owned"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

// TableName name of our table
func (Owner) TableName() string {
	return "owners"
}

// NewOwner constructor
func NewOwner(lastName, firstName, emailAddress, phoneNumber, birthDate, typeOfPropertyOwned string, age int, noOfPropertyOwned int64) *Owner {
	return &Owner{
		LastName:            lastName,
		FirstName:           firstName,
		EmailAddress:        emailAddress,
		PhoneNumber:         phoneNumber,
		Birthdate:           birthDate,
		TypeOfPropertyOwned: typeOfPropertyOwned,
		Age:                 age,
		NoOfPropertyOwned:   noOfPropertyOwned,
	}
}

// Create creates a new owner
func (o *Owner) Create() error {
	err := db.Debug().Model(&o).Create(&o).Error
	if err != nil {
		return err
	}
	return err
}

// GetOwnerByID returns an owner specified by the ID
func GetOwnerByID(id int64) (*Owner, error) {
	var owner Owner
	if id < 0 {
		return nil, errors.New("ID is required")
	}
	err := db.Debug().Model(&owner).Where("id=?", id).Scan(&owner).Error
	return &owner, err
}

// GetOwners returns all owners
func GetOwners(limit int64, page int64) (*[]Owner, error) {
	var owners []Owner
	st := page * limit
	err := db.Debug().Model(&Owner{}).Order("id desc").Limit(int(limit)).Offset(int(st)).Scan(&owners).Error
	return &owners, err
}

// Update updates the owner
func (o *Owner) Update() error {
	return db.Debug().Model(&o).Where("id=?", o.ID).Update(&o).Error
}

// InsertFakeData creates fake data in database
func InsertFakeData() error {
	faker, err := faker.New("en")
	if err != nil {
		return errors.New("Errors creating fake data")
	}

	owner := NewOwner(faker.LastName(), faker.FirstName(), faker.SafeEmail(), faker.PhoneNumber(), "April 7, 1996", "Room", 21, 3)
	owner.Create()
	return nil
}
