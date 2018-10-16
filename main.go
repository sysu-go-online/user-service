package main

import (
	"os"

	flag "github.com/spf13/pflag"
	"github.com/sysu-go-online/service-end/router"
)

func main() {
	var PORT = os.Getenv("PORT")
	if len(PORT) == 0 {
		PORT = "8080"
	}
	var port = flag.StringP("port", "p", PORT, "Define the port where user service runs")
	flag.Parse()

	s := router.GetServer()
	s.Run(":" + *port)
}
