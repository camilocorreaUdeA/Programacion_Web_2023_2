# CSS

## Otros selectores y combinaciones de selectores

### Selector por atributo

En CSS podemos especificar a que elemento están dirigidas las reglas de estilo sacando ventaja de la presencia de un atributo o de su valor.

En la lista a continuación podemos observar las distintas formas que nos permiten utilizar los atributos de los elementos HTML como selectores para las reglas de estilo de CSS:

1. <code>[attribute]</code>: Este selector aplica el estilo en los elementos que tengan un atrbuto nombrado "attribute".
   
```css
    /* Esta regla se aplica a todos los elementos que tengan el atributo "type" */
    [type]{
      color:rgb(28, 123, 98);
    }
```
2. <code>[attribute="value"]</code>: Este selector aplica el estilo en los elementos que tengan un atrbuto nombrado "attribute" cuyo valor exacto sea "value".

```css
   /* Esta regla se aplica a todos los elementos que tengan el atributo "type" con el valor "external-link" */
  [type="external-link"]{
     font-family: Courier;      
  }
```
3. <code>[attribute~="value"]</code>: Este selector aplica el estilo en los elementos que tengan un atrbuto nombrado "attribute" y que alguno de sus valores (valores separados por espacio) sea exactamente "value".

```html
<a href="https://misitioweb.com" type="third-party external-link">Mi Sitio Web</a>
```

```css
   /* Esta regla se aplica a todos los elementos que tengan el atributo "type" y que tenga un valor "external-link" entre el listado de valores del atributo */
  [type~="external-link"]{
     font-family: Courier;      
  }
```

4. <code>[attribute|="value"]</code>: Este selector aplica el estilo en los elementos que tengan un atrbuto nombrado "attribute" y que su valor sea exactamente "value" o que empiece por "value" seguido de guión "value-".

```html
<a href="https://misitioweb.com" type="external-link">Mi Sitio Web</a>
```

```css
   /* Esta regla se aplica a todos los elementos que tengan el atributo "type" y que tenga un valor "external-link" o que empiece por "external-" */
  [type|="external"]{
     font-family: Courier;      
  }
```

5. <code>[attribute^="value"]</code>: Este selector aplica el estilo en los elementos que tengan un atrbuto nombrado "attribute" y que su valor empiece por "value".

```html
<a href="https://misitioweb.com" type="external-link">Mi Sitio Web</a>
```

```css
   /* Esta regla se aplica a todos los elementos que tengan el atributo "type" y que cuyo valor empiece por "external" */
  [type^="external"]{
     font-family: Courier;      
  }
```
6. <code>[attribute$="value"]</code>: Este selector aplica el estilo en los elementos que tengan un atrbuto nombrado "attribute" y que su valor termine en "value".

```html
<a href="https://misitioweb.com" type="external-link">Mi Sitio Web</a>
```

```css
   /* Esta regla se aplica a todos los elementos que tengan el atributo "type" y que cuyo valor termine en "link" */
  [type$="link"]{
     font-family: Courier;      
  }
```
7. <code>[attribute*="value"]</code>: Este selector aplica el estilo en los elementos que tengan un atrbuto nombrado "attribute" y que su valor contenga "value" en cualquier parte.

```html
<a href="https://misitioweb.com" type="third-party-link">Mi Sitio Web</a>
```

```css
   /* Esta regla se aplica a todos los elementos que tengan el atributo "type" y que cuyo valor contenga "-party-" */
  [type*="-party-"]{
     font-family: Courier;      
  }
```
8. <code>[attribute="value" s]</code>: Este selector aplica el estilo en los elementos que tengan un atrbuto nombrado "attribute" y cuyo su valor sea "value". Pero tiene sensibilidad a diferenciar el uso de mayúsculas y minúsculas en el valor del atributo

```html
<a href="https://misitioweb.com" type="third-party-link">Mi Sitio Web</a>
```

```css
   /* Esta regla se aplica a todos los elementos que tengan el atributo "type" y que cuyo valor contenga "Party" pero tomando en cuenta que empieza con mayúscula */
  [type*="Party" s]{
     font-family: Courier;      
  }
```
9. <code>[attribute="value" i]</code>: Este selector aplica el estilo en los elementos que tengan un atrbuto nombrado "attribute" y cuyo su valor sea "value". Ignora la sensibilidad a diferenciar el uso de mayúsculas y minúsculas en el valor del atributo.

