package util

import "github.com/goddtriffin/helmet"

func HelmetHandler() *helmet.Helmet {
	return helmet.Default()
}
