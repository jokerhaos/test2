package main

import (
	"encoding/json"
	"fmt"
)

type App struct {
	Id string `json:"id"`
}

type Org struct {
	Name string `json:"name"`
}

type AppWithOrg struct {
	App
	Org
}

func main() {
	data := []byte(`
        {
            "id": "k34rAT4",
            "name": "My Awesome Org"
        }
    `)

	var b AppWithOrg

	json.Unmarshal(data, &b)
	fmt.Println("==============================")
	fmt.Printf("%#v", b)
	fmt.Println()

	a := AppWithOrg{
		App: App{
			Id: "k34rAT4",
		},
		Org: Org{
			Name: "My Awesome Org",
		},
	}
	data, _ = json.Marshal(a)
	fmt.Println("==============================")
	fmt.Println(string(data))
}
