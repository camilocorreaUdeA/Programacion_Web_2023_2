# Manos a la obra

A poner en practica lo que has aprendido de Javascript

## Ejercicio propuesto

Implemente una sencilla aplicación web aprovechando los datos que proporciona la API pública de Pokémon.

### Requerimientos:

<ol>
  <li>Implemente una aplicación que permite averiguar datos de los pokémons extraídos directamente de la API pública de Pokémon.</li>
  <li>La aplicación debe cargar una página inicial o home en la que hay una barra de búsqueda para ingresar el nombre del pokémon que se desea consultar (ver Figura 1).</li>
  <li>Si la consulta en el home falla por alguna razón, usted debe indicar ese suceso por algún medio al usuario (sea un pop-up del navegador o un mensaje que se carga directamente en la página). Una falla muy común es que el usuario ingresa mal el nombre del pokémon (mal escrito, mala ortografía, etc.), o bien simplemente puede darse una falla en la petición a la API.</li>
  <li>Si la consulta es exitosa entonces se debe cargar la información del pokémon en el home como se observa en la Figura 2. Debe tener en cuenta:</li>
  <ol>
    <li>Los detalles o información del pokémon son a gusto del desarrollador pero son obligatorios el nombre del pokémon, la imagen, una descripción, un corto listado de habilidades o movimientos y la siguiente evolución (este último campo solo se debe mostrar cuando aplique, es decir, no se debe cargar si el pokémon no tiene una evolución)</li>
    <li>Se debe agregar un botón para evolucionar tal y como se muestra en la Figura 2. Este botón sólo debe aparecer en la página cuando el pokémon puede evolucionar, de lo contrario no se debe cargar ese botón.</li>
    <li>Al dar clic en el botón de evolución se debe cargar la información del pokémon evolucionado. Se debe seguir viendo de forma similar a como se muestra en la Figura 2.</li>
    <li>En la esquina superior derecha se debe tener un enlace al home. Al dar clic en este enlace se debe cargar el home tal y como se muestra en la Figura 1. Es decir, se deja de observar cualquier información de un pokémon que esté cargada en ese momento.</li>
  </ol>
</ol>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/14787ea5-5fb0-46b2-9e83-f6fd136562cd)

Figura 1. Sitio principal o home de la aplicación.


![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/cfbe2872-0f77-409f-abd5-42b308d20a84)

Figura 2. Aspecto de la aplicación cuando se cargan los detalles de un pokémon.

A continuación puede encontrar una breve descripción de algunos endpoints con los que cuenta la aplicación. Estos endpoints le proporcionan la información necesaria para completar la actividad.

Nombre de la API: PokéAPI

[PokéAPI](https://pokeapi.co/)

Endpoints:

<ol>
  <li>
    Pokemon: https://pokeapi.co/api/v2/pokemon/{nombre_pokemon}
    Una solicitud de tipo GET a este endpoint devuelve un json que contiene datos relevantes del pokemon por el que se está consultando.

El campo id trae un número entero que en ocasiones se podrá utilizar como parámetro de consulta (path param) en las URLs de otros endpoints de la API.

El campo name, contiene el nombre del pokemon.

En el objeto sprites hay URLs a imágenes o “thumbnails” del pokemon.

Los campos abilities y moves contienen habilidades y movimientos característicos del pokémon.
  </li>
  <li>
  Species: https://pokeapi.co/api/v2/pokemon-species/{id_pokemon}/
    Este endpoint en particular retorna dos campos de interés:

evolution_chain: es un objeto que en el campo url almacena una URL a un endpoint en el que se puede averiguar si el Pokémon puede evolucionar y cuáles serían esas posibles evoluciones.

flavor_text_entries: es un arreglo de objetos en los que se puede encontrar un campo flavor_text que contiene una corta descripción del pokémon en distintos idiomas.
  </li>
  <li>
  Evolution Chain: https://pokeapi.co/api/v2/evolution-chain/{id}

  Este endpoint permite averiguar la posible evolución de un pokémon.
En el objeto retornado existe un campo evolves_to que su vez contiene un objeto species en el que se puede consultar el nombre de la siguiente evolución y la URL al endpoint Species de la evolución.    
  </li>
</ol>

NOTA: Las figuras que se muestran son solo mockups de referencia. Por favor explote al máximo su creatividad y estilice la aplicación de la mejor manera posible.

ENTREGA: Tal cual lo ha hecho para las demás actividades. Un enlace al repositorio y otro al despliegue de la aplicación en Github. 
