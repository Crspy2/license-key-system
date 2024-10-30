package database

import (
	"go.jetify.com/typeid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type LogModel struct {
	ID          string    `gorm:"unique;primaryKey"`
	Object      string    `gorm:"<-:create;not null"`
	Title       string    `gorm:"<-:create;not null"`
	Description string    `gorm:"<-:create;not null"`
	OccurredAt  time.Time `gorm:"<-:create;not null"`

	StaffID string
	Staff   StaffModel `gorm:"foreignKey:StaffID;references:ID"`
}

func (pm *LogModel) TableName() string {
	return "logs"
}

func (pm *LogModel) BeforeCreate(tx *gorm.DB) error {
	pm.ID = typeid.Must(typeid.WithPrefix("log")).String()
	pm.OccurredAt = time.Now()
	return nil
}

type Log struct {
	db *gorm.DB
}

func newLog(db *gorm.DB) *Log {
	return &Log{db: db}
}

func (s *Log) schema() error {
	return s.db.AutoMigrate(&LogModel{})
}

func (s *Log) Get(id string) (*LogModel, error) {
	var license LogModel

	err := s.db.
		Preload(clause.Associations).
		Where(&LogModel{ID: id}).
		First(&license).
		Error

	if err != nil {
		return nil, err
	}

	return &license, nil
}

func (s *Log) List() ([]LogModel, error) {
	var logs []LogModel

	err := s.db.Find(&logs).Error
	if err != nil {
		return nil, err
	}

	return logs, nil
}

func (s *Log) LogEvent(staffId, object, title, description string, occurredAt time.Time) (*LogModel, error) {
	log := LogModel{
		Object:      object,
		Title:       title,
		Description: description,
		OccurredAt:  occurredAt,
		StaffID:     staffId,
	}

	err := s.db.Create(&log).Error
	if err != nil {
		return nil, err
	}

	return &log, nil
}
