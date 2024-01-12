package checker

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/taverok/proxy-checker-example/pkg/db/mysql"
	"github.com/taverok/proxy-checker-example/service/checker/config"
	"github.com/taverok/proxy-checker-example/service/checker/proxy"
)

type App struct {
	Cfg          *config.Config
	DB           *sql.DB
	ProxyRepo    *proxy.Repo
	ProxyService *proxy.Service
}

func NewApp() (*App, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	app := App{Cfg: cfg}

	// db
	db, err := mysql.NewMysql(cfg.DB)
	if err != nil {
		return nil, err
	}
	app.DB = db

	// repo
	app.ProxyRepo = &proxy.Repo{DB: db}

	// service
	app.ProxyService = &proxy.Service{
		Cfg:  cfg,
		Repo: app.ProxyRepo,
	}

	// handler

	return &app, nil
}

func (app *App) Listen() {
	router := &mux.Router{}
	router.HandleFunc("/health", Health)

	slog.Info(fmt.Sprintf("Starting server on port %d", app.Cfg.Server.Port))
	s := http.Server{
		Addr:         fmt.Sprintf(":%d", app.Cfg.Server.Port),
		Handler:      router,
		ReadTimeout:  time.Duration(app.Cfg.Server.Timeout) * time.Second,
		WriteTimeout: time.Duration(app.Cfg.Server.Timeout) * time.Second,
		IdleTimeout:  time.Duration(app.Cfg.Server.Timeout) * time.Second,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func (app *App) Shutdown() {
	err := app.DB.Close()
	if err != nil {
		slog.Error(err.Error())
	}
}

func Health(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "OK")
}
