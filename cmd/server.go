/*
 * IM server for ezim.
 * Wang Ruichao (793160615@qq.com)
 */

package main

import (
	"flag"
	"net/http"

	"github.com/wangggong/ezim/config"
	"github.com/wangggong/ezim/ctrler"

	RESTful "github.com/emicklei/go-restful"
)

func init() {
}

func route(ws *RESTful.WebService) {
	ws.Route(ws.GET("/room/{rid}").To(ctrler.GetRoomInfo))
	ws.Route(ws.POST("/room").To(ctrler.CreateRoom))
	ws.Route(ws.DELETE("/room").To(ctrler.DeleteRoom))

	ws.Route(ws.GET("/user/{uid}").To(ctrler.GetUserInfo))
	ws.Route(ws.POST("/user").To(ctrler.CreateUser))
	ws.Route(ws.DELETE("/user").To(ctrler.DeleteUser))

	ws.Route(ws.GET("/room/{rid}/user").To(ctrler.GetRoomUsers))
	ws.Route(ws.POST("/room/{rid}/user/{uid}").To(ctrler.AddUser))
	ws.Route(ws.DELETE("/room/{rid}/user/{uid}").To(ctrler.DeleteRoomUser))

	ws.Route(ws.POST("/online/user/{uid}").To(ctrler.Online))
	ws.Route(ws.POST("/offline/user/{uid}").To(ctrler.Offline))

	ws.Route(ws.GET("/msg/{rid}").To(ctrler.GetMsg))
	ws.Route(ws.POST("/msg/{rid}").To(ctrler.SendMsg))

	ws.Route(ws.GET("/healthcheck").To(ctrler.HealthCheck))
}

func main() {
	var filename string
	flag.StringVar(&filename, "f", "", "Configuration file.")
	flag.Parse()
	config.LoadConfig(filename)

	ws := new(RESTful.WebService)
	route(ws)
	RESTful.Add(ws)

	http.ListenAndServe(config.Config.HTTPPort, nil)
}
