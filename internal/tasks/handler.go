package tasks

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/josepostiga/godoit/internal/tasks/repositories"
	"net/http"
	"os"
)

func response(c *fiber.Ctx, data interface{}, status int) error {
	c.Status(status)

	if data == nil {
		return nil
	}

	return c.JSON(&fiber.Map{
		"data": data,
	})
}

func index(c *fiber.Ctx) error {
	tasks, _ := repositories.NewRepository(os.Getenv("DATABASE_DRIVER")).FindAll()

	return response(c, tasks, fiber.StatusOK)
}

func store(c *fiber.Ctx) error {
	var t *repositories.Task

	if err := json.NewDecoder(bytes.NewReader(c.BodyRaw())).Decode(&t); err != nil {
		return err
	}

	err := repositories.NewRepository(os.Getenv("DATABASE_DRIVER")).Create(t)
	if err != nil {
		return response(c, &fiber.Map{"error": err.Error()}, fiber.StatusBadRequest)
	}

	return response(c, t, fiber.StatusCreated)
}

func update(c *fiber.Ctx) error {
	var t *repositories.Task
	id, _ := c.ParamsInt("id")

	if err := json.NewDecoder(bytes.NewReader(c.BodyRaw())).Decode(&t); err != nil {
		return err
	}

	t.Id = id

	err := repositories.NewRepository(os.Getenv("DATABASE_DRIVER")).Update(t)
	if err != nil {
		return response(c, &fiber.Map{"error": err.Error()}, fiber.StatusBadRequest)
	}

	return response(c, t, fiber.StatusOK)
}

func show(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	t, err := repositories.NewRepository(os.Getenv("DATABASE_DRIVER")).FindById(id)
	if err != nil {
		return response(c, &fiber.Map{"error": err.Error()}, fiber.StatusNotFound)
	}

	return response(c, t, fiber.StatusOK)
}

func delete(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	err := repositories.NewRepository(os.Getenv("DATABASE_DRIVER")).Delete(id)
	if err != nil {
		return response(c, &fiber.Map{"error": err.Error()}, fiber.StatusNotFound)
	}

	return response(c, nil, http.StatusNoContent)
}
