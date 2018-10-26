package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/libgoost/gweb"
)

var cfgPath = flag.String("c", "config.json", "config file path")

func main() {
	flag.Parse()
	cfg := &gweb.Config{}

	data, err := ioutil.ReadFile(*cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, cfg)
	if err != nil {
		log.Fatal(err)
	}
	route := gweb.NewService(cfg)
	log.Fatal(http.ListenAndServe(cfg.ListenPort, route))
}
