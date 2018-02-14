package main

import (
	_ "techblog/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)

func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:      "sessionId",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  "",
	}
	beego.GlobalSessions, _ = session.NewManager("memory", sessionConfig)
	go beego.GlobalSessions.GC()
}

func main() {
	beego.Run()
}
