package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/rs/zerolog/log"
)

type SchoolJSON struct {
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	Location string `json:"location"`
	Address  string `json:"address"`
	Ranking  string `json:"ranking"`
}

type School struct {
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	Location string `json:"location"`
	Address  string `json:"address"`
	Ranking  string `json:"ranking"`
	Logo     string `json:"logo"`
}

func getSchools(host string) []School {
	var folder = "./school-logos"
	log.Info().Msg("Called getSchools")

	schoolJson, err := ioutil.ReadFile("./schools.json")
	if err != nil {
		log.Log().Err(err)
	}

	var schools []SchoolJSON
	if err := json.Unmarshal(schoolJson, &schools); err != nil {
		log.Log().Err(err)
	}

	var newSchools []School
	for _, school := range schools {
		newSchools = append(newSchools, School{
			Name:     school.Name,
			Slug:     school.Slug,
			Location: school.Location,
			Address:  school.Address,
			Ranking:  school.Ranking,
			Logo:     host + "/logo/" + GetLogoUrl(folder, school.Slug) + ".png",
		})
	}

	return newSchools
}
