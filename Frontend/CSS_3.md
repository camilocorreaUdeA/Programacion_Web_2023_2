# CSS

## Maquetación de los elementos en el documento HTML (CSS Layout)

### Flujo normal del documento HTML

Por defecto y a menos de que hayas introducido manualmente algún cambio a la maquetación del documento HTML usando CSS los elementos de un sitio web se maquetan o posicionan utilizando el flujo normal o flujo por defecto del navegador. Este flujo normal asegura que el contenido de la página web sea <i>leible</i> para cualquier navegador web o dispositivo de navegación asistida, como un lector de pantalla.

La pregunta ahora es ¿Cómo posiciona los elementos en pantalla el flujo normal? Y la respuesta está en un concepto que ya vimos con anterioridad que es el modelo de cajas. Es decir, los elementos individuales de un documento HTML se posicionan de tal forma que se tenga en estricta consideración las propiedades del modelo de cajas (padding, border, margin) que tenga el elemento HTML asociadas a su contenido.

Ahora, primero que todo debemos hacer una distinción entre los elementos HTML separandolos en dos categorías a saber:

<ul>
  <li>Elementos de bloque (<i>Block-level elements</i>)</li>
  <li>Elementos de línea (<i>Inline-level elements</i>)</li>
</ul>

Los elementos de bloque se ubican en pantalla en una línea exclusiva ocuapndo por completo el ancho de línea del elemento padre (si el elemento padre es el body, entonces será todo el ancho de pantalla), y se apilan uno por debajo del otro de manera sucesiva y de acuerdo al orden de aparición en el documento HTML.

```html
<body>
<div>
  <h1>Aplicaciones Web</h1>
  <p>She was infatuated with color. She didn't have a favorite color per se, but she did have a fondness for teals and sea greens. 
  You could see it in the clothes she wore that color was an important part of her overall style. She took great pride that color 
  flowed from her and that color was always all around her. That is why, she explained to her date sitting across the table, that 
  she could never have a serious relationship with him due to the fact that he was colorblind.</p>
</div>
</body>
```
```css
h1{
  font-size:70px;
  background-color:lightgreen;
  border:2px solid black;
}
p{
  font-size:25px;
  background-color:lightblue;
  border:2px solid black;
}
```
![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/66cd900c-2cf4-43e3-b97e-789bf5ae0433)

En el ejemplo vemos 2 elementos de bloque: <code>h1</code> y <code>p</code>, podemos observar lo que se había comentado anteriormente, que ocupan todo el ancho del elemento contenedor (en este caso un elemento <code>div</code>), están apilados uno sobre otro de acuerdo con el orden de aparición en el documento HTML y su dimensión vertical se ajusta al tamaño del contenido (observe que el párrafo tiene más dimensión vertical que el encabezado). Además se puede observar la separación entre los dos elementos (para esto hemos resaltado el border) como consecuencia de que por defecto los elemento de bloque tienen la propiedad margin.

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/647c2f8e-9ed9-4456-9f5d-fe0c3d7a27f6)
<p>Se puede observar en la imagen que el margin por defecto asignado al elemento <code>h1</code> es de 46.9px y solo en las direcciones arriba y abajo.</p>

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/318f88d3-58b3-489d-bfa1-ada25707ec65)
<p>Mientras que el margin por defecto del elemento <code>p</code> es de 25px y también solo en las direcciones arriba y abajo.</p>

Otro detalle interesante que se destaca es como el margin del elemento <code>h1</code> llega hasta el límite del borde del elemento <code>p</code>, es decir la separación real entre ambos elementos son 46.9px y no una suma de los margin de los dos elementos como se podría llegar a pensar. A ese comportamiento en la separación de dos elementos de bloque contiguos se le conoce como <i>margin collapse</i> o colapso de márgenes, y es simplemente que el margin con mayor dimensión absorbe al de menor dimensión para convertirse en la separación efectiva entre los dos elementos de bloque.

Por otra parte, los elementos de línea no son apilados uno tras otro en una nueva línea de la pantalla para cada uno, por el contrario son ubicados uno tras otro en la misma línea hasta que llenan todo el ancho del elemento contenedor y eventualmente pasan a una nueva línea cuando sobrepasan dicho ancho. Por defecto las dimensiones de un elemento de línea coinciden con las dimensiones del contenido del mismo, aún así es posible modificarlas a través de CSS (cuando hacer eso tiene algún sentido).

