package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/rs/zerolog/log"
)

type BankJSON struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	Code string `json:"code"`
	USSD string `json:"ussd"`
}

type Bank struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	Code string `json:"code"`
	USSD string `json:"ussd"`
	Logo string `json:"logo"`
}

func SayHello(name string) string {
	return "Hello " + name
}

func getBanks(host string) []Bank {
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
			Name: bank.Name,
			Slug: bank.Slug,
			Code: bank.Code,
			USSD: bank.USSD,
			Logo: host + "/logo/" + GetUrl(bank.Slug) + ".png",
		})
	}

	return newBanks

}
