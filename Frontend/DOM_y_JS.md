# DOM y Javascript

## Document Object Model (DOM)

### ¿Qué es el DOM?

El <i>Document Object Model</i> o modelado del documento HTML como un objeto es un representación en forma de árbol de todos los contenidos de un documento HTML. El árbol ilustra las relaciones que existen entre los elementos HTML que componen el documento HTML. Explicado de una manera más sencilla, el árbol es un gráfico de un conjunto de nodos conectados de forma jerárquica, es decir nodos padres conectados a sus nodos hijos, donde los nodos son los elementos HTML del documento. Luego, cada rama del árbol indica los niveles de anidamiento que existen en un elemento html.

```html
<!DOCTYPE html>
<html>
 <head>
  <title>Mi web</title>
 </head>
 <body>
  <header>
   <h1>Bienvenido a mi sitio web!</h1>
  </header>
  <main>
   <section>
    <article>
     <h2>Acerca de mí...</h2>
     <img src="img/yo.jpg" alt="Yo programando xD" \>
     <p>Soy un desarrollador web con 5 años de experiencia, mis tecnologías favoritas son...</p>
    </article>
   </section>
  </main>
 </body>
</html>
```
![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/3880aabd-561a-425a-b553-17587e42ca6a)

El árbol es solo una representación gráfica de lo que es el DOM, lo realmente importante es que el DOM modela el documento HTML como un objeto (y también a todos los nodos pertenecientes al árbol del DOM), ¿se les hace común ese término?, si pensaste en el concepto de objeto de la programación orientada a objetos estás en la ruta indicada. Al modelar el documento HTML como un objeto el DOM nos proporciona propiedades y métodos que podemos aprovechar para acceder a los elementos HTML del documento y de esa forma podamos manipular, insertar, actualizar y remover elementos HTML del documento con ayuda de un lenguaje de programación. Con Javascript por supuesto, recuerda que el navegador web está en capacidad de correr código escrito en Javascript.

![tomado de https://www.w3schools.com/js/js_htmldom.asp](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/f25880b0-15ec-42c7-9726-246e45fbfb97)


### Propiedades y métodos del DOM

Para acceder al nodo principal de un documento HTML utilizamos la propiedad <code>document.documentElement</code>, este nodo contiene a todos los elementos del documento HTML, es decir, esta propiedad nos devuelve un objeto del elemento <code>html</code> del documento. Recuerde que todos los elementos del documento descienden de este elemento y por encima de el no hay ningún otro elemento, por eso mismo es el nodo principal del documento. Abra una consola de Javascript en su navegador y ejecute el siguiente comando para que verifique esta propiedad.

```js
console.log(document.documentElement)
```
A continuación vamos a proporcionar un listado de las propiedades del DOM de uso más común, puede ver un listado más completo en este [enlace](https://linuxhint.com/html-dom-methods-and-properties/)

Propiedades de los nodos del DOM:
<ul>
  <li><b>innerHTML:</b>Esta propiedad permite acceder y manipular el contenido de un elemento HTML</li>
  <li><b>nodeName:</b></li>
  <li><b>nodeValue:</b></li>
  <li><b>nodeType:</b></li>
  <li><b>parentNode:</b></li>
  <li><b>parentElement:</b></li>
  <li><b>style:</b></li>
  <li><b>tagName:</b></li>
  <li><b>textContent:</b></li>
  <li><b>firstChild:</b></li>
  <li><b>lastChild:</b></li>
  <li><b>className:</b></li>
  <li><b>childNodes:</b></li>
  <li><b>attributes:</b></li>
  <li><b>childElementCount:</b></li>
  <li><b>id:</b></li>
</ul>