```html
<a href="https://misitioweb.com" type="third-Party-link">Mi Sitio Web</a>
```

```css
   /* Esta regla se aplica a todos los elementos que tengan el atributo "type" y que cuyo valor contenga "-party-" ignorando que use mayúsculas o minúsculas */
  [type*="-party-" i]{
     font-family: Courier;      
  }
```

### Agrupación de selectores

Cuando hay una regla que se puede (o se debe) aplicar a varios elementos al tiempo pero no hay un selector de los que hemos visto que los agrupe a todos al tiempo en un mismo conjunto, y además en aras de evitar crear reglas duplicadas para los distintos elementos, existe la posibilidad de agrupar varios selectores en la definición de una regla. Solo es especificar los selectores separados por coma (<code>,</code>), la regla luego aplica a todos los selectores que se han especificado.

```css
strong,
em,
.my-class,
[type]{
  background-color:yellow;
}
```

### Pseudo clases y pseudo elementos

Una pseudo clase es un selector que sirve para identificar elementos que se encuentran en un estado en específico. Ya sea porque cumplen con una condición determinada con respecto a otros elementos del documento HTML, por ejemplo ser el primer elemento <code>li</code> de un elemento <code>ul</code>. O bien, porque se encuentran en un estado resultado de la interacción del usuario con el elemento, por ejemplo cuando se pasa el puntero del mouse sobre un enlace (<code>hoover</code>).

Luego, estas pseudo clases actúan como si se agregara implícitamente una clase que permite identificar de forma más específica al elemento en cuestión aprovechando su estado, resultado de su relación con otros elementos o de la interacción del usuario.

Un selector por pseudo clase se contruye con el nombre del elemento objetivo seguido de la pseudo clase, anteponiento dos puntos (<code>:</code>) al nombre de la pseudo clase.

Un ejemplo por situación o estado del elemento dentro del documento HTML

