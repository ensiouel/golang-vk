package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ensiouel/golang-vk"
)

var (
	Token = flag.String("token", "", "access token")
)

func main() {
	flag.Parse()

	session, err := vk.NewSession(vk.AuthConfig{Token: *Token}, true)
	if err != nil {
		panic(err)
	}

	var wall *vk.Wall
	err = session.Call("wall.get", nil).
		Set("owner_id", 1).
		Set("offset", 0).
		Set("count", 100).
		Execute(&wall)
	if err != nil {
		log.Println(err)
	}

	for _, post := range wall.Items {
		fmt.Println(post.Text)
	}

}
