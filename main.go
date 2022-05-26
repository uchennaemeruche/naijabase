package main

import (
	"errors"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
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

	app := fiber.New()

	app.Static("/logo/", "./public")

	app.Use(cors.New())

	app.Get("/hello", func(c *fiber.Ctx) error {
		greeting := SayHello("Uchenna" + host)

		return c.JSON(greeting)
	})

	app.Get("/banks", func(c *fiber.Ctx) error {
		banks := getBanks(host)
		return c.Status(fiber.StatusOK).JSON(banks)
	})

	app.Get("/schools", func(c *fiber.Ctx) error {
		schools := getSchools(host)
		return c.Status(fiber.StatusOK).JSON(schools)
	})

	// app.Use(logger.New(logger.Config{ // add Logger middleware with config
	//     Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	// }))

	if err := app.Listen(":" + port); err != nil {
		log.Info().Msg(host)
		log.Log().Err(err)
	}

}
