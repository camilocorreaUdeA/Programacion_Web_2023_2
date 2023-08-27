# HTML

## Hipervínculos en un documento HTML

<p>Los enlaces o hipervínculos son la esencia de la WWW son los que permiten establecer referencias entre distintos hipertextos. Es decir, son los que permiten que desde un sitio web podamos acceder a otro con solo dar clic en un enlace.</p>

<p>Para agregar enlaces a nuestros documentos HTML utilizamos el elemento <code>a</code>, <i>anchor</i> en español <i>ancla</i>, pero debemos tener en cuenta que no va a funcionar a menos de que especifiquemos un atributo en particular, en este caso es el atributo <code>href</code> o <i>hypertext reference</i> (referencia a hipertexto) donde especificamos el documento HTML que vamos a visitar luego de dar clic en el enlace.</p>

```html
  <a href=”http://www.enlace.com”>Este es un hipervinculo</a>
```
<p>Un enlace definido como acabamos de ver va a abrir la nueva página en la misma pestaña del navegador de la página en la que estamos en ese momento. Si queremos que en enlace abra la página en una pestaña nueva debemos agregar el atributo <code>target</code>code> con el valor <code>_blank</code>code></p>

```html
<a href=”http://www.enlace.com” target=”_blank”>Este abre en una nueva pestaña</a>
```
Existen 3 tipos de enlaces que podemos crear:

1. Enlaces a recursos externos que están disponibles a través de Internet.
2. Enlaces a recursos internos que están en los directorios donde se almacena el documento HTML que los referencia.
3. Enlaces a ubicaciones específicas dentro de un documento HTML (anclas).

Para los enlaces del primer tipo utilizamos lo que se conoce como referencias absolutas o dicho de otro modo la URL completa del recurso. Como se vio anteriormente en la sección de Historia y Conceptos una URL es su forma más básica está compuesta de la siguiente forma: <code>protocolo://nombre-dominio/ruta-recurso</code>

```html
<a href=”https://www.codecademy.com/learn” target=”_blank”>Aprenda HTML</a>
```
Para el segundo tipo vamos a proceder a pasar la ruta relativa dentro de los directorios de nuestro servidor donde se encuentra el recurso al que estamos haciendo referencia.

1. Para recursos que están en el mismo directorio que nuestro documento HTML simplemente pasamos el nombre del archivo.

```html
<a href=”acerca.html” target=”_blank”>Acerca de nosotros</a>
```
2. Para los que están en otro directorio que se encuentra al nivel del documento HTML pasamos el nombre del directorio (y otros directorios internos, si es el caso) y el nombre del archivo.

```html
<a href=”cursos/2023/participantes.html” target=”_blank”>Lista estudiantes del curso</a>
```
3. Y finalmente para recursos que están en directorios por fuera del directorio donde se encuentra el documento HTML vamos a utilizar la notación “../” para navegar a través de los directorios.

```html
<a href=”../ejercicios/web.html” target=”_blank”>Práctica HTML</a>
```
En lo que respecta al tercer tipo es similar a los anteriores con la condición de que debe existir un ancla al punto específico del documento que se busca referenciar. Esto se logra con un atributo id del elemento referenciado.

Por ejemplo, queremos referenciar un párrafo en particular dentro del mismo documento HTML, entonces el elemento &lt;p&gt; debe tener un atributo <code>id</code> que funcionará como ancla.

```html
<p id=”nombre-ancla”>Contenido del párrafo…</p>
```
Y el enlace debe construirse de la siguiente manera utilizando el símbolo # y el nombre del ancla tal cual se definió:

```html
<a href=”#nombre-ancla”>Referencia al párrafo en específico</a>
```
Buenas prácticas a la hora de crear hipervínculos:
<ul>
  <li>Siempre es una buena idea incluir palabras clave en el texto de los hipervínculos para que describan de manera efectiva lo que se está tratando de vincular.</li>
  <li>Por ningún motivo use la URL como texto del hipervínculo.</li>
  <li>Evite usar palabras como “hipervínculo”, “enlace”, “link” o “clic aquí”  en el texto del hipervínculo.</li>
  <li>Trate de que los textos del hipervínculo sean breves y concisos.</li>
</ul>

## Multimedia

### Imágenes

Para agregar una imagen o un gif animado a un documento HTML se utiliza el elemento vacío <code>img</code> acompañado como mínimo de los siguientes atributos:
<ul>
  <li>src: es la ruta a la imagen, sin este atributo no se puede cargar ninguna imagen en el navegador. Su utilización es exactamente igual a como vimos recientemente para el atributo href de los hipervínculos. Podemos indicar la ruta absoluta a una imagen que es un recurso en la web como: https://images.unsplash.com/photo-1582568742002-499ac33758ce, o bien puede ser una imagen alojada en un directorio de nuestro servidor y en ese caso utilizamos una ruta relativa a la imagen.</li>
  <li>alt: es un texto alternativo que se usa para dar una descripción en detalle de la imagen cuando por algún motivo no se carga en la página web.</li>
  <li>width y height: estos atributos especifican las dimensiones de la imagen en píxeles. Se recomienda utilizar estos atributos para evitar la carga desordenada de la página web tan pronto se van cargando las imágenes ya que estos atributos reservan el espacio que van a ocupar las imágenes al momento de carga de la página web y de esa manera se evita ese movimiento molesto del contenido de la página conforme se van cargando las imágenes.</li>
