# Consumiendo APIs con Javascript

Para implementar una REST API, con todas las operaciones CRUD que revisamos con anterioridad, Javascript tiene a disposición las siguiente dos opciones:
<ol>
  <li><b>Fetch API</b>: Es una interfaz nativa de Javascript que permite hacer peticiones HTTP a un servidor. Utiliza las <code>Promises</code> para permitir un flujo de ejecución asincrónico (esto se debe a que no podemos predecir cuánto tiempo puede tomar la petición en completarse). De modo general esta interfaz ofrece un método global <code>fetch</code> al cual se le pasa la URL del endpoint al que se va a hacer la petición y un conjunto de opciones que permiten especificar parámetros como el método, los encabezados y el cuerpo (cuando se requiere) que permiten la construcción de la petición HTTP.</li>
  <li><b>Axios</b>: Es un cliente HTTP proporcionado en una librería de terceros de Javascript. Funciona de manera similar a Fetch API en cuánto a la construcción de peticiones HTTP, es decir, requiere de una URL y un conjunto de parámetros para la confección de la petición y utiliza <code>Promises</code> para un flujo de ejecución asincrónico</li>
</ol>

## Fetch API


## Axios API
