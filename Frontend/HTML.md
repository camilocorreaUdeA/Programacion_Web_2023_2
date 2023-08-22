# HTML

## Introducción

<p>El <i>front-end</i> o capa de presentación de una aplicación web hace referencia al software que se desarrolla usando el conjunto de tecnologías <b>HTML</b>, <b>CSS</b> y <b>Javascript</b>. Con <b>HTML</b> se define la estructura de la presentación del contenido de la página web, <b>CSS</b> permite dar estilo y personalizar la parte visual de la página, mientras que <b>Javascript</b> permite agregar lógica y funcionalidad dinámica.</p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/ba588505-7830-4823-b38d-ff8a7ba9c4be)

<p>A partir de esta presentación y las siguientes vamos a enfocarnos en aprender HTML, <em>HyperText Markup Language</em> (Lenguaje de Marcado de Hipertexto), la anatomía general de una página web, el conjunto de etiquetas que componen el lenguaje HTML y cómo utilizarlas para estructurar el contenido de un sitio web.</p>

<p>HTML no es un lenguaje de programación ya que no permite codificar lógica de programación ni describir algoritmos o funcionalidades, por el contrario HTML solo se encarga de definir la forma en que se presentan los contenidos de un sitio web, y para ese fin utiliza un conjunto de etiquetas que permiten marcar o señalizar el contenido del sitio web para que el navegador sepa cómo mostrarlo en pantalla.</p>

<p>Cuando nos referimos a la forma de estructurar el contenido estamos hablando de muchas cosas: dividir el texto en párrafos, agregar estilo al texto (cursiva, negrilla, resaltado, etc.), definir hipervínculos, agregar imágenes, videos y animaciones, definir listas, tabular información, definir el esquema de la página web, etc.</p>

<p>Para poder definir esa estructura HTML permite la marcación del texto con etiquetas que el navegador interpreta para poder visualizar dicha estructura en pantalla. Dichas etiquetas son la unidad fundamental de HTML, permiten encerrar o delimitar parte del contenido para crear elementos HTML (que vamos a definir más adelante) que son los que dan la estructura y estilo al contenido.</p>

<p>De acuerdo con el tipo de elemento que permite crear se puede agrupar las etiquetas de HTML en las siguientes categorías:</p>

<ul>
  <li><b>Raíz principal</b>: Sólo hay una etiqueta en esta categoría y define el elemento de mayor jerarquía en el documento HTML.</li>
  <li><b>Documentación y metadatos</b>: Los metadatos proporcionan información de la página web, por ejemplo los estilos y scripts necesarios para renderizar la página en el navegador.</li>
  <li><b>Cabeza y cuerpo de la página</b>: Las etiquetas de esta categoría engloban en sus interior los elementos que se despliegan o no en el navegador como por ejemplo el texto y los metadatos respectivamente.</li>
  <li><b>Organización del contenido en secciones</b>: Permiten organizar el contenido en partes lógicas permitiendo crear un esquema general de la página web proporcionando elementos que permiten identificar secciones particulares del contenido.</li>
  <li><b>Contenido textual</b>: Permiten organizar en bloques o secciones el contenido que se ubica en el cuerpo del documento HTML.</li>
  <li><b>Semánticas de texto</b>: Permiten definir el significado, estructura o estilo de cualquier pieza de texto del contenido.</li>
  <li><b>Imágenes y multimedia</b>: Brindan soporte a distintos recursos multimedia como imágenes, video y audio.</li>
  <li><b>Contenido incrustado</b>: Permiten incrustar contenido externo en un punto específico del documento HTML, contenido como otras páginas web, contenido interactivo basado en plugins, etc.</li>
  <li><b>Scripting</b>: Para crear contenido dinámico e interfaces a otras aplicaciones web.</li>
  <li><b>Tabular contenido</b>: Permiten crear y manejar contenido en forma de tablas.</li>
  <li><b>Formas</b>: Son los elementos que permiten la creación de formularios que un usuario del sitio puede llenar para enviarlos a la aplicación web.</li>
  <li><b>Demarcar ediciones</b>: Permiten indicar partes donde el texto ha sufrido alteraciones.</li>
  <li><b>SVG y MathML</b>: Permiten incrustar contenido SVG y MathML directamente en el documento HTML.</li>
  <li><b>Elementos interactivos</b>: Ayudan en la creación de objetos interactivos de interfaz de usuario.</li>
  <li><b>Componentes web</b>: Permiten agregar contenido relacionado con la tecnología HTML de Web Components.</li>
</ul>

## Anatomía de un documento HTML

### Elementos HTML

<p>Un elemento dentro de un documento HTML está compuesto de la siguiente manera:</p>
<ol>
  <li><b>La etiqueta de apertura</b>: Consiste en el nombre del elemento encerrado entre 
paréntesis angulares (<>). Esta etiqueta marca el lugar a partir del cual el elemento empieza a hacer efecto.
</li>
  <li><b>El contenido</b>:Es el contenido como tal del elemento, por ejemplo el texto en un párrafo. </li>
  <li><b>La etiqueta de cierre</b>: Es exactamente igual a la etiqueta de apertura con la única diferencia de que tiene un forward slash (/) antes del nombre del elemento. Marca el punto hasta donde es válido el efecto del elemento.</li>
