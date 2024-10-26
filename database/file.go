package database

import (
	"go.jetify.com/typeid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type FileModel struct {
	Id        string `gorm:"unique;primaryKey" json:"id"`
	ProductId uint
	Filename  string
	FilePath  string
	CreatedAt time.Time
	UpdatedAt time.Time

	Product  ProductModel `gorm:"foreignKey:ProductId;references:Id" json:"product"`
	Uploader StaffModel   `gorm:"foreignKey:UploaderId;references:Id" json:"uploader"`
}

func (pm *FileModel) TableName() string {
	return "files"
}

func (pm *FileModel) BeforeCreate(tx *gorm.DB) error {
	pm.Id = typeid.Must(typeid.WithPrefix("file")).String()
	pm.CreatedAt = time.Now()
	return nil
}

func (pm *FileModel) AfterUpdate(tx *gorm.DB) error {
	pm.UpdatedAt = time.Now()
	return nil
}

type File struct {
	db *gorm.DB
}

func newFile(db *gorm.DB) *Product {
	return &Product{db: db}
}

func (f *File) schema() error {
	err := f.db.AutoMigrate(FileModel{})
	if err != nil {
		return err
	}

	return nil
}

func (f *File) GetById(id string) (*FileModel, error) {
	var prd FileModel

	err := f.db.
		Preload(clause.Associations).
		Where(&FileModel{Id: id}).
		First(&prd).
		Error

	if err != nil {
		return nil, err
	}

	return &prd, nil
}

func (f *File) GetAll() ([]FileModel, error) {
	var prd []FileModel

	err := f.db.
		Find(&prd).
		Error

	if err != nil {
		return nil, err
	}

	return prd, nil
}

func (f *File) Create(name, passwordHash string) (*FileModel, error) {
	user := StaffModel{
		Name:         name,
		PasswordHash: passwordHash,
	}

	var err error
	err = f.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return nil, nil
}
