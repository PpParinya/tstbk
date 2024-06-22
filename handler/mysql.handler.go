package handler

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/tst/backend/model/entity"
)

func txtFile(fileName string) []string {

	file, err := os.Open("./assets/textfile/"+fileName+".txt")

	if err != nil {
		fmt.Println("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()

	var arrPlate []string
	for _, each_ln := range text {
		// fmt.Println(each_ln)
		arrPlate = append(arrPlate, each_ln)
	}

	return arrPlate
}


func GetCarBrand(ctx *fiber.Ctx) error {

	var brands []entity.Base_car_brands

	brandStr := txtFile("carBrands")
	for i := 0; i < len(brandStr); i++ {
		s := strings.Split(brandStr[i], "|")

		s1, _ := strconv.Atoi(s[1]);

		brands = append(brands, entity.Base_car_brands{ Car_brand_id: s1, Brand_name_en: s[2] })
	}
	
	return ctx.JSON(brands)
}

func GetProvinces(ctx *fiber.Ctx) error {

	var provinces []entity.Base_provinces

	brandStr := txtFile("provinces")
	for i := 0; i < len(brandStr); i++ {
		s := strings.Split(brandStr[i], "|")

		s1, _ := strconv.Atoi(s[1])
		s2, _ := strconv.Atoi(s[2]);
		s5, _ := strconv.Atoi(s[5]);


		provinces = append(provinces, entity.Base_provinces{ Province_id: s1, Country_id: s2, Province_name_en: s[3], Province_name_th: s[4], Thai_dlt_code: s5 })
	}


	return ctx.JSON(provinces)

}

