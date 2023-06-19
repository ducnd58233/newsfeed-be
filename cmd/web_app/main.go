package main

import (
	"flag"
	"log"

	"github.com/ducnd58233/newsfeed-be/configs"
	"github.com/ducnd58233/newsfeed-be/internal/app/web_app"
	"github.com/ducnd58233/newsfeed-be/internal/app/web_app/service"
)

var path = flag.String("config", "config.yml", "config path for this service")

func main() {
	flag.Parse()
	conf, err := configs.GetWebConfig(*path)
	log.Println(conf)
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}
	webSvc, err := service.NewWebService(conf)
	if err != nil {
		log.Fatalf("failed to init service: %v", err)
	}
	web_app.WebController{
		WebService: *webSvc,
		Port:       conf.Port,
	}.Run()
	log.Println("web app started.")
}
