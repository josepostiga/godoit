package tasks

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/josepostiga/godoit/internal"
	"github.com/josepostiga/godoit/internal/tasks/repositories"
	"net/http"
)

func index(c *fiber.Ctx) error {
	var r = repositories.New()
	var tasks, _ = r.FindAll()

	return responses.New(c, tasks, fiber.StatusOK)
}

func store(c *fiber.Ctx) error {
	var t *repositories.Task
	var r = repositories.New()

	if err := json.NewDecoder(bytes.NewReader(c.BodyRaw())).Decode(&t); err != nil {
		return err
	}

	err := r.Create(t)
	if err != nil {
		return responses.New(c, &fiber.Map{"error": err.Error()}, fiber.StatusBadRequest)
	}

	return responses.New(c, t, fiber.StatusCreated)
}

func update(c *fiber.Ctx) error {
	var t *repositories.Task
	var r = repositories.New()
	var id, _ = c.ParamsInt("id")

	if err := json.NewDecoder(bytes.NewReader(c.BodyRaw())).Decode(&t); err != nil {
		return err
	}

	t.Id = id

	err := r.Update(t)
	if err != nil {
		return responses.New(c, &fiber.Map{"error": err.Error()}, fiber.StatusBadRequest)
	}

	return responses.New(c, t, fiber.StatusOK)
}

func show(c *fiber.Ctx) error {
	var id, _ = c.ParamsInt("id")
	var r = repositories.New()

	t, err := r.FindById(id)
	if err != nil {
		return responses.New(c, &fiber.Map{"error": err.Error()}, fiber.StatusNotFound)
	}

	return responses.New(c, t, fiber.StatusOK)
}

func delete(c *fiber.Ctx) error {
	var r = repositories.New()
	var id, _ = c.ParamsInt("id")

	err := r.Delete(id)
	if err != nil {
		return responses.New(c, &fiber.Map{"error": err.Error()}, fiber.StatusNotFound)
	}

	return responses.New(c, nil, http.StatusNoContent)
}

func status(c *fiber.Ctx) error {
	var r = repositories.New()
	var id, _ = c.ParamsInt("id")

	err := r.ToggleStatus(id)
	if err != nil {
		return responses.New(c, &fiber.Map{"error": err.Error()}, fiber.StatusNotFound)
	}

	return responses.New(c, nil, fiber.StatusOK)
}
