package app

import (
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"os"
)

func Connect() error {
	_, err := mgo.Dial(os.Getenv("MONGO_HOST"))
	return errors.Wrap(err, "failed connecting to mongo")
}
