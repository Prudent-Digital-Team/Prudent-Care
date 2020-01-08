package routes

import (
	"api/models"
	"api/utils"
	"encoding/json"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"

	"github.com/labstack/echo"
)

type Services struct {
	Env    *utils.Enviroment
	Prefix string
}

// Init ...
func (s *Services) Init(prefix string, env *utils.Enviroment) error {
	s.Env = env
	s.Prefix = prefix

	g := env.Echo.Group(prefix)
	g.GET("/services", s.GetAll)
	g.GET("/services/:id", s.getService)

	rt := g.Group("",
		GatePass(CreatePass(
			models.RoleAdmin,
			http.MethodPost,
			http.MethodPut,
			http.MethodGet,
			http.MethodDelete,
		), s.Env.Log),
	)
	rt.POST("/services", s.addServices)
	rt.POST("/services/:id", s.editServices)
	rt.DELETE("/services/:id", s.deleteServices)

	return nil

}

func (s *Services) GetAll(c echo.Context) error {
	res := new(utils.Response)
	service := new(models.Services)

	// var err error

	page, err := utils.ToNum(c.QueryParam("page"))

	perPage, err := utils.ToNum(c.QueryParam("perPage"))

	if err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	limit := (page - 1) * perPage

	if err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	list, count, err := service.All(s.Env.DB, perPage, limit)

	for _, items := range list {
		items.Image, _ = json.Marshal(struct{}{})

	}

	resp := models.ServicesListArray{Items: list, Total: count}

	if err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	res.Set("data", resp)
	return c.JSON(http.StatusOK, res)

}

func (s *Services) getService(c echo.Context) error {
	res := &utils.Response{}

	slug := c.Param("id")

	page := &models.Services{
		URL: slug,
	}

	err := page.GETBYSTRING(s.Env.DB, slug)
	if err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	res.SetData(page)
	return c.JSON(http.StatusOK, res)

}

func (s *Services) addServices(c echo.Context) error {
	response := new(utils.Response)
	page := new(models.Services)
	form := struct {
		Name    string          `json:"name" validate:"required"`
		URL     string          `json:"url" validate:"required"`
		Content string          `json:"content" validate:"required"`
		Image   json.RawMessage `json:"image" validate:"required"`
	}{}

	err := c.Bind(&form)
	if err != nil {
		response.SetError(err)
		return c.JSON(http.StatusOK, response)
	}
	if err = c.Validate(&form); err != nil {
		response.SetError(err)
		return c.JSON(http.StatusOK, response)
	}
	img, err := utils.ImageHandler(form.Image, s.Env)
	form.Image, _ = json.Marshal(img)

	copier.Copy(page, form)

	if err != nil {
		response.SetError(err)
		return c.JSON(http.StatusOK, response)
	}

	err = utils.Transact(s.Env.DB, func(db *gorm.DB) error {
		if err := page.Create(db); err != nil {
			return err
		}
		return nil
	})

	response.Set("Services", form)
	return c.JSON(http.StatusOK, response)
}

func (s *Services) editServices(c echo.Context) error {
	response := new(utils.Response)
	page := new(models.Services)
	id, _ := utils.ToNum(c.Param("id"))

	// page := &models.Services{
	// 	ID: id,
	// }
	form := struct {
		Name    string          `json:"name" validate:"required"`
		URL     string          `json:"url" validate:"required"`
		Content string          `json:"content" validate:"required"`
		Image   json.RawMessage `json:"image" validate:"required"`
	}{}

	if err := page.GET(s.Env.DB, id); err != nil {
		response.SetError(err)
		return c.JSON(http.StatusOK, response)
	}

	if err := c.Bind(&form); err != nil {
		response.SetError(err)
		return c.JSON(http.StatusOK, response)
	}

	page.ID = id

	if err := c.Validate(form); err != nil {
		response.SetError(err)
		return c.JSON(http.StatusOK, response)
	}

	copier.Copy(page, form)

	if err := page.UPDATE(s.Env.DB); err != nil {
		response.SetError(err)
		return c.JSON(http.StatusOK, response)
	}

	response.SetData(struct{ Message string }{"done"})
	return c.JSON(http.StatusOK, response)

}

//Delete ...
func (pg *Services) deleteServices(c echo.Context) error {
	response := new(utils.Response)
	id, _ := utils.ToNum(c.Param("id"))
	page := models.Services{ID: id}

	err := utils.Transact(pg.Env.DB, func(db *gorm.DB) error {
		if err := page.DELETE(db); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		response.SetError(err)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response.SetData(struct{ message string }{"done"})
	return c.JSON(http.StatusOK, response)

}
