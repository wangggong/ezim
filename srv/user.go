/*
 * User logic of the service.
 * Wang Ruichao (793160615@qq.com)
 */

package srv

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"sync"
	"time"

	"github.com/btcsuite/btcutil/base58"
	"github.com/pborman/uuid"
	"qiniupkg.com/x/log.v7"
)

const (
	statusOnline  = "online"
	statusOffline = "offline"
)

var (
	userMtx sync.RWMutex
)

// GetUserInfo returns the info of the given user ID.
func GetUserInfo(UID string) (userSt, error) {
	userMtx.RLock()
	defer userMtx.RUnlock()

	var user userSt
	err := userCol().Find(bson.M{"_id": UID}).One(&user)
	if err != nil {
		log.Errorf("got error when getting user info of %v: %v", UID, err)
	}

	return user, err
}

// CreateUser creates a user.
func CreateUser(username, passwd string) (userSt, error) {
	userMtx.Lock()
	defer userMtx.Unlock()

	UID := fmt.Sprintf("user_%v", uuid.New())
	user := userSt{
		UID:     UID,
		Usrname: username,
		Passwd:  base58.Encode([]byte(passwd)),
		Ct:      time.Now().Unix(),
	}

	err := userCol().Insert(user)
	if err != nil {
		log.Errorf("got error when creating user: %v", err)
	}

	return user, err
}

// DeleteUser deletes a user by user ID and password.
func DeleteUser(UID, passwd string) error {
	userMtx.Lock()
	defer userMtx.Unlock()

	if ok, err := validUserPasswd(UID, passwd); err != nil {
		return err
	} else if !ok {
		return fmt.Errorf("invalid passwd for %v: %v", UID, passwd)
	}

	err := userCol().Remove(bson.M{"_id": UID})
	if err != nil {
		log.Errorf("got error when deleting user %v: %v", UID, err)
	}

	return err
}

// Online is the service logic for login.
func Online(UID, passwd string) error {
	if ok, err := validUserPasswd(UID, passwd); err != nil {
		return err
	} else if !ok {
		return fmt.Errorf("invalid passwd for %v: %v", UID, passwd)
	}
	userCol().Upsert(bson.M{"_id": UID}, bson.M{"$set": bson.M{
		"status": statusOnline,
		"lit":    time.Now().Unix(),
	}})

	return nil
}

// Offline is the service logic for login.
func Offline(UID, passwd string) error {
	if ok, err := validUserPasswd(UID, passwd); err != nil {
		return err
	} else if !ok {
		return fmt.Errorf("invalid passwd for %v: %v", UID, passwd)
	}
	userCol().Upsert(bson.M{"_id": UID}, bson.M{"$set": bson.M{
		"status": statusOffline,
		"lot":    time.Now().Unix(),
	}})

	return nil
}

// GetMsg shows the message in the given room.
func GetMsg(RID, UID, passwd, upasswd string, ct int64) error {
	// TODO
	return nil
}

// SendMsg sends the message to the given room for the user.
func SendMsg(RID, UID, passwd, upasswd string) error {
	// TODO
	return nil
}
