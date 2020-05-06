/*
 * Utility functions of the service.
 * Wang Ruichao (793160615@qq.com)
 */

package srv

import (
	"github.com/btcsuite/btcutil/base58"
	"gopkg.in/mgo.v2/bson"
	"qiniupkg.com/x/log.v7"
)

func validRoomToken(RID, token string) (bool, error) {
	roomMtx.RLock()
	var room roomSt
	err := roomCol().Find(bson.M{"_id": RID}).One(&room)
	if err != nil {
		log.Errorf("cannot find room %v: %v", RID, err)
		return false, err
	}
	roomMtx.RUnlock()

	if base58.Encode([]byte(token)) != room.Token {
		log.Errorf("invalid token: RID %v, token %v", RID, token)
		return false, nil
	}
	return true, nil
}

func validRoomPasswd(RID, passwd string) (bool, error) {
	roomMtx.RLock()
	var room roomSt
	err := roomCol().Find(bson.M{"_id": RID}).One(&room)
	if err != nil {
		log.Errorf("cannot find room %v: %v", RID, err)
		return false, err
	}
	roomMtx.RUnlock()

	if base58.Encode([]byte(passwd)) != room.Passwd {
		log.Errorf("invalid passwd: RID %v, passwd %v", RID, passwd)
		return false, nil
	}
	return true, nil
}

func validUserPasswd(UID, passwd string) (bool, error) {
	userMtx.RLock()
	var user userSt
	err := userCol().Find(bson.M{"_id": UID}).One(&user)
	if err != nil {
		log.Errorf("cannot find user %v: %v", UID, err)
		return false, err
	}
	userMtx.RUnlock()

	if base58.Encode([]byte(passwd)) != user.Passwd {
		log.Errorf("invalid passwd: UID %v, passwd %v", UID, passwd)
		return false, nil
	}
	return true, nil
}
