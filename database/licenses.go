package database

import (
	"database/sql"
	"go.jetify.com/typeid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type LicenseModel struct {
	ID               string        `gorm:"unique;primaryKey"`
	Duration         time.Duration `gorm:"not null"`
	TimesCompensated uint          `gorm:"default:0"`
	HoursCompensated uint          `gorm:"default:0"`
	Activation       sql.NullTime
	Expiration       sql.NullTime
	CreatedAt        time.Time
	UpdatedAt        time.Time

	UserID    *uint64
	User      *UserModel `gorm:"foreignKey:UserID;references:ID"`
	ProductID string
	Product   ProductModel `gorm:"foreignKey:ProductID;references:ID"`
}

func (pm *LicenseModel) TableName() string {
	return "licenses"
}

func (pm *LicenseModel) BeforeCreate(tx *gorm.DB) error {
	pm.ID = typeid.Must(typeid.WithPrefix("key")).String()
	pm.CreatedAt = time.Now()
	return nil
}

type License struct {
	db *gorm.DB
}

func newLicense(db *gorm.DB) *License {
	return &License{db: db}
}

func (s *License) schema() error {
	return s.db.AutoMigrate(&LicenseModel{})
}

func (s *License) Get(key string) (*LicenseModel, error) {
	var license LicenseModel

	err := s.db.
		Preload(clause.Associations).
		Where(&LicenseModel{ID: key}).
		First(&license).
		Error

	if err != nil {
		return nil, err
	}

	return &license, nil
}

func (s *License) Create(productId string, duration time.Duration) (*LicenseModel, error) {
	license := LicenseModel{
		ProductID: productId,
		Duration:  duration,
	}

	err := s.db.Create(&license).Error
	if err != nil {
		return nil, err
	}

	return &license, nil
}

func (s *License) Redeem(key string, userId uint64) (*LicenseModel, error) {
	license, err := s.Get(key)
	if err != nil {
		return nil, err
	}

	license.UserID = &userId
	license.Activation = sql.NullTime{Time: time.Now(), Valid: true}
	license.Expiration = sql.NullTime{Time: time.Now().Add(license.Duration), Valid: true}

	if err := s.db.Save(license).Error; err != nil {
		return nil, err
	}

	return license, nil
}
