package models

type Clientes struct {
	// Id         int64  `json:"id"sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Nombre     string `json:"nombre"`
	Apellido_1 string `json:"apellido_1"`
	Apellido_2 string `json:"apellido_2"`
	Edad       int64  `json:"edad"`
}

type Clientes_udp struct {
	Id         int64  `json:"id"`
	Nombre     string `json:"nombre"`
	Apellido_1 string `json:"apellido_1"`
	Apellido_2 string `json:"apellido_2"`
	Edad       int64  `json:"edad"`
}

/* type Clientes_basic struct {
	Apellido_2 string `json:"apellido_2"`
	Edad       int64  `json:"edad"`
}
*/
