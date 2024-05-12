package main

import (
	"fmt"
	"net"
	"time"

	"github.com/Gophercraft/core/grunt"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	var grunt_client_options grunt.ClientOptions
	grunt_client_options.Endpoint = "192.168.0.252:3724"
	grunt_client_options.IP = net.IPv4(127, 0, 0, 1)
	grunt_client_options.LinkInterval = 5 * time.Second
	grunt_client_options.Build = 5875

	client := grunt.NewClient(&grunt_client_options)
	client.Credentials("user1", "password")

	client.Handle(grunt.ClientConnecting, func(c *grunt.Client) {
		fmt.Println("connecting")
	})
	client.Handle(grunt.ClientConnected, func(c *grunt.Client) {
		fmt.Println("connected")
	})
	client.Handle(grunt.ClientSuccess, func(c *grunt.Client) {
		fmt.Println("successfully logged in")
	})
	client.Handle(grunt.ClientGotRealmlist, func(c *grunt.Client) {
		fmt.Println("Got realmlist")
		fmt.Println(spew.Sdump(c.RealmList()))
	})

	if err := client.Run(); err != nil {
		panic(err)
	}
}
