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
func Add(numbers: string): int
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
func TestAdd (t *testing.T) {
 //Arrange
 //Act
 //Assert
} 
```
Validamos que la prueba falle (Red step):

```sh
$ go test
```
![pantallazo-dia1-red-step](./doc/assets/pantallazo-dia1-red-step.png)

Agregamos el código necesario para pasar la prueba

```golang
func Add(numbers: string): int {
    //TODO
}
```

validamos que la prueba sea exitosa (Green step)

```sh
$ go test
```
![pantallazo-dia1-green-step](./doc/assets/pantallazo-dia1-green-step.png)

¡A refactorizar!

```golang
func Add(numbers: string): int {
    //TODO
}
```

 ### Día 2

 Habilita la función Add para que pueda manejar una cantidad desconocidad de números.

**Solución:**

Escribiendo la prueba, importante tener en cuenta el patrón triple AAA:

```golang
func TestAdd (t *testing.T) {
 //Arrange
 //Act
 //Assert
} 
```
Validamos que la prueba falle (Red step):

```sh
$ go test
```
![pantallazo-dia2-red-step](./doc/assets/pantallazo-dia2-red-step.png)

Agregamos el código necesario para pasar la prueba

```golang
func Add(numbers: string): int {
    //TODO
}
```

validamos que la prueba sea exitosa (Green step)

```sh
$ go test
```
![pantallazo-dia2-green-step](./doc/assets/pantallazo-dia2-green-step.png)

¡A refactorizar!

```golang
func Add(numbers: string): int {
    //TODO
}
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
