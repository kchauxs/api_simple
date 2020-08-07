package user

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

var storage Storage

func init() {
	storage = make(map[string]*Model)
	u := &Model{
		FirstName: "admin",
		Email:     "admin@gmail.com",
		Password:  "pass120012",
	}

	storage.Create(u)
}

//Model Instancias
type Model struct {
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

//Storage objeto que simula la base datos
type Storage map[string]*Model

//Create crea usuarios
func (s Storage) Create(m *Model) *Model {
	s[m.Email] = m
	return s[m.Email]
}

//GetAll todos los usuarios
func (s Storage) GetAll() Storage {
	return s
}

//GetAllPaginate .
func (s Storage) GetAllPaginate(l, p int) []*Model {
	us := make([]*Model, 0, len(s))
	for _, v := range s {
		us = append(us, v)
		fmt.Println(v)
	}
	fmt.Println(us)
	offset := l*p - l
	r := us[offset : l*p]
	return r
}

//GetByEmail .
func (s Storage) GetByEmail(e string) *Model {
	if v, ok := s[e]; ok {
		return v
	}

	return nil
}

//Delete .
func (s Storage) Delete(e string) {
	delete(s, e)
}

//Update .
func (s Storage) Update(e string, z *Model) *Model {
	s[e] = z
	return s[e]
}

//Login .
func (s Storage) Login(e, p string) *Model {
	for _, v := range s {
		if v.Email == e && v.Password == p {
			return v
		}
	}

	return nil
}

//Claim  .
type Claim struct {
	Usuario Model
	jwt.StandardClaims
}
