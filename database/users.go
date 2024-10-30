package database

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserModel struct {
	ID           uint64 `gorm:"unique;primaryKey"`
	Name         string `gorm:"unique"`
	PasswordHash sql.NullString
	HWID         sql.NullString
	Banned       bool `gorm:"default:false"`

	Licenses []LicenseModel `gorm:"foreignKey:UserID;references:ID"`
}

func (pm *UserModel) TableName() string {
	return "users"
}

type User struct {
	db *gorm.DB
}

func newUser(db *gorm.DB) *User {
	return &User{db: db}
}

func (s *User) schema() error {
	return s.db.AutoMigrate(&UserModel{})
}

func (s *User) GetById(id uint64) (*UserModel, error) {
	var user UserModel

	err := s.db.
		Preload(clause.Associations).
		Where(&UserModel{ID: id}).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *User) GetAll() ([]UserModel, error) {
	var users []UserModel

	err := s.db.
		Find(&users).
		Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *User) Create(name, passwordHash string) (*UserModel, error) {
	user := UserModel{
		Name:         name,
		PasswordHash: sql.NullString{String: passwordHash, Valid: true},
	}

	err := s.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *User) ResetPassword(id uint64) (*UserModel, error) {
	user, err := s.GetById(id)
	if err != nil {
		return nil, err
	}

	if user.Banned {
		return nil, errors.New("user is banned")
	}

	if !user.PasswordHash.Valid {
		return nil, errors.New("user does not have a password set")
	}

	user.PasswordHash = sql.NullString{Valid: false}
	if err := s.db.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *User) ResetHWID(id uint64) (*UserModel, error) {
	user, err := s.GetById(id)
	if err != nil {
		return nil, err
	}

	if user.Banned {
		return nil, errors.New("user is banned")
	}

	if !user.HWID.Valid {
		return nil, errors.New("user does not have a hardware id set")
	}

	user.HWID = sql.NullString{Valid: false}
	if err := s.db.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *User) Ban(id uint64) (*UserModel, error) {
	user, err := s.GetById(id)
	if err != nil {
		return nil, err
	}

	user.Banned = true
	if err := s.db.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *User) UnBan(id uint64) (*UserModel, error) {
	user, err := s.GetById(id)
	if err != nil {
		return nil, err
	}

	user.Banned = false
	if err := s.db.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Authenticate For use in a REST api or go GUI application. Will authenticate users based on their name and password.
func (s *User) Authenticate(name, password string) (*UserModel, error) {
	var user UserModel

	err := s.db.
		Preload(clause.Associations).
		Where(&StaffModel{Name: name}).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	if user.Banned {
		return nil, errors.New("you are banned from using this service")
	}

	if !user.PasswordHash.Valid {
		passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
		if err != nil {
			return nil, err
		}

		user.PasswordHash = sql.NullString{String: string(passwordBytes), Valid: true}
		s.db.Save(&user)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash.String), []byte(password))
	if err != nil {
		return nil, errors.New("passwords did not match")
	}

	return &user, nil
}
