package env

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

// GetEnv returns value from env, fallback to defaultVal if empty.
func GetEnv(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}

// LogInfo writes info message to storage/logs (daily): storage/logs/2006-01-02/info.log
func LogInfo(message string) {
	now := time.Now()
	dateDir := now.Format("2006-01-02")
	dir := filepath.Join("storage", "logs", dateDir)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Printf("LogInfo mkdir: %v", err)
		return
	}
	filePath := filepath.Join(dir, "info.log")
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("LogInfo open: %v", err)
		return
	}
	defer file.Close()
	log.SetOutput(file)
	log.Printf("INFO: %s %s", now.Format("2006-01-02 15:04:05"), message)
}

// LogError writes error message to storage/logs (daily): storage/logs/2006-01-02/error.log
func LogError(message string) {
	now := time.Now()
	dateDir := now.Format("2006-01-02")
	dir := filepath.Join("storage", "logs", dateDir)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Printf("LogError mkdir: %v", err)
		return
	}
	filePath := filepath.Join(dir, "error.log")
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("LogError open: %v", err)
		return
	}
	defer file.Close()
	log.SetOutput(file)
	log.Printf("ERROR: %s %s", now.Format("2006-01-02 15:04:05"), message)
}