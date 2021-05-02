package user

import (
	"github.com/JIeeiroSst/go-app/config"
	"github.com/JIeeiroSst/go-app/internal/user/delivery/http"
	"github.com/JIeeiroSst/go-app/internal/user/usercase"
	"github.com/JIeeiroSst/go-app/pkg/jwt"
	"github.com/JIeeiroSst/go-app/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Server struct {
	db *gorm.DB
	cfg *config.Config
	jwt *jwt.TokenUser
	hash *utils.Hash
}

func NewServer(db *gorm.DB, cfg *config.Config, jwt *jwt.TokenUser, hash *utils.Hash) *Server{
	return &Server{
		db:db,
		cfg:cfg,
		jwt:jwt,
		hash:hash,
	}
}

func (s *Server) Run() error {
	userCase :=usercase.NewUserCase(s.db,*s.hash,*s.jwt,*s.cfg)
	e:=echo.New()
	http.NewUserHTTP(e, userCase)
	e.Logger.Fatal(s.cfg.Server.UserPost)
	return nil
}