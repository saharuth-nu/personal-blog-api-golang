package config

import (
	"blog-api/core/repositories"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"golang.org/x/term"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {

	width := 80
	width, _, _ = term.GetSize(int(os.Stdout.Fd()))

	fmt.Println(strings.Repeat("-", width))

	sql, _ := fc()
	fmt.Printf("%v\n", sql)
}

func InitDatabase() *gorm.DB {

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Bangkok",
		Env.DBHost, Env.DBUser, Env.DBPasswd, Env.DBName, Env.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: &SqlLogger{},
		DryRun: false,
	})

	if err != nil {
		panic(err)
	}

	if err = db.AutoMigrate(
		repositories.Article{},
	); err != nil {
		panic(err)
	}

	return db
}
