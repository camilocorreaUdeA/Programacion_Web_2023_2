# Manos a la obra
A poner en práctica lo que has aprendido de desarrollo backend

## Ejercicio propuesto
Vamos a implementar una aplicación CRUD (Create, Read, Update, Delete) básica sobre una API REST (GET, POST, PUT/PATCH, DELETE).

### Estructura del proyecto
Para la implementación de esta aplicación vamos a definir una estructura de proyecto que nos permita dividir o separar los componentes de la aplicación de acuerdo con su funcionalidad específica. Dicho de otro modo, vamos a separar las responsabilidades de la aplicación, y cada una de estas responsabilidades tendrá su propio paquete dentro de la aplicación.
<ul>
  <li><b>Repositorio o capa de acceso a la base de datos:</b> Está responsabilidad ya está implementada y lista para ser utilizada. En la siguiente ruta https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/tree/main/Backend/practica/repository se encuentra el paquete que debe agregar a su proyecto; allí están las implementaciones de las distintas operaciones básicas sobre una base de datos SQL.</li><br>
  <li><b>Modelos de datos o capa de las entidades de la base de datos: En esta capa se definen los modelos de datos que se mapean a la base de datos, es decir los objetos que se guardan, se modifican y se consultan en la base de datos. En este paquete deben ir las estructuras de datos utilizadas tanto para presentar/recibir los datos al/desde front-end como para que la capa de controladores puedan convertirlos para la capa de repositorio.</b></li><br>  
  <li><b>Controladores o capa de lógica de la aplicación:</b> En este paquete se definen las funciones que implementan la lógica de cada uno de los endpoints de la API. Aquí se debe implementar una capa intermedia que comunica a las capas de repositorio y de handlers. Se encarga de traducir las solicitudes HTTP en operaciones sobre la base de datos y de encapsular los retornos de la base de datos al formato de respuesta de los endpoints.</li><br>  
  <li><b>Handlers o capa de funcionalidad de los endpoints:</b> En este paquete se deben implementar las funciones que van a atender las solicitudes hechas a la aplicación. Cada endpoint de la REST API debe estar asociado a una función manejadora o <i>handler</i> que se encarga de validar la solicitud y de enrutarla a la función correspondiente en la capa de controladores.</li>
</ul>

La estructura del proyecto debería ser similar a la que se presenta a continuación:

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/22d4d74b-d189-4746-bf05-1b7c2a7a1113)

### Requerimientos de la aplicación

