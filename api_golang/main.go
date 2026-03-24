package main

import (
	"api_golang/database"
	"api_golang/routes"
	"api_golang/utils/env"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Chỉ lấy env từ file .env (không đọc .env.example)
	_ = godotenv.Load(".env")

	db, err := database.ConnectMySQL()
	if err != nil {
		env.LogError(fmt.Sprintf("Database: %v", err))
		log.Fatalf("Database: %v", err)
	}
	defer db.Close()

	// Chạy migration bằng lệnh: go run . migrate
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		if err := database.MigrateUsers(db); err != nil {
			env.LogError(fmt.Sprintf("Migration: %v", err))
			log.Fatalf("Migration: %v", err)
		}
		env.LogInfo("Migration done. Exit.")
		return
	}

	// Khi chạy server, vẫn tự migrate nếu chưa có bảng
	if err := database.MigrateUsers(db); err != nil {
		env.LogError(fmt.Sprintf("Migration: %v", err))
		log.Fatalf("Migration: %v", err)
	}

	port := env.GetEnv("APP_PORT", "8008")
	addr := ":" + port
	r := routes.SetupRouter(db)

	env.LogInfo("Server listening on " + addr)
	env.LogInfo("API base: " + addr + "/api")
	if err := r.Run(addr); err != nil {
		env.LogError(fmt.Sprintf("Server failed: %v", err))
		log.Fatalf("Server failed: %v", err)
	}
}