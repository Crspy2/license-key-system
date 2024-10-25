package database

import (
	"errors"
	"go.jetify.com/typeid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type Permission int64

const (
	ApproveStaffPermission Permission = 1 << iota // 1 << 0 == 1
	HWIDResetPermission                           // 1 << 1 == 2
	PassResetPermission                           // 1 << 2 == 4
	ViewKeysPermission                            // 1 << 3 == 8
	KeyGenPermission                              // 1 << 4 == 16
	CompensatePermission                          // 1 << 5 == 32
	StatusChangePermission                        // 1 << 6 == 64
	ProductsPermission                            // 1 << 7 == 128
	UserActionsPermission                         // 1 << 8 == 256
	OffsetsPermission                             // 1 << 9 == 512
)

var permissionNames = map[Permission]string{
	ApproveStaffPermission: "Approve Staff",
	HWIDResetPermission:    "Reset HWID",
	PassResetPermission:    "Reset Password",
	ViewKeysPermission:     "View License Keys",
	KeyGenPermission:       "Generate Keys",
	CompensatePermission:   "Compensate Keys",
	StatusChangePermission: "Change Product Status",
	ProductsPermission:     "Manage Products",
	UserActionsPermission:  "Manage User Actions",
	OffsetsPermission:      "Set/Edit Offsets",
}

type StaffModel struct {
	Id           string         `gorm:"unique;primaryKey" json:"id"`
	Name         string         `gorm:"unique" json:"name"`
	PasswordHash string         `json:"password_hash"`
	Perms        int64          `gorm:"type:bigint" json:"perms"`
	Approved     bool           `gorm:"default:false" json:"approved"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	Sessions     []SessionModal `gorm:"unique;foreignKey:StaffId" json:"sessions"`
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

func (sm *StaffModel) HasPermission(permission Permission) bool {
	return sm.Perms&int64(permission) != 0
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

func (sm *StaffModel) UpdatePermissions(permissions ...Permission) {
	var newPerms int64
	for _, permission := range permissions {
		newPerms |= int64(permission)
	}
	sm.Perms = newPerms
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

	staff.UpdatePermissions(HWIDResetPermission, PassResetPermission, KeyGenPermission)

	s.db.Save(&staff)

	return nil
}

func (s *Staff) AddPermission(id string, permission Permission) error {
	var staff StaffModel

	err := s.db.Where(&StaffModel{Id: id}).First(&staff).Error
	if err != nil {
		return err
	}
	staff.Perms |= int64(permission) // Set the permission bit

	s.db.Save(&staff)

	return nil
}

func (s *Staff) RemovePermission(id string, permission Permission) error {
	var staff StaffModel

	err := s.db.Where(&StaffModel{Id: id}).First(&staff).Error
	if err != nil {
		return err
	}
	staff.Perms &= ^int64(permission) // Clear the permission bit

	s.db.Save(&staff)

	return nil
}

func (s *Staff) SetPermissions(id string, permissions ...Permission) error {
	var staff StaffModel

	err := s.db.Where(&StaffModel{Id: id}).First(&staff).Error
	if err != nil {
		return err
	}

	var newPerms int64
	for _, permission := range permissions {
		newPerms |= int64(permission)
	}
	staff.Perms = newPerms

	s.db.Save(&staff)

	return nil
}
