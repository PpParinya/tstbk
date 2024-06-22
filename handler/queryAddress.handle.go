package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func QueryAddress(ctx *fiber.Ctx) error {

	// userID := ctx.Locals("userID").(string)
	lat := ctx.Query("lat")
	lon := ctx.Query("lon")

	// url := "http://10.12.1.50:8080/api/user/-9223372036854771827/address?lat=13.656456947&lon=100.543205261&lang=en"
	url := "http://10.12.1.50:8080/api/user/-9223372036854771827/address?lat=" + lat + "&lon=" + lon + "&lang=th"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(body))

	// return ctx.JSON(string(body))
	return ctx.JSON(fiber.Map{
		"address": string(body),
	})
}

