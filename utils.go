// Licensed under the Open Software License version 3.0

package nekos

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const neko = "https://nekos.life/api/v2"

func req(endpoint string) Response {
	req := fmt.Sprintf("%s/%s", neko, endpoint)
	res, err := http.Get(req)
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
