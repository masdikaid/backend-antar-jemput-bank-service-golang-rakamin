package api

import (
	"backend-a-antar-jemput/internal/contract"
	"backend-a-antar-jemput/tools/helper"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetProvince(c *fiber.Ctx) error {
	var provinces []*contract.Province
	var err error
	res, err := http.Get("http://www.emsifa.com/api-wilayah-indonesia/api/provinces.json")
	if err != nil {
		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, err.Error())

	}
	data, err := ioutil.ReadAll(res.Body)
	json.Unmarshal(data, &provinces)
	if err != nil {
		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, err.Error())
	}

	return helper.JsonResponseOkBuilder(c, fiber.StatusOK, "Success", provinces)
}

func GetCity(c *fiber.Ctx) error {
	var cities []*contract.City
	var err error
	selectProvince := contract.Province{}
	err = c.BodyParser(&selectProvince)
	if err != nil {
		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, err.Error())
	}

	res, err := http.Get("http://www.emsifa.com/api-wilayah-indonesia/api/regencies/" + selectProvince.Id + ".json")
	if err != nil {
		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, err.Error())

	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, err.Error())
	}
	json.Unmarshal(data, &cities)
	return helper.JsonResponseOkBuilder(c, fiber.StatusOK, "Success", cities)
}

func GetDistrict(c *fiber.Ctx) error {
	var districts []*contract.District
	var err error
	selectCity := contract.City{}
	err = c.BodyParser(&selectCity)
	if err != nil {
		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, err.Error())
	}

	res, err := http.Get("http://www.emsifa.com/api-wilayah-indonesia/api/districts/" + selectCity.Id + ".json")
	if err != nil {
		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, err.Error())

	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return helper.JsonResponseFailBuilder(c, fiber.StatusBadRequest, err.Error())
	}
	json.Unmarshal(data, &districts)
	return helper.JsonResponseOkBuilder(c, fiber.StatusOK, "Success", districts)
}
