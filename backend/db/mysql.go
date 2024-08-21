package db

import (
	"fmt"
	"reflect"
	"time"

	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type database struct {
	// dbType 数据库类型
	dbType string

	// dbName 数据库库名
	dbName string

	// username 用户名
	username string

	// password 密码
	password string

	// host 数据库链接
	host string

	// port 数据库端口
	port string

	// tablePrefix 表前缀
	tablePrefix string

	// maxConns 最大连接数
	maxConns int

	// idelConns
	idelConns int
}

type option func(*database)

// WithDBType 数据类型
func WithDBType(dbType string) option {
	return func(d *database) { d.dbType = dbType }
}

// WithDBName 数据库库名
func WithDBName(dbName string) option {
	return func(d *database) { d.dbName = dbName }
}

// WithUsername 用户名
func WithUsername(username string) option {
	return func(d *database) { d.username = username }
}

// WithPassword 密码
func WithPassword(password string) option {
	return func(d *database) { d.password = password }
}

// WithPost 数据库链接
func WithHost(host string) option {
	return func(d *database) { d.host = host }
}

// WithPost 端口
func WithPort(port string) option {
	return func(d *database) { d.port = port }
}

// WithTablePrefix 表前缀
func WithTablePrefix(tablePrefix string) option {
	return func(d *database) { d.tablePrefix = tablePrefix }
}

// WithMaxConns 数据库最大连接数
func WithMaxConns(maxConns int) option {
	return func(d *database) { d.maxConns = maxConns }
}

// WithIdelConns 连接池最大连接数
func WithIdelConns(idelConns int) option {
	return func(d *database) { d.idelConns = idelConns }
}

// NewMysql 创建 mysql db 连接
func NewMysql(opts ...option) *gorm.DB {
	m := &database{}

	for _, opt := range opts {
		opt(m)
	}

	// 多种数据库的时候考虑使用策略模式
	switch m.dbType {
	case "mysql":

		//稳定后关闭debug模式
		DefaultLogger := logger.Default
		//DefaultLogger = DefaultLogger.LogMode(logger.Info)

		// 时区改为 UTC 时区
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
			m.username, m.password, m.host, m.port, m.dbName)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: DefaultLogger,
		})
		if err != nil {
			fmt.Printf("mysql connect err :%v\n", err)
			panic(err)
		}

		//注入自定义tracer
		// if err := db.Use(otelgorm.NewPlugin()); err != nil {
		// 	panic(err)
		// }

		// 设置连接池
		sqlDb, _ := db.DB()
		sqlDb.SetMaxOpenConns(m.maxConns)
		sqlDb.SetMaxIdleConns(m.idelConns)
		sqlDb.SetConnMaxLifetime(time.Hour)
		fmt.Println("mysql connect success")
		_ = db.Callback().Create().Before("gorm:create").Register("create_time_stamp", updateTimestampForCreateCallback)
		_ = db.Callback().Update().Before("gorm:update").Register("update_time_stamp", updateTimestampForUpdateCallback)
		return db
	default:
		fmt.Println("database type unsupport")
		return nil
	}
}

func WithDB(db *gorm.DB) {
	DB = db
}

func CloseDB() {
	defer func() {
		db1, _ := DB.DB()
		db1.Close()
	}()
}

func updateTimestampForCreateCallback(db *gorm.DB) {
	if db.Statement.Schema != nil {
		now := time.Now().UTC()
		for _, v := range db.Statement.Schema.Fields {
			switch db.Statement.ReflectValue.Kind() {
			case reflect.Slice, reflect.Array:
				for i := 0; i < db.Statement.ReflectValue.Len(); i++ {
					_, isZero := v.ValueOf(db.Statement.Context, db.Statement.ReflectValue.Index(i))
					if v.DBName == "create_time" && isZero {
						_ = v.Set(db.Statement.Context, db.Statement.ReflectValue.Index(i), now)
					}
					if v.DBName == "update_time" && isZero {
						_ = v.Set(db.Statement.Context, db.Statement.ReflectValue.Index(i), now)
					}
				}
			case reflect.Struct:
				_, isZero := v.ValueOf(db.Statement.Context, db.Statement.ReflectValue)
				if v.DBName == "create_time" && isZero {
					db.Statement.SetColumn("create_time", now, true)
				}
				if v.DBName == "update_time" && isZero {
					db.Statement.SetColumn("update_time", now, true)
				}
			}
		}
	}
}

func updateTimestampForUpdateCallback(db *gorm.DB) {
	if db.Statement.Schema != nil {
		now := time.Now().UTC()
		for _, v := range db.Statement.Schema.Fields {
			switch db.Statement.ReflectValue.Kind() {
			case reflect.Slice, reflect.Array:
				for i := 0; i < db.Statement.ReflectValue.Len(); i++ {
					if v.DBName == "update_time" {
						_ = v.Set(db.Statement.Context, db.Statement.ReflectValue.Index(i), now)
					}
				}
			case reflect.Struct:
				if v.DBName == "update_time" {
					db.Statement.SetColumn("update_time", now, true)
				}
			}
		}
	}
}
