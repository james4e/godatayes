package godatayes

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"strings"
)

var (
	domain = "api.wmcloud.com"
	token  = ""
)

// Init is for setting the token
func Init(t string) {
	token = t
}

func getData(path string) (string, error) {
	path = "/data/v1" + path
	url := "https://" + domain + path
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("request fail:", err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	if strings.Contains(path, ".csv?") {
		out, err := GBKToUTF8(body)
		if err == nil {
			return string(out), nil
		}
		return "", err
	}

	return string(body), nil

}
