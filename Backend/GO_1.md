# Go
![2](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/e131e81c-9a0d-4db2-8d40-99119ebe006e)
## El lenguaje de programación Go (Golang)

Go inició como un proyecto de bajo perfil que rapidamente escalo hasta convertirse en un proyecto ambicioso patrocinado por Google. A finales de 2011 se liberó Go como un proyecto de código abierto y luego el 28 de Marzo de 2012 se lanzó la primera versión estable del lenguaje y a partir de ese momento se procura liberar una nueva versión cada medio año.

"...El objetivo del proyecto Go era eliminar la lentitud y la falta de coordinación en el proceso de desarrollo de software en Google, y por tanto hacer el proceso más productivo y escalable. El lenguaje fue diseñado por y para gente que escribe, lee, depura y mantiene grandes proyectos de software..." -Rob Pike-

Al crear el lenguaje de programación Go la idea principal era combinar las mejores características que están presentes en otros lenguajes, entre ellas:
<ul>
  <li>Fácil de integrar con las herramientas actuales para disparar la productividad</li>
  <li>Alta eficiencia y tipado estático</li>
  <li>Alto desempeño para aplicaciones de networking y un aprovechamiento total de las arquitecturas multi-core</li>
</ul>

Go es un lenguaje especialmente apto para las siguientes aplicaciones y sistemas:
<ul>
  <li>Servicios distribuidos interconectados: Las aplicaciones de red dependen fuertemente de la concurrencia, la cual es manejada de forma simple y eficiente por característcas nativas del lenguaje como los canales y las go-rutinas</li>
  <li>Desarrollo nativo en la nube: Las apliaciones nativas de la nube son más fáciles de construir con las características que Go ofrece para concurrencia y networking, además de la portabilidad entre sistemas operativos y arquitecturas que ofrece de forma nativa el lenguaje.</li>
  <li>Remplazo para infraestructura existente: Go es un lenguaje ideal para re-escribir aplicaciones, permite mejorar la seguridad en el manejo de memoria, permite el despliegue multi-plataforma y una base de código limpia que es de más fácil mantenimiento en el futuro.</li>
  <li>Utilidades y herramientas independientes: Los programas en Go complilan en binarios que practicamente no requieren de dependencias externas. Esto hace a Go ideal para la creación de herramientas y utilidades que se pueden poner en producción con prontitud y que pueden ser empaquetadas para su redistribución.</li>
</ul>

Tal vez Go no es lo sufiecientemente apto en los siguientes casos:
<ul>
  <li>Interfaces gráficas o GUIs (aunque hay paquetes como Fyne y Wails que tienen un desempeño de aceptable a bueno)</li>
  <li>Desarrollo de sistemas embebidos y desarrollo de bajo nivel en general: Existe un branch de Go conocido como TinyGO que permite programar dispositivos de hardware. Go no está diseñado para la creación de drivers o interfaces de propósito general de entrada/salida de datos (comunicaión con periféricos y hardware externo)</li>
</ul>

Ventajas de Go con respecto a otros lenguajes de programación:
<ul>
<li>Fácil de usar y de leer: La sintaxis de Go es simple, entre 25 y 35 palabras clave conforman el lenguaje. Y además su curva de aprendizaje lo hace bastante accesible a programadores novatos.</li>
<li>Amplia librería estándar: Una amplia y variada librería estándar viene empaquetada en el lenguaje, lo cual evita el uso de librerías de terceros para similares propósitos. La librería es sofisticada pero no compleja ni complicada, y el manejo de dependencias de Go evita riesgo de conflictos entre nombres de fiunciones.</li>
<li>Seguridad: Gracias a lo simple del código, a que es un lenguaje de tipado estático y al garbage collector incluido en el runtime del lenguaje.</li>
<li>Soporte: Además del soporte de ingenieros de Google cuenta con una siempre creciente coumunidad.</li>
<li>Concurrencia y paralelismo: Incluidos en el lenguaje en librerías con funciones para estos propósitos.</li>
<li>Manejo de dependencias: El sistema de manejo de dependencias está incluido en el lenguaje facilitando la inclusión de paquetes de terceros y el posterior empaquetamiento en el binario de la aplicación.</li>  
</ul>

