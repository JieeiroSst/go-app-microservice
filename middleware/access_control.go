package middleware

import (
	"errors"
	"fmt"
	"github.com/JIeeiroSst/go-app/pkg/jwt"
	cacheErr "github.com/allegro/bigcache"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"

	"github.com/labstack/echo/v4"
	"strings"

	"github.com/JIeeiroSst/go-app/pkg/bigcache"
)

type Authorization struct {
	cache bigcache.Cache
	jwt jwt.TokenUser
}


func (a *Authorization) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		bearToken:=c.Request().Header.Get("Authorization")
		if bearToken == ""{
			return c.JSON(401,"Authentication failure: Token not provided")
		}
		strArr := strings.Split(bearToken, " ")
		message,err:=a.jwt.ParseToken(strArr[1])
		if err!=nil{
		 	return c.JSON(401,message)
		}
		sessionId, _ := c.Cookie("current_subject")
		sub,err:=a.cache.Get(sessionId)
		if errors.Is(err, cacheErr.ErrEntryNotFound) {
			return c.JSON(401,"user hasn't logged in yet")
		}
		c.Set("current_subject", string(sub))
		return next(c)
	}
}

func (a *Authorization) Authorize(obj string, act string, adapter persist.Adapter,next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, _ := c.Cookie("current_subject")
		val, existed := a.cache.Get(cookie)
		if existed!=nil {
			return c.JSON(401,"user hasn't logged in yet")
		}
		ok, err := enforce(string(val), obj, act, adapter)
		if err != nil {
			return c.JSON(500,"error occurred when authorizing user")
		}
		if !ok {
			return c.JSON(500,"forbidden")
		}
		return next(c)
	}
}

func enforce(sub string, obj string, act string, adapter persist.Adapter) (bool, error) {
	enforcer, err := casbin.NewEnforcer("pkg/conf/rbac_model.conf", adapter)
	if err != nil {
		return false, fmt.Errorf("failed to create casbin enforcer: %w", err)
	}
	err = enforcer.LoadPolicy()
	if err != nil {
		return false, fmt.Errorf("failed to load policy from DB: %w", err)
	}
	ok, err := enforcer.Enforce(sub, obj, act)
	return ok, err
}
