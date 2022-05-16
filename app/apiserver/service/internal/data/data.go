package data

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/tanmingqian/go-tam/app/apiserver/service/internal/conf"
	mysqlRepo "github.com/tanmingqian/go-tam/app/apiserver/service/internal/data/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, mysqlRepo.NewUserRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db  *gorm.DB
	log *log.Helper
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	return &Data{
		db:  NewDB(c, logger),
		log: log.NewHelper(log.With(logger, "module", "apiserver/data")),
	}, cleanup, nil
}

func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "apiserver/data/gorm"))

	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		conf.Database.Username,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Database,
		true,
		"Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil
	}

	sqlDB.SetMaxOpenConns(int(conf.Database.MaxOpenConnections))

	sqlDB.SetConnMaxLifetime(time.Duration(conf.Database.MaxConnectionLifeTime))

	sqlDB.SetMaxIdleConns(int(conf.Database.MaxIdleConnections))

	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatal(err)
	}
	return db
}

func (d *Data) GetDBIns() *gorm.DB {
	return d.db
}