</ul>

```html
<img
  src="images/medellin-city.jpg"
  alt="Una vista panorámica de la ciudad de Medellín, se observa el centro de la ciudad con sus edificios y los barrios enclavados en la ladera oriental"
  width="400"
  height="341" />
```
Para agregar un título, leyenda o pie de foto a la imagen puedes utilizar el elemento <code>figure</code> en conjunto con el elemento <code>figcaption</code>. El primero es simplemente un contenedor, no necesariamente de imágenes, que proporciona una semántica que ayuda a dar estructura a todo el contenido de la página. Mientras que el segundo elemento sirve para dar una explicación o descripción  del contenido del elemento figure.

```html
<figure>
  <img
    src="images/medellin-city.jpg"
    alt="Una vista panorámica de la ciudad de Medellín, se observa el centro de la ciudad con sus edificios y los barrios enclavados en la ladera oriental"
    width="400"
    height="341" />
  <figcaption>
    Ciudad de Medellín. Fotografía tomada de Periódico Alma Mater.
  </figcaption>
</figure>
```
### Videos

Existen diferentes alternativas para agregar videos a un documento HTML, nosotros solo veremos las dos de uso más extendido.

Elemento <code>video</code>

Este elemento permite incrustar video en un documento HTML de forma ágil e intuitiva, de una manera muy similar a como vimos que se insertan imágenes. En su forma más simple vamos a utilizar el elemento video para especificar la fuente del video, agregar una interfaz para controlar la reproducción y el volumen del audio del video y un mensaje para el usuario en caso de que por algún motivo no se pueda reproducir el video.

```html
<video src="mi-video.mp4" controls>
  <p>
    Upss, al parecer tu navegador no puede reproducir este video.
    <a href="mi-video.mp4">Puedes verlo aquí</a>
  </p>
</video>
```
<ul>
  <li>src: Es la ruta donde se encuentra el video, aplica la misma explicación que usamos para este mismo atributo en el elemento <code>img</code>.</li>
  <li>controls: Es un atributo booleano, por eso no tiene valor asociado, que sirve para indicarle al navegador que agregue al video una interfaz con controles para la reproducción del video (reproducir, pausar, adelantar, etc.) y para el volumen del audio.</li>
  <li>Párrafo opcional para fallas: Es un mensaje que se muestra cuando la reproducción del video falla y se puede aprovechar para brindar una alternativa para la reproducción del video (un enlace externo por lo general).</li>
</ul>

Si por alguna razón el navegador no soporta el formato del video, existe un elemento HTML que se puede anidar en el elemento video para tener disponibles otras fuentes de backup en distintos formatos y que probablemente si sean soportadas por ese navegador. Dicho elemento es source (un elemento vacío) que tiene un atributo src, ya largamente discutida su funcionalidad hasta el momento, y otro atributo type para especificar el formato del video.

```html
<video controls>
  <source src="mi-video.mp4" type=”video/mp4” />
  <source src="mi-video.webm" type=”video/webm” /> 
  <p>
    Upss, al parecer tu navegador no puede reproducir este video.
    <a href="mi-video.mp4">Puedes verlo aquí</a>
  </p>
</video>
```
Existen otros atributos adicionales que se pueden agregar al elemento video y que pueden ser de utilidad dependiendo del contexto en el que se incrusta el video:
<ul>
  <li>width y height: para indicar las dimensiones del video en el navegador. Funciona siempre y cuando se conserve la razón ancho/alto del video (aspect ratio).</li>
  <li>autoplay: Booleano. Si está presente indica que el video inicie reproducción de manera automática.</li>
  <li>loop: Booleano. Hace que el video se reproduzca en un ciclo continuo.</li>
  <li>muted: Booleano. Hace que por defecto el audio del video esté muteado.</li>
  <li>poster: Permite indicar la ruta a una imagen que se despliega antes de que el video sea reproducido.</li>
  <li>preload: Permite hacer buffering del video cuando se le asigna el valor “auto”.</li>
</ul>

```html
<video 
  controls
  width=”400
  height=”400”
  autoplay
  loop
  muted
  preload=”auto”
  poster=”anuncio.jpeg”>
  <source src="mi-video.mp4" type=”video/mp4” />
  <source src="mi-video.webm" type=”video/webm” /> 
  <p>
    Upss, al parecer tu navegador no puede reproducir este video.
    <a href="mi-video.mp4">Puedes verlo aquí</a>
  </p>
</video>
```
Este elemento permite incrustar en tu sitio web contenido de terceros (documentos HTML, videos, aplicativos, plugins, etc.) sobre el que no se tiene control directo y se quiere aprovechar tal cual está hecho por ejemplo videos de Youtube, mapas de Google Maps, sistemas de comentarios como el de Disqus, modales con trinos de Twitter o posts de Facebook, etc.

Se recomienda el uso del elemento iframe sólo cuando sea estrictamente necesario y en muy contadas ocasiones ya que puede acarrear problemas de seguridad a tu sitio web.

Para incrustar un video de Youtube vaya al video y de clic en la opción compartir y luego en la opción insertar y copie el código que allí le proporcionan.







