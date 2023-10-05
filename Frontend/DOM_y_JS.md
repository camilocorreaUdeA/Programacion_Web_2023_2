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
  <li><b>innerHTML: </b>Esta propiedad permite acceder y manipular el contenido de un elemento HTML.</li>
  <li><b>nodeName: </b>Devuelve el nombre del nodo.</li>
  <li><b>nodeValue: </b>Devuelve o permite acceder al valor del nodo.</li>
  <li><b>nodeType: </b>Devuelve el tipo de nodo.</li>
  <li><b>parentNode: </b>Devuelve el nodo padre del elemento.</li>
  <li><b>parentElement: </b>Funciona igual que el anterior pero devuelve nulo cuando el nodo padre no es un elemento HTML.</li>
  <li><b>style: </b>Retorna un objeto que permite acceder a las reglas de estilo del elemento HTML (declaradas en línea con el atributo style)</li>
  <li><b>tagName: </b>Devuelve el nombre de la etiqueta del elemento HTML.</li>
  <li><b>textContent: </b>Devuelve y premite cambiar el texto del contenido de un elemento HTML. Puede simplificarse a <i>text</i></li>
  <li><b>innerText: </b>Funciona igual al anterior.</li>
  <li><b>firstChild: </b>Devuelve un objeto que permite acceder a las propiedades del primer nodo descendiente del elemento.</li>
  <li><b>lastChild: </b>Igual que el anterior pero en el caso del último descendiente</li>
  <li><b>className: </b>Permite conocer el valor del atributo <i>class</i> del elemento</li>
  <li><b>childNodes: </b>Devuelve una lista con todos los nodos descendientes del elemento</li>
  <li><b>attributes: </b>Devuelve una lista con todos los atributos del elemento.</li>
  <li><b>childElementCount: </b>Devuelve el número total de elementos que son descendientes directos del elemento.</li>
  <li><b>id: </b>Devuelve el valor del atributo <i>id</i> del elemento.</li>
</ul>

Métodos del DOM:
<ul>
 <li><b></b></li>
 <li><b></b></li> 
 <li><b></b></li> 
 <li><b></b></li> 
 <li><b></b></li> 
 <li><b></b></li> 
 <li><b></b></li> 
 <li><b></b></li> 
 <li><b></b></li> 
 <li><b></b></li> 
 <li><b></b></li> 
</ul>