<ol>
  <li>Escoja un objeto de datos para trabajar. Puede ser cualquiera de su preferencia, por ejemplo artículos de una tienda online, empleados, estudiantes, platillos de gastronomía, animales domésticos y/o salvajes, etc. Luego de escoger el objeto de datos debe definir los campos que lo van a constituir, por ejemplo, si escogió libros los campos deberían ser tal vez: título, autor, edición, título alternativo, país, año de publicación, arte de carátula/portada, premios, ¿es best seller?, idiomas en los que ha sido traducido, etc. Su objeto de datos debe contar por lo menos con 7 campos distintos (incluya campos de distintos tipos de dato).</li><br>
  <li>Implemente una base de datos SQL (preferiblemente en https://www.elephantsql.com/) para su objeto de datos. Debe crear una tabla con las columnas correspondientes a cada uno de los campos del objeto de datos. Agregue datos a la tabla para que la aplicación cuente con datos por defecto.</li><br>
<li>La aplicación debe implementar endpoints o rutas para las 4 operaciones básicas sobre la base de datos. Es decir, la aplicación debe permitir insertar nuevas filas a la tabla en la base de datos, actualizar una fila (indexada por un campo id), leer las columnas de una fila determinada por el campo id, leer un listado de columnas de la tabla (aplica paginación) y también eliminar filas de la tabla (indexando por el campo id)</li><br>
  <li>La aplicación debe estar soportada por una REST API. Debe utilizar los métodos HTTP para implementar las distintas operaciones sobre la base de datos y además debe utilizar el formato JSON para el encapsulamiento de los datos que ingresan/salen de la aplicación. Cada endpoint de la aplicación debe estar asociado a una operación sobre la base de datos.</li><br>
  <li>Los endpoints de la aplicación deben ser atendidos por <i>handlers</i> en los que se enruta la solicitud a la función de procesamiento correspondiente (controlador) y se confeccionan las respuestas de la API. Los <i>handlers</i> se encargan de hacer las conversiones entre los cuerpos de las solicitudes que están en formato JSON y el modelo de datos que es una estructura de Go y viceversa. También se encargan de extraer los parámetros que vengan en la URL de la solilcitud (id, limite, offset, etc.). Y cualquier tipo de validación adicional de la solicitud debe hacerse allí.</li><br>
  <li>En las funciones de la capa de controladores se debe ejecutar el procesamiento de las solicitudes. Se deben hacer las conversiones de datos, desde el modelo y hacia el formato en que los reciben las funciones del paquete de repositorio (mapas o parámetros de funciones). En los controladores se deben definir las consultas a la base de datos (<i>SQL queries</i>) en el lenguaje SQL. Y luego de completado el procesamiento de la solicitud debe retornar los resultados a la capa de <i>handlers</i> para que allí se confeccione la respuesta a la petición.</li><br>
  <li>Implemente una sencilla UI (interfaz de usuario) de prueba en la que pueda probar la funcionalidad de cada uno de los endpoints de la aplicación. Esta interfaz de usuario es una página web sencilla desarrollada con HTML, CSS y Javascript en la cual hay opciones disponibles para ejecutar una prueba de funcionamiento sobre cada uno de los endpoints de manera individual.</li>
</ol>

### Detalles de implementación

__Capa de modelo de datos__
<ol>
  <li>El modelo es un <code>struct</code> que debe contar con mínimo 7 campos. Cada campo corresponde a una columna de la tabla en la base de datos</li>
  <li>Debe incluir un campo adicional de tipo <code>int64</code> para almacenar un <code>id</code> que permitira identificar cada objeto de datos de los demás que están almacenados en la tabla de la base de datos. Este campo <code>id</code> se utiliza para leer y actualizar filas de la tabla de forma individual (dos de las operacinoes de la aplicación)</li>
  <li>Cada campo del modelo debe ir acompañado por un <code>tag</code> o etiqueta en la que se asocia el campo a una columna de la tabla de la base de datos, es decir, el nombre usado en el <code>tag</code> debe conincidir exactamente con el nombre de la columna en la tabla. El <code>tag</code> se construye de la siguiente manera: <code>`db:"nombre_columna"`</code></li>
 </ol>
 Su modelo de datos debería verse similar al del ejemplo a continuación:

 ![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/881646c9-c808-4873-a9e3-aa0640c4c045)

__Capa de controladores__
<ol>
  <li>El controlador es una estructura que tiene un campo que es un objeto del tipo de la interfaz <code>Repository</code> exportada por el paquete <i>repository</i> (https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/tree/main/Backend/practica/repository)</li>
  <li>El controlador debe tener asociados los métodos que implementan cada una de las operaciones CRUD (Create, Read, Update, Delete) sobre la base de datos.</li>
  <li>El controlador debe actuar como una capa mediadora entre los <i>handlers</i> y el <i>repositorio</i>, por tanto sus métodos deben convertir los datos recibidos en la solicitud y transformarlos acorde con el formato en que deben ser pasados a las funciones de la capa de repositorio.</li>
  <li>Los métodos del controlador deben utilizar el respectivo método de la capa de repositorio. Esos métodos deben ser invocados a través del objeto que está en el campo de la estructura controlador (objeto del tipo de la interfaz <code>Repository</code>)</li>
  <li>Métodos a implementar:
    <ul>
      <li>Create: Es un método que permite insertar una nueva fila (registro) a la tabla en la base de datos. Debe retornar el id (valor de la columna id) del nuevo registro y un valor de error paa los casos en que el método falla al insertar en la base de datos. Debe recibir como parámetro de entrada el objeto de la solicitud (http.Request) para de allí poder extraer el cuerpo que contiene los datos que se van a ingresar a la tabla.</li>
      <li>Read: Este método permite consultar un registro indexando por el valor de la columna id. Debe retornar un slice de bytes que contiene el objeto de datos retornado por la capa de repositorio codificado en JSON. También debe retornar un valor de error para propagar cualquier fallo que ocurra al tratar de leer la tabla.</li>
      <li>List: Es un método similar al anterior pero este permite leer varios registros consecutivos de la tabla. Desde la capa de handlers recibe como parámetros la cantidad de registros que desea leer (limit, page size, etc.) y a partir de que fila de la tabla quiere leer (offset). El método debe retornar los registros codificados en JSON (un slice de byte) y un valor de error por si ocurre alguna falla al consultar la tabla.</li>
      <li>Update: Con este método se actualiza un registro (fila) específico de la tabla (indexado por el valor de la columna id). Luego de decodificar el objeto de datos del cuerpo de la petición debe actualizar solamente los campos que contengan algún valor (valor distinto al zero-value del tipo de dato) para actualizar. Esté método debe retornar un valor de error en el que se indique si la operación falló o si por el contrario se ejecutó con éxito.</li>
      <li>Delete: Es el método utilizado para remover o borrar una fila de la tabla de la base de datos. Este método debe recibir el id de la fila que se quiere eliminar y debe retornar un valor de error para indicar el éxito o fallo de la operación.</li>            
    </ul>
    <li>NewController: Esta función NO debe ser un método del controlador. Es una función que permite crear una instancia del controlador, recibe como parámetro un objeto del tipo de la interfaz <code>Repository</code> exportada por el paquete <i>repository</i> que se utiliza para incializar el campo presente en la estructura controlador. Esta función retorna la instancia con el campo inicializado</li>
  </li>
  <li>En este paquete también se definen las consultas (queries) que utilizaran los métodos contra la base de datos a través de la capa de repositorio.</li>
</ol>

__Capa de handlers__
<ol>
  <li>Esta capa implementa los métodos que atienden directamente las rutas o <i>endpoints</i> expuestos desde la aplicación hacia el exterior. Son métodos de una estructura <i>Handler</i></li>
  <li>Implemente un método <i>handler</i> por cada uno de los <i>endpoints</i> de la aplicación. Cada <i>endpoint</i> debe cubrir una operación en base de datos (CRUD)</li>
  <li>Los <i>handlers</i> reciben como parámetros de entrada un objeto del tipo <code>http.ResponseWriter</code> y otro objeto del tipo puntero <i>http.Request</i></li>
  <li>Los <i>handlers</i> se encargan de pasar los parámetros de la solicitud (y el cuerpo de la solicitud cuando aplica) a los métodos de la capa de controladores. Para el caso de operaciones como lectura, lista, actualización y eliminación deben pasar un parámetro <code>id</code> pasado en la URL del <i>endpoint</i>. Para las operaciones de creación y actualización deben pasar el cuerpo de la solicitud que es donde vienen los datos para almacenar en la base de datos.</li>
  <li>Los <i>handlers</i> deberían implementar unas validaciones básicas a los parámetros de la URL. Por ejemplo: que el <i>id</i> sea efectivamente un número (no contenga letras ni caracteres especiales), que no sea un número negativo ni cero, etc. Si alguna validación falla se debe retornar un status <i>Bad Request</i> como respuesta a la solicitud. (Esas validadciones aplican también para los parámetros para la cantidad de datos y el offset del método List)</li>
  <li>Cuando la operación se completa con éxito el <i>handler</i> debe retornar un status <i>OK</i> a excepción de cuando se reponde a una solicitud de tipo <i>POST</i> para la que se debe responder con el status <i>Created</i>.</li>
  <li>Para las operaciones Read y List se debe retornar el cuerpo de la respuesta a la solicitud. Este cuerpo es recibido desde la capa de controladores (slice de bytes).</li>
  <li>Los <i>handlers</i> pueden implementarse como <code>http.HandlerFunc</code> o bien como <code>http.Handler</code></li>
</ol>

__Main__

El proyecto debe contar con un archivo fuente <code>main.go</code> en el que se implementa y se ejecuta un servidor web (en una función main). Los <i>endpoints</i> y sus respectivos <i>handlers</i> deben asociarse al servidor.

En este archivo se deben crear las instancias de la interfaz <code>Repository</code>, del objeto de la estructura del controlador y del objeto de la estructura de los handlers.

Y también es en este archivo que se definen las rutas (URLs) de los <i>endpoints</i>.

### Información adicional
<ul>
  <li>Puede utilizar los paquetes externos que considere convenientes. (Routers, multiplexers, frameworks web, etc.)</li>
  <li>Hablando de paquetes externos, para poder utilizar el paquete <i>repository</i> debe instalar en su proyecto un driver para SQL que puede encontrar en este repositorio: https://github.com/jmoiron/sqlx</li>
</ul>




