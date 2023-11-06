package errcode

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func IsMongoRecordNotFound(err error) bool {
	return errors.Is(err, mongo.ErrNoDocuments)
}

func IsGormRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
