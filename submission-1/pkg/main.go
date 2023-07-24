package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type text string
type User struct {
	id      int
	name    string
	address string
	job     string
}
type Student struct {
	user   User
	reason string
}

func main() {
	argsInput := os.Args[1]

	users := [3]Student{
		{user: User{id: 1, name: "Haikal", address: "bekasi", job: "programer"}, reason: "ingin jadi programer"},
		{user: User{id: 2, name: "Wawan", address: "jakarta", job: "programer"}, reason: "ingin belajar"},
		{user: User{id: 3, name: "Andi", address: "depok", job: "programer"}, reason: "gabut"},
	}

	for _, element := range users {

		if element.user.name == argsInput {
			a := map[string]interface{}{
				"nama":      element.user.name,
				"alamat":    element.user.address,
				"pekerjaan": element.user.job,
				"alasan":    element.reason}
			bs, _ := json.Marshal(a)
			fmt.Println(string(bs))
			break
		}

	}
	// myLog := logger.MyLogger{
	// 	ServiceName: "asdsada",
	// 	ServiceType: "",
	// }

	// fmt.Println(logger.Name)

	// fmt.Println(myLog)
}
