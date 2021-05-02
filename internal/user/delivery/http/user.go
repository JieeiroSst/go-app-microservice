package http

import (
	"fmt"
	userModel "github.com/JIeeiroSst/go-app/internal/models/user"
	"github.com/JIeeiroSst/go-app/internal/user"
	"github.com/JIeeiroSst/go-app/pkg/bigcache"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"time"
)

type UserHttp struct {
	userCase user.UserUserCase
	cache bigcache.Cache
}

func NewUserHTTP(e *echo.Echo,userCase user.UserUserCase)*UserHttp{
	h:=&UserHttp{
		userCase: userCase,
	}

	userGroup:=e.Group("/user")
	userGroup.POST("/login",h.Signup)
	userGroup.POST("/signup",h.Signup)

	return h
}

func (u *UserHttp) Login(e echo.Context) error {
	username,password:=e.FormValue("username"),e.FormValue("password")
	for iter:=u.cache.Iterator();iter.SetNext();{
		info, err := iter.Value()
		if err != nil {
			continue
		}
		if string(info.Value()) == username {
			u.cache.Delete(info.Key())
			log.Printf("forced %s to log out\n", username)
			break
		}
	}
	user:=userModel.Users{
		Username: username,
		Password: password,
	}
	token:=u.userCase.Login(user)
	if token == ""{
		return e.JSON(500,"no such account")
	}
	uId, err := uuid.NewRandom()
	if err != nil {
		log.Info(fmt.Errorf("failed to generate UUID: %w", err))
	}
	sessionId := fmt.Sprintf("%s-%s", uId.String(), username)
	err = u.cache.Set(sessionId,[]byte(username))
	if err != nil {
		log.Info(fmt.Errorf("failed to store current subject in cache: %w", err))
	}
	cookie:=&http.Cookie{
		Name:       "current_subject",
		Value:      "",
		Path:       "/api",
		Domain:     "",
		Expires:    time.Time{},
		RawExpires: "",
		MaxAge:     0,
		Secure:     false,
		HttpOnly:   false,
		SameSite:   0,
		Raw:        "",
		Unparsed:   nil,
	}
	e.SetCookie(cookie)
	return e.JSON(200,map[string]string{
		"message":"logged in successfully",
		"data":token,
	})
}

func (u *UserHttp) Signup(e echo.Context) error {
	username,password:=e.FormValue("username"),e.FormValue("password")
	user:=userModel.Users{
		Username: username,
		Password: password,
	}
	info:=u.userCase.SignUp(user)
	if info == ""{
		return e.JSON(500,"")
	}
	return e.JSON(200,map[string]string{
		"data":info,
		"mesage":"signup in successfully",
	})
}