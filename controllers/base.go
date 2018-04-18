package controllers

import (
	"github.com/astaxie/beego/session"
	"github.com/astaxie/beego"
	"net/http"
)

var globalSessions *session.Manager

func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:"gosessionid",
		EnableSetCookie: true,
		Gclifetime:3600,
		Maxlifetime: 3600,
		Secure: false,
		CookieLifeTime: 3600,
		ProviderConfig: "./tmp",
	}
	globalSessions, _ = session.NewManager("memory",sessionConfig)
	go globalSessions.GC()
}

type BaseController struct {
	beego.Controller
	Username string
}

func (c *BaseController)GetUsername(w http.ResponseWriter, r *http.Request) string {
	sess, _ := globalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	return sess.Get("username").(string)
}

func (c *BaseController)SetUsername(w http.ResponseWriter, r *http.Request,username string)  {
	sess, _ := globalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	sess.Set("username",username)
}