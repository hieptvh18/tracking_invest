package main

import (
	"api_golang/database"
	"api_golang/routes"
	"api_golang/util"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Chỉ lấy env từ file .env (không đọc .env.example)
	_ = godotenv.Load(".env")

	db, err := database.OpenMySQL()
	if err != nil {
		util.LogError(fmt.Sprintf("Database: %v", err))
		log.Fatalf("Database: %v", err)
	}
	defer db.Close()

	// Chạy migration bằng lệnh: go run . migrate
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		if err := database.MigrateUsers(db); err != nil {
			util.LogError(fmt.Sprintf("Migration: %v", err))
			log.Fatalf("Migration: %v", err)
		}
		util.LogInfo("Migration done. Exit.")
		return
	}

	// Khi chạy server, vẫn tự migrate nếu chưa có bảng
	if err := database.MigrateUsers(db); err != nil {
		util.LogError(fmt.Sprintf("Migration: %v", err))
		log.Fatalf("Migration: %v", err)
	}

	port := util.GetEnv("APP_PORT", "8080")
	addr := ":" + port
	r := routes.SetupRouter(db)

	util.LogInfo("Server listening on " + addr)
	util.LogInfo("API base: " + addr + "/api")
	if err := r.Run(addr); err != nil {
		util.LogError(fmt.Sprintf("Server failed: %v", err))
		log.Fatalf("Server failed: %v", err)
	}
}
