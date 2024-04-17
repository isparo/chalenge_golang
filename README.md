# chalenge_golang

## Preguntas:

Responde las siguientes preguntas:
1.- ¿Qué son “shorts variables”?
2.- ¿Cuando puedes utilizar “short variables” y cuando no?
3.- ¿Qué significa inferencia de tipos de datos?
4.- ¿Puede una constante declararse de manera corta como son “short variables”?
5.- ¿Qué es “defer”?
6.- ¿Qué son los “pointer”?
7 .- ¿Qué es “struct”?
     - Es un tipo de dato que 
8.- ¿Qué es “goroutine”?
    - Una goroutine nos ayuda a generar procesos que corrent en un thread 



## Ejercicio:
Realiza el siguiente ejercicio, comparte la liga del repositorio de tu prueba.
1. Crear un servicio de registro de usuario que reciba como parámetros usuario, correo,
teléfono y contraseña. - listo

2. El servicio deberá validar que el correo y telefono no se encuentren registrados, de lo
contrario deberá retornar un mensaje “el correo/telefono ya se encuentra registrado”.

3. Deberá validar que la contraseña sea de 6 caracteres mínimo y 12 máximo y contener
al menos una mayúscula, una minúscula, un carácter especial (@ $ o &) y un número.
4. Validar que el teléfono sea a 10 dígitos y el correo tenga un formato válido.


5. Crear un servicio login que reciba como parámetros usuario o correo y contraseña.
6. El servicio debe devolver un token jwt.
7. Deberá validar que el usuario o correo y contraseña sean válidos, de lo contrario
retorna un mensaje “usuario / contraseña incorrectos”.
1. En ambos servicios se deberá validar que todos los parámetros solicitados vayan en el
cuerpo de la petición, de lo contrario retorna un mensaje con el campo faltante.

Ejemplo login:
{
“correo”: “prueba@gmail.com”
}

La respuesta debe ser:
“Falta el campo contraseña”
Notas: Se consideran buenas prácticas, documentación en código, pruebas unitarias y
arquitectura.

----------
## Unit testing
-- For this example I'm using gomock in order to generate the mock files

-- I'm using test tables to define the unit test

## Generate swagger
swag init -d cmd,internal --parseDependency --parseInternal --parseDepth 2


## Go to API documentation (swagger):
http://localhost:8080/swagger/index.html

## How to run it:

For this exaple I'm using `docker-compose`in order to configure a local deployment that going to run
the `user` service and the local database.

start service and database:
`docker-compose up`

stop service and database:
`docker-compose down`

