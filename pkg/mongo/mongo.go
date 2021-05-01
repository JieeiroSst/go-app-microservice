package mongo

import (
	"github.com/JIeeiroSst/go-app/config"
	mgo "gopkg.in/mgo.v2"
	"log"
	"time"
)

type Session struct {
	session *mgo.Session
}

func NewSession(c *config.MongoConfig) *Session {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{c.MongoDBHosts},
		Timeout:  60 * time.Second,
	})
	if err != nil {
		log.Fatalf("[ConnectDB]: %s\n", err)
	}
	session.SetMode(mgo.Monotonic, true)
	return &Session{session}
}

func (s *Session) Copy() *mgo.Session {
	return s.session.Copy()
}

func (s *Session) Close() {
	if s.session != nil {
		s.session.Close()
	}
}
