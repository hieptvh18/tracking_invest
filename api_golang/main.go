package main

import (
	"api_golang/database"
	"api_golang/routes"
	"api_golang/util"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Chỉ lấy env từ file .env (không đọc .env.example)
	_ = godotenv.Load(".env")

	db, err := database.OpenMySQL()
	if err != nil {
		log.Fatalf("Database: %v", err)
	}
	defer db.Close()

	// Chạy migration bằng lệnh: go run . migrate
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		if err := database.MigrateUsers(db); err != nil {
			log.Fatalf("Migration: %v", err)
		}
		log.Println("Migration done. Exit.")
		return
	}

	// Khi chạy server, vẫn tự migrate nếu chưa có bảng
	if err := database.MigrateUsers(db); err != nil {
		log.Fatalf("Migration: %v", err)
	}

	port := util.GetEnv("APP_PORT", "8080")
	addr := ":" + port
	r := routes.SetupRouter(db)

	log.Printf("Server listening on http://localhost%v", addr)
	log.Printf("API base: http://localhost%v/api", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
