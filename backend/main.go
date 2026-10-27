package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GitLab CI/CD Backend Version 1.0")
}

func healthHandler(dbHost string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := net.DialTimeout("tcp", dbHost+":5432", 2*time.Second)
		if err != nil {
			http.Error(w, "backend OK, database DOWN", http.StatusServiceUnavailable)
			return
		}
		conn.Close()
		fmt.Fprint(w, "backend OK, database OK")
	}
}

func main() {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "db"
	}

	http.HandleFunc("/", home)
	http.HandleFunc("/health", healthHandler(dbHost))
	http.ListenAndServe(":8080", nil)
}
