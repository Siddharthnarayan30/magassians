package util

import (
	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load(".env")
	if err != nil {
		Log.Error().Msg("Error in loading .env file")
	}

}
