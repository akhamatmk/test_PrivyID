package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// return all env
	return os.Getenv(key)
}

func main() {

	var enviroment string

	// read flag in env uses , and set default env if flag not fiiled
	flag.StringVar(&enviroment, "env", "DEVELOPMENT", "Setting Env")
	flag.Parse()

	enviroment = strings.ToUpper(enviroment)

	// set env to development if env not PRODUCTION
	if enviroment != "PRODUCTION" {
		enviroment = "DEVELOPMENT"
	}

	// godotenv package
	port := goDotEnvVariable(fmt.Sprintf("%s_PORT", enviroment))
	appName := goDotEnvVariable(fmt.Sprintf("%s_APP_NAME", enviroment))
	debug := goDotEnvVariable(fmt.Sprintf("%s_DEBUG_MODE", enviroment))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s := fmt.Sprintf("Welcome %s this with Enviroment %s", appName, enviroment)

		// function just for print log message and status code if debug true
		if debug == "true" {
			printLog(http.StatusOK, s)
		}

		// print payload
		fmt.Fprintf(w, s)
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		// function just for print log message and status code if debug true
		if debug == "true" {
			printLog(http.StatusOK, "pong")
		}

		// print payload
		fmt.Fprintf(w, "pong")
	})

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// function just for print log message and status code if debug true
		if debug == "true" {
			printLog(http.StatusOK, fmt.Sprintf("%q", start))
		}

		// print payload
		fmt.Fprintf(w, "%s", start)
	})

	fmt.Println("It Will be run in port ", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

// function just for print log message and status code if debug true
func printLog(statusCode int, message string) {
	fmt.Println("")
	fmt.Println("Status Code : ", statusCode)
	fmt.Println("Message :", message)
	fmt.Println("")
}
