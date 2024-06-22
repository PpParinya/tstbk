package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func NameAddress(lat string, lon string) string {

	// userID := ctx.Locals("userID").(string)

	// url := "http://10.12.1.50:8080/api/user/-9223372036854771827/address?lat=13.656456947&lon=100.543205261&lang=en"
	url := "http://10.12.1.50:8080/api/user/-9223372036854771827/address?lat=" + lat + "&lon=" + lon + "&lang=th"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	// fmt.Println(string(body))

	// return ctx.JSON(string(body))
	return string(body)
}