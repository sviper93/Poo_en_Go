package course

import "fmt"

type Course struct { // Como el nombre de la estructura y el resto de campos
	Name    string // empiezan con mayúsculas son exportados
	Price   float64
	IsFree  bool   // al cambiarle "IsFree" por "isFree" hacemos que el campo sea
	UserIDs []uint // no exportado y por ende, no puede ser accedido por nadie
	Classes map[uint]string
}

func (c *Course) ChangePrice(price float64) { // Estos también son exportados
	c.Price = price
}

func (c *Course) PrintClasses() { // Estos también son exportados
	text := "Las clases son: "
	for _, class := range c.Classes {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}
