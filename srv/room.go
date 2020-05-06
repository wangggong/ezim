/*
 * Room logic of the service.
 * Wang Ruichao (793160615@qq.com)
 */

package srv

import (
	"fmt"
	"sync"
	"time"

	"github.com/btcsuite/btcutil/base58"
	"github.com/pborman/uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"qiniupkg.com/x/log.v7"
)

var (
	roomMtx    sync.RWMutex
	roomUsrMtx sync.RWMutex
)

// GetRoomInfo returns the info of the given room ID.
func GetRoomInfo(RID string) (roomSt, error) {
	roomMtx.RLock()
	defer roomMtx.RUnlock()

	var room roomSt
	err := roomCol().Find(bson.M{"_id": RID}).One(&room)
	if err != nil {
		log.Errorf("got error when getting room info of %v: %v", RID, err)
	}

	return room, err
}

// CreateRoom creates a room.
func CreateRoom(token, passwd string) (roomSt, error) {
	roomMtx.Lock()
	defer roomMtx.Unlock()

	RID := fmt.Sprintf("room_%v", uuid.New())
	room := roomSt{
		RID:    RID,
		Token:  base58.Encode([]byte(token)),
		Passwd: base58.Encode([]byte(passwd)),
		Ct:     time.Now().Unix(),
	}

	err := roomCol().Insert(room)
	if err != nil {
		log.Errorf("got error when creating room: %v", err)
	}

	return room, err
}

// DeleteRoom deletes a room by room ID and token.
func DeleteRoom(RID, token string) error {
	if ok, err := validRoomToken(RID, token); err != nil {
		return err
	} else if !ok {
		return fmt.Errorf("invalid token for %v: %v", RID, token)
	}

	roomMtx.Lock()
	err := roomCol().Remove(bson.M{"_id": RID})
	if err != nil {
		log.Errorf("got error when deleting room %v: %v", RID, err)
	}
	roomMtx.Unlock()

	return err
}

// GetRoomUsers returns the users of the room.
func GetRoomUsers(RID, passwd string) ([]*userSt, error) {
	var users []*userSt

	if ok, err := validRoomPasswd(RID, passwd); err != nil {
		return users, err
	} else if !ok {
		return users, fmt.Errorf("invalid passwd for %v: %v", RID, passwd)
	}

	var records []bson.M
	roomUsrMtx.RLock()
	roomUsrCol().Find(bson.M{"rid": RID}).All(records)
	roomUsrMtx.RUnlock()

	var uids []string
	for _, rec := range records {
		uid, _ := rec["uid"].(string)
		uids = append(uids, uid)
	}

	userMtx.RLock()
	userCol().Find(bson.M{"_id": bson.M{"$in": uids}}).All(users)
	userMtx.RUnlock()

	return users, nil
}

// AddUser adds the given user to the room.
func AddUser(RID, UID, passwd string) (err error) {
	if ok, err := validRoomPasswd(RID, passwd); err != nil {
		return err
	} else if !ok {
		return fmt.Errorf("invalid passwd for %v: %v", RID, passwd)
	}

	var user userSt
	err = userCol().Find(bson.M{"_id": UID}).One(&user)
	if err != nil {
		log.Errorf("got error when find user %v: %v", UID, err)
		return
	}

	var rec bson.M
	roomUsrMtx.Lock()
	defer roomUsrMtx.Unlock()
	err = roomUsrCol().Find(bson.M{"uid": UID, "rid": RID}).One(&rec)
	if err == mgo.ErrNotFound {
		roomUsrCol().Insert(bson.M{"uid": UID, "rid": RID,
			"ct": time.Now().Unix()})
	}

	return
}

// DeleteRoomUser deletes the given user from the room.
func DeleteRoomUser(RID, UID, passwd, upasswd string) (err error) {
	if ok, err := validRoomPasswd(RID, passwd); err != nil {
		return err
	} else if !ok {
		return fmt.Errorf("invalid passwd for %v: %v", RID, passwd)
	}

	if ok, err := validUserPasswd(UID, upasswd); err != nil {
		return err
	} else if !ok {
		return fmt.Errorf("invalid passwd for %v: %v", RID, passwd)
	}

	roomUsrMtx.Lock()
	roomUsrCol().Remove(bson.M{"uid": UID, "rid": RID})
	roomUsrMtx.Unlock()

	return
}
