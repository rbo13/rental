package models

import (
	"log"
	"time"
)

// Tenant object
type Tenant struct {
	ID          int64     `json:"id" gorm:"AUTO_INCREMENT"`
	AccountID   int64     `json:"account_id" gorm:"index"`
	LastName    string    `json:"last_name"`
	FirstName   string    `json:"first_name"`
	PhoneNumber string    `json:"phone_number"`
	Birthdate   string    `json:"birthdate"`
	Gender      string    `json:"gender"`
	Age         int       `json:"age"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TableName ...
func (Tenant) TableName() string {
	return "tenants"
}

// GetTenantByID ....
func GetTenantByID(id int64) (*Tenant, error) {
	var tenant Tenant
	err := db.Debug().Model(&tenant).Where("id=?", id).Scan(&tenant).Error
	return &tenant, err
}

// GetTenantByAccountID ...
func GetTenantByAccountID(accountID int64) (*Tenant, error) {
	var tenant Tenant
	err := db.Debug().Model(&tenant).Where("account_id=?", accountID).Scan(&tenant).Error
	return &tenant, err
}

// GetTenants ...
func GetTenants(limit int64, page int64) ([]Tenant, error) {
	var tenants []Tenant
	var err error
	st := page * limit
	err = db.Debug().Model(&Tenant{}).Order("id desc").Limit(int(limit)).Offset(int(st)).Scan(&tenants).Error
	return tenants, err
}

// GetTenantsUsingStoredProcedure ...
func GetTenantsUsingStoredProcedure() ([]Tenant, error) {
	var tenants []Tenant
	err := db.Debug().Raw("CALL `rental`.`GetAllTenants`()").Scan(&tenants).Error
	if err != nil {
		log.Printf("ERROR DUE TO: %v", err)
		return nil, err
	}
	return tenants, nil
}

// NewTenant contructor
func NewTenant(accountID int64, lastName, firstName, phoneNumber, gender, birthDate string) *Tenant {
	return &Tenant{
		AccountID:   accountID,
		LastName:    lastName,
		FirstName:   firstName,
		PhoneNumber: phoneNumber,
		Gender:      gender,
		Birthdate:   birthDate,
	}
}

// Create creates new record of tenants
func (t *Tenant) Create() error {
	err := db.Debug().Model(&t).Create(&t).Error
	if err != nil {
		return err
	}
	return err
}

// Update updates the tenant
func (t *Tenant) Update() error {
	return db.Debug().Model(&t).Where("account_id=?", t.AccountID).Update(&t).Error
}
