package main

import (
	"fmt"
	"github.com/maddyonline/sleepy"
	"net/http"
	"net/url"
)

type ReplayResource struct{}

type Item struct{}

func (rr ReplayResource) Get(values url.Values, headers http.Header) (int, interface{}, http.Header) {
	fmt.Println(values)
	return 200, values, http.Header{"Content-type": {"application/json"}}
}

func (item Item) Get(values url.Values, headers http.Header) (int, interface{}, http.Header) {
	items := []string{"item1", "item2"}
	data := map[string][]string{"items": items}
	return 200, data, http.Header{"Content-type": {"application/json"}}
}

func main() {
	item := new(Item)
	api := sleepy.NewAPI()
	api.AddResource(item, "/items")

	rr := new(ReplayResource)
	api.AddResource(rr, "/replay")
	api.Start(3000)
}
