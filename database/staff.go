package database

import (
	"errors"
	"fmt"
	"go.jetify.com/typeid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type Permission int64

const (
	DefaultPermission Permission = 1 << iota
	HWIDResetPermission
	PassResetPermission
	CompensationPermission
	ChangeStatusPermission
	ManageProductsPermission
	KeyGenPermission
	ViewLogsPermission
	ManageUsersPermission
	ManageStaffPermission
)

var permissionNames = map[Permission]string{
	DefaultPermission:        "Default",
	HWIDResetPermission:      "HWIDReset",
	PassResetPermission:      "PasswordReset",
	CompensationPermission:   "Compensate",
	ChangeStatusPermission:   "ProductStatus",
	ManageProductsPermission: "ManageProducts",
	KeyGenPermission:         "GenerateKeys",
	ViewLogsPermission:       "ViewLogs",
	ManageUsersPermission:    "ManageUsers",
	ManageStaffPermission:    "ManageStaff",
}

type Role = int32

const (
	StaffRole Role = iota
	SeniorStaffRole
	LeadStaffRole
	DevRole
	OwnerRole
)

type StaffModel struct {
	ID           string `gorm:"unique;primaryKey"`
	Name         string `gorm:"unique"`
	Role         Role   `gorm:"default:0"`
	PasswordHash string
	Perms        Permission `gorm:"type:bigint"`
	Approved     bool       `gorm:"default:false"`
	CreatedAt    time.Time
	UpdatedAt    time.Time

	Sessions []SessionModel `gorm:"foreignKey:StaffID"`
	Logs     []LogModel     `gorm:"foreignKey:StaffID"`
}

func (sm *StaffModel) TableName() string {
	return "staff"
}

func (sm *StaffModel) BeforeCreate(tx *gorm.DB) error {
	sm.ID = typeid.Must(typeid.WithPrefix("staf")).String()
	sm.CreatedAt = time.Now()
	return nil
}

func (sm *StaffModel) HasPermission(permission Permission) bool {
	return sm.Perms&permission != 0 || sm.Role == OwnerRole
}

func (sm *StaffModel) GetPermissionNames() []string {
	var perms []string
	for perm, name := range permissionNames {
		if sm.HasPermission(perm) {
			perms = append(perms, name)
		}
	}
	return perms
}

func (sm *StaffModel) GetRoleText() string {
	switch sm.Role {
	case 0:
		return "Staff"
	case 1:
		return "Senior Staff"
	case 2:
		return "Lead Staff"
	case 3:
		return "Developer"
	case 4:
		return "Owner"
	}
	return "Staff"
}

func (sm *StaffModel) HasHigherPermissions(otherStaff StaffModel) bool {
	return sm.Perms > otherStaff.Perms
}

func (sm *StaffModel) HasHigherRole(otherStaff StaffModel) bool {
	return sm.Role > otherStaff.Role
}

type Staff struct {
	db *gorm.DB
}

func newStaff(db *gorm.DB) *Staff {
	return &Staff{db: db}
}

func (s *Staff) schema() error {
	return s.db.AutoMigrate(&StaffModel{})
}

func (s *Staff) get(id string) (*StaffModel, error) {
	var staff StaffModel

	err := s.db.
		Where(&StaffModel{ID: id}).
		First(&staff).
		Error

	if err != nil {
		return nil, err
	}

	return &staff, nil
}

func (s *Staff) GetById(id string) (*StaffModel, error) {
	var staff StaffModel

	err := s.db.
		Preload(clause.Associations).
		Where(&StaffModel{ID: id}).
		First(&staff).
		Error

	if err != nil {
		return nil, err
	}

	return &staff, nil
}

func (s *Staff) GetByName(name string) (*StaffModel, error) {
	var staff StaffModel

	err := s.db.
		Where(&StaffModel{Name: name}).
		First(&staff).
		Error

	if err != nil {
		return nil, err
	}

	return &staff, nil
}

func (s *Staff) List() ([]StaffModel, error) {
	var staff []StaffModel

	err := s.db.Find(&staff).
		Order("created_at asc").
		Error
	if err != nil {
		return nil, err
	}

	return staff, nil
}

func (s *Staff) Create(name, passwordHash string) (*StaffModel, error) {
	staff := StaffModel{
		Name:         name,
		PasswordHash: passwordHash,
	}

	err := s.db.Create(&staff).Error
	if err != nil {
		return nil, err
	}

	return &staff, nil
}

func (s *Staff) Authenticate(name, password string) (*StaffModel, error) {
	staff, err := s.GetByName(name)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(staff.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("passwords did not match")
	}

	return staff, nil
}

func (s *Staff) SetAccess(id string, approved bool) (*StaffModel, error) {
	staff, err := s.get(id)
	if err != nil {
		return nil, err
	}

	staff.Approved = approved
	if approved && !staff.HasPermission(DefaultPermission) {
		staff.Perms |= DefaultPermission | HWIDResetPermission | PassResetPermission
	} else if !approved {
		staff.Perms = 0
	}

	if err := s.db.Save(staff).Error; err != nil {
		return nil, err
	}

	return staff, nil
}

func (s *Staff) AddPermission(id string, permission Permission) error {
	staff, err := s.get(id)
	if err != nil {
		return err
	}
	staff.Perms |= permission

	return s.db.Save(staff).Error
}

func (s *Staff) RemovePermission(id string, permission Permission) error {
	staff, err := s.get(id)
	if err != nil {
		return err
	}
	staff.Perms &= ^permission

	return s.db.Save(staff).Error
}

func (s *Staff) SetPermissions(id string, permissions []Permission) (*StaffModel, error) {
	staff, err := s.get(id)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return nil, err
	}

	var newPerms Permission
	for _, permission := range permissions {
		newPerms |= permission
	}
	staff.Perms = newPerms

	if err := s.db.Save(staff).Error; err != nil {
		fmt.Println("Error 2: ", err)
		return nil, err
	}

	return staff, nil
}

func (s *Staff) SetRole(id string, role Role) (*StaffModel, error) {
	staff, err := s.get(id)
	if err != nil {
		return nil, err
	}

	staff.Role = role
	if err := s.db.Save(staff).Error; err != nil {
		return nil, err
	}

	return staff, nil
}