```html
<ul>
   <li>First element</li>  <!-- Este elemento recibira los efectos de la regla de estilos -->
   <li>Second element</li>
   <li>Third element</li>
</ul>
```
```css
ul li:first-child{
  font-weight: bold;
}
```
Un ejemplo por interacción del usuario con el elemento
```html
<!-- La psedo clase tiene sentido al pasar el puntero del mouse sobre el enlace -->
<a href="https://www.udea.edu.co">Universidad de Antioquia</a> 
```
```css
a:hoover{
  font-weight: bold;
}
```
Puedes ver una lista completa de pseudo clases disponibles en CSS aquí: [Pseudoclases](https://www.w3schools.com/css/css_pseudo_classes.asp)

Los pseudo elementos a diferencia de las pseudo clases no observan del estado actual del elemento para agregarle implicitamente una clase, sino que permiten insertar nuevos elementos HTML a través de CSS. Para construir un selector con pseudo elementos se utiliza doble dos puntos (<code>::</code>) seguido del pseudo elemento que se pretende utilizar.

A continuación la lista de pseudo elementos disponibles en CSS:

1. <code>::after</code>: Inserta contenido después de un elemento. Por ejemplo: <code>p::after</code> inserta contenido después del elemento <code>&lt;p&gt;</code>. Este selector se acostumbra a usar en conjunto con la propiedad <code>content</code> para generar contenido desde CSS.
2. <code>::before</code>: Inserta contenido antes de un elemento. Por ejemplo: <code>p::before</code> inserta contenido antes del elemento <code>&lt;p&gt;</code>. Este selector se acostumbra a usar en conjunto con la propiedad <code>content</code> para generar contenido desde CSS.
```html
<!doctype html>
<html>
  <head>
    <style>
      .texto::before {
        content:"Una línea de texto agregada por CSS antes del contenido original. ";
      }
    </style>
  </head>
  <body>
    <p class="texto">Línea de texto original.</p>
  </body>
</html>
```
3. <code>::first-letter</code>: Selecciona la primera letra de un elemento de texto. Por ejemplo: <code>h1::first-letter</code> selecciona la primera letra del elemento <code>&lt;h1&gt;</code>
4. <code>::first-line</code>: Selecciona el primer renglón de un elemento <code>&lt;p&gt;</code>. Ejemplo: <code>p::first-line</code>
5. <code>::marker</code>: Selecciona las viñetas o marcadores de los elementos <code>&lt;li&gt;</code>. Ejemplo: <code>li::markers</code>
6. <code>::selection</code>: Aplica la regla de estilo CSS al texto seleccionado con el mouse. Ejemplo:
```html
<!doctype html>
<html>
  <head>
    <style>
      ::selection {
        color: rgb(250, 5, 21);
        background-color: rgb(65, 65, 190);
      }
    </style>
  </head>
  <body>
    <h2>Ejemplo del pseudo elemento ::selection</h2>
    <p>Doing the reverse is, of course, also possible using essentially the same technique. That is, marshaling a Go struct into a JSON array is done by defining a MarshalJSON method.</p>
  </body>
</html>
```
### Combinaciones

Selector de descendientes (<i>descendant combinator</i>)

Este selector se construye como la combinación de dos selectores separados por espacio. La regla de estilos se aplica al último selector de la combinación siempre y cuando se cumpla que el otro selector de la combinación es su ancestros: ancestro -> descendiente.

```html
<ul>
   <li>First element</li>
      <ol>
         <li>first element</li>
         <li>second element</li>
      </ol>
   <li>Second element</li>
   <li>Third element</li>
</ul>
```
```css
/* selector de descendientes*/
ul li{
  border: 2px solid blue;
}
```

Selector de descendientes directos o hijos (child combinator)

Funciona exactamente igual al anterior con la diferencia de que solo se seleccionan los descendientes directos de los ancestros y se omiten descendientes de menor nivel. Para combinar los selectores se utiliza el símbolo <code>&gt;</code>.

```html
<ul>
   <li>First element</li>
      <ol>
         <li>first element</li>
         <li>second element</li>
      </ol>
   <li>Second element</li>
   <li>Third element</li>
</ul>
```
```css
/* selector de descendientes directos*/
ul > li{
  border: 2px solid blue;
}
```
Selector de hermano contiguo o adyacente (next/adjacent sibling combinator) 

Este selector combina dos selector y aplica la regla de estilos al segundo selector siempre y cuando este precedido por un elemento que cumpla con el primer selector. El símbolo para combinar los selectores es <code>+</code> en este caso.

```html
<article>
    <h1>El encabezado</h1>
    <p>Primer párrafo del artículo</p>
    <p>Segundo párrafo del artículo</p>
</article>
```
```css
h1+p {
    font-weight: bold;
    background-color: #333;
    color: #fff;
    padding: .5em;
}  
```
Selector de hermanos en general (general sibling combinator)

Funciona de la misma manera que el selector anterior pero no se limita solamente al elemento "hermano" contiguo o adyacente sino a todos los demás que pueden no necesariamente ser contiguos. El simbolo utilizado para este selector es <code>~</code>.

```html
<article>
    <h1>El encabezado</h1>
    <p>Primer párrafo del artículo</p>
    <a href="https://udearroba.udea.edu.co">Ude@</a>
    <p>Segundo párrafo del artículo</p>
</article>
```
```css
h1~p {
    font-weight: bold;
    background-color: #333;
    color: #fff;
    padding: .5em;
}  
```
Selectores compuestos

CSS permite crear selectores que resultan de hacer combinaciones de otros selectores, esto con el único objetivo de crear reglas más específicas, es decir, dirigidas a elementos en particular dentro del documento HTML. Para crear estos selectores compuestos se combinan o encadenan dos o más selectores omitiendo el uso de cualquier símbolo como los que acabamos de ver para los otros selectores combinados.

Por ejemplo si quisieramos crear una regla específica para hipervínculos de un sitio web que tienen el mismo valor en su atributo <code>class</code>.

```html
<p class="udea">El himno de la Universidad de Antioquia tiene letra del poeta antioqueño Edgar Poe Restrepo y música adaptada por el maestro José María Bravo Márquez...</p>
<a href="https://www.udea.edu.co" class="udea">Sitio web Universidad de Antioquia</a>
<a href="https://udearroba.udea.edu.co" class="udea">Ude@</a>
<a href="https://ingenieria.udea.edu.co">Faculta de Ingeniería</a>
```
```css
a.udea {
    font-weight: bold;
    background-color: white;
    color: rgb(28, 100, 2);
    padding: .5em;
}

a.udea:hover{
    font-weight: bold;
    background-color: rgb(28, 100, 2);
    color: #fff;
    padding: .5em;
}    
```

## La cascada de CSS

### Especificidad

### Herencia

### Orden de las reglas



