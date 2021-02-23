# Tutorial: TDD Kata con Golang

![logo-go](./doc/assets/logo-go.png)

## Introducción

En este tutorial se va a construir un pequeño Kata basado en el siguiente reto (https://osherove.com/tdd-kata-1/) útilizando golang.

## Calculadora 

La idea de este ejercicio es aplicarlo diariamente utilizando TDD (patrón red-green-refactor) para acostumbrarnos como desarrolladores a escribir primero las pruebas y luego el código de negocio. Para ello es fundamental hacer el ejercicio diario con una dedicación de al menos 15 min.

_Recomendaciones_

1. Siempre concentrate en hacer una tarea a la vez.
2. Trata de realizar siempre que puedas un trabajo incremental.
3. Asegurate de validar las entradas **Correctas**, ni hay
 necesidad de validar entradas invalidas para este ejercicio.

 El código completo de este ejercicio lo puedes encontrar [aquí](./code/).

 ### Contenido

 - [Día 1](#día-1)
 - [Día 2](#día-2)
 - [Día 2](#día-3)

 ### Día 1

 Crear una calculadora simple con entradas tipo ```string``` a través de una función con la siguiente firma:

 ```golang
func Add(numbers string) int
 ```

- El método debe tomar dos números, separados por comas y retornar su suma, ejemplo de valores de entrada: ```"", "1", "1,2"```.
- Para cadenas vacias debe retornar ```0```.

_Concejos iniciales_:

- Comienza con el caso de prueba más simple de una cadena vacía y pasa a uno y dos números.
- Recuerda resolver las cosas de la forma más sencilla posible para que te obligues a escribir pruebas en las que no pensaste.
- Recuerda refactorizar después de cada prueba aprobada.

**Solución:**

Escribiendo la prueba, importante tener en cuenta el patrón triple AAA:

```golang
func TestAdd(t *testing.T) {
	//Arrange
	expected := 0
	//Act
	outcome := Add("")
	//Assert
	if outcome != expected {
		t.Errorf("Test with empty value failed")
	}
}
```
Validamos que la prueba falle (Red step):

```sh
$ go test
```
```
.\calculator_test.go:9:13: undefined: Add
```

Agregamos el código necesario para pasar la prueba

```golang
func Add(numbers string) int {
	return 0
}
```

Validamos con `go test` que la prueba sea exitosa (Green step)

```
PASS
ok      _.../code/day1  0.237s
```

¡A refactorizar!

Modificamos la prueba para el escenario en el que se recibe un número, sin olvidar el patron trple AAA.

Para agregar un nuevo escenario podemos hacer uso de la herramienta [**_subtest_**](https://golang.org/pkg/testing/#hdr-Subtests_and_Sub_benchmarks) que provee el paquete **_testing_**. En momentos es útil agrupar las pruebas alrededor de "algo" y luego tener subpruebas que describen diferentes escenarios.

En este momento agregaremos un nuevo escenario para validar cuando recibimos un solo número en la función `Add`

```golang
func TestAdd(t *testing.T) {

	t.Run("empty value", func(t *testing.T) {
		//Arrange
		expected := 0
		//Act
		outcome := Add("")
		//Assert
		if outcome != expected {
			t.Errorf("Test with empty value failed")
		}
	})

	t.Run("one value", func(t *testing.T) {
		//Arrange
		expected := 2
		//Act
		outcome := Add("2")
		//Assert
		if outcome != expected {
			t.Errorf("Test with one value failed")
		}
	})
}
```

Validamos con `go test` que la prueba sea fallida nuevamente (Red step)

```
--- FAIL: TestAdd (0.00s)
    --- FAIL: TestAdd/one_value (0.00s)
        calculator_test.go:25: Test with one value failed
FAIL
exit status 1
```

Ahora que tenemos una prueba fallida bien escrita, modificamos el código agregando un `if` para validar cuando no se reciben números en la cadena de entrada, también utilizamos el paquete [**_strconv_**](https://golang.org/pkg/strconv/) para convertir el número que se recibe en string a int.

```golang
import "strconv"

func Add(numbers string) int {
	if len(numbers) == 0 {
		return 0
	} else {
		number, _ := strconv.Atoi(numbers)
		return number
	}
}
```

Validamos con `go test` que la prueba sea exitosa (Green step)

```
PASS
ok      _.../code/day1        0.266s
```

¡A refactorizar de nuevo!

Por último agregaremos un nuevo escenario en nuetro test para validar cuando recibimos dos números separados por una coma.

```golang
t.Run("two values", func(t *testing.T) {
    //Arrange
    expected := 6
    //Act
    outcome := Add("4,2")
    //Assert
    if outcome != expected {
        t.Errorf("Test with one value failed")
    }
})
```

Validamos con `go test` que la prueba sea fallida nuevamente (Red step)

```
--- FAIL: TestAdd (0.00s)
    --- FAIL: TestAdd/two_values (0.00s)
        calculator_test.go:36: Test with one value failed
FAIL
exit status 1
```

Modificamos la función `Add` para agregar 
Ahora que tenemos nuevamente una prueba fallida bien escrita, modificamos el código para agregar la lógica de suma de dos números, podemos seguir utilizando el paquete **_strconv_**.

```golang
func Add(numbers string) int {
	if len(numbers) == 0 {
		return 0
	} else if len(numbers) == 1 {
		number, _ := strconv.Atoi(numbers)
		return number
	} else {
		firstNumber, _ := strconv.Atoi(numbers[:1])
		secondNumber, _ := strconv.Atoi(numbers[2:])
		return firstNumber + secondNumber
	}
}
```

Validamos con `go test` que la prueba sea exitosa nuevamente (Green step)

```
PASS
ok      _.../code/day1        0.186s
```

 ### Día 2

 Habilita la función Add para que pueda manejar una cantidad desconocidad de números.

**Solución:**

Escribiendo la prueba, importante tener en cuenta el patrón triple AAA:

```golang
t.Run("unknown values", func(t *testing.T) {
	//Arrange
	expected := 30
	//Act
	outcome := Add("1,6,7,9,3,4")
	//Assert
	if outcome != expected {
		t.Errorf("Test with unknown value failed")
	}
})
```
Validamos con `go test` que la prueba falle (Red step)

```
--- FAIL: TestAdd (0.00s)
    --- FAIL: TestAdd/unknown_values (0.00s)
        calculator_test.go:47: Test with unknown value failed
FAIL
exit status 1
FAIL    _.../code/day2        0.180s
```

Agregamos el código necesario para pasar la prueba

```golang
func Add(numbers string) int {
	if len(numbers) == 0 {
		return 0
	} else if len(numbers) == 1 {
		number, _ := strconv.Atoi(numbers)
		return number
	} else {
		splitNumbers := strings.Split(numbers, ",")
		result := 0
		for i := 0; i < len(splitNumbers); i++ {
			number, _ := strconv.Atoi(splitNumbers[i])
			result += number
		}
		return result
	}
}
```

validamos que la prueba sea exitosa (Green step)

```
PASS
ok      _.../code/day2        0.255s
```

¡A refactorizar!

En este momento no hay mucho que refactorizar, sin embargo, si vemos nuestro código ya no es necesaria la validación cuando recibimos un solo número, con el ciclo que acabamos de crear estamos cubriendo ese escenario, por esta razón podemos eliminarlo.

```golang
func Add(numbers string) int {
	if len(numbers) == 0 {
		return 0
	} else {
		splitNumbers := strings.Split(numbers, ",")
		result := 0
		for i := 0; i < len(splitNumbers); i++ {
			number, _ := strconv.Atoi(splitNumbers[i])
			result += number
		}
		return result
	}
}
```

Finalmente validamos nuestro set de pruebas unitarias, con `go test -v` podemos verificar la ejecución de cada uno de los escenarios planteados en nuestras pruebas unitarias.

```
=== RUN   TestAdd
=== RUN   TestAdd/empty_value
=== RUN   TestAdd/one_value
=== RUN   TestAdd/two_values
=== RUN   TestAdd/unknown_values
--- PASS: TestAdd (0.00s)
    --- PASS: TestAdd/empty_value (0.00s)
    --- PASS: TestAdd/one_value (0.00s)
    --- PASS: TestAdd/two_values (0.00s)
    --- PASS: TestAdd/unknown_values (0.00s)
PASS
ok      _.../code/day2        0.182s
```

### Día 3

Permite que el método Add maneje nuevas líneas entre números (en lugar de comas).
la siguiente entrada está bien: ```"1\n2,3"``` (será igual a 6)_
la siguiente entrada **NO** está bien: ```"1,\n"``` (no es necesario probarlo, solo aclarar)_

**Solución:**

Escribiendo la prueba, importante tener en cuenta el patrón triple AAA:

```golang
func TestAdd (t *testing.T) {
 //Arrange
 //Act
 //Assert
} 
```
Validamos que la prueba falle (Red step), y agregamos el código necesario para pasar la prueba

```golang
func Add(numbers: string): int {
    //TODO
}
```

validamos que la prueba sea exitosa (Green step) y empezamos a refactorizar!

```golang
func Add(numbers: string): int {
    //TODO
}
```

NOTA: faltan los días 4,5 a partir de allí se crea un servicio rest que exponga las funcionalidades de la calculadora. A partir de estos dos puntos tampoco sería necesario poner los pantallazos del red, green step
