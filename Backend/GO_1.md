# Go
![2](https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/e131e81c-9a0d-4db2-8d40-99119ebe006e)
## El lenguaje de programación Go (Golang)

Go inició como un proyecto de bajo perfil que rapidamente escalo hasta convertirse en un proyecto ambicioso patrocinado por Google. A finales de 2011 se liberó Go como un proyecto de código abierto y luego el 28 de Marzo de 2012 se lanzó la primera versión estable del lenguaje y a partir de ese momento se procura liberar una nueva versión cada medio año.

"...El objetivo del proyecto Go era eliminar la lentitud y la falta de coordinación en el proceso de desarrollo de software en Google, y por tanto hacer el proceso más productivo y escalable. El lenguaje fue diseñado por y para gente que escribe, lee, depura y mantiene grandes proyectos de software..." -Rob Pike-

Al crear el lenguaje de programación Go la idea principal era combinar las mejores características que están presentes en otros lenguajes, entre ellas:
<ul>
  <li>Fácil de integrar con las herramientas actuales para disparar la productividad</li>
  <li>Alta eficiencia y tipado estático</li>
  <li>Alto desempeño para aplicaciones de networking y un aprovechamiento total de las arquitecturas multi-core</li>
</ul>

Go es un lenguaje especialmente apto para las siguientes aplicaciones y sistemas:
<ul>
  <li>Servicios distribuidos interconectados: Las aplicaciones de red dependen fuertemente de la concurrencia, la cual es manejada de forma simple y eficiente por característcas nativas del lenguaje como los canales y las go-rutinas</li>
  <li>Desarrollo nativo en la nube: Las apliaciones nativas de la nube son más fáciles de construir con las características que Go ofrece para concurrencia y networking, además de la portabilidad entre sistemas operativos y arquitecturas que ofrece de forma nativa el lenguaje.</li>
  <li>Remplazo para infraestructura existente: Go es un lenguaje ideal para re-escribir aplicaciones, permite mejorar la seguridad en el manejo de memoria, permite el despliegue multi-plataforma y una base de código limpia que es de más fácil mantenimiento en el futuro.</li>
  <li>Utilidades y herramientas independientes: Los programas en Go complilan en binarios que practicamente no requieren de dependencias externas. Esto hace a Go ideal para la creación de herramientas y utilidades que se pueden poner en producción con prontitud y que pueden ser empaquetadas para su redistribución.</li>
</ul>

Tal vez Go no es lo sufiecientemente apto en los siguientes casos:
<ul>
  <li>Interfaces gráficas o GUIs (aunque hay paquetes como Fyne y Wails que tienen un desempeño de aceptable a bueno)</li>
  <li>Desarrollo de sistemas embebidos y desarrollo de bajo nivel en general: Existe un branch de Go conocido como TinyGO que permite programar dispositivos de hardware. Go no está diseñado para la creación de drivers o interfaces de propósito general de entrada/salida de datos (comunicaión con periféricos y hardware externo)</li>
</ul>

Ventajas de Go con respecto a otros lenguajes de programación:
<ul>
<li>Fácil de usar y de leer: La sintaxis de Go es simple, entre 25 y 35 palabras clave conforman el lenguaje. Y además su curva de aprendizaje lo hace bastante accesible a programadores novatos.</li>
<li>Amplia librería estándar: Una amplia y variada librería estándar viene empaquetada en el lenguaje, lo cual evita el uso de librerías de terceros para similares propósitos. La librería es sofisticada pero no compleja ni complicada, y el manejo de dependencias de Go evita riesgo de conflictos entre nombres de fiunciones.</li>
<li>Seguridad: Gracias a lo simple del código, a que es un lenguaje de tipado estático y al garbage collector incluido en el runtime del lenguaje.</li>
<li>Soporte: Además del soporte de ingenieros de Google cuenta con una siempre creciente coumunidad.</li>
<li>Concurrencia y paralelismo: Incluidos en el lenguaje en librerías con funciones para estos propósitos.</li>
<li>Manejo de dependencias: El sistema de manejo de dependencias está incluido en el lenguaje facilitando la inclusión de paquetes de terceros y el posterior empaquetamiento en el binario de la aplicación.</li>  
</ul>

