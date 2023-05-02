package variables

import (
	"fmt"
	"strconv"
	"time"
)

/*Nota: Las variables declaradas por fuera de una funcion son de tipo global y hay de dos tipos
La que primera en mayuscula -> Se puede acceder desde cualquier archivo o funcion
la que es en minuscula -> Solo se puede accer en la funcion de ese mismo archivo
*/

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Pedro"
	Estado = true
	Sueldo = 1577.66
	Fecha = time.Now()

	cadena := fmt.Sprintf(`
	Nombre: %v,
	Estado: %v,
	Sueldo: %v,
	Fecha: %v
	`, Nombre, Estado, Sueldo, Fecha)

	fmt.Println(cadena)
}

func ConviertoATexto(number int) (bool, string) {
	texto := strconv.Itoa(number)
	return true, texto
}
