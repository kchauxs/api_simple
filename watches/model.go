package blouse

var storage Storage

func init() {
	storage = make(map[string]*Model)
}

//Model ..
type Model struct {
	Marca  string `json:"marca"`
	Precio int    `json:"precio"`
	Color  string `json:"color"`
}

//Storage ..
type Storage map[string]*Model

//Create ..
func (s Storage) Create(m *Model) *Model {
	s[m.Marca] = m
	return s[m.Marca]
}

//GetAll ..
func (s Storage) GetAll() Storage {
	return s
}

//GetByMarca ..
func (s Storage) GetByMarca(m string) *Model {
	if v, ok := s[m]; ok {
		return v
	}

	return nil
}

//Delete ..
func (s Storage) Delete(m string) {
	delete(s, m)
}

//Update ..
func (s Storage) Update(m string, z *Model) {
	s[m] = z
}
