package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Error struct {
	Message string
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	err := godotenv.Load(".env.development")
	if err != nil {
		log.Fatal().Msg("Error loading enviroment variables")
		log.Info().Msg("Error loading enviroment variables")
	}
	port, existPort := os.LookupEnv("PORT")

	host, existHost := os.LookupEnv("HOST")

	if !existPort || !existHost {
		log.Fatal().Msg("Port or Host not set in environment variable")
		log.Error().Err(errors.New("Port or Host not set in environment variable")).Msg("")
	}

	log.Info().Msg(port)

	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	route := mux.NewRouter()

	route.PathPrefix("/logo/").Handler(http.StripPrefix("/logo/", http.FileServer(http.Dir("./logos"))))

	route.HandleFunc("/hello", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")

		greeting := SayHello("Uchenna")
		_ = json.NewEncoder(res).Encode(greeting)

	})
	route.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")

		banks := getBanks(host)
		_ = json.NewEncoder(res).Encode(banks)

	})

	route.NotFoundHandler = http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusNotFound)

		res.Header().Set("Content-Type", "application/json")

		_ = json.NewEncoder(res).Encode(Error{Message: "Oops!! Unavailable route"})
	})

	// handler := cors.Default().Handler(route)
	handler := cors.AllowAll().Handler(route)

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		fmt.Print(host)
		log.Log().Err(err)
	}
}

func GetUrl(bankSlug string) string {
	var files []string
	f, err := os.Open("./bank-logos")
	if err != nil {
		log.Log().Err(err)
	}

	fileInfo, err := f.Readdir(0)

	_ = f.Close()

	if err != nil {
		log.Log().Err(err)
	}

	for _, file := range fileInfo {
		if file.Name() == ".DS_store" {
			continue
		}
		files = append(files, file.Name())
	}

	_, found := findFile(files, bankSlug+".png")

	if found {
		return bankSlug
	}

	return "default-image"
}

func findFile(slice []string, value string) (int, bool) {
	for i, item := range slice {
		if item == value {
			return i, true
		}

	}
	return -1, false
}
