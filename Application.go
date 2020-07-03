package gosimpleweb

import (
	"database/sql"
	"github.com/gentwolf-shen/gohelper/config"
	"github.com/gentwolf-shen/gohelper/dict"
	"github.com/gentwolf-shen/gohelper/endless"
	"github.com/gentwolf-shen/gohelper/ginhelper"
	"github.com/gentwolf-shen/gohelper/gomybatis"
	"github.com/gentwolf-shen/gohelper/logger"
	"github.com/gentwolf-shen/goweb/interceptor"
	"github.com/gentwolf-shen/goweb/statik"
	"github.com/gin-gonic/gin"
	"path"
	"runtime"
	"strings"
)

type Application struct {
	address       string
	engine        *gin.Engine
	dbConnections map[string]*sql.DB
}

func New() *Application {
	app := &Application{}
	app.init()
	return app
}

func (this *Application) init() {
	cfg, err := config.LoadDefault()
	if err != nil {
		panic("load default config error: " + err.Error())
	}

	this.address = cfg.Web.Port

	dict.EnableEnv = true
	_ = dict.LoadDefault()

	logger.LoadDefault()

	if len(cfg.Db) > 0 {
		this.openDb(cfg.Db)
	}

	if cfg.Web.IsDebug {
		this.engine = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		this.engine = gin.New()
	}

	this.engine.Use(ginhelper.AllowCrossDomainAll())
	this.engine.Use(this.auth())

	runtime.GOMAXPROCS(runtime.NumCPU())
}

func (this *Application) Register(register func(app *Application)) *Application {
	register(this)
	return this
}

func (this *Application) Run(arr []string) *Application {

	if err := endless.ListenAndServe(this.address, this.engine); err != nil {
		logger.Error(err)
	}
	return this
}

func (this *Application) ShutdownHook(hook func()) {
	hook()

	this.closeDb()
}

func (this *Application) openDb(cfg map[string]config.DbConfig) {
	this.dbConnections = make(map[string]*sql.DB, len(cfg))
	var err error

	for name, c := range cfg {
		this.dbConnections[name], err = sql.Open(c.Type, c.Dsn)
		if err != nil {
			logger.Errorf("init database %s %v", name, err)
		} else {
			this.dbConnections[name].SetMaxIdleConns(c.MaxIdleConnections)
			this.dbConnections[name].SetMaxOpenConns(c.MaxOpenConnections)

			files := statik.ReadDir("/mapper/" + name)
			for _, file := range files {
				filename := "/mapper/" + name + "/" + file.Name()
				key := strings.TrimSuffix(path.Base(file.Name()), path.Ext(file.Name()))
				gomybatis.SetMapper(this.dbConnections[name], key, string(statik.Read(filename)))
			}
		}
	}
}

func (this *Application) closeDb() {
	for _, db := range this.dbConnections {
		_ = db.Close()
	}
}

func (this *Application) GetDb(name string) *sql.DB {
	return this.dbConnections[name]
}

func (this *Application) GetWebEngine() *gin.Engine {
	return this.engine
}

func (this *Application) auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		bl := interceptor.Valid(c)
		logger.Debugf("auth status %s -> %v", c.Request.URL.Path, bl)
		if !bl {
			ginhelper.ShowNoAuth(c)
			c.Abort()
		}
	}
}
