package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Response struct {
	CurrentTime string `json:"current_time"`
}

func main() {

	//Logger code for logging
	logFile, err := os.OpenFile("api.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
		return
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	log.Println("Server starting...")

	// Connecting to mysql database assignment
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/assignment")
	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
	}
	defer db.Close()

	// Checking if database connection is working
	if err := db.Ping(); err != nil {
		log.Fatalf("Error connecting database: %v\n", err)
	}

	// Endpoint to get current time of toronto.
	http.HandleFunc("/currentTime", func(w http.ResponseWriter, r *http.Request) {
		loc, err := time.LoadLocation("America/Toronto")
		if err != nil {
			log.Printf("Failed to load timezone: %v\n", err)
			http.Error(w, "Failed to load timezone", http.StatusInternalServerError)
			return
		}
		currentTime := time.Now().In(loc)

		// Inserting time to database.
		_, err = db.Exec("INSERT INTO time_log (timestamp) VALUES (?)", currentTime)
		if err != nil {
			log.Printf("Failed to log time to database: %v\n", err)
			http.Error(w, "Failed to log time", http.StatusInternalServerError)
			return
		}

		log.Printf("Time logged to database: %s\n", currentTime.Format(time.RFC3339))

		response := Response{CurrentTime: currentTime.Format(time.RFC3339)}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		log.Println("Successfully returned current time")
	})

	// Endpoint to get list of logged times.
	http.HandleFunc("/listTimes", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT timestamp FROM time_log")
		if err != nil {
			log.Printf("Failed to fetch data from database: %v\n", err)
			http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Initializing a list variable to store the times.
		var list []string
		for rows.Next() {
			var rawTime []uint8
			if err := rows.Scan(&rawTime); err != nil {
				log.Printf("Error scanning row: %v\n", err)
				http.Error(w, "Error processing data", http.StatusInternalServerError)
				return
			}

			timeString := string(rawTime)
			parsedTime, err := time.Parse("2006-01-02 15:04:05", timeString)
			if err != nil {
				log.Printf("Error parsing time: %v\n", err)
				http.Error(w, "Error parsing time data", http.StatusInternalServerError)
				return
			}

			list = append(list, parsedTime.Format(time.RFC3339))
		}

		if err := rows.Err(); err != nil {
			log.Printf("Error iterating rows: %v\n", err)
			http.Error(w, "Error reading data", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(list)
	})

	serverAddress := ":8080"
	log.Printf("Server running on http://localhost%s\n", serverAddress)
	if err := http.ListenAndServe(serverAddress, nil); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}
