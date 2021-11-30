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

func req(target string) Response {
	reqUrl := fmt.Sprintf("%s/%s", nekoUrl, target)
	res, err := http.Get(reqUrl)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data Response
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
