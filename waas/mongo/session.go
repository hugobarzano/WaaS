package mongo

import (
	"gopkg.in/mgo.v2"
	"os"
)

type Session struct {
	session *mgo.Session
}

var mongodb string;

func NewSession() (*Session,error) {
	mongodb = os.Getenv("MONGO_URI")
	if mongodb == "" {
		mongodb ="mongodb://mongodb,127.0.0.1:27017"
	}
	session, err := mgo.Dial(mongodb)
	if err != nil {
		return nil,err
	}
	return &Session{session}, err
}

func(s *Session) Copy() *Session {
	current:=&Session{s.session.Copy()}
	s.session.Close()
	return current
}

func(s *Session) GetCollection(db string, col string) *mgo.Collection {
	return s.session.DB(db).C(col)
}

func(s *Session) Close() {
	if(s.session != nil) {
		s.session.Close()
	}
}

func(s *Session) DropDatabase(db string) error {
	if(s.session != nil) {
		return s.session.DB(db).DropDatabase()
	}
	return nil
}
