// Licensed under the Open Software License version 3.0

package nekos

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const neko = "https://nekos.life/api/v2"

func req(endpoint string) (Response, error) {
	req := fmt.Sprintf("%s/%s", neko, endpoint)
	res, err := http.Get(req)
	if err != nil {
		return Response{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Response{}, err
	}

	var data Response
	err = json.Unmarshal(body, &data)
	if err != nil {
		return Response{}, err
	}
	return data, nil
}
