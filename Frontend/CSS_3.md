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

En el ejemplo vemos 2 elementos de bloque: <code>h1</code> y <code>p</code>, podemos observar lo que se había comentado anteriormente de que ocupan todo el ancho del elemento contenedor (en este caso un elemento <code>div</code>), están apilados uno sobre otro de acuerdo con el orden de aparición en el documento HTML. Además se puede observar la separación entre los dos elementos (para esto hemos resaltado el border) como consecuencia de que por defecto los elemento de bloque tienen la propiedad margin.

![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/647c2f8e-9ed9-4456-9f5d-fe0c3d7a27f6)
Se puede observar en la imagen que el margin por defecto asignado al elemento <code>h1</code> es de 46.9px y solo en las direcciones arriba y abajo.
![image](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/318f88d3-58b3-489d-bfa1-ada25707ec65)
Mientras que el margin por defecto del elemento <code>p</code> es de 25px y también solo en las direcciones arriba y abajo.

Otro detalle interesante que se destaca es como el margin del elemento <code>h1</code> llega hasta el límite del borde del elemento <code>p</code>, es decir la separación real entre ambos elementos son 46.9px y no una suma de los margin de los dos elementos como se podría llegar a pensar. A ese comportamiento en la separación de dos elementos de bloque contiguos se le conoce como <i>margin collapse</i> o colapso de márgenes, y es simplemente que el margin con mayor dimensión absorbe al de menor dimensión para convertirse en la separación efectiva entre los dos elementos de bloque.





