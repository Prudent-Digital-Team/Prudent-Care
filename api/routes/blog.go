package routes

import (
	"api/models"
	"api/utils"
	"encoding/json"
	// "fmt"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/rs/xid"
)

type Blog struct {
	Env    *utils.Enviroment
	Prefix string
}

// Init ...
func (pg *Blog) Init(prefix string, env *utils.Enviroment) error {
	pg.Env = env
	pg.Prefix = prefix

	g := env.Echo.Group(prefix)

	//  Plain routes all users have access
	g.GET("/blog", pg.GetAll)
	g.GET("/blog/:id", pg.Get)

	rt := g.Group("",
		GatePass(CreatePass(
			models.RoleAdmin,
			http.MethodPost,
			http.MethodPut,
			http.MethodGet,
			http.MethodDelete,
		), pg.Env.Log),
	)

	// Admin Routes only

	rt.POST("/blog", pg.Create)
	rt.DELETE("/blog/:id", pg.Delete)
	rt.POST("/blog/:id", pg.Update)
	return nil
}

func (pg *Blog) GetAll(c echo.Context) error {
	res := new(utils.Response)
	pages := new(models.Blog)

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

	list, count, err := pages.All(pg.Env.DB, perPage, limit)

	maxlength := 200;
	for idx, item := range(list){
		if (idx != 0){
			out, err := utils.TruncateHtml([]byte(item.Content), maxlength, "....")
			if err != nil{
				log.Info(err)
			}
			item.Content = string(out) 
		}
	}

	resp := models.BlogListArray{Items: list, Total: count}

	if err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	res.Set("data", resp)
	return c.JSON(http.StatusOK, res)
}


// Get...
func (pg *Blog) Get(c echo.Context) error {
	res := new(utils.Response)
	blog := new(models.Blog)
	slug := c.Param("id")

	// squery := fmt.Sprintf("%%%s%%", slug)

	err := blog.GETBYSLUG(pg.Env.DB, slug)
	if err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	res.SetData(blog)
	return c.JSON(http.StatusOK, res)

}

// Create...
func (pg *Blog) Create(c echo.Context) error {
	response := new(utils.Response)
	page := new(models.Blog)
	form := struct {
		UUID       string          `json:"uuid"`
		Title      string          `json:"title" validate:"required"`
		Content    string          `json:"content" validate:"required"`
		Slug       string          `json:"slug" validate:"required"`
		Author     string          `json:"author" validate:"required"`
		CategoryID int             `json:"category_id" validate:"required"`
		Image      json.RawMessage `json:"image" validate:"required"`
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

	form.UUID = xid.New().String()

	img, err := utils.ImageHandler(form.Image, pg.Env)

	if err != nil {
		form.Image, _ = json.Marshal(struct{}{})
	}

	form.Image, _ = json.Marshal(img)

	copier.Copy(page, form)

	if err != nil {
		response.SetError(err)
		return c.JSON(http.StatusOK, response)
	}

	err = utils.Transact(pg.Env.DB, func(db *gorm.DB) error {
		if err := page.Create(db); err != nil {
			return err
		}
		return nil
	})

	response.Set("Blog", form)
	return c.JSON(http.StatusOK, response)

}

func (pg *Blog) Delete(c echo.Context) error {
	response := new(utils.Response)
	id, _ := utils.ToNum(c.Param("id"))

	page := models.Blog{ID: id}

	err := page.GET(pg.Env.DB, id)

	if err != nil {
		response.SetError(err)
		return c.JSON(http.StatusOK, response)
	}

	err = utils.Transact(pg.Env.DB, func(db *gorm.DB) error {
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

func (pg *Blog) Update(c echo.Context) error {
	response := new(utils.Response)
	page := new(models.Blog)
	id := c.Param("id")
	form := struct {
		Title      string          `json:"title" validate:"required"`
		Content    string          `json:"content" validate:"required"`
		Slug       string          `json:"slug" validate:"required"`
		Author     string          `json:"author" validate:"required"`
		CategoryID int             `json:"category_id" validate:"required"`
		Image      json.RawMessage `json:"image" validate:"required"`
	}{}

	if err := page.GETBYUUID(pg.Env.DB, id); err != nil {
		response.SetError(err)
		return c.JSON(http.StatusOK, response)
	}

	if err := c.Bind(&form); err != nil {
		response.SetError(err)
		return c.JSON(http.StatusOK, response)
	}
	page.UUID = id

	if err := c.Validate(form); err != nil {
		log.Info("err", err)
		response.SetError(err)
		return c.JSON(http.StatusOK, response)
	}

	img, err := utils.ImageHandler(form.Image, pg.Env)

	if err != nil {
		form.Image, _ = json.Marshal(struct{}{})
	}
	form.Image, _ = json.Marshal(img)

	copier.Copy(page, form)

	if err := page.UPDATE(pg.Env.DB); err != nil {
		response.SetError(err)
		return c.JSON(http.StatusOK, response)
	}

	response.SetData(struct{ Message string }{"done"})
	return c.JSON(http.StatusOK, response)

}
