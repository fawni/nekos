// Licensed under the Open Software License version 3.0

package nekos

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const nekoUrl = "https://nekos.life/api/v2/"

func DoRequest(target string) (res Responses) {
	reqUrl := fmt.Sprintf("%s/%s", nekoUrl, target)
	resp, err := http.Get(reqUrl)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data Responses
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
