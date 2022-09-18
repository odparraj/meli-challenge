# Meli Challenge

## Corriendo la aplicacion
Este comando creará los servicios necesarios para iniciar la app:
```bash
# start
$ docker-compose -f docker-compose.yml up -d
```
Se puede verificar la ejecución de los containers (**meli-web**, **meli-mysql**) en el Docker Desktop dashboard.

## Deteniendo la aplicación
Para detener la aplicacion simplemente ejecutamos el siguiente comando:
```bash
# stop
$ docker-compose -f docker-compose.yml down
```

## Importando archivos
Nuestro servicio de importacion de archivos asume que los datos estan guardados previamente en el sistema de almacenamiento, por ende debemos copiar dentro de la carpeta **storage/**: los archivos que queremos importar, en la carpeta en cuestion ya tenemos 3 archivos de prueba, el endpoint al que debemos apuntar es:
<br>
<br>
<br>

|Endpoint|http://127.0.0.1:8000/proccess-file|
|-|-|
|Method|POST|
|Body:||
|file|nombre del archivo a importar: required,string|
|decoder|decodificación a usar opciones: jsonln , txt , csv|
|deimiter| aplica para los decoders txt y csv, es opcional|

<br>
<br>

### *Ejemplo csv*
```json
{
    "file": "technical_challenge_data.csv",
    "decoder" : "csv",
    "delimiter" : ","
}
```
### *Ejemplo txt*

```json
{
    "file": "technical_challenge_data.txt",
    "decoder" : "text",
    "delimiter" : ";"
}
```
### *Ejemplo jsonln*
```json
{
    "file": "technical_challenge_data.jsonln",
    "decoder" : "jsonln"
}
```
## Ejemplo de resuesta

Nuestro servicio procesa el archivo con una premisa no bloqueante, es decir si una linea falla se debe continuar a la siguiente sin detener el proceso, la respuesta contiene dos claves:
**errors** un array de string con la descripcion de cada error encontrado
**message** un mensaje personalizado, por defecto contiene un mensaje para describir el tiempo de ejecución:

### *Ejemplo de ejecución sin errores encontrados*

```json
{
    "errors": [],
    "message": "execution time: 5.39 seconds"
}
```
### *Ejemplo de ejecución con errores encontrados en algunas lineas*

```json
{
    "errors": [
        "error proccessing line 3, message: item not found [MLA693105237]",
        "error proccessing line 5, message: item not found [MLA655915616]",
        "error proccessing line 99, message: item not found [MLA711759823]",
        "error proccessing line 26, message: item not found [MLA723145127]",
        "error proccessing line 53, message: item not found [MLA727165201]",
        "error proccessing line 43, message: item not found [MLB1284334463]"
    ],
    "message": "execution time: 5.39 seconds"
}
```
### *Ejemplo de ejecución con un nombre de archivo no existente*
```json
{
    "errors": [
        "open storage/technical_challenge_data.csvs: no such file or directory"
    ],
    "message": "fails reading file"
}
```

## Respuestas Al Desafío Teórico

### [Ir a las respuestas](https://github.com/odparraj/meli-challenge/blob/master/theoretical-challenge.md)