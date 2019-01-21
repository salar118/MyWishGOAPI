package mongo

import (
	"log"

	config "../config"
	mgo "gopkg.in/mgo.v2"
)

//Session contains mongo session
type MongoSession struct {
	session *mgo.Session
}

//NewMongoSession establish a mongo session
func NewMongoSession(config *config.MongoConfig) (*MongoSession, error) {
	s := MongoSession{}
	//var err error
	session, err := mgo.Dial(config.IP)
	if err != nil {
		log.Println("No mongodb session")
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	s.session = session
	return &s, nil
}
