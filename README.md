# spike-testing-golang
En este repo se hace una revisión de los metodos, recomendaciones y buenas practicas para realizar pruebas automatizadas a código golang. Es un análisis exploratorio con el fin de tener un insumo para aterrizar los lineamientos de golang.

Por: 
- @gioguzman-adl
- @alejandro56664-adl

## Contenido

- [1. Objetivos](#1-objetivos)
- [2. Introducción](#2-introducción)
- [3. Buenas prácticas](#3-buenas-prácticas)
- [4. Bibliotecas disponibles](#4-bibliotecas-disponibles)
- [5. Tutorial](#5-tutorial)
- [6. Conclusiones](#6-conclusiones)
- [7. Fuentes](#7-fuentes)

## 1. Objetivos

- Realizar un análisis exploratorio sobre las buenas practicas para pruebas automáticas documentadas y reportadas por la comunidad de Go.
- Construir una tabla en donde se sintetice las mejores practicas encontradas que puedan servir de base para el trabajo dentro de las mesas de trabajo.
- Presentar una primera aproximación a las buenas prácticas definidas en este documento para su rapida evaluación a través de un tutorial de 30 minutos.

## 2. Introducción

Este trabajo se realiza con el fin de sentar las bases o al menos realizar una primera aproximación hacia la consolidación de una guía de estilo y buenas prácticas para el desarrollo de pruebas automatizadas en golang dentro de las mesas de trabajo de ADL. Para ello se realiza un análisis exploratorio en la red de las guías de estilo publicadas por varias compañías que actualmente usan golang para sus desarrollos, esto con el fin de partir sobre una base solida. Se exploran principios, recomendaciones, patrones, antipatrones, herramientas, bibliotecas y demas que puedan servir a los desarrolladores a realizar su trabajo de una manera mas rápida y efectiva. Vale la pena aclarar que el alcance para este trabajo se reduce a la construcción de pruebas automáticas, también se excluye modelos de programación como TDD (_Test Driven Development_) o BDD (_Behavior Driven Development_) ya que suponen herramientas de un nivel mas alto y pueden aplicarse independientemente del lenguaje de programación, aúnque no excluye que las herramientas y consejos detallados mas adelante estén relacionados con estas prácticas. Para concluir se presenta de manera sintentizada la información encontrada para la evaluación del lector, también se presenta un tutorial que sirve como punto de partida para la asimilación de los conceptos tratados a lo largo del documento.

Es importante aclarar al lector que suponemos que tiene unas bases minimas de programación, por lo que no nos dentendremos a explicar algunos conceptos, sin embargo si considera que es necesario aclarar algún tema o justificar mejor lo descrito en este documento, sientase libre de comunicarnoslo, estaremos muy felices de leer sus opiniones y recomendaciones.

## 3. Buenas prácticas

Generalmente cuando se esta estudiando un nuevo lenguaje de programación no basta con conocer la sintaxis, semantica o incluso la filosofía de diseño y desarrollo que este nos propone, también se hace necesario conocer las recomendaciones que naturalmente surgen dentro de las comunidades alrededor del lenguaje para escribir código fuente que sea facil de entender y compartir dentro de la comunidad, esto es importante porque pasamos mas tiempo leyendo código que escribiendolo (*ref*) y generalmente estas recomendaciones nos ayudan a realizar nuestro trabajo de una manera mas eficaz. 


### 3.1 Principios

Búscar recomendaciones para la escritura de pruebas a nivel general.

_Dependencias_

En general siempre es buena práctica mantener las dependencias al minimo. La introducción de una nueva dependencia debe argumentarse y tener en cuenta tanto la compatibilidad de licencias, tener un analisis de seguridad para evitar inyectar vulnerabilidades. También dependiendo del tipo de solución (Ejemplo un microservicio o función _serverless_) puede tener muchas dependencias a bibliotecas externas puede tener un impacto considerable en el rendimiento (*ref*).

_Salida de las pruebas con información útil_

Es importante garantizar la mantenibilidad de nuestro código a través de pruebas que provean una retroalimentación adecuada cuando estas fallan. Esto puede reducir tiempos de revisón en el código tratando de entender la razón por la cuál una prueba falló. Siempre debe procurarse por presentar información clara y concisa sobre la(s) causa(s) que generarón el fallo, enfocandose siempre en presentar la causa raíz. Esto no se logra de la noche a la mañana y no existe una receta mágica para lograrlo, sin embargo es valiosa la retroalimentación que puedan brindar los pares cuando se realizan las sesiones de revisión de código.

_Código limpio_

Aunque las pruebas no hacen parte explicitamente de lo que se considera códgio de producción, es importante mantener la legibilidad del código de pruebas ya que hacen parte sustancial del proceso del software. (_#TODO complementar_)

_Principios F.I.R.S.T_

- **F**ast
- **I**ndependent
- **R**epeteable:
- **S**elf-validating
- **T**imely

_Enfocarse en pruebas que agreguen valor_

El tiempo y la energía de los desarrolladores son recursos limitados dentro de las mesas de trabajo por lo tanto es fundamental enfocarse en aquellas pruebas que agreguen valor al negocio:

- pruebas que validen reglas de negocio.
- Se puede comenzar por la funcionalidad principal y así sucesivamente ir cubriendo funcionalidades

### 3.2 Anatómia de una prueba en Go

En esta sección se va revisar la estructura básica de una prueba en golang. Como norma general se crean un archivo de pruebas con el nombre de la "cosa" que se va a probar seguido de el texto '\_test'. Ejemplo, supongamos que tenemos una entidad llamada  _string_ y definimos su comportamiento en un archivo 'string.go', tendrá por tanto su correspondiente 'string_test.go'.


Ejemplo de una prueba en go:
```golang
//adder_test.go
package main

import "testing"

ffunc TestWallet(t *testing.T) {

    //Arrange
    wallet := Wallet{}
    want := 10
    
    //Act
    wallet.Deposit(10)
    
    got := wallet.Balance()
    
    //Assert
    if got != want {
        t.Errorf("got %d want %d", got, want)
    }
}
```

En esta porción de código se puede observar lo siguiente:

- Se importa el módulo "testing". 
- El método de prueba es de acceso público (las primera letra en mayúscula lo indica mas info [aquí]() e inicia con la palabra _Test_.
- Recibe por parámetro un puntero a un objeto tipo '_testing.T_' el cuál sirve de _hook_ para tener acceso a diferentes útilidades de pruebas. En este caso para acceder al método '_Errorf_' que permite imprimir un mensaje de error en pantalla indicando la razón del fallo de la prueba.
- La prueba tiene una estructura _Preparar, Actuar y Afirmar_ que detallará en la siguiente sección.
__Nota__: como se describía al principio de esta sección, los archivos con código de pruebas y código de producción generalmente se encuentran en la misma carpeta, es importante que ambos archivos pertenezcan al mismo paquete y que al estructurar su código tenga cuidado de poner __archivos de código del mismo paquete en la misma carpeta_. (*ref*)

### 3.3 Patrón AAA

Este patrón aunque mas conocido del mundo Java, permite estructurar las pruebas una forma clara y concisa. Divide el código de la prueba en tres partes:

- _Arrange_ (Preparación): en esta etapa se prepara todo lo que se requiere para la ejecucíon de la prueba, configurar y disponibilizar mocks, stubs, datos de prueba, etc.

- _Act_ (Actuar): una vez se tiene todo listo para realizar la prueba se ejecuta el método o función que se va a probar. 

- _Assert_ (Afirmar): Cómo ya se ejecutó la porción de código que se desea probar, se procede a la validación de las afirmaciones necesarias para comprobar el correcto funcionamiento de la prueba. En esta etapa se debe prestar particular atención en la claridad de los mensajes que se le va a presentar al desarrollador.

Ejemplo sin patrón AAA

```golang

```

Con patrón AAA

```golang

```

También es importante tener en cuenta la repetibilidad de las pruebas (Principio FIRST)

No repetible

```golang

```

Repetible

```golang

```

### 3.4 Pruebas de tipo Benchmark

Otra de las caracteristicas que incluye de manera nativa Golang son las pruebas de desempeño tipo _Benchmark_ que permite tener una idea del desempeño que tiene el código escrito y poder identificar en etapas tempranas de desarrollo posibles puntos de contención que dificulten el escalamiento a futuro.

```golang
package main

import "testing"

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
```

TODO: Profundizar mas sobre este tema.

### 3.4 Usando _Subtests_ y _Sub-benchmarks_

https://blog.golang.org/subtests

### 3.5 _Table Driven Tests_

Para la realización de pruebas en Golang usando el patrón _Table Driven Test_ es necesario resaltar algunas funciones que pueden ser bastante útiles a la hora de crear pruebas con las herramientas nativas del lenguaje.

- t.Run
- t.Help

Ejemplo:

```golang

```


https://github.com/quii/learn-go-with-tests/blob/main/structs-methods-and-interfaces.md

https://github.com/golang/go/wiki/TableDrivenTests


### 3.6 _Mocking_

Hacer una pequeña introducción a la inyección de dependencias en go
https://github.com/quii/learn-go-with-tests/blob/main/mocking.md

### 3.6 Pruebas y concurrencia

https://github.com/quii/learn-go-with-tests/blob/main/concurrency.md

TODO: Describir las buenas practicas recomendadas por la comunidad para realizar pruebas automatizadas en lenguaje go.

## 4. Bibliotecas disponibles

TODO: incluir tabla comparativa donde se muestren las herramientas mas utilizadas ventajas/desventejas.

| *Nombre* | *Descripción* | *Ventaja* | *Desventajas*|
|--|--|--|--|
| [Testify](https://github.com/stretchr/testify) | TODO | TODO | TODO |
| [HttpExpect](https://github.com/gavv/httpexpect) | TODO | TODO | TODO |

## 5. Tutorial

TODO: Tutorial donde se pone en practica un ejercicio de TDD usando Golang donde se demuestra el uso de las buenas practicas y las bibliotecas presentadas.

## 6. Conclusiones

## 7. Fuentes

- https://github.com/bahlo/go-styleguide#testing
- https://wiki.crdb.io/wiki/spaces/CRDB/pages/181371303/Go+coding+guidelines
- https://github.com/quii/learn-go-with-tests

- Advanced testing in go https://youtu.be/S1O0XI0scOM
- Advanced testing in golang https://youtu.be/he_q_-Ih37U
- Automatic Mockfiles generation https://youtu.be/l6q2xa2fEjg



