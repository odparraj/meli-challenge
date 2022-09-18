# Respuestas Al Desafío Teórico

## Procesos, hilos y corrutinas

### Un caso en el que usarías procesos para resolver un problema y por qué.

-  **Rta:** Los procesos son usados para tareas intensivas de procesamiento de datos que son aisladas, se pueden ejecutar en paralelo. Analizar grandes volúmenes de datos rápidamente puede ser un ejemplo. Por lo general las tareas en procesos necesitan CPU para realizar cálculos la mayor parte del tiempo.

### Un caso en el que usarías threads para resolver un problema y por qué.

-  **Rta:** los threads son un punto intermedio entre los procesos y las corrutinas, lo usaría para programas que interactúen con la UI, de esta manera se evita el bloqueo de la interfaz

### Un caso en el que usarías corrutinas para resolver un problema y por qué.


-  **Rta:** las corrutinas se utilizas para tareas intensivas de E/S, una tarea normal pasa la mayor parte del tiempo esperando de recursos como disco, red, RAM. Un caso de uso podria ser realizar muchas consultas de red sin neecesidad de procesar los datos

## Optimización de recursos del sistema operativo
### Si tuvieras 1.000.000 de elementos y tuvieras que consultar para cada uno de ellos información en una API HTTP. ¿Cómo lo harías? Explicar.

- **Rta:** Realizaría la tarea utilizando corrutinas porque se requieren muchas operaciones de E/S, en este caso de red, el programa casi no requiere uso de CPU

## Análisis de complejidad

### Dados 4 algoritmos A, B, C y D que cumplen la misma funcionalidad, con complejidades O(n^2), O(n^3), O(2^n) y O(n log n), respectivamente, ¿Cuál de los algoritmos favorecerías y cuál descartarías en principio? Explicar por qué.


- ### **Rta:** Favorecería el algoritmo de complejidad O(n log n) porque es mas eficiente, descartaria el O(2^n) porque es el menos eficiente, la complejidad se puede expresar como el numero de operaciones requeridas para ejecutar el algoritmo

### Asume que dispones de dos bases de datos para utilizar en diferentes problemas a resolver. La primera llamada AlfaDB tiene una complejidad de O(1) en consulta y O(n2) en escritura. La segunda llamada BetaDB que tiene una complejidad de O(log n) tanto para consulta, como para escritura. ¿Describe en forma sucinta, qué casos de uso podrías atacar con cada una?

- **AlfaDB** sería optimo para operaciones de solo lectura (casos de uso donde  tenemos multiples clientes de lectura y pocas ocasiones de actualización de datos ej: sistemas de reportes que se actualiza mensualmente)
- **BetaDB** es optimo para operaciones de lectura y escritura (casos de uso donde los clientes requieren leer y escribir datos de forma frecuent ej: E-Commerce)