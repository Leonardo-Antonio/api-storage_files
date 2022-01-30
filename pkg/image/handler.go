package image

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/Leonardo-Antonio/api-storage_files/pkg/response"
	"github.com/aidarkhanov/nanoid"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	store *redis.Client
}

func (h *handler) Save(ctx *fiber.Ctx) error {
	body := new(body)
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.Err(err.Error(), nil))
	}

	buf, err := base64.StdEncoding.DecodeString(body.ImageB64)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.Err(err.Error(), nil))
	}

	if err := h.writeFile(body.Name, buf); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.Err(err.Error(), nil))
	}

	statusDel := h.store.Del(context.TODO(), "images")
	if statusDel.Err() != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.Err(err.Error(), nil))
	}

	return ctx.Status(http.StatusCreated).JSON(response.Success("saved", nil))
}

func (h *handler) GetAll(ctx *fiber.Ctx) error {

	images, err := h.store.Get(context.TODO(), "images").Result()
	if err != redis.Nil {
		result := new([]infoImage)
		json.Unmarshal([]byte(images), result)
		return ctx.Status(http.StatusOK).JSON(response.Success("ok", result))
	}

	infoImages, err := h.readFiles(ctx.BaseURL())
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.Err(err.Error(), nil))
	}

	buf, err := json.Marshal(&infoImages)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.Err(err.Error(), nil))
	}

	statusSet := h.store.Set(context.TODO(), "images", buf, time.Hour*24)
	if statusSet.Err() != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.Err(err.Error(), nil))
	}

	return ctx.Status(http.StatusOK).JSON(response.Success("ok", infoImages))
}

func (h *handler) Remove(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	if err := os.Remove(fmt.Sprintf("static/images/%s", name)); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.Err(err.Error(), nil))
	}

	return ctx.Status(http.StatusOK).JSON(response.Success("removed", name))
}

func (h *handler) writeFile(name string, imageB64 []byte) error {
	if err := ioutil.WriteFile(
		fmt.Sprintf("static/images/%s___%s", nanoid.New(), name),
		imageB64, 0777,
	); err != nil {
		return err
	}

	return nil
}

func (h *handler) readFiles(baseURL string) (*[]infoImage, error) {
	fileInfo, err := ioutil.ReadDir("static/images")
	if err != nil {
		return nil, err
	}

	infoImages := new([]infoImage)
	for _, info := range fileInfo {
		image := infoImage{
			Name:         info.Name(),
			Src:          fmt.Sprintf("%s%s/images/public/%s", baseURL, os.Getenv("VERSION_API"), info.Name()),
			Size:         info.Size(),
			Modification: info.ModTime(),
		}
		*infoImages = append(*infoImages, image)
	}

	return infoImages, nil
}
