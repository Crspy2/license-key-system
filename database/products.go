package database

import (
	"go.jetify.com/typeid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type ProductModel struct {
	Id              string    `gorm:"unique;primaryKey" json:"id"`
	Name            string    `gorm:"unique" json:"name"`
	Status          string    `gorm:"not null" json:"status"`
	StatusChangedAt time.Time `json:"status_changed_at"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`

	Files []FileModel `gorm:"foreignKey:ProductId" json:"files"`
	// TODO: Add license keys as foreign key array for product (for compensation)
}

func (pm *ProductModel) TableName() string {
	return "products"
}

func (pm *ProductModel) BeforeCreate(tx *gorm.DB) error {
	pm.Id = typeid.Must(typeid.WithPrefix("prod")).String()
	pm.CreatedAt = time.Now()
	return nil
}

func (pm *ProductModel) AfterUpdate(tx *gorm.DB) error {
	pm.UpdatedAt = time.Now()
	return nil
}

type Product struct {
	db *gorm.DB
}

func newProduct(db *gorm.DB) *Product {
	return &Product{db: db}
}

func (s *Product) schema() error {
	err := s.db.AutoMigrate(ProductModel{})
	if err != nil {
		return err
	}

	return nil
}

func (s *Product) GetById(id string) (*ProductModel, error) {
	var prd ProductModel

	err := s.db.
		Preload(clause.Associations).
		Where(&ProductModel{Id: id}).
		First(&prd).
		Error

	if err != nil {
		return nil, err
	}

	return &prd, nil
}

func (s *Product) GetAll() ([]ProductModel, error) {
	var prd []ProductModel

	err := s.db.
		Find(&prd).
		Error

	if err != nil {
		return nil, err
	}

	return prd, nil
}

func (s *Product) Create(name, passwordHash string) (*ProductModel, error) {
	user := StaffModel{
		Name:         name,
		PasswordHash: passwordHash,
	}

	var err error
	err = s.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return nil, nil
}
