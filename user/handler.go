package user

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/kchauxs/api_simple/response"
	"github.com/labstack/echo"
)

func Create(c echo.Context) error {
	u := &Model{}
	err := c.Bind(u)
	if err != nil {
		r := response.Model{
			MensajeError: response.MensajeError{
				Codigo:    "U001",
				Contenido: "El objeto usuario no tiene la estructura correcta",
			},
		}

		return c.JSON(http.StatusBadRequest, r)
	}

	u = storage.Create(u)
	r := response.Model{
		MensajeOK: response.MensajeOK{
			Codigo:    "U200",
			Contenido: "Creado correctamente",
		},
		Data: u,
	}

	return c.JSON(http.StatusCreated, r)
}

func Update(c echo.Context) error {
	u := &Model{}
	email := c.Param("email")
	err := c.Bind(u)
	if err != nil {
		r := response.Model{
			MensajeError: response.MensajeError{
				Codigo:    "U001",
				Contenido: "El objeto usuario no tiene la estructura correcta",
			},
		}

		return c.JSON(http.StatusBadRequest, r)
	}

	u = storage.Update(email, u)
	r := response.Model{
		MensajeOK: response.MensajeOK{
			Codigo:    "U201",
			Contenido: "Actualizado correctamente",
		},
		Data: u,
	}

	return c.JSON(http.StatusOK, r)
}

func Delete(c echo.Context) error {
	email := c.Param("email")

	storage.Delete(email)
	r := response.Model{
		MensajeOK: response.MensajeOK{
			Codigo:    "U202",
			Contenido: "Borrado correctamente",
		},
	}

	return c.JSON(http.StatusOK, r)
}

func GetByEmail(c echo.Context) error {
	email := c.Param("email")
	u := storage.GetByEmail(email)
	if u == nil {
		r := response.Model{
			MensajeError: response.MensajeError{
				Codigo:    "U003",
				Contenido: "El email no se encuentra",
			},
		}

		return c.JSON(http.StatusNotFound, r)
	}

	rsrc := strings.Split(c.Request().RequestURI, email)[0]
	rsrc = rsrc[:len(rsrc)-1]

	n1 := response.Navegacion{
		Descripcion: "Self",
		Link:        c.Request().RequestURI,
	}
	n2 := response.Navegacion{
		Descripcion: "Resource",
		Link:        rsrc,
	}
	ns := make([]response.Navegacion, 0)
	ns = append(ns, n1)
	ns = append(ns, n2)

	r := response.Model{
		MensajeOK: response.MensajeOK{
			Codigo:    "U204",
			Contenido: "Consultado correctamente",
		},
		Data: struct {
			Data       interface{}           `json:"data"`
			Navegacion []response.Navegacion `json:"navegacion"`
		}{
			u,
			ns,
		},
	}

	return c.JSON(http.StatusOK, r)
}

func GetAll(c echo.Context) error {
	us := storage.GetAll()

	r := response.Model{
		MensajeOK: response.MensajeOK{
			Codigo:    "U205",
			Contenido: "Consultado correctamente",
		},
		Data: us,
	}

	return c.JSON(http.StatusOK, r)
}

func GetAllPaginate(c echo.Context) error {
	l := c.QueryParam("limit")
	p := c.QueryParam("page")

	limit, err := strconv.Atoi(l)
	if err != nil {
		limit = 1
	}
	page, err := strconv.Atoi(p)
	if err != nil {
		page = 1
	}

	us := storage.GetAllPaginate(limit, page)

	r := response.Model{
		MensajeOK: response.MensajeOK{
			Codigo:    "U205",
			Contenido: "Consultado correctamente",
		},
		Data: us,
	}

	return c.JSON(http.StatusOK, r)
}

//Login .
func Login(c echo.Context) error {
	u := &Model{}
	err := c.Bind(u)
	if err != nil {
		r := response.Model{
			MensajeError: response.MensajeError{
				Codigo:    "E001",
				Contenido: "Formato incorrecto",
			},
		}

		return c.JSON(http.StatusBadRequest, r)
	}

	d := storage.Login(u.Email, u.Password)
	if d == nil {
		r := response.Model{
			MensajeError: response.MensajeError{
				Codigo:    "L001",
				Contenido: "Usuario o password incorrectos",
			},
		}

		return c.JSON(http.StatusBadRequest, r)
	}
	d.Password = ""
	token, err := generateJWT(*d)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "no se pudo generar el token")
	}
	type logueo struct {
		Usuario Model
		Token   string
	}

	l := logueo{
		*d,
		token,
	}

	r := response.Model{
		MensajeOK: response.MensajeOK{
			Codigo:    "O001",
			Contenido: "Logueado ok",
		},
		Data: l,
	}

	return c.JSON(http.StatusOK, r)
}

// getTokenFromAuthorizationHeader busca el token del header Authorization
func getTokenFromAuthorizationHeader(r *http.Request) (string, error) {
	ah := r.Header.Get("Authorization")
	if ah == "" {
		return "", errors.New("el encabezado no contiene la autorización")
	}

	// Should be a bearer token
	if len(ah) > 6 && strings.ToUpper(ah[0:6]) == "BEARER" {
		return ah[7:], nil
	}
	return "", errors.New("el header no contiene la palabra Bearer")

}

// getTokenFromURLParams busca el token de la URL
func getTokenFromURLParams(r *http.Request) (string, error) {
	ah := r.URL.Query().Get("authorization")
	if ah == "" {
		return "", errors.New("la URL no contiene la autorización")
	}

	return ah, nil
}
