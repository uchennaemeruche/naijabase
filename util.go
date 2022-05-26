package main

import (
	"os"

	"github.com/rs/zerolog/log"
)

func GetLogoUrl(folder, slug string) string {
	var files []string
	f, err := os.Open(folder)
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

	_, fileExist := findFile(files, slug+".png")

	if !fileExist {
		return "image-placeholder"
	}

	return slug
}

func findFile(slice []string, value string) (int, bool) {
	for i, item := range slice {
		if item == value {
			return i, true
		}

	}
	return -1, false
}
