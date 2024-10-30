package database

import (
	"database/sql"
	"errors"
	"go.jetify.com/typeid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type Status = string

var (
	Development Status = "Development"
	Testing     Status = "Testing"
	Maintenance Status = "Under Maintenance"
	Operational Status = "Operational"
)

type ProductModel struct {
	ID        string `gorm:"unique;primaryKey"`
	Name      string `gorm:"unique"`
	Status    Status `gorm:"default:Development"`
	PausedAt  sql.NullTime
	CreatedAt time.Time
	UpdatedAt time.Time

	Licenses []LicenseModel `gorm:"foreignKey:ProductID;references:ID"`
}

func (pm *ProductModel) TableName() string {
	return "products"
}

func (pm *ProductModel) BeforeCreate(tx *gorm.DB) error {
	pm.ID = typeid.Must(typeid.WithPrefix("prod")).String()
	pm.CreatedAt = time.Now()
	return nil
}

type Product struct {
	db *gorm.DB
}

func newProduct(db *gorm.DB) *Product {
	return &Product{db: db}
}

func (s *Product) schema() error {
	return s.db.AutoMigrate(&ProductModel{})
}

func (s *Product) Get(id string) (*ProductModel, error) {
	var prd ProductModel
	err := s.db.
		Preload(clause.Associations).
		Where(&ProductModel{ID: id}).
		First(&prd).
		Error

	if err != nil {
		return nil, err
	}

	return &prd, nil
}

func (s *Product) List() ([]ProductModel, error) {
	var prd []ProductModel
	err := s.db.Find(&prd).Error
	if err != nil {
		return nil, err
	}

	return prd, nil
}

func (s *Product) Create(name string) (*ProductModel, error) {
	product := ProductModel{Name: name}
	err := s.db.Create(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (s *Product) Delete(id string) (*ProductModel, error) {
	product, err := s.Get(id)
	if err != nil {
		return nil, err
	}

	err = s.db.Delete(&ProductModel{}, id).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Product) SetStatus(id string, status Status) (*ProductModel, error) {
	product, err := s.Get(id)
	if err != nil {
		return nil, err
	}

	pastStatus := product.Status
	product.Status = status

	if (status == Testing || status == Maintenance) && !product.PausedAt.Valid {
		product.PausedAt = sql.NullTime{Time: time.Now(), Valid: true}
	}

	if err := s.db.Save(product).Error; err != nil {
		return nil, err
	}

	if status == Operational && (pastStatus == Maintenance || pastStatus == Testing) && product.PausedAt.Valid {
		_, err = s.CompensateKeys(id, time.Since(product.PausedAt.Time))
		if err != nil {
			return nil, err
		}
		product.PausedAt = sql.NullTime{Valid: false}

		if err := s.db.Save(product).Error; err != nil {
			return nil, errors.New("product status updated, but failed to compensate for lost time")
		}
	}

	return product, nil
}

func (s *Product) CompensateKeys(id string, duration time.Duration) (*ProductModel, error) {
	tx := s.db.Begin()

	product, err := s.Get(id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, license := range product.Licenses {
		license.Expiration = sql.NullTime{Time: license.Expiration.Time.Add(duration), Valid: true}
		if err := tx.Save(&license).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return product, nil
}
