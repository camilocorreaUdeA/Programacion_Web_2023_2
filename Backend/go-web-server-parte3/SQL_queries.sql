/* Query para crear una tabla */

CREATE TABLE nombre_tabla (
    columna1 tipo_data (INT, BOOL, TEXT.......)
    columna2 tipo_data
    ...
    columna3 tipo_data
);

/* Query para eliminar tabla */

DROP TABLE nombre_tabla;

/* Query para insertar filas (items de datos) */

INSERT INTO nombre_tabla (col1, col2,...., coln) VALUES (valor_col1, valor_col2,...., valor_coln) RETURNING col1;

/* Query para leer columnas (o filas completas) de una tabla */

SELECT col1, col2,...., coln FROM nombre_tabla LIMIT 100 OFFSET 2500; /* lee toda las filas de la tabla */

SELECT col1, col2,...., coln FROM nombre_tabla WHERE condicion; /* lee una o varias filas que cumplen la condicion WHERE */

ORDER BY ASC DESC /* ordenar */
GROUP BY /* agrupar */
LIMIT /* cantidad de resultados para leer */
OFFSET /* decir a partir de que fila leer */

/* Query para actualizar el valor de una o varias columnas de una fila */

UPDATE nombre_tabla SET col1=val1, col3=val3 WHERE condicion;

/* Query para eliminar filas */

DELETE FROM nombre_tabla WHERE condicion;












