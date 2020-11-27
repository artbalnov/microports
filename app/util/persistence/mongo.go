package persistence

import (
	"fmt"

	"github.com/globalsign/mgo"
)

func GetMongoSession(mongoURL string) (*mgo.Session, error) {
	if mongoURL == "" {
		return nil, fmt.Errorf("mongo URL is empty")
	}

	session, err := mgo.Dial(mongoURL)
	if err != nil {
		return nil, err
	}

	return session, nil
}