Utilidades de Go para los desarrolladores:
<ul>
  <li>Generación automática de documentacipon del código.</li>
  <li>Herramientas de testing y benchmarking nativas incluídas en el lenguaje.</li>
  <li>Detección de condiciones de carrera para apliaciones concurrentes.</li>
  <li>Un sistema de tipos liviano pero suficiente para clasificar y diferenciar datos y objetos.</li>
  <li>Concurrencia nativa para la ejecución de tareas en segundo plano.</li>
  <li>Garbage collector para el uso racional de memoria.</li>
  <li>Compilación multiplaforma que permite desplegar en cualquier arquitectura y sistema operativo.</li>
  <li>Buen desempeño y velocidad en comparación con otros lenguajes de programación similares.</li>
  <li>Amplia librería estándar, especializada en funciones de concurrencia y networking.</li>
  <li>Punteros para manejo eficiente de memoria.</li>
</ul>

## Instalación de Go

Puedes revisar el [Getting started](https://go.dev/doc/tutorial/getting-started) oficial.
También puedes seguir las instrucciones oficiales de [instalación](https://go.dev/doc/install)

En general el proceso de instalación de Go es bastante intuitivo:
<ul>
  <li>Instala el IDE de tu preferencia, recomiendo VSCode pero Goland también está bien.</li>
  <li>Descarga Go [aquí](https://go.dev/dl/). Allí están los instaladores para los distintos sistemas operativos.</li>
  <li>Instalación manual de Go en tu máquina:
    <ul>
      <li>Linux: Ejecute en la terminal: <code>rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.3.linux-amd64.tar.gz</code>, luego setea la variable de ambiente <code>PATH</code> para que también apunte al directorio bin de Go: <code>export PATH=$PATH:/usr/local/go/bin</code></li>
      <li>Con MacOS y Windows descargue el instalador correspondiente, ejecutelo y verifique la instalación con el comando <code>go version</code> en la terminal.</li>
    </ul>
  </li>
  <li>Otras formas de instalar Go:
  <ul>
    <li>MacOS: Instale <b>homebrew</b> en su sistema y una vez instalado ejecute en la terminal el comando <code>brew install go</code></li>
    <li>Windows: Instale <b>chocolatey</b> y una vez instalado ejecute en la terminal el comando <code>choco install golang</code></li>
  </ul>
    </li>
  <li>Verifique la instalación ejecutando el comando <code>go version</code></li>
</ul>

Una vez hayas instalado Go en tu máquina y si ya tienes VSCode instalado también (¡si no, entonces instalalo!). Abre VSCode y luego presiona la combinación de teclas <code>Ctrl + Shif + P</code> y en el menú que se despliega escoge la opción <code>Go: Install/Update Tools</code>. Espere que se complete la instalación de las herramientas.

<img width="603" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/8a9b7c5d-b8b5-4dd4-82a0-c7926d2bd2c0">

Luego instale la extensión de Go para facilitar la codificación ya que trae funciones como autocompletado, indentación automática y verificación de errores de sintaxis, etc.

<img width="655" alt="image" src="https://github.com/camilocorreaUdeA/Programacion_Web_2023_2/assets/42076547/aa49bf57-b453-45e1-a61f-adb722405ca3">

## Go modules

<i>Go modules</i> es el sistema de administración de dependencias en Go, el cual facilita el acceso a la información de versiones y la inclusión de paquetes externos en una aplicación de Go.

Un modulo en Go es una colección de paquetes almacenados en un árbol de archivos (directorio) donde un archivo llamado <i>go.mod</i> es el nodo raíz del árbol.

El archivo <i>go.mod</i> define la ruta del modulo que se usa como la ruta de <i>import</i> del directorio raíz. También indica la versión de Go usada para compilar el proyecto y el conjunto de dependencias (otros modulos) que se importan y se utilizan en el proyecto.

## Creación de un proyecto en Go




