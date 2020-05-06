/*
 * HTTP handlers for ezim.
 * Wang Ruichao (793160615@qq.com)
 */

package ctrler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/wangggong/ezim/srv"

	RESTful "github.com/emicklei/go-restful"
)

// GetRoomInfo is the handler of getting room info.
func GetRoomInfo(r *RESTful.Request, w *RESTful.Response) {
	RID := r.PathParameter("rid")
	info, err := srv.GetRoomInfo(RID)
	if err != nil {
		w.WriteErrorString(http.StatusBadRequest, err.Error())
	}
	w.WriteAsJson(basicResp{Ret: 1, Data: info})
}

// CreateRoom is the handler of creating room.
func CreateRoom(r *RESTful.Request, w *RESTful.Response) {
	token, _ := r.BodyParameter("token")
	passwd, _ := r.BodyParameter("passwd")
	info, err := srv.CreateRoom(token, passwd)
	if err != nil {
		w.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteAsJson(basicResp{Ret: 1, Data: info})
}

// DeleteRoom is the handler of deleting room.
func DeleteRoom(r *RESTful.Request, w *RESTful.Response) {
	RID, _ := r.BodyParameter("rid")
	token, _ := r.BodyParameter("token")
	err := srv.DeleteRoom(RID, token)
	if err != nil {
		w.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteAsJson(basicResp{Ret: 1})
}

// GetUserInfo is the handler of getting user info.
func GetUserInfo(r *RESTful.Request, w *RESTful.Response) {
	UID := r.PathParameter("uid")
	info, err := srv.GetUserInfo(UID)
	if err != nil {
		w.WriteErrorString(http.StatusBadRequest, err.Error())
	}
	w.WriteAsJson(basicResp{Ret: 1, Data: info})
}

// CreateUser is the handler of creating user.
func CreateUser(r *RESTful.Request, w *RESTful.Response) {
	username, _ := r.BodyParameter("username")
	passwd, _ := r.BodyParameter("passwd")
	info, err := srv.CreateUser(username, passwd)
	if err != nil {
		w.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteAsJson(basicResp{Ret: 1, Data: info})
}

// DeleteUser is the handler of deleting user.
func DeleteUser(r *RESTful.Request, w *RESTful.Response) {
	UID, _ := r.BodyParameter("uid")
	passwd, _ := r.BodyParameter("passwd")
	err := srv.DeleteUser(UID, passwd)
	if err != nil {
		w.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteAsJson(basicResp{Ret: 1})
}

// GetRoomUsers is the handler of getting users in room.
func GetRoomUsers(r *RESTful.Request, w *RESTful.Response) {
	RID := r.PathParameter("rid")
	passwd, _ := r.BodyParameter("passwd")
	users, err := srv.GetRoomUsers(RID, passwd)
	if err != nil {
		w.WriteErrorString(http.StatusBadRequest, err.Error())
	}
	w.WriteAsJson(basicResp{Ret: 1, Data: users})
}

// AddUser is the handler of add user for given room.
func AddUser(r *RESTful.Request, w *RESTful.Response) {
	RID, _ := r.BodyParameter("rid")
	UID, _ := r.BodyParameter("uid")
	passwd, _ := r.BodyParameter("passwd")
	err := srv.AddUser(RID, UID, passwd)
	if err != nil {
		w.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteAsJson(basicResp{Ret: 1})
}

// DeleteRoomUser is the handler of deleting user in room.
func DeleteRoomUser(r *RESTful.Request, w *RESTful.Response) {
	RID := r.PathParameter("rid")
	UID := r.PathParameter("uid")
	passwd, _ := r.BodyParameter("passwd")
	upasswd, _ := r.BodyParameter("user_passwd")
	err := srv.DeleteRoomUser(RID, UID, passwd, upasswd)
	if err != nil {
		w.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteAsJson(basicResp{Ret: 1})
}

// Online is the handler of login.
func Online(r *RESTful.Request, w *RESTful.Response) {
	UID := r.PathParameter("uid")
	passwd, _ := r.BodyParameter("user_passwd")
	err := srv.Online(UID, passwd)
	if err != nil {
		w.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteAsJson(basicResp{Ret: 1})
}

// Offline is the handler of logout.
func Offline(r *RESTful.Request, w *RESTful.Response) {
	UID := r.PathParameter("uid")
	passwd, _ := r.BodyParameter("user_passwd")
	err := srv.Offline(UID, passwd)
	if err != nil {
		w.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteAsJson(basicResp{Ret: 1})
}

// GetMsg is the handler for getting message of the given room.
func GetMsg(r *RESTful.Request, w *RESTful.Response) {
	RID := r.PathParameter("rid")
	UID := r.PathParameter("uid")
	passwd, _ := r.BodyParameter("passwd")
	upasswd, _ := r.BodyParameter("user_passwd")
	sct, _ := r.BodyParameter("ct")
	ct, _ := strconv.Atoi(sct)
	err := srv.GetMsg(RID, UID, passwd, upasswd, int64(ct))
	if err != nil {
		w.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteAsJson(basicResp{Ret: 1})
}

// SendMsg is the handler for sending message of the given room.
func SendMsg(r *RESTful.Request, w *RESTful.Response) {
	RID := r.PathParameter("rid")
	UID := r.PathParameter("uid")
	passwd, _ := r.BodyParameter("passwd")
	upasswd, _ := r.BodyParameter("user_passwd")
	err := srv.SendMsg(RID, UID, passwd, upasswd)
	if err != nil {
		w.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteAsJson(basicResp{Ret: 1})
}

// HealthCheck is handler for health check.
func HealthCheck(r *RESTful.Request, w *RESTful.Response) {
	w.WriteAsJson(basicResp{
		Ret:  1,
		Data: map[string]int64{"timestamp": time.Now().Unix()},
	})
}