</ol>
<p>Así las cosas, un elemento HTML es una etiqueta de apertura seguida del contenido y este a vez seguido de una etiqueta de cierre.</p>
<img width="601" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/cfc5c650-7067-477e-bd89-3f2de64562fc">

### Anidación de elementos HTML
<p>HTML permite ubicar elementos al interior de otros elementos, a esto se le conoce como anidamiento. Es importante tener en cuenta que en el momento de anidar se debe conservar el orden en la ubicación de las etiquetas de cierre para que el navegador no vaya a tener problemas al tratar de presentar el contenido.</p>

```html
<p>Este es un <em>texto</em> de nuestro sitio web</p>
```
### Elementos HTML vacíos
<p>Son elementos que no siguen el patrón del que venimos hablando etiqueta de apertura, contenido y etiqueta de cierre y es porque no tienen contenido, entonces no tiene sentido tener etiquetas para limitarlo, de ahí el nombre de elementos vacíos.</p>

```html
<img src="imagenes/web/icon.png" alt="icono UI"/>
<input type="text"/>
```
<p>No es estrictamente necesario agregar el forward slash (/) al final, antes del paréntesis angular de cierre, pero por consistencia y compatibilidad con XML se suele agregar siempre que sea posible.</p>

### Atributos de un elemento HTML
<p>Los elementos HTML pueden tener atributos que proporcionan información adicional acerca del elemento, estos atributos no influyen en el efecto directo del elemento sobre el contenido.</p>

<p>Para definir un atributo en un elemento HTML se deja un espacio luego del nombre del elemento y se agrega el nombre del atributo seguido de un signo de igual (=) y luego entre comillas se agrega el valor del atributo.</p>

<img width="602" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/a14d26f7-0dc4-469a-8443-2c4ee4f71d4c">

<p>Existen algunos atributos conocidos como atributos <b><i>booleanos</i></b> que son aquellos que solo aceptan un valor que coincide en ser el mismo nombre del atributo, de modo que es aceptable expresar el atributo sin otorgarle valor alguno.</p>

```html
<input type="text" disabled="disabled"/>
<!-- estos dos elementos input son equivalentes -->
<input type="text" disabled/>
```
<p>Un elemento puede tener múltiples atributos, lo importante es separar los atributos por un espacio para poder diferenciarlos de los otros. No olvidar tampoco las comillas en los valores de los atributos para evitar problemas con los atributos a continuación.</p>

### Estructura básica de un documento HTML
<p>La combinación de elementos HTML individuales permite conformar un documento HTML o página web, sin embargo existe un esquema general que permite definir la estructura básica de un documento HTML que podemos ver a continuación:</p>

```html
<!doctype html>
<html lang="eng-US">
  <head>
    <meta charset="utf-8" />
    <title>Mi sitio web</title>
  </head>
  <body>
    <p>Bienvenido a mi sitio web</p>
  </body>
</html>
```
<p>El <i>doctype</i></p>
<p>Todos los documentos HTML comienzan con la declaración <code>!doctype html</code>”. El propósito de esta etiqueta es indicar al navegador la versión del lenguaje HTML que debería usarse para desplegar el documento. Nosotros no vamos a complicarnos y la vamos a utilizar tal y como la vemos actualmente ya que por defecto indica que se use la versión más reciente del leguaje, es decir, HTML5.</p>

<p>Elemento <i>html</i></p>

<p>El elemento HTML proporciona la raíz principal del documento HTML, es decir, que cualquier otro elemento que haga parte del documento debe ir anidado al interior de este. A este elemento es común verlo acompañado del atributo <b>lang</b> que sirve para mejorar la accesibilidad del documento ya que le indica a los lectores de pantalla utilizados por personas con problemas o limitaciones visuales el idioma en el que está la página web.</p>

<p>Elemento <i>head</i></p>

<p>Este elemento actúa como un contenedor para todos aquellos elementos que se van a incluir en el documento HTML pero que no hacen parte del contenido que se despliega en el navegador. Estos elementos incluyen las palabras claves y descripciones para los motores de búsqueda y otras aplicaciones web, las referencias a las hojas de estilo en CSS y scripts en Javascript que va a utilizar el documento, el título para la pestaña del navegador y la información para la sección de favoritos del navegador y el conjunto de caracteres que puede utilizar el texto del documento  entre otros.</p>

<p>Elemento <i>body</i></p>

<p>Es el elemento que contiene todo el contenido que se va a desplegar en la página web y sera visible en el navegador, esto incluye texto, hipervínculos, imágenes, multimedia, etc.</p>

<p>Comentarios en un documento HTML</p>

<p>Para agregar comentarios se utilizan los marcadores <!-- y --> para delimitar en su interior el comentario.</p>
<p></p>
<p></p>
<p></p>
<p></p>
<p></p>
<p></p>

