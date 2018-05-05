package main

import (
	"encoding/gob"
	"os"
	"time"

	"github.com/kangbb/ccrsystem/core/controller"
	flag "github.com/spf13/pflag"
)

func main() {
	gob.Register(time.Time{})

	var PORT = os.Getenv("PORT")
	if len(PORT) == 0 {
		PORT = "8080"
	}
	var port = flag.StringP("port", "p", PORT, "Define the port where server runs")
	flag.Parse()

	s := controller.GetServer()
	s.Run(":" + *port)
}
