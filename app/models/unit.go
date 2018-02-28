package models

import "time"

// Unit ...
type Unit struct {
	ID         int64     `json:"id" gorm:"AUTO_INCREMENT"`
	TenantID   int64     `json:"tenant_id" gorm:"index"`
	UnitPrice  int64     `json:"unit_price"`
	UnitType   string    `json:"unit_type"`
	UnitStatus bool      `json:"unit_status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// TableName ...
func (Unit) TableName() string {
	return "units"
}

// NewUnit unit constructor
func NewUnit(tenantID, unitPrice int64, unitType string, unitStatus bool) *Unit {
	return &Unit{
		TenantID:   tenantID,
		UnitPrice:  unitPrice,
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

// GetUnitByTenantID gets unit by the tenant id
func GetUnitByTenantID(tenantID int64) (*Unit, error) {
	var unit Unit
	err := db.Debug().Model(&unit).Where("tenant_id=?", tenantID).Scan(&unit).Error
	return &unit, err
}
