package util

// import (
// 	"database/sql"
// 	"fmt"

// 	"github.com/go-ini/ini"
// )

//连接数据库
// func Db() (*sql.DB, error) {
// 	cfg, err := ini.Load("./config/config.ini")
// 	if err != nil {
// 		fmt.Println("ini.Load(\"./config/config.ini\") error")
// 		return nil, err
// 	}
// 	cfgMysqlSection := cfg.Section("mysql")
// 	DatebaseHost := cfgMysqlSection.Key("Host").String()
// 	DatebasePort := cfgMysqlSection.Key("Port").String()
// 	DatebaseName := cfgMysqlSection.Key("DatebaseName").String()
// 	DatebasePassword := cfgMysqlSection.Key("Password").String()
// 	DatebaseUsername := cfgMysqlSection.Key("Username").String()

// 	db, err := sql.Open("mysql", DatebaseUsername+":"+DatebasePassword+"@tcp("+DatebaseHost+":"+DatebasePort+")/"+DatebaseName)

// 	if err != nil {
// 		fmt.Println("sql open error")
// 		return nil, err
// 	}
// 	return db, nil
// }

import (
	"fmt"

	"github.com/go-ini/ini"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	cfg, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("ini.Load(\"./config/config.ini\") error")
		panic(err)
	}
	cfgMysqlSection := cfg.Section("mysql")
	DatebaseHost := cfgMysqlSection.Key("Host").String()
	DatebasePort := cfgMysqlSection.Key("Port").String()
	DatebaseName := cfgMysqlSection.Key("DatebaseName").String()
	DatebasePassword := cfgMysqlSection.Key("Password").String()
	DatebaseUsername := cfgMysqlSection.Key("Username").String()
	openStr := DatebaseUsername + ":" + DatebasePassword + "@tcp(" + DatebaseHost + ":" + DatebasePort + ")/" + DatebaseName + "?charset=utf8mb4&parseTime=True&loc=Local"

	Db, err = gorm.Open(mysql.Open(openStr), &gorm.Config{})
	if err != nil {
		fmt.Println("gorm open error")
		panic(err)
	}
}
