package controller

import (
	"encoding/json"
	"fiberever/biz/dokter"
	"fiberever/model"
	"strconv"

	"github.com/alisholihindev/go-lib"
	"github.com/gofiber/fiber/v2"
)

// GetDokter godoc
// @tags Dokter
// @Accept  json
// @Produce  json
// @Param id path int true "Dokter ID"
// @Success 200 {object} lib.OutputFormat{Data=model.Dokter}
// @Security BearerAuth
// @Router /dokter/{id} [get]
func GetDokter(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	dokter, err := dokter.Get(id)
	if err != nil {
		lib.CommonLogger().Error(err)
		return c.Status(500).JSON(lib.ResponseFormat(false, err.Error(), nil, nil))
	}
	return c.JSON(lib.ResponseFormat(true, "ok", dokter, nil))
}

// AddDokter godoc
// @tags Dokter
// @Accept  mpfd
// @Produce  json
// @Param nama formData string false  "Nama"
// @Param Alamat formData string false  "Alamat"
// @Success 200 {object} lib.OutputFormat{Data=model.Dokter}
// @Security BearerAuth
// @Router /dokter/ [post]
func AddDokter(c *fiber.Ctx) error {
	a := model.Dokter{}
	if err := c.BodyParser(&a); err != nil {
		lib.CommonLogger().Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(lib.ResponseFormat(false, err.Error(), nil, nil))
	}

	app, err := dokter.Add(a)
	if err != nil {
		lib.CommonLogger().Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(lib.ResponseFormat(false, err.Error(), nil, nil))
	}

	return c.JSON(lib.ResponseFormat(true, "Created", app, nil))
}

// GetDokterAll godoc
// @tags Dokter
// @Accept  json
// @Produce  json
// @Param page query string false "1"
// @Param limit query string false "20"
// @Param order query string false "ID DESC, StageName, Name ASC"
// @Success 200 {object} lib.OutputFormat{Data=lib.Paginator{Records=[]model.Dokter}}
// @Security BearerAuth
// @Router /dokter [get]
func GetDokterAll(c *fiber.Ctx) error {
	var q model.Dokter
	queryString := c.Params("query", "{}")
	page, _ := strconv.Atoi(c.Params("page", "1"))
	limit, _ := strconv.Atoi(c.Params("limit", "20"))
	order := lib.ToSnakeCase(c.Params("order", "ID desc"))

	if err := json.Unmarshal([]byte(queryString), &q); err != nil {
		lib.CommonLogger().Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(lib.ResponseFormat(false, err.Error(), nil, nil))
	}
	dokter, err := dokter.GetAll(q, page, limit, order)
	if err != nil {
		lib.CommonLogger().Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(lib.ResponseFormat(false, err.Error(), nil, nil))
	}
	return c.JSON(lib.ResponseFormat(true, "ok", dokter, nil))

}

// EditDokter godoc
// @tags Dokter
// @accept  json
// @produce  json
// @Param dokter_id path int true "Dokter ID"
// @Param Dokter body model.Dokter true "Analysis"
// @success 200 {object} lib.OutputFormat{data=[]model.Dokter}
// @security BearerAuth
// @Router /dokter/{dokter_id} [put]
func UpdateDokter(c *fiber.Ctx) error {
	ID, _ := strconv.Atoi(c.Params("dokter_id"))

	var param model.Dokter
	if err := c.BodyParser(&param); err != nil {
		lib.CommonLogger().Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(lib.ResponseFormat(false, err.Error(), nil, nil))
	}

	a, err := dokter.Update(ID, &param)
	if err != nil {
		lib.CommonLogger().Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(lib.ResponseFormat(false, err.Error(), nil, nil))
	}
	return c.JSON(lib.ResponseFormat(true, "ok", a, nil))
}

// DeleteDokter godoc
// @tags Dokter
// @Accept  json
// @Produce  json
// @Param id path int true "dokter ID"
// @Success 200 {object} lib.OutputFormat{Data=model.Dokter}
// @Security BearerAuth
// @Router /dokter/{id} [delete]
func DeleteDokter(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	b, err := dokter.Delete(id)

	if err != nil {
		lib.CommonLogger().Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(lib.ResponseFormat(false, err.Error(), nil, nil))
	}

	return c.JSON(lib.ResponseFormat(true, "ok", b, nil))
}