Utilidades de Go para los desarrolladores:
<ul>
  <li>Generación automática de documentacipon del código.</li>
  <li>Herramientas de testing y benchmarking nativas incluídas en el lenguaje.</li>
  <li>Detección de condiciones de carrera para apliaciones concurrentes.</li>
  <li>Un sistema de tipos liviano pero suficiente para clasificar y diferenciar datos y objetos.</li>
  <li>Concurrencia nativa para la ejecución de tareas en segundo plano.</li>
  <li>Garbage collector para el uso racional de memoria.</li>
  <li>Compilación multiplaforma que permite desplegar en cualquier arquitectura y sistema operativo.</li>
  <li>Buen desempeño y velocidad en comparación con otros lenguajes de programación similares.</li>
  <li>Amplia librería estándar, especializada en funciones de concurrencia y networking.</li>
  <li>Punteros para manejo eficiente de memoria.</li>
</ul>

## Instalación de Go

Puedes revisar el [Getting started](https://go.dev/doc/tutorial/getting-started) oficial.
También puedes seguir las instrucciones oficiales de [instalación](https://go.dev/doc/install)

En general el proceso de instalación de Go es bastante intuitivo:
<ul>
  <li>Instala el IDE de tu preferencia, recomiendo VSCode pero Goland también está bien.</li>
  <li>Descarga Go en https://go.dev/dl/. En ese sitio están los instaladores para los distintos sistemas operativos.</li>
  <li>Instalación manual de Go en tu máquina:
    <ul>
      <li>Linux: Ejecute en la terminal: <code>rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.3.linux-amd64.tar.gz</code>, luego setea la variable de ambiente <code>PATH</code> para que también apunte al directorio bin de Go: <code>export PATH=$PATH:/usr/local/go/bin</code></li>
      <li>Con MacOS y Windows descargue el instalador correspondiente, ejecutelo y verifique la instalación con el comando <code>go version</code> en la terminal.</li>
    </ul>
  </li>
  <li>Otras formas de instalar Go:
  <ul>
    <li>MacOS: Instale <b>homebrew</b> en su sistema y una vez instalado ejecute en la terminal el comando <code>brew install go</code></li>
    <li>Windows: Instale <b>chocolatey</b> y una vez instalado ejecute en la terminal el comando <code>choco install golang</code></li>
  </ul>
    </li>
  <li>Verifique la instalación ejecutando el comando <code>go version</code></li>
</ul>

Una vez hayas instalado Go en tu máquina y si ya tienes VSCode instalado también (¡si no, entonces instalalo!). Abre VSCode y luego presiona la combinación de teclas <code>Ctrl + Shif + P</code> y en el menú que se despliega escoge la opción <code>Go: Install/Update Tools</code>. Espere que se complete la instalación de las herramientas.

<img width="603" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/8a9b7c5d-b8b5-4dd4-82a0-c7926d2bd2c0">

Luego instale la extensión de Go para facilitar la codificación ya que trae funciones como autocompletado, indentación automática y verificación de errores de sintaxis, etc.

<img width="655" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/aa49bf57-b453-45e1-a61f-adb722405ca3">

## Go modules

<i>Go modules</i> es el sistema de administración de dependencias en Go, el cual facilita el acceso a la información de versiones y la inclusión de paquetes externos en una aplicación de Go.

Un modulo en Go es una colección de paquetes almacenados en un árbol de archivos (directorio) donde un archivo llamado <i>go.mod</i> es el nodo raíz del árbol.

El archivo <i>go.mod</i> define la ruta del modulo que se usa como la ruta de <i>import</i> del directorio raíz. También indica la versión de Go usada para compilar el proyecto y el conjunto de dependencias (otros modulos) que se importan y se utilizan en el proyecto.

## Creación de un proyecto en Go

Para crear un proyecto de Go:
<ol>
  <li>Crear un directorio para alojar el proyecto.</li>
  <li>Acceda al directorio con VSCode.</li>
  <li>Abra una Terminal (Terminal -> New Terminal)</li>
  <li>En la terminal ejecute el siguiente comando <code>go mod init &lt;<i>ruta-al-modulo</i>&gt;</code> Donde <i>ruta-al-modulo</i> es el nombre del modulo del proyecto, por convención de acostumbra que esta ruta sea la misma URL al repositorio en github.com donde se alojará el proyecto, por ejemplo: <i>github.com/usuario123/repo-proyecto</i></li>

  ![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/3cd3a8b7-cafd-431f-a3f9-7dbebdc04846)
  
  <li>Verifique que un archivo <i>go.mod</i> haya sido agregado al directorio raíz del proyecto.</li>  
</ol>

## Anatomía básica de un archivo fuente de Go (archivos con extensión <i>.go</i>)

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/196c1dd0-032a-4e7b-9cda-def7e0ca9d15)

<ul>
  <li>Línea package &lt;<i>nombre-package</i>&gt;: Esta línea es la única parte obligatoria de la anatomía y debe ir en todos los archivos fuente de una aplicación de Go, ella indica a que paquete (package) pertenece el archivo fuente. Si el archivo fuente contiene una función <i>main</i> (<i>entry point</i> a la aplicación) entonces el nombre del paquete debe ser <i>main</i>, y por tanto la línea debe quedar <code>package main</code></li>
  <li>Conjunto de <i>imports</i>: En esta sección se hace referencia a los paquetes de los cuales se importan variables, tipos o funciones en el archivo de fuente actual. Se puede importar tanto paquetes locales (es decir, paquetes que están dentro del mismo proyecto) o paquetes externos (son paquetes creados por terceros que se agregan como modulos al proyecto) que se descargan direactmente del repositorio donde están alojados. La sintaxis incluye la palabra <i>import</i> seguida por el conjunto de los nombres y/o las rutas a los paquetes que se van a importar, encerrados entre paréntesis. Dichos paréntesis no son necesarios cuando es un solo paquete el que se importa.</li>
  <li>Declaraciones de variables, constantes y tipos: Por convención se sugiere hacer declaraciones de variables y constantes y definiciones de tipos justo después de la sección de imports y antes de la definición de cualquier función.</li>
  <li>Definición de funciones (lógica): A continuación de las declaraciones de variables y constantes y de las definiciones de tipos se codifican las funciones que hacen parte del archivo fuente. Si el paquete es main, entonces se sugiere que la función main se defina antes de cualquier otra función.</li>
</ul>

## Jerarquía dentro de un módulo en Go

Hasta el momento hemos hablado de los conceptos de módulo, paquete y archivo fuente pero sin dar una definción precisa de cada uno y sin explicar cuál sería la relación entre estos conceptos, pues bien, a continuación dejamos las definiciones pertinentes:
<ul>
  <li><b>Archivo fuente:</b> Es un archivo amoldado por el patrón de anatomía básica que vimos anteriormente. Todos los archivos fuente en Go tienen la extensión .go y son simplemente el lugar en el que se agrupan un conjunto de funciones, declaraciones de variables y definiciones de tipos que están inter-relacionados y que tienen el objetivo de implementar unas funcionalidades que componen una lógica de negocio o algoritmo en específico.</li>
  <li><b>Paquete (package):</b> Un paquete es una agrupación de archivos fuente dentro de un mismo directorio, dichos archivos deben compartir el mismo nombre de paquete en la línea package. Un paquete permite exportar a otros paquetes (o módulos) las variables, constantes, tipos y funciones definidos en los archivos fuente que lo conforman.</li>
  <li><b>Módulo:</b> Es un conjunto o colección formado por uno o más paquetes almacenados en un árbol de archivos y en cuyo directorio raíz existe un archivo go.mod que contiene el nombre del módulo (ruta al repositorio), la versión del lenguaje con la que fue creado y las posibles dependencias que pueda tener con otros módulos externos. Un módulo en Go es en cierto grado lo que comunmente se denomina librería o biblioteca en otros lenguajes de programación.</li>
</ul>

## Tipos de datos en Go, declaración de variables, constantes, funciones y definición de tipos

### Tipos nativos de Go

Numéricos

```go
/* enteros */
uint8 (uint, byte), uint16, uint32, uint64
int8 (int, rune), int16, int32, int64

/* decimales */
float32, float64

/* complejos */
complex64, complex128
```
Booleano

```go
bool
```
Texto y cadenas de caracteres

```go
string
```
Punteros: Los punteros son variables que permiten almacenar direcciones de memoria o referencias de otras variables (no punteros) del mismo tipo. Los punteros cuentan con los operadores <code>*</code> y <code>&</code> para averiguar por el valor almacenado en la referencia y para acceder a una referencia.

```go
/* puntero a una variable de tipo uint64 */
var ptr *uint64
```

Colecciones de datos

Arrays: Son colecciones de datos de un mismo tipo que se almacenan de forma consecutiva en memoria y para las que se define el tamaño en el momento mismo de la declaración. Un tipo array está compuesto tanto por el tipo de dato de los elementos como por la longitud de arreglo mismo. Un array es estático en memoria, esto quiere decir que no pueden cambiar de tamaño en tiempo de ejecución, conservan el tamaño asignado en la declaración durante el ciclo de vida de la aplicación.

```go
/* arreglo de 10 elementos de tipo int */
var array [10]int
```
Slices: Son colecciones similares a los arrays pero con la diferencia de que son dinámicos en memoria, es decir, pueden cambiar de tamaño en tiempo de ejecución, se les puede retirar o agregar elementos. Hay que tener cuidado cuando al agregar elementos se sobrepasa la capacidad del slice ya que esto implica una relocalización del slice en memoria, por tanto invalidando los punteros y copias del slice.

```go
/* slice */
var slice []string
```
Mapas: Un mapa en Go es una colección no ordenada de elementos de cierto tipo a los que se puede acceder o indexar mediante un conjunto de claves que también son de algún tipo. Son similares a los mapas o diccionarios de otros lenguajes de programación, es decir, son colecciones de pares clave-valor, donde el par puede ser del mismo tipo o de tipos distintos, y una clave es única e irrepetible en el mapa. Las claves pueden ser de cualquier tipo a excepción de los tipos array, slice, mapa o función.

```go
/* mapa cuyas claves son de tipo string y sus valores de tipo int */
var mapa map[string]int
```
Funciones, interfaces y estructuras

Funciones: Las funciones en Go son tratadas como tipos de datos. Un tipo función describe el conjunto de funciones que comparten el mismo conjunto de parámetros de entrada (cantidad y tipos de datos) y tipos de retorno.

```go
/* variable de tipo funcion para almacenar funciones que reciben un dato string y retornan un slice de int64 y un error */
var fun func(string)([]int64, error)
```
Interfaces: Las interfaces son contratos que definen un conjunto de métodos (funciones) que deden ser implementados por cualquier tipo que quiera cumplir el contrato de la interfaz (implementar la interfaz). Un tipo que implementa una interfaz automáticamente adquiere el tipo definido por la interfaz. Por lo tanto, una variable del tipo de una interfaz es capaz de almacenar un valor de cualquier otro tipo que implemente la interfaz (que implemente los métodos definidos en la interfaz). Todos los tipos en Go implementan una interfaz en común que es la interfaz vacía (empty interface) definida como <code>interface{}</code> y también con el alias <code>any</code>.

```go
/* una interfaz que define unos métodos para ser implementados por otros tipos */
type MiInterfaz interface {
   MetodoUno()
   MetodoDos(string)error
}
```
Estructuras: Las estructuras son secuencias de datos de distintos tipos (inclusive de otros tipos de estructuras o de interfaces), llamados campos de la estructura. Cada campo tiene un nombre y un tipo. Las estructuras permiten que se les asocien funciones, conocidas como métodos de la estructura, lo cual las hace similares a las clases de otros lenguajes de programación. Se puede acceder a los campos y métodos de una estructura mediante una variable o instancia del tipo de la estructura. Se puede decir que las estructuras son tipos personalizados por el usuario en Go.

```go
/* una estructura en Go con sus campos */
type MiEstructura struct {
   Nombre          string
   Edad            uint
   Email           string
   Caracteristicas []string
}
```
### Declaración de variables en Go

Variables globales de un paquete

Las variables globales de un paquete son aquellas visibles desde cualquier función definida en el paquete. Siempre se declaran por fuera de cualquier función y en la sección del ar archivo fuente dedicado a la declaración de variables y definición de tipos y funciones.

Las variables globales se declaran siempre con la palabra clave <code>var</code> seguida del nombre de la variable y opcionalmentedel tipo de la variable o de una inicialización dejando al compilador la tarea de deducir el tipo.

```go
/* distintas formas de declarar variables globales en Go */
var cadenaDeTexto string /* indicando el tipo y sin inicialización explícita, el valor inicial de cadenaDeTexto será "" */
var contadorDeEventos uint = 10 /* indicando el tipo e inicializando de manera explícita */
var dataSource = "digital" /* sin indicar el tipo y dejando su deducción al compilador a partir del valor asignado */
var x, y, z float32 /* declarando un conjunto de variables del mismo tipo */
var a, b, c = false, 100, "hello"  /* declarando un conjunto de variables de distinto tipo */
/* declarando un conjunto de variables de distintos tipos */
var (
  isSafe = true
  numArray []float64
  times int32 = 1
)
```
Variables locales de una función

Las variables locales son aquellas que se declaran al interior de las funciones y solo tienen validez dentro de la función en la que se declaran (scope). Se pueden declarar de la misma manera que se declaran las variables globales y adicionalmente se pueden declarar usando el operador walrus dejando que el compilador deduzca el tipo.

```go
func myFunc() {
  myLocalVar := "una_var_local"
}
```
Constantes

Las constantes son literales asociados a un valor que se conserva durante toda la ejecución de la aplicación. Las constantes pueden ser globales o locales según se requiera. Y para su declaración aplican las mismas reglas que para la declaración de variables globales y locales, a excepción de que se usa siempre la palabra clave <code>const</code>, siempre se deben inicializar en el momento de su declaración y no se pueden declarar utilizando el operador <code>:=</code> (operador walrus).

Nota: No se pueden declarar constantes de los tipos array, slice, map, struct ni interface.

```go
/* distintas formas de declarar constantes en Go */
const dataCode uint = 111 /* indicando el tipo */
const newLabelCode = "tag_label"  /* permitiendo deducción del tipo de la constante */
const x, y, z = 1, 2, 3 /* declarando un conjunto de constantes del mismo tipo */
/* declarando un conjunto de constantes de distintos tipos */
const (
  isAnalog = false
  numPossibleDataEvents = 15
  DigitalDatasource = "generic_digital"
)
```
Declaración de punteros

Para declarar un puntero se utiliza el operador <code>*</code> y para almacenar una referencia en el puntero se utiliza el operador <code>&</code>. No se pueden declarar punteros constantes ni punteros a constantes. La aritmética de punteros no existe en Go. Pueden declararse tanto punteros globales como locales y por lo tanto también aplica el operador walrus y la deducción de tipo para punteros.

Nota: Existe el operador <code>new</code> para inicializar un puntero al momento de su declaración. Inicializar se entiende como asignar un valor distinto a nulo (<code>nil</code>) al puntero.

```go
/* distintas formas de declarar punteros en Go */
var ptrOne *int /* puntero a variables de tipo int */
var ptrTwo = &someVariable  /* deducción del tipo puntero por inicialización */
var ptrThree = new(MyStruct) /* inicialización del puntero a instancias del tipo MyStruct con el operador new */

func myFunc() {
  myLocalVar := 50
  myLocalPtr := &myLocalVar /* deducción del tipo de puntero a *int */
}
```
Declaración de arrays, slices y maps



Definición e instanciación de structs

Definición de funciones

Definición de interfaces






