package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Format struct {
	Status string `json:"status"`
	Data User `json:"data"`
}

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Profession string `json:"profession"`
	Avatar     string `json:"avatar"`
	Role       string `json:"role"`
}

func GetUserByID(id int) (Format,error) {
	
	urlUserService := os.Getenv("URL_SERVICE_USERS")
	url:= urlUserService + "/users/" + strconv.Itoa(id) + "/detail"
	
	var user Format
	
	err := GetJson(url, &user)

	if err != nil {

		return user, err

	}

	return user, nil

}

func GetJson(url string, target interface{}) error {
	
	resp, err := http.Get(url)
	fmt.Println(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}