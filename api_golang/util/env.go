package util

import (
	"os"
)

// GetEnv returns value from env, fallback to defaultVal if empty.
func GetEnv(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}

// func exec log info to storage/logs/*.log
func logInfo(message string) {
	now := time.Now().Format("2006-01-02 15:04:05")
	filePath := fmt.Sprintf("storage/logs/%s.log", now)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Printf("INFO: %s %s", now, message)
}

// func exec log error to storage/logs/*.log
func logError(message string) {
	now := time.Now().Format("2006-01-02 15:04:05")
	filePath := fmt.Sprintf("storage/logs/%s.log", now)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Printf("ERROR: %s %s", now, message)
}