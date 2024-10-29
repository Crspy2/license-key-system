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
	DeleteUserPermission
	KeyGenPermission
	ManageStaffPermission
)

var permissionNames = map[Permission]string{
	DefaultPermission:      "Default",
	HWIDResetPermission:    "HWIDReset",
	PassResetPermission:    "PasswordReset",
	CompensationPermission: "Compensate",
	ChangeStatusPermission: "ProductStatus",

	ManageProductsPermission: "ManageProducts",

	DeleteUserPermission:  "DeleteUsers",
	KeyGenPermission:      "GenerateKeys",
	ManageStaffPermission: "ManageStaff",
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
	Id           string         `gorm:"unique;primaryKey" json:"id"`
	Name         string         `gorm:"unique" json:"name"`
	Image        string         `json:"image"`
	Role         Role           `gorm:"default:0" json:"role"`
	PasswordHash string         `json:"password_hash"`
	Perms        Permission     `gorm:"type:bigint" json:"perms"`
	Approved     bool           `gorm:"default:false" json:"approved"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	Sessions     []SessionModal `gorm:"foreignKey:StaffId" json:"sessions"`
}

func (sm *StaffModel) TableName() string {
	return "staff"
}
func (sm *StaffModel) BeforeCreate(tx *gorm.DB) error {
	sm.Id = typeid.Must(typeid.WithPrefix("staf")).String()
	sm.CreatedAt = time.Now()
	return nil
}

func (sm *StaffModel) AfterUpdate(tx *gorm.DB) error {
	sm.UpdatedAt = time.Now()
	return nil
}

func (sm *StaffModel) HasPermission(permission Permission) bool {
	return sm.Perms&permission != 0
}

func (sm *StaffModel) GetPermissionNames() []string {
	var perms []string
	for perm, name := range permissionNames {
		if sm.HasPermission(perm) { // Use the HasPermission method
			perms = append(perms, name)
		}
	}
	return perms
}

func (sm *StaffModel) UpdatePermissions(permissions ...Permission) Permission {
	var newPerms Permission
	for _, permission := range permissions {
		newPerms |= permission
	}
	return newPerms
}

func (sm *StaffModel) HasHigherPermissions(otherStaff StaffModel) bool {
	if sm.Role > otherStaff.Role {
		return true
	} else if sm.Role == otherStaff.Role {
		if sm.Perms > otherStaff.Perms {
			return true
		}
	}
	return false
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

	err = bcrypt.CompareHashAndPassword([]byte(staff.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("passwords did not match")
	}

	return &staff, nil
}

func (s *Staff) SetStaffAccess(id string, approved bool) (*StaffModel, error) {
	var staff StaffModel

	err := s.db.
		Where(&StaffModel{Id: id}).
		First(&staff).
		Error

	fmt.Println(staff)

	if err != nil {
		return nil, err
	}

	fmt.Println(approved)
	staff.Approved = approved
	fmt.Println(staff)

	if approved {
		staff.Perms = staff.UpdatePermissions(DefaultPermission, HWIDResetPermission, PassResetPermission)
	} else {
		staff.Perms = staff.UpdatePermissions(0)
	}

	s.db.Where(&StaffModel{Id: id}).Save(&staff)
	//s.db.Save(&staff)

	return &staff, nil
}

func (s *Staff) AddPermission(id string, permission Permission) error {
	var staff StaffModel

	err := s.db.Where(&StaffModel{Id: id}).First(&staff).Error
	if err != nil {
		return err
	}
	staff.Perms |= permission // Set the permission bit

	s.db.Save(&staff)

	return nil
}

func (s *Staff) RemovePermission(id string, permission Permission) error {
	var staff StaffModel

	err := s.db.Where(&StaffModel{Id: id}).First(&staff).Error
	if err != nil {
		return err
	}
	staff.Perms &= ^permission // Clear the permission bit

	s.db.Save(&staff)

	return nil
}

func (s *Staff) SetPermissions(id string, permissions []Permission) (*StaffModel, error) {
	var staff StaffModel

	err := s.db.Where(&StaffModel{Id: id}).First(&staff).Error
	if err != nil {
		return nil, err
	}

	var newPerms Permission
	for _, permission := range permissions {
		newPerms |= permission
	}
	staff.Perms = newPerms

	s.db.Save(&staff)

	return &staff, nil
}

func (s *Staff) SetRole(id string, role Role) (*StaffModel, error) {
	var staff StaffModel

	err := s.db.Where(&StaffModel{Id: id}).First(&staff).Error
	if err != nil {
		return nil, err
	}

	staff.Role = role
	s.db.Save(&staff)

	return &staff, nil
}
