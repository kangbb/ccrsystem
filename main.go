package main

import (
	"encoding/gob"
	"os"
	"time"

	flag "github.com/spf13/pflag"
	"github.com/urfave/negroni"
)

func main() {
	gob.Register(time.Time{})

	var PORT = os.Getenv("PORT")
	if len(PORT) == 0 {
		PORT = "8080"
	}
	var port = flag.StringP("port", "p", PORT, "Define the port where server runs")
	flag.Parse()

	// Start a server.
	r := routers.GetRouters()
	s := negroni.Classic()
	s.UseHandler(r)
	s.Run(":" + *port)
}
