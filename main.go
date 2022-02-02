package main

import (
	"fmt"
	"github.com/greenteabiscuit/gomock-api-server/infrastructure/router"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func main() {
	rr := router.Router{}
	// データベースコネクション実行


	db, _ := gorm.Open(mysql.Open("root:root@tcp(localhsot:33306)/sample-db?parseTime=true&loc=Asia%2FTokyo"))

	defer func() {
		conn, err := db.DB()
		if err != nil {
			fmt.Println("error in close DB: %w", err)
		}
		if err := conn.Close(); err != nil {
			fmt.Println("error in close DB: %w", err)
		}
	}()

	r := rr.InitRouter(db)
	r.Run(os.Getenv("HTTP_HOST") + ":" + os.Getenv("HTTP_PORT")) // listen and serve on 0.0.0.0:8080
}
