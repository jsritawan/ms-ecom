package storage

import (
	"context"
	"fmt"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"github.com/jsritawan/ms-ecom/internal/config"
)

type (
	IStorage[K ModelType] interface {
		FindById(ctx context.Context, id int) (data K, err error)
		FindAll(ctx context.Context) (data []K, err error)
	}

	AbstractStorage[K ModelType] struct {
		db        *gorm.DB
		tableName string
	}

	Storage struct {
		db *gorm.DB
	}

	ModelType interface{}
)

const (
	dsnWithoutSSLFormat = "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s"
	dsnWithSSLFormat    = dsnWithoutSSLFormat + " sslcert=%s sslkey=%s sslrootcert=%s"
)

func New(db *config.DBConfig) *Storage {
	var dsn string
	if strings.EqualFold(db.SSLMode, "disable") {
		dsn = fmt.Sprintf(dsnWithoutSSLFormat, db.Host, db.Username, db.Password, db.Name, db.Port, db.SSLMode, db.Timezone)
	} else {
		dsn = fmt.Sprintf(dsnWithSSLFormat, db.Host, db.Username, db.Password, db.Name, db.Port, db.SSLMode, db.Timezone, db.SSLCert, db.SSLKey, db.SSLRootCert)
	}

	var (
		log      gormlogger.Interface
		gormConn *gorm.DB
		err      error
	)

	log = gormlogger.Default.LogMode(gormlogger.Info)

	gormConn, err = openPostgreSQLConnection(dsn, log)
	if err != nil {
		errCount := 0
		for i := 1; i <= 3; i++ {
			log.Info(context.TODO(), fmt.Sprintf("Try connecting to PostgreSQL... [%d]", i))
			time.Sleep(6 * time.Second)
			gormConn, err = openPostgreSQLConnection(dsn, log)
			if err != nil {
				errCount++
				continue
			}
			break
		}
		if errCount == 3 {
			panic(err)
		}
	}

	conn, err := gormConn.DB()
	if err != nil {
		panic(err)
	}
	conn.SetMaxIdleConns(db.MaxIdleConns)
	conn.SetConnMaxIdleTime(db.MaxIdleTime * time.Second)
	conn.SetMaxOpenConns(db.MaxOpenConns)
	conn.SetConnMaxLifetime(db.MaxLifeTime * time.Second)

	return &Storage{
		db: gormConn,
	}
}

func (s *Storage) HeathCheck() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Ping()
}

func (s *AbstractStorage[K]) FindById(ctx context.Context, id int) (data K, err error) {
	err = s.db.WithContext(ctx).Table(s.tableName).First(&data, id).Error
	return data, err
}

func (s *AbstractStorage[K]) FindAll(ctx context.Context) (data []K, err error) {
	err = s.db.WithContext(ctx).Table(s.tableName).Find(&data).Error
	return data, err
}

func openPostgreSQLConnection(dsn string, log gormlogger.Interface) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: log,
	})
}
