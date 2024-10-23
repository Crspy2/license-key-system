package database

type table interface {
	schema() error
}
