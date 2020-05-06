/*
 * Mongo configuration of the service.
 * Wang Ruichao (793160615@qq.com)
 */

package srv

import (
	"github.com/wangggong/ezim/config"

	"gopkg.in/mgo.v2"
	"qiniupkg.com/x/log.v7"
)

var session *mgo.Session

func mustDial(mongoURL string) *mgo.Session {
	var err error
	session, err := mgo.Dial(mongoURL)
	if err != nil {
		log.Fatalf("cannot dial to mongodb: %v", err)
	}
	return session
}

func mongo() *mgo.Session         { return session }
func roomCol() *mgo.Collection    { return session.DB("im").C("room") }
func userCol() *mgo.Collection    { return session.DB("im").C("user") }
func roomUsrCol() *mgo.Collection { return session.DB("im").C("room_user") }

func init() {
	session = mustDial(config.Config.MongoURL)
}
