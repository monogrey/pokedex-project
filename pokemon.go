package main

//type pType int

type Pokemon struct {
	ID   int    `json: "id"`
	Name string `json: "name"`
	Desc string `json: "desc"`
	/* to be isnerted: pokemon types
	likely to the format of:
	TypeI int `json: "typei"`
	TypeII int `json: "typeii"`
	*/
}

const (
	fire = iota
	water
	grass
)
