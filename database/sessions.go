package database

import (
	"go.jetify.com/typeid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type SessionModel struct {
	ID        string `gorm:"unique;primaryKey"`
	CsrfToken string `gorm:"unique"`
	IpAddress string `gorm:"unique"`
	UserAgent string
	ExpiresAt time.Time

	StaffID string
	Staff   StaffModel `gorm:"foreignKey:StaffID;references:ID"`
}

func (sm *SessionModel) BeforeCreate(tx *gorm.DB) error {
	sm.ID = typeid.Must(typeid.WithPrefix("sess")).String()
	sm.CsrfToken = typeid.Must(typeid.WithPrefix("csrf")).String()
	return nil
}

func (sm *SessionModel) TableName() string {
	return "sessions"
}

type Session struct {
	db *gorm.DB
}

func newSessions(db *gorm.DB) *Session {
	return &Session{db: db}
}

func (s *Session) schema() error {
	return s.db.AutoMigrate(&SessionModel{}) // Direct return for readability
}

func (s *Session) Get(id string) (*SessionModel, error) {
	var session SessionModel

	err := s.db.
		Preload(clause.Associations).
		Where(&SessionModel{ID: id}).
		First(&session).
		Error
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (s *Session) ListUserSessions(id string) ([]SessionModel, error) {
	var sessions []SessionModel

	err := s.db.
		Preload(clause.Associations).
		Where(&SessionModel{StaffID: id}).
		Where("expires_at > ?", time.Now()).
		Find(&sessions).
		Error

	if err != nil {
		return nil, err
	}

	return sessions, nil
}

func (s *Session) Create(session *SessionModel) error {
	err := s.db.Create(session).Error
	if err != nil {
		return err
	}

	err = s.db.Preload(clause.Associations).
		Where(&SessionModel{ID: session.ID}).
		First(session).
		Error

	return err
}

func (s *Session) Delete(session *SessionModel) error {
	return s.db.Delete(session).Error
}

func (s *Session) DeleteByIP(ip string) error {
	return s.db.Where("ip_address = ?", ip).Delete(&SessionModel{}).Error
}
