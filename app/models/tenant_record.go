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

// GetTenantRecordByID gets by id
func GetTenantRecordByID(id int64) (*TenantRecord, error) {
	var tr TenantRecord
	err := db.Debug().Model(&tr).Where("id=?", id).Scan(&tr).Error
	return &tr, err
}
