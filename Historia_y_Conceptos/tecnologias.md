### ¿Cómo funciona la Web?

## ¡Y así funciona Internet!

<img width="599" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/84f7aeb7-d36b-4643-b7d7-6f7bf7748213">

### Capas del modelo de comunicación

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/89b0e275-529c-429a-bd63-1a09cf47c2d6)

<b>Aplicación</b>

Esta capa la utilizan las aplicaciones de red, que así se denomina a las aplicaciones que dependen de estar conectadas a Internet para su correcto funcionamiento (navegadores web, mensajería instantánea, clientes de correo electrónico, música y entretenimiento, etc.)

Para que los mensajes (datos) producidos por las aplicaciones de red puedan transmitirse en la red se necesita de un conjunto de protocolos que dependiendo de la naturaleza de la aplicación proporcionen una interfaz entre la capa de aplicación y las capas subyacentes del modelo de comunicación. Los protocolos más conocidos son HTTP, HTTPS, FTP, SFTP, POP3, SMTP y estos en conjunto con otros forman la base de servicios de red como transferencia de archivos, navegación web, correo electrónico, mensajería instantánea, etc.

<b>Transporte</b>

La capa de transporte garantiza la integridad y confiabilidad de la comunicación mediante la segmentación de los mensajes de datos, el control de flujo y errores de la comunicación.

En el proceso de segmentación se divide los mensajes provenientes de la capa de aplicación en unidades de datos más pequeños llamados segmentos que están caracterizados por puertos de origen y destino para saber desde y hacia qué aplicación de red van dirigidos los datos y un número de secuencia para reconstruir los mensajes segmentados en el destino.

En el proceso de control de flujo la capa de transporte controla la cantidad de datos que se transmiten de forma que se aproveche en su totalidad la máxima tasa de datos entre origen y destino sin incurrir en errores o pérdidas de información.

En el proceso de control de errores la capa de transporte utiliza un esquema conocido como Automatic Repeat Request para solicitar por segmentos faltantes que se hayan pérdido en tránsito y mediante un esquema de checksum verifica que los segmentos no hayan sido modificados o corrompidos.

Los protocolos de la capa de transporte son Transmission Control Protocol (TCP) y User Datagram Protocol (UDP). TCP está orientado a la conexión mientras que UDP no.

<b>Red o Internet</b>

La capa de red es la encargada de transmitir los segmentos de la capa de transporte a un computador destino ubicado en una red distinta a la del origen de los datos.

Esta capa se encarga del direccionamiento lógico (utilizando direcciones IP), la determinación de la mejor ruta y posterior enrutamiento de los datos. Hacen parte de la capa de red los protocolos de enrutamiento OSPF, RIP, IS-IS, BGP entre otros. Y también protocolos como ICMP y ARP.

<b>Enlace (Acceso al medio)</b>

En la capa de enlace se hace la transmisión de tramas de datos entre dos dispositivos ubicados dentro de la misma red. Las direcciones físicas de las tarjetas de red (NIC), también conocidas como direcciones MAC o de acceso al medio.

Existen estrategias de acceso al medio como CSMA/CD y TDMA que la capa de enlace aprovecha para controlar cómo las tramas de datos se ubican y reciben en el medio de transmisión que puede ser confinado como cables de cobre y fibra óptica o no confinado como el aire libre o el vacío.

<b>Física</b>

Los bits que componen una trama de datos son convertidos en señales eléctricas, de luz o radiadas por la capa física dependiendo del medio por el que se vayan a transmitir los datos. La capa física se encarga de la modulación y codificación de las señales electromagnéticas que representan las secuencias binarias de las tramas de datos.

<img width="608" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/bb531161-2e1e-45ef-b8a9-168ae12cc5a5">

### TCP/IP

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/3ddffc4d-a1c9-4c54-8fe4-aa3dfa446f94)

<b>El protocolo de Internet (IP)</b>: Este protocolo es el encargado de enrutar o entregar los paquetes de datos entre una aplicación o dispositivo que es el origen de los datos y un destinatario (otra aplicación o dispositivo). El protocolo se asegura de que los paquetes lleguen al destino correcto, para este fin define un conjunto de reglas y formatos a los cuales deben ajustarse las aplicaciones y dispositivos para poder comunicarse e intercambiar datos en una red o a través de una interconexión de redes. 

<b>Protocolo de control de la transmisión (TCP)</b>: TCP es un protocolo orientado a la conexión que asegura comunicaciones íntegras y confiables entre los dispositivos (aplicaciones) interconectados en una red. Este protocolo proporciona todas las funciones necesarias para el transporte confiable de los datos en una red, siendo la más relevante el mecanismo llamado Positive Acknowledgment with Re-transmission (PAR) que sirve para hacer la re-transmisión automática de datos que se corrompen o se pierden mientras que están en tránsito a través de la red.

<b>3-way Handshake (Apretón de manos de 3 vías)</b>

