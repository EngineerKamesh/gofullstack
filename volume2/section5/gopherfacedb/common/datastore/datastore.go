package datastore

import (
	"errors"

	"github.com/EngineerKamesh/gofullstack/volume2/section5/gopherfacedb/models"
)

type Datastore interface {
	CreateUser(user *models.User) error
	GetUser(username string) (*models.User, error)
	Close()
}

const (
	MYSQL = iota
	MONGODB
	REDIS
)

func NewDatastore(datastoreType int, dbConnectionString string) (Datastore, error) {

	switch datastoreType {
	case MYSQL:
		return NewMySQLDatastore(dbConnectionString)
	case MONGODB:
		return NewMongoDBDatastore(dbConnectionString)
	case REDIS:
		return NewRedisDatastore(dbConnectionString)
	default:
		return nil, errors.New("The datastore you specified does not exist!")
	}

}
