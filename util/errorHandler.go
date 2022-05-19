package util

func CheckError(err error) {

	defer recover()

	if err != nil {
		Log.Fatal().Err(err).Msgf("server error detected")
		panic(err)
	}

}
