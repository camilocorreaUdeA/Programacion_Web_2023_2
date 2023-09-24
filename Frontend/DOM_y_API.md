# DOM, APIs y otras cosas

## DOM

## ¿Qué es una API?

![Tomado de https://blog.back4app.com/](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/aa6dd98f-35f4-4602-b597-2ff4c48b707d)

Una Interfaz de Programación de Aplicación (Application Programming Interface) es en general un conjunto de funciones, protocolos o funcionalidades encapsuladas en subrutinas que permiten la comunicación e integración entre dos artefactos o productos de software distintos (apliaciones, servicios web, sistemas operativos, drivers, etc.) que no cuentran entre ellos con una interfaz nativa de interacción.

Habitualmente se le conoce a la API como el contrato establecido entre un proveedor de información y un usuario de información, en el que se establece la información requerida de parte el usuario y hacia el proveedor, o en otros términos la llamada o solicitud. Y también la información que va del proveedor hacia el usuario, es decir la respuesta.

De ese modo la API actúa como un mediador o capa intermedia que simplifica la comunicación entre el usuario y el proveedor sin necesidad de que ninguno tenga conocimiento de la implementación interna del otro. Por ejemplo cuando usted instala una aplicación en su computador o dispositivo móvil y esta hace uso de varios periféricos del sistema, en ese caso su aplicación no necesita saber como acceder y manipular los recursos disponibles en el sistema sino que es este último el que a través del sistema operativo proporciona una serie de funciones que le permiten a la aplicación conectar con los periféricos sin tener que manipularlos directamente, y por el contrario es el sistema operativo quien de acuerdo con las solicitudes de la aplicación va a comunicarse con los drivers del hardware del sistema para acceder a los dispositivos.

![Tomado de https://www.techspot.com/article/2237-what-is-api/](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/8f2b13d2-2f1b-49b4-8c16-47f906544728)

El uso de APIs simplifica en gran medida el desarrollo de nuevas aplicaciones y productos de software y además genera ahorros en tiempo y dinero al permitir que la integración con otras aplicaciones y artefactos de software sea transparente, homogenea y no requiera de la intervención de interfaces y adaptadores hechos a la medida que en el mayor número de casos no existen y tienen que ser desarrollados antes que la misma aplicación.

Por ejemplo, suponga que usted desea desarrollar una aplicación para aplicar filtros a las fotografías que los usuarios se toman con la cámara de su IPhone, y además estás fotos se suben directamente a la cuenta de Instagram del usuario une vez el usuario lo ha solicitado. Piense por un momento cuánto tiempo tomaría crear una aplicación de estas características si usted estuviera obligado a tener que crear una pieza de software para manipular la cámara del Iphone y otra adicional para la subida y publicación de las fotos en Instagram. Afortunadamente eso no es necesario ya que lo único en lo que usted debe concentrarse es en el desarrollo de la aplicación porque el sistema operativo del IPhone le proporciona las funciones de alto nivel que usted puede utilizar para comunicarse con la cámara del dispositivo. Y por el otro lado, la aplicación de Instagram cuenta con una web API para que con a tavés de una simple solicitud HTTP usted subir y publicar las fotos sin preocuparse del funcionamiento interno de la aplicación de Instagram. 

Clasificación de las APIs

Hata este punto hemos dado una definición general al concepto de API y comentado algunas situaciones ejemplo para ilustrar más clramente el concepto. Ahora daremos un vistazo a la clasificación de las APIs y vamos a empezar por una clasificación general en la cual dividiremos las APIs en dos grandes categorías: Locales y Remotas.

Las APIs locales son las que hemos venido comentando en los ejemplos anteriores, son estás interfaces que actúan como puentes que comunican a dos artefactos de software que están contenidas o instaladas dentro de un mismo sistema. Estás APis no requieren de un sistema auxiliar de conexión a distancia entre quien usuario y proveedor sino que al contrario ambos coinciden en compartir un bus de datos físico mediante el cual se pueden comunicar de forma directa. Por ejemplo cuando el mouse envía las señales de pulsado de botones o del sensor de moviemiento al sistema operativo para que se vea el cursor moverse y dar clic en pantalla. En este caso puntual el mouse no requiere de ningún sistema auxiliar que le ayude a encontrar la ruta al sistema operativo para enviar su información, y es obviamente porque tanto los drivers del mouse y la pantalla como el sistema operativo están instalados en el mismo hardware y pueden comunicarse utilizando los medios de comunicación física ques este último ofrece.

Esta categoría de APIs no es de nuestro interés para los alcances y objetivos del curso.

Las APIs remotas en cambio si requieren de un agente externo que proporcione los medios y capacidades para que las aplicaciones puedan llevar a cabo el proceso de comunicación, y esto se debe en gran medida a que las aplicaciones residen en sistemas distintos, entonces antes de comunicarse a través de la API se debe establecer una conexión que debe ser mantenida por el tiempo de la comunicación. El ejemplo más común aquí es cuando usted va a navegar en un sitio web, su navegador web está instalado en su computador mientras que el sitio web está alojado en un servidor que no es más que otro computador localizado en otra ubicación geográfica distinta. Por tanto, para que su navegador pueda solicitar contenidos al servidor web requiere de establecer una conexión remota al servidor para luego poder hacer la solicitud a través de la API.

Esta categoría de APIs es la que nos interesa para los fines de este curso.

Clasificación de las APIs remotas en términos de acceso:
<ul>
  <li><b>Open/Public API:</b> SOn las APIS que están disponibles de manera pública y no existen restricciones para utilizarlas. Públicas no necesariamente significa que su uso sea gratuito.</li>
  <li><b>Partner API:</b> Son APIs que no están disponibles de forma pública y por tanto se necesita contar con permisos o licencias para utilizarlas. Dentro de esta categoría estpan las APIs que un socio de negocio proporciona para interactuar con sus productos de software.</li>
  <li><b>Internal/Private API:</b> Son APIs para usar solo entre componentes internos del sistema y por tanto no se exponen al público. Un buen ejemplo es una API que permite a un servicio de un sistema distribuido hacer una ejecución remota de un procedimiento que está localizado en otro servicio del mismo sistema al que está conectado mediante una red de comunicación en común.</li>
  <li><b>Composite API:</b> Son APIs que se modelan como una combinación de distintos datos y APIs. Por lo general son secuencias de tareas que corren de manera sincrónica para incrementar la velociad de ejecución de un proceso al poder agrupar y ejecutar varias solicitudes en un solo llamado a la API. Además mejora el desempeño del cliente ya que en vez de hacer varios llamados a distintas APIs y esperar individualmente las respuestas de cada una puede hacer un solo llamado en el que se hacen todas las solicitudes y se recibe una sola respuesta.</li>
</ul>

Clasificación de las APIs remotas de acuerdo con su arquitectura:
<ul>
  <li><b>API REST:</b> Representational State Trasnfer permite que un cliente ejecute acciones en un servidor utilizando los métodos HTTP (GET, POST, PUT, PATCH y DELETE) El servidor se encarga de proporcionar "endpoints" para que el cliente pueda manipular distintos tipos de datos. La característica principal de REST es que se no conserva el estado entre solicitudes, o sea que el servidor no conserva datos del cliente y las distintas solicitudes de cliente a servidor son independientes y carecen de relación. En otras palabras, que se pueda ejecutar o no una solicitud no depende de la ejecución de solicitudes pasadas ni va a influir en la ejecución de solicitudes futuras.</li>
  <li><b>API GraphQL:</b> El lenguaje GraphQL permite desarrollar APIs que ofrecen acceso a datos vinculados que se pueden pensar como un grafo de entidades, por lo tanto se pueden crear solicitudes fuertemente tipadas en las que se espcifica de manera exacta los datos que se necesitan. Luego, mediante un solo "endpoint" expuesto el servidor se obtienen los datos, al contrario de otros tipos de APIs en las que se necesitaría consultar múltiples "endpoints" para obtener los mismos datos.</li>
  <li><b>API SOAP:</b> SOAP es el acrónimo de Simple Object Access Protocol que es un estándar de mensajería que se basa en el formato XML para la construcción de los mensajes de solicitud y respuesta. SOAP utiliza el patrón de RPC (Remote Procedure Call) donde la comunicación entre cliente y sevidor se modela como el llamado a funciones o métodos a los que se les pasan parámetros específicos para que se retorne un resultado con los datos esperados.</li>
  <li><b>API RPC:</b> La arquitectura RPC propone un enfoque en el cual cargas de trabajo remotas se ejecutan como si fuera de manera local, esta característca hace que RPC sea muy útil en la construcción de servicios distribuidos ya que una aplicación se puede componer de rutinas o funciones que internamente hacen llamados a funciones que se ejecutan en otros servicios externos. Esto también significa una gran felxibilidad en comparación con arquitecturas como REST poque se puede ejecutar practicamente cualquier función y solicitar cualquier tipo de acción sobre los datos, mientras que REST está restringido a los métodos ofrecidos por el protocolo HTTP.</li>
</ul>

Para los fines y objetivos de este curso nos interesa enfocarnos solo en las APIs REST así que desde este punto y en adelante cada vez que usemos el término API nos estaremos refiriendo específicamente a REST API.

## REST APIs

Las REST APIs comunican un cliente y un servidor a través de peticiones/solicitudes HTTP, en casi todas las ocasiones para ejecutar operaciones como crear, leer, actualizar y eliminar recursos en una base de datos. Por ejemplo para leer/obtener un recurso a través de una REST API se tendría que usar el método HTTP GET, para crear un nuevo recurso el método POST (o PUT), para actulizar están disponibles los métodos PUT y PATCH y para eliminar un recurso está el método DELETE.

Hay que tener en cuenta que cuando hablamos de obtener o leer un recurso no estamos haciendo referencia al recurso como tal sino a una representación del mismo, de ahí viene el mismo acrónimo de REST, Representational State Transfer, transferencia del estado de representación. Y una representación de un recurso es el estado de ese recurso en un momento en particular. 

La representación del recurso puede ser presentada por la API en distintos formatos, practicamente en cualquiera, entre las más populares se incluyen JSON (Javascript Object Notation), HTML, XML, Python, PHP, texto plano, entre otras. Siendo la más popular de todas JSON ya que es más entendible para el lector y más fácil de decodificar para una máquina y además es agnóstico del lenguaje de programación.

En REST APIs que estén bien diseñadas se incluye el uso de encabezados y parámetros en las solicitudes, ya que permiten incluir información importante como metadatos, credenciales de autenticación y autorización, las rutas a los recursos o URIs, datos para caching y cookies entre otros. Y en el caso de las respuestas también se considera importante incluir los códigos de estado de la solicitud HTTP.

A continuación listamos algunos principios de diseño de APIs REST

<ul>
  <li><b>Interfaces uniformes:</b> Todas las solicitudes al mismo recurso deben ser iguales sin importar desde dónde se haga la solicitud. Se debe asegurar que cualquier dato pertenece a un solo identificador uniforme de recursos URI. Se recomienda también que los recursos no sean muy grandes pero que contengan la información que necesita el cliente.</li>
  <li><b>Desacoples entre cliente y servidor: El cliente y el servidor deben ser sistemas independientes. La única información que el cliente debe conocer del servidor son las URIs a los recursos y aparte de los "endpoints" proporcionados por el servidor no debe tener ningún otro canal de interacción con este. Y por el lado del servidor, este no debe poder modificar nada en la aplicación del cliente y su única función de cara al cliente es la de responder con los datos solicitados por este.</b></li>
  <li><b>Ausencia de estado en las solicitudes:</b> Cada solicitud requiere compartir toda la información necesaria para su procesamiento. No deben existir sesiones de lado del servidor y este no debe almacenar ningún dato del cliente.</li>
  <li><b>Caché-abilidad:</b> Para incrementar el desempeño del lado del cliente y la escalabilidad del lado del servidor se pueden utilizar cachés de los recursos en el lado del servidor. El servidor debe informar al cliente que recursos pueden ser cacheables.</li>
  <li><b>Arquitectura de capas:</b> En el canal de comunicación entre cliente y servidor pueden existir muchas capas intermediarias, por tanto se deben diseñar las APIs para que tanto cliente como servidor no tengan en cuenta si en realidad se están comunicando de manera directa con una apliación intermediaria.</li>
  <li><b>Código bajo demanda:</b> En caso de que una API se utilice para enviar código que pueda ser ejecutado se debe restringir está ejecucuión a que solo sea bajo demanda, solicitud o autorización explícita.</li>
</ul>

Hemos mencionado en varias ocasiones dos conceptos para los cuales no hemos dado una definición concreta. Veamoslas a continuación:

JSON:

Endpoint:
