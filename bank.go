package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/rs/zerolog/log"
)

type BankJSON struct {
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	Class     string `json:"class"`
	SortCode  string `json:"sort_code"`
	USSDCode  string `json:"ussd"`
	SwiftCode string `json:"swift_code"`
	Website   string `json:"website"`
}

type Bank struct {
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	Class     string `json:"class"`
	SortCode  string `json:"sort_code"`
	USSDCode  string `json:"ussd_code"`
	SwiftCode string `json:"swift_code"`
	Website   string `json:"website"`
	Logo      string `json:"logo"`
}

func SayHello(name string) string {
	return "Hello " + name
}

func getBanks(host string) []Bank {
	var folder = "./bank-logos"
	log.Info().Msg("Called getBanks")
	bankJson, err := ioutil.ReadFile("./banks.json")
	if err != nil {
		log.Log().Err(err)
	}

	var banks []BankJSON

	if err := json.Unmarshal(bankJson, &banks); err != nil {
		log.Log().Err(err)
	}
	var newBanks []Bank
	for _, bank := range banks {
		newBanks = append(newBanks, Bank{
			Name:      bank.Name,
			Slug:      bank.Slug,
			Class:     bank.Class,
			SortCode:  bank.SortCode,
			SwiftCode: bank.SwiftCode,
			USSDCode:  bank.USSDCode,
			Website:   bank.Website,
			Logo:      host + "/logo/" + GetLogoUrl(folder, bank.Slug) + ".png",
		})
	}

	return newBanks

}
