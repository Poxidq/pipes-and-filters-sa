package main

import (
	"paf/config"
	"paf/services/api"
	"paf/services/filter"
	"paf/services/publish"
	"paf/services/screaming"
	"paf/shared"
)

func main() {
	config.LoadConfig()

	api := &api.ApiService
	filter := &filter.Filter{}
	screaming := &screaming.Screaming{}
	publish := &publish.Publish{}
	screaming.SetNext(publish)
	filter.SetNext(screaming)
	api.SetNext(filter)

	api.Execute(&shared.Process{})
}
