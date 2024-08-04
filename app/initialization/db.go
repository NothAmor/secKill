package initialization

import (
	"fmt"
	"secKill/app/common"
	"secKill/app/entity"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDb() {
	dbConfig := common.Config.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)

	var err error
	common.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	common.DB.AutoMigrate(
		&entity.User{},
	)

	DB, err := common.DB.DB()
	if err != nil {
		panic(err)
	}

	DB.SetMaxIdleConns(dbConfig.MaxIdleConns)
	DB.SetMaxOpenConns(dbConfig.MaxOpenConns)
	DB.SetConnMaxLifetime(time.Duration(dbConfig.ConnMaxLifetime) * time.Second)
	DB.SetConnMaxIdleTime(time.Duration(dbConfig.MaxIdleTime) * time.Second)
}