Este es el proceso para establecer una conexión entre dos actores, un dispositivo (aplicación) abierto-activo conocido como el cliente y otros en modo abierto-pasivo conocido como el servidor.

En el primer estado de este proceso el cliente envía al servidor un segmento TCP denominado SYN (Synchronize Sequence Number) mediante el cual le informa que desea iniciar una comunicación con el número de secuencia indicado en el segmento.

En el segundo paso el servidor responde al cliente con un segmento SYN-ACK, donde responde a la secuencia enviada por el cliente y a su vez propone una secuencia para comenzar el intercambio de datos.

En el tercer y último paso el cliente responde al servidor con un segmento ACK que continúa la secuencia numérica enviada por el servidor, es en este momento en que se considera que la comunicación ha sido establecida.

El 3-way handshake no es un proceso exclusivo de la fase de establecimiento de la conexión. Es en realidad el pilar fundamental para que el intercambio de datos en TCP sea realmente confiable ya que cualquier dispositivo va a reenviar los datos hasta no recibir un acuso de recibo o ACK de parte del otro dispositivo involucrado en la comunicación.

En la capa de transporte se verifica la integridad de los datos mediante un algoritmo de verificación de errores ejecutado por el receptor de los datos. Si el receptor considera que los datos recibidos han sido alterados entonces procede a descartar ese segmento de modo que el dispositivo origen de los datos debe enviar de nuevo hasta que reciba un acuso de recibo positivo.

<b>Proceso completo de transferencia de datos en TCP</b>

El proceso de intercambio de datos en TCP está compuesto en general por 3 fases:

Fase 1: En esta ocurre el 3-way handshake que permite que los dispositivos se conozcan y establezcan una relación de confianza.

Fase 2: Los dispositivos intercambian información usando las secuencias numéricas acordadas en la primera fase y los acuso de recibo correspondientes.

Fase 3: Una vez se termina el intercambio de datos el servidor procede a cerrar la conexión utilizando un segmento FIN al cual el cliente debe dar acuso de recibo y a su vez el servidor debe dar acuso de recibo de la contestación del cliente para dar por terminada la conexión.

### DNS (Domain Name System)

DNS funciona semejante a como funcionaban los antiguos directorios telefónicos, solo que en este caso un nombre de dominio está asociado a una dirección IP que lo identifica en Internet.

Para los seres humanos es más fácil recordar nombres de dominio (también llamados direcciones web) que direcciones IP, luego, para visitar una página web simplemente ingresamos la dirección del sitio web en la barra de búsqueda del navegador.

<ol>
<li>El navegador revisa en su caché interna si ya tiene una dirección IP asociada al nombre de dominio.</li>
<li>El sistema operativo revisa si hay un match en la memoria caché del sistema.</li>
<li>El sistema operativo hace una consulta al “Resolver” que es un servicio que provee el proveedor de servicios de Internet o ISP. El Resolver también revisa su caché para encontrar un match.</li>
<li>El Resolver hace una petición el Root Name Server que responde con la dirección IP del servidor de Top-Level Domain (TLD) que tiene información acerca de los dominios de segundo nivel asociados a los dominios .COM, .EDU, .GOV, .ORG, etc.</li>
<li>Si el nombre de dominio tiene un subdominio entonces el servidor de dominios de segundo nivel proporciona la información acerca de los servidores de subdominio (authoritative name servers) asociados a ese dominio de segundo nivel. Ejemplo: blog.logrocket.com</li>
<li>El servidor de subdominio (authoritative name server) responde al Resolver con la dirección IP asociada a la dirección web.</li>
<li>El Resolver actualiza su caché y devuelve la dirección IP al cliente que la solicitó.</li>
<li>El cliente actualiza su caché, pasa la dirección al navegador y este actualiza su propia caché y hace la petición web a la dirección IP.</li>
</ol>

<img width="603" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/9351239d-bc1e-4775-8d99-7ffd641a6aec">

### ¿Y qué es el hipertexto?

<img width="387" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/f1c96d75-610b-45e4-9041-9bfa83195c03">

El Hipertexto es simplemente texto que contiene un enlace (hiperenlace) que conecta con otro texto. El hipertexto es la piedra angular y motor de la web. El hipertexto está intrínsecamente ligado a la web y sin este, está última no podría funcionar tal y como la conocemos hoy en día.

El hipertexto permite enlazar la información de una forma no lineal o secuencial pero asociativa al mismo tiempo. Ese tipo de enlazamiento permite que quien lea un hipertexto tenga la posibilidad de hacerlo en el orden que mejor le parezca. Además, permite que los enlaces en los textos (o documentos) puedan enlazar a cualquier otro tipo de recurso (aplicaciones, audios, videos, animaciones, etc.) u otros textos de diferentes autores (enlaces externos) distribuidos en distintos lugares.

De acuerdo con lo anterior el hipertexto se constituye en una forma única de comunicación humana que posibilita nuevas maneras para crear significado.

