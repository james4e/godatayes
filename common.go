package godatayes

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"strings"

	iconv "github.com/djimenez/iconv-go"
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
	out := make([]byte, len(body))
	if strings.Contains(path, ".csv?") {
		iconv.Convert(body, out, "gb2312", "utf-8")
		return string(out), nil
	} else {
		return string(body), nil
	}

}
