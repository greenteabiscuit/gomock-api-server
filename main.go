package main

import (
	"fmt"
	"github.com/greenteabiscuit/gomock-api-server/infrastructure/router"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	rr := router.Router{}
	// データベースコネクション実行

	db, _ := gorm.Open(mysql.Open("root:root@tcp(localhost:33306)/sample-db?parseTime=true&loc=Asia%2FTokyo"))

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
	r.Run("localhost:8081")
}
