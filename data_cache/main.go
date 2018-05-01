package main

import (
	"github.com/patrickmn/go-cache"
	"time"
	"fmt"
)

var (
	c = cache.New(10 * time.Second, 10 * time.Second)
)
func main() {
	//c.Set("data", "123456789", cache.DefaultExpiration)
	for {
		fmt.Println(provideData())
		time.Sleep( 1* time.Second)
	}
}

func provideData() string {
	data, found := c.Get("data")
	if found {
		return data.(string)
	} else {
		println("making cache ...")
		c.Set("data", "123456789", cache.DefaultExpiration)
		data, found := c.Get("data")
		if found {
			return data.(string)
		} else {
			return ""
		}
	}
}