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

Las APIs remotas en cambio si requieren de un agente externo que proporcione los medios y capacidades para que las aplicaciones puedan llevar a cabo el proceso de comunicación, y esto se debe en gran medida a que las aplicaciones residen en sistemas distintos, entonces antes de comunicarse a través de la API se debe establecer una conexión que debe ser mantenida por el tiempo de la comunicación. El ejemplo más común aquí es cuando usted va a navegar en un sitio web, su navegador web está instalado en su computador mientras que el sitio web está alojado en un servidor que no es más que otro computador localizado en otra ubicación geográfica distinta. Por tanto, para que su navegador pueda solicitar contenidos al servidor web requiere de establecer una conexión remota al servidor que le permita hacer la solicitud a través de la web API.

Esta categoría de APIs es la que nos interesa para los fines de este curso.

Clasificación de las APIs de acuerdo con su arquitectura:
<ul>
  <li></li>
  <li></li>
  <li></li>
  <li></li>
  <li></li>
</ul>


