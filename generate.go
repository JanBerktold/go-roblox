package main

import (
	"os/exec"
	"bytes"
	"fmt"
	"os"
)

type ApiSite struct {
	URL string
	Name string
}

var apisites []ApiSite = []ApiSite{
	ApiSite{
		Name: "inventory",
		URL: "https://inventory.roblox.com/docs/json/v1",
	},
	ApiSite{
		Name: "groups",
		URL: "https://groups.roblox.com/docs/json/v1",
	},
	ApiSite{
		Name: "develop",
		URL: "https://develop.roblox.com/docs/json/v1",
	},
	ApiSite{
		Name: "billing",
		URL: "https://billing.roblox.com/docs/json/v1",
	},
	ApiSite{
		Name: "avatar",
		URL: "https://billing.roblox.com/docs/json/v1",
	},
	ApiSite{
		Name: "auth",
		URL: "https://auth.roblox.com/docs/json/v1",
	},
}

func main() {
	for _, site := range apisites {

		if err := os.Mkdir(site.Name, 0777); err != nil {
			panic(err)
		}

		cmd := exec.Command("swagger", "generate", "client", "/f", site.URL, "/t", site.Name, "/name", site.Name)
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &out
		if err := cmd.Run(); err != nil {
			fmt.Println(out.String())
			panic(err)
		} else {
			fmt.Println(out.String())

		}
	}
}