# CSS

## Otros selectores, combinaciones y especificidad

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

Cuando hay una regla que se puede (o se debe) aplicar a varios elementos al tiempo pero no hay un selector de los que hemos visto que los agrupe a todos al tiempo en un mismo conjunto, y además en aras de evitar crear reglas duplicadas para los distintos elementos, existe la posibilidad de agrupar varios selectores en la definición de una regla. Solo es especificar los selectores separados por coma (,), la regla luego aplica a todos los selectores que se han especificado.

```css
strong,
em,
.my-class,
[type]{
  background-color:yellow;
}
```

### Pseudo clases y pseudo elementos


