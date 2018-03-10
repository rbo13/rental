package models

import (
	"time"
)

// TenantRecord our simple model for tenant's record
type TenantRecord struct {
	ID                  int64     `json:"id" gorm:"AUTO_INCREMENT"`
	TenantID            int64     `json:"tenant_id" gorm:"index"`
	UnitID              int64     `json:"unit_id" gorm:"index"`
	OwnerID             int64     `json:"owner_id" gorm:"index"`
	FirstName           string    `json:"first_name"`
	LastName            string    `json:"last_name"`
	UnitType            string    `json:"unit_type"`
	StartDate           string    `json:"start_date"`
	EndDate             string    `json:"end_date"`
	TenantStatus        bool      `json:"tenant_status"`
	PaymentStatus       bool      `json:"payment_status"`
	TotalAmountPaid     int64     `json:"total_amount_paid"`
	TotalPendingPayment int64     `json:"total_pending_payment"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

// TableName table name
func (TenantRecord) TableName() string {
	return "tenant_records"
}

// NewTenantRecord constructor
func NewTenantRecord() *TenantRecord {
	return &TenantRecord{}
}

// Create creates new record of tenant
func (tr *TenantRecord) Create() error {
	err := db.Debug().Model(&tr).Create(&tr).Error
	return err
}

// Update updates the tenant record
func (tr *TenantRecord) Update() error {
	return db.Debug().Model(&tr).Where("id=?", tr.ID).Update(&tr).Error
}

// GetTenantRecordByID gets by id
func GetTenantRecordByID(id int64) (*TenantRecord, error) {
	var tr TenantRecord
	err := db.Debug().Model(&tr).Where("id=?", id).Scan(&tr).Error
	return &tr, err
}

// GetTenantRecordByOwnerID gets by owner id
func GetTenantRecordByOwnerID(ownerID int64) (*TenantRecord, error) {
	var tr TenantRecord
	err := db.Debug().Model(&tr).Where("owner_id=?", ownerID).Scan(&tr).Error
	return &tr, err
}

// GetRecordsByOwnerID returns an arry of records
func GetRecordsByOwnerID(ownerID int64) ([]TenantRecord, error) {
	var tr []TenantRecord
	err := db.Debug().Model(&TenantRecord{}).Where("owner_id=? and tenant_status=0", ownerID).Order("id desc").Scan(&tr).Error
	return tr, err
}
