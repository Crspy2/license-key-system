package database

import (
	"go.jetify.com/typeid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type SessionModal struct {
	Id        string    `gorm:"unique;primaryKey" json:"id"`
	CsrfToken string    `gorm:"unique" json:"csrf_token"`
	StaffId   string    `json:"staff_id"`
	IpAddress string    `gorm:"unique" json:"ip_address"`
	UserAgent string    `json:"user_agent"`
	ExpiresAt time.Time `json:"expires_at"`

	Staff StaffModel `gorm:"foreignKey:StaffId;references:Id" json:"staff"`
}

func (sm *SessionModal) BeforeCreate(tx *gorm.DB) error {
	sm.Id = typeid.Must(typeid.WithPrefix("sess")).String()
	sm.CsrfToken = typeid.Must(typeid.WithPrefix("csrf")).String()
	return nil
}

func (sm *SessionModal) TableName() string {
	return "sessions"
}

type Session struct {
	db *gorm.DB
}

func newSessions(db *gorm.DB) *Session {
	return &Session{db: db}
}

func (s *Session) schema() error {
	err := s.db.AutoMigrate(SessionModal{})
	if err != nil {
		return err
	}

	return nil
}

func (s *Session) Get(id string) (*SessionModal, error) {
	var session SessionModal

	err := s.db.
		Preload(clause.Associations).
		Where(&SessionModal{
			Id: id,
		}).
		First(&session).
		Error
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (s *Session) GetUserSessions(id string) ([]SessionModal, error) {
	var sessions []SessionModal

	err := s.db.
		Preload(clause.Associations).
		Where(&SessionModal{StaffId: id}).
		Where("expires_at > ?", time.Now()).
		Find(&sessions).
		Error

	if err != nil {
		return nil, err
	}

	return sessions, nil
}

func (s *Session) Create(session *SessionModal) error {
	var err error
	err = s.db.Create(session).Error
	if err != nil {
		return err
	}

	err = s.db.Preload(clause.Associations).Where(&SessionModal{
		Id: session.Id,
	}).First(session).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *Session) Delete(id string) error {
	sessionMatch := &SessionModal{
		Id: id,
	}

	err := s.db.
		Preload(clause.Associations).
		Where(sessionMatch).
		Delete(sessionMatch).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (s *Session) DeleteByIP(ip string) error {
	sessionMatch := &SessionModal{
		IpAddress: ip,
	}

	err := s.db.
		Preload(clause.Associations).
		Where(sessionMatch).
		Delete(sessionMatch).
		Error

	if err != nil {
		return err
	}

	return nil
}