<b>¿Cómo funciona el hipertexto en una página web?</b>

Para crear hipertexto se necesita primero de un hiperenlace, este último se puede incrustar en una página web utilizando el lenguaje <strong>HTML</strong> (<i>Hypertext Markup Language</i>) el cual tiene una etiqueta llamada “ancla” (anchor) que permite asociar una URL a un texto. De este modo se incrusta en el texto una referencia digital a otra página web y al dar clic al texto se accede a la página referenciada.

Al hacer clic en el hipertexto se da inicio a una petición al servidor que aloja el sitio web, utilizando el protocolo <strong>HTTP</strong>, generalmente, una petición <strong>GET</strong> que permite descargar desde el servidor al navegador el contenido de la página web (el código HTML).

<img width="597" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/9e121ca5-ee23-41b2-a57a-e32cd87a28c3"><br>
a: Etiqueta ancla (anchor) de HTML<br>
href: referencia a un contenido o recurso en la web

### URL: Uniform Resource Locator (Localizador Uniforme de Recursos)

<img width="569" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/27125788-876e-47b4-8ce6-7c82749be4af">

Una URL es la dirección única de un recurso alojado en la web. Cuando se hace una petición a un servicio web se debe especificar una URL para que el servidor web pueda identificar las características del servicio que es solicitado.

<b>Partes de una URL</b>

Esquema: Se refiere al protocolo que se utiliza para hacer la petición. En el caso de hipertextos se utiliza el protocolo HTTP (Hypertext Transfer Protocol).

Nombre de dominio: Es la dirección web que identifica al recurso en la web y que está asociada a una única dirección IP.

Ruta al recurso: Identifica el lugar exacto dentro del servidor donde se encuentra el recurso en específico que se está solicitando. Dicho de otra manera es la ruta a través de los directorios donde está almacenado el recurso (la página web).

Parámetros de la solicitud: Son opcionales, se señalizan con el símbolo de interrogación (?) y se separan con el símbolo ampersand (&). Son parejas de clave y valor que se utilizan para solicitar al servidor que realice acciones específicas sobre los resultados que debería retornar la petición.

Los usos más comunes de los parámetros de solicitud son:

<ol>
<li>Rastreo (tracking): Son parámetros pasivos, es decir que son agregados automáticamente en la solicitud y no directamente desde el navegador. Se utilizan para rastrear detalles del cliente que hace la solicitud, como por ejemplo si pulsó algún enlace para llegar a la página, la geolocalización, etc.</li>
<li>Ordenamiento (sorting): Son parámetros activos que permiten indicarle al servidor que ordene los resultados en un orden determinado.</li>
<li>Búsqueda (searching): Para indicarle al servidor que utilice su motor interno de búsqueda.</li>
<li>Identificación (identifying): Para retornar resultados que cumplan con una característica determinada que los diferencia de otros.</li>
<li>Paginación (pagination): Sirven para dividir largas listas de resultados en bloques o series de menor longitud.</li>
<li>Traducción (translation): Para indicar el idioma en el que se quiere recibir los resultados o el recurso.</li>
<li>Filtrado (filtering): Para retornar solo resultados que cumplan con un criterio específico.</li>
</ol>

<img width="589" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/0e3494c8-fd4e-4cd4-884a-10f1915902fb">

Ancla: Esta parte de la URL también es opcional y sirve para ubicarse en un punto en específico del contenido de la página web. Vale la pena mencionar que dicho punto debe estar señalizado por una etiqueta HTML llamada ancla nombrada (named anchor) para poder ser alcanzable. Se utiliza el símbolo numeral (#) para indicar el nombre del ancla en la URL.

### HTTTP: HyperText Transfer Protocol (Protocolo para la transferencia de hipertexto)

HTTP es un protocolo (sistema de reglas, formatos y estándares que definen cómo se deben compartir/intercambiar datos en redes de computadores) de Internet que se utiliza para interactuar con recursos que están alojados en servidores que están interconectados en la red.

HTTP es un protocolo para conexiones bajo arquitectura cliente-servidor, donde el cliente, usualmente un navegador web, realiza una solicitud a un servidor web que almacena un recurso que le interesa consultar que puede ser desde un documento HTML, una base de datos, una imagen, un video, etc.

En HTTP los clientes y servidores se comunican mediante el intercambio de mensajes individuales. A los mensajes enviados por el cliente se les conoce usualmente como peticiones (requests) mientras que a los mensajes enviados por el servidor, en respuesta a las peticiones del cliente, se les conoce simplemente como respuestas (responses).

Para mostrar la página web en pantalla el navegador web (cliente) envía una petición solicitando un documento HTML al servidor. Cualquier otro script presente (o referenciado) en el documento HTML ejecutado por el navegador y que necesite hacer solicitudes adicionales al servidor para actualizar el contenido de la página web lo hará a través de nuevas peticiones HTTP al servidor.

