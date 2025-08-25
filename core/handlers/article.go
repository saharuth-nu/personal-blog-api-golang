package handlers

import (
	"blog-api/core/models"
	"blog-api/core/services"
	"blog-api/utils"

	"github.com/gofiber/fiber/v2"
)

type articleHandler struct {
	articleSrv services.ArticleService
}

func NewArticleHandler(articleSrv services.ArticleService) articleHandler {
	return articleHandler{articleSrv: articleSrv}
}

func (h articleHandler) GetArticles(c *fiber.Ctx) error {

	response, err := h.articleSrv.GetArticles()
	if err != nil {
		return utils.ErrorFormat(c, fiber.StatusInternalServerError, err.Error())
		// return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// return c.Status(fiber.StatusOK).JSON(response)
	return utils.SuccessFormat(c, fiber.StatusOK, "", response)
}

func (h articleHandler) GetArticleByUID(c *fiber.Ctx) error {

	param := struct {
		UID string `params:"uid"`
	}{}
	if err := c.ParamsParser(&param); err != nil {
		// return fiber.NewError(fiber.ErrBadRequest.Code, err.Error())
		return utils.ParamParserFail(c)
	}

	response, err := h.articleSrv.GetArticleByUID(param.UID)
	if err != nil {
		// return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		return utils.ErrorFormat(c, fiber.StatusInternalServerError, err.Error())
	}

	// return c.Status(fiber.StatusOK).JSON(response)
	return utils.SuccessFormat(c, fiber.StatusOK, "", response)
}

func (h articleHandler) CreateArticle(c *fiber.Ctx) error {

	b := new(models.NewArticleRequest)

	if err := c.BodyParser(b); err != nil {
		// return fiber.NewError(fiber.ErrBadRequest.Code, err.Error())
		return utils.BodyParserFail(c)
	}

	request := models.NewArticleRequest{
		Title:   b.Title,
		Content: b.Content,
		Tag:     b.Tag,
	}

	err := h.articleSrv.CreateArticle(request)
	if err != nil {
		// return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		return utils.ErrorFormat(c, fiber.StatusInternalServerError, err.Error())
	}

	// return c.Status(fiber.StatusCreated).JSON(response)
	return utils.SuccessFormat(c, fiber.StatusOK, "success", nil)
}

func (h articleHandler) DeleteArticleByUID(c *fiber.Ctx) error {
	param := struct {
		UID string `params:"uid"`
	}{}
	if err := c.ParamsParser(&param); err != nil {
		// return fiber.NewError(fiber.ErrBadRequest.Code, err.Error())
		return utils.ParamParserFail(c)
	}

	err := h.articleSrv.DeleteArticleByUID(param.UID)
	if err != nil {
		// return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		return utils.ErrorFormat(c, fiber.StatusInternalServerError, err.Error())
	}

	// return c.Status(fiber.StatusOK).JSON(response)
	return utils.SuccessFormat(c, fiber.StatusOK, "success", nil)
}

func (h articleHandler) UpdateArticleByUID(c *fiber.Ctx) error {
	param := struct {
		UID string `params:"uid"`
	}{}
	if err := c.ParamsParser(&param); err != nil {
		// return fiber.NewError(fiber.ErrBadRequest.Code, err.Error())
		return utils.ParamParserFail(c)
	}

	b := new(models.NewArticleRequest)

	if err := c.BodyParser(b); err != nil {
		// return fiber.NewError(fiber.ErrBadRequest.Code, err.Error())
		return utils.BodyParserFail(c)
	}

	request := models.NewArticleRequest{
		Title:   b.Title,
		Content: b.Content,
		Tag:     b.Tag,
	}

	err := h.articleSrv.UpdateArticleByUID(param.UID, request)
	if err != nil {
		// return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		return utils.ErrorFormat(c, fiber.StatusInternalServerError, err.Error())
	}

	// return c.Status(fiber.StatusOK).JSON(response)
	return utils.SuccessFormat(c, fiber.StatusOK, "success", nil)

}
