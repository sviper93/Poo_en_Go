package main

import (
	"fmt"
)

func main() {

	Go := &course.Course{
		"Go desde cero",
		12.34,
		false,
		[]uint{12, 56, 89},
		map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Maps",
		},
	}

	Go.PrintClasses()
	Go.ChangePrice(67.12)
	fmt.Println(Go.Price)

}

/*

En Go no tenemos los modificadores de acceso como public, protected, private como en otros lenguajes
pero sí tenemos los identificadores exportados y no exportados, quienes nos permiten indicarle a Go
que queremos y no exportar, de esta manera, podemos restringir el acceso a ciertos métodos o características
al usuario que trabaje con nuestro paquete.

Son identificadores exportados son todos aquellos cuyo nombre empiece con mayúscula, los métodos y nombres
de campos cuya primera letra esté en mayúscula también son exportados.

1. Creamos la carpeta "course" y dentro movemos "course.go", además cambiamos el paquete a course.
2. Creamos el modulo a "course" para poder acceder a sus propiedades y métodos desde acá.
*/
