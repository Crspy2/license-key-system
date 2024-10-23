package database

import (
	"errors"
	"github.com/crspy2/license-panel/app/internal-http/utils"
	"go.jetify.com/typeid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type Permission string

// Staff have an array of permissions which indicate what they are allowed to do
const (
	ApproveStaffPermission Permission = "approve" // Approve registered Staff so they can login

	HWIDResetPermission Permission = "hwid" // Ability to reset HWIDs
	PassResetPermission Permission = "pass" // Ability to reset account passwords
	ViewKeysPermission  Permission = "keys" // Ability to view user license keys and search for keys

	KeyGenPermission     Permission = "keygen" // Ability to generate Keys
	CompensatePermission Permission = "comp"   // Ability to compensate license keys

	StatusChangePermission Permission = "status"   // Ability to edit products and upload files for products
	ProductsPermission     Permission = "products" // Ability to change product status
	UserActionsPermission  Permission = "Staff"    // Ability to change user account information
	OffsetsPermission      Permission = "offsets"  // Ability to set and edit offsets
)

type StaffModel struct {
	Id           string         `gorm:"unique;primaryKey" json:"id"`
	Name         string         `json:"name"`
	PasswordHash string         `json:"password_hash"`
	Perms        []Permission   `gorm:"type:text" json:"perms"`
	Approved     bool           `gorm:"default:false" json:"approved"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	Sessions     []SessionModal `gorm:"foreignKey:StaffId" json:"sessions"`
}

func (sm *StaffModel) TableName() string {
	return "staff"
}
func (sm *StaffModel) BeforeCreate(tx *gorm.DB) error {
	sm.Id = typeid.Must(typeid.WithPrefix("staff")).String()
	sm.CreatedAt = time.Now()
	return nil
}

func (sm *StaffModel) AfterUpdate(tx *gorm.DB) error {
	sm.UpdatedAt = time.Now()
	return nil
}

type Staff struct {
	db *gorm.DB
}

func newStaff(db *gorm.DB) *Staff {
	return &Staff{db: db}
}

func (s *Staff) schema() error {
	err := s.db.AutoMigrate(StaffModel{})
	if err != nil {
		return err
	}

	return nil
}

func (s *Staff) GetById(id string) (*StaffModel, error) {
	var user StaffModel

	err := s.db.
		Preload(clause.Associations).
		Where(&StaffModel{Id: id}).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Staff) GetByName(name string) (*StaffModel, error) {
	var user StaffModel

	err := s.db.
		Preload(clause.Associations).
		Where(&StaffModel{Name: name}).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Staff) GetAll() ([]StaffModel, error) {
	var staff []StaffModel

	err := s.db.
		Find(&staff).
		Error

	if err != nil {
		return nil, err
	}

	return staff, nil
}

func (s *Staff) Create(name, passwordHash string) (*StaffModel, error) {
	user := StaffModel{
		Name:         name,
		PasswordHash: passwordHash,
	}

	var err error
	err = s.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Staff) ApproveStaff(id string) error {
	var staff StaffModel

	err := s.db.
		Where(&StaffModel{Id: id}).
		First(&staff).
		Error

	if err != nil {
		return err
	}

	staff.Approved = true
	staff.Perms = []Permission{HWIDResetPermission, PassResetPermission, ViewKeysPermission}
	s.db.Save(&staff)

	return nil
}

func (s *Staff) Authenticate(name, password string) (*StaffModel, error) {
	var staff StaffModel

	err := s.db.
		Preload(clause.Associations).
		Where(&StaffModel{Name: name}).
		First(&staff).
		Error

	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(password, staff.PasswordHash) {
		return nil, errors.New("passwords did not match")
	}

	return &staff, nil
}

func (s *Staff) SetPermissions(id string, perms []Permission) (*StaffModel, error) {
	var staff StaffModel

	err := s.db.
		Preload(clause.Associations).
		Where(&StaffModel{Id: id}).
		First(&staff).
		Error

	if err != nil {
		return nil, err
	}

	staff.Perms = perms

	err = s.db.Save(&staff).Error
	if err != nil {
		return nil, err
	}

	return &staff, nil
}
