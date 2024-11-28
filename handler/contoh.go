package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ContohHandler(c *fiber.Ctx) error {
	zParam := c.Params("z")
	var x int = 2
	var y int = 2
	var sum int = x + y
	z, err := strconv.Atoi(zParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "z harus diisi dengan angka",
		})
	}
	if z != sum {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Jawaban salah",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Jawaban benar",
	})
}