```html
<body>
  <div>
  <a href="#">Hola</a>
  <a href="#">Mundo</a>
  <a href="#">Aplicaciones</a>
  <a href="#">Web</a>
  <a href="#">Telecomunicaciones</a>
  <strong>she explained to her date sitting across the table</strong>
  </div>
</body>
```
```css
strong{
  font-size:25px;
  border:2px solid black;
  background-color:grey;
}

a{
  font-size:25px;
  border:2px solid black;
  background-color:lightpink;
}
```
![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/abed4d6c-0a19-4c23-a3ed-1161bee594ae)

En la imagen a continuación se observa como los elementos de línea van llenando paulatinamente el ancho del elemento contenedor.

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/31a0f06f-f0b1-4c00-90b3-d08efb75e251)

En las siguientes imagenes se puede observar que no tienen margenes asignadas por defecto (ni padding ni margin) y que el tamaño del elemento se ajusta al del contenido.

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/144b83ca-62c8-4147-9bb3-15fe2fb6527b)

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/120407cd-4a31-4604-bd0e-e0a03650e2c3)

En este [enlace](https://www.w3schools.com/html/html_blocks.asp) puede ver listas separadas de los elementos de bloque y de línea de HTML.

### Contenedores <code>div</code> y <code>span</code>

<code>div</code>

<p>El elemento <code>div</code> es un elemento de bloque que actúa como contenedor genérico para otros elementos. A diferencia de los elementos semánticos de HTML <code>div</code> no expresa el significado de su contenido y en parte se debe a que su función es simplemente la de ser un elemento auxiliar que permita dividir la página web en distintos bloques que pueden ser estilizados con CSS individualmente.</p>

```html
<body>
  <div class="bloque-1"></div>
  <div class="bloque-2"></div>
  <div class="bloque-3"></div>
</body>
```
```css
.bloque-1{
  background-color:#4061c4;
  height:200px;
  border:2px solid black;
}

.bloque-2{
  background-color:rgba(251, 203, 80, 0.54);
  height:200px;
  border:2px solid black;
}

.bloque-3{
  background-color:rgba(180, 45, 4, 0.52);
  height:200px;
  border:2px solid black;
}
```
![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/c3765efe-0d88-423a-943a-071785578a6f)

<code>span</code>

<p>El elemento <code>span</code> también es un contenedor genérico para agrupar otros elementos y estilizarlos. Pero <code>span</code></p> es un elemento de línea, por tanto es más útil para agrupar textos u otros elementos de línea como los hipervínculos (elemento <code>a</code>).</p>

```html
<body>
  <p>She was infatuated with color. She didn't have a favorite color per se, but she did have a fondness for teals and sea greens. You could see it in the clothes she wore that color was an important part of her overall style. She took great pride that <span id="text">color flowed from her and that color was always all around her.</span> That is why, she explained to her date sitting across the table, that she could never have a serious relationship with him due to the fact that he was colorblind.</p>
</body>
```
```css
#text:hover{
  background-color:lightpink;
  border:2px solid black;
  font-size:35px;
  font-style:italic;
}
```

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/d100349b-cb99-4afa-a645-0a3cf2a69c9a)

### Propiedad <code>display</code>

La propiedad <code>display</code> sirve para determinar si un elemento es de bloque o de línea. Modificando esta propiedad se puede lograr que un elemento que por naturaleza es de bloque se convierta en un elemento de línea y vicerversa, y además nos va a ser también de utilidad para construir elementos de cuadrícula (<i>grid</i>) y elementos de caja flexible (<i>flex</i>).

Esta propiedad también puede tener el valor <i>inline-block</i> que permite que los elementos se comporten como elementos de línea a los que se les puede asignar dimensiones de ancho y alto (<i>width</i> y <i>height</i>). Recuerde que los elementos de línea ignoran esas propiedades porque ajustan sus dimensiones a las del contenido.

Posibles valores para la propiedad <code>display</code>:
<ul>
  <li><code>block</code>:</li>
  <li><code>inline</code>:</li>
  <li><code>inline-block</code>:</li>
  <li><code>grid</code>:</li>
  <li><code>flex</code>:</li>
</ul>

### Grid







