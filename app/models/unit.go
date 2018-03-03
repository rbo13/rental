package models

import "time"

// Unit ...
type Unit struct {
	ID         int64     `json:"id" gorm:"AUTO_INCREMENT"`
	OwnerID    int64     `json:"owner_id" gorm:"index"`
	UnitPrice  int64     `json:"unit_price"`
	UnitType   string    `json:"unit_type"`
	UnitNumber int64     `json:"unit_number"`
	UnitStatus bool      `json:"unit_status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// TableName ...
func (Unit) TableName() string {
	return "units"
}

// NewUnit unit constructor
func NewUnit(ownerID, unitPrice, unitNumber int64, unitType string, unitStatus bool) *Unit {
	return &Unit{
		OwnerID:    ownerID,
		UnitPrice:  unitPrice,
		UnitNumber: unitNumber,
		UnitType:   unitType,
		UnitStatus: unitStatus,
	}
}

// Create creates new record of unit
func (u *Unit) Create() error {
	err := db.Debug().Model(&u).Create(&u).Error
	if err != nil {
		return err
	}
	return err
}

// GetUnitByOwnerID gets unit by the tenant id
func GetUnitByOwnerID(ownerID int64) (*Unit, error) {
	var unit Unit
	err := db.Debug().Model(&unit).Where("owner_id=?", ownerID).Scan(&unit).Error
	return &unit, err
}

// GetUnits gets all the unit
func GetUnits() ([]Unit, error) {
	var units []Unit
	var err error
	err = db.Debug().Model(&Unit{}).Order("id desc").Scan(&units).Error
	return units, err
}
