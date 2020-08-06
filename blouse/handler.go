package blouse

import (
	"net/http"

	"github.com/kchauxs/api_simple/response"
	"github.com/labstack/echo"
)

func Create(c echo.Context) error {
	m := &Model{}
	err := c.Bind(m)
	if err != nil {
		r := response.Model{
			MensajeError: response.MensajeError{
				"E102",
				"El objeto blusa está mal enviado",
			},
		}
		return c.JSON(http.StatusBadRequest, r)
	}

	d := storage.Create(m)
	r := response.Model{
		MensajeOK: response.MensajeOK{
			"A001",
			"blusa creado correctamente",
		},
		Data: d,
	}
	return c.JSON(http.StatusCreated, r)
}

func GetAll(c echo.Context) error {
	m := &Model{}
	err := c.Bind(m)
	if err != nil {
		r := response.Model{
			MensajeError: response.MensajeError{
				"E102",
				"El objeto blusa está mal enviado",
			},
		}
		return c.JSON(http.StatusBadRequest, r)
	}

	d := storage.GetAll()
	r := response.Model{
		MensajeOK: response.MensajeOK{
			"A111",
			"todas las blusas",
		},
		Data: d,
	}
	return c.JSON(http.StatusOK, r)
}
