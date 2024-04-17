# chalenge_golang

## Preguntas:

## Responde las siguientes preguntas:

1.- ¿Qué son “shorts variables”?

- Es una forma de declarar las variables en la cual no se declara su tipo y el compilador lo infiere por medio del valor, ejemplo:
     
     myVar1 := "the value"
     myVar2 := 0
     myVar := false
     
2.- ¿Cuando puedes utilizar “short variables” y cuando no?

   - Para declarar y asignar un valor que ya sepamos que va a tener la variable
   - En for loops para definir las condiciones `for i := 0; i <5; i++ {}` 
   - cuando tenemos un bloque var() y vamos para declarar variables esto nos ayuda a hacerlo mas legible
       
            var(
               myVar1 = "the value"
               myVar2 = 0
               myVar = false
            )
       
   - Cuando usamos variables locales ya sea dentro de condiciones, loops, funciones, goroutines, etc.

3.- ¿Qué significa inferencia de tipos de datos?

 - Es cuando el compilador en este caso Golang, se hace el uso de las short variables para inferir el tipo de la variable por medio del valor, ejemplo:

          myVar1 = "the value"  --- string
          myVar2 = 0  ---- int
          myVar = false ---- bool
     
4.- ¿Puede una constante declararse de manera corta como son “short variables”?

- No como tal ya que no pdríamos hacer `const myConst := "val"` pero si podemos definirla `const myConst = "val"` lo que hace que el compilador infiera el tipo de dato.

5.- ¿Qué es “defer”?

   - Es una instruccion que nos ayuda a ejecutar alguna logica cuando la funcion haya terminado, ejemplo:

     func myFunc() {
          fmt.Println("esto")

          defer func(){
            fmt.Println("esto al finalizar")
          }()

          fmt.Println("esto tambien")
     }

6.- ¿Qué son los “pointer”?

   - Son variables que almacenan el valor en una direccion de memoria, a diferencia de una variable sin puntero que almacena el valor, cuando se usa un puntero este asigna el valor a un espacio de memoria.
   - Los punteros se definen con el uso de asterisco *, ejemplo: `var myvar *int` para asignarle una referencia se hace usando Ampersand & `myPointer = &variable`.
   - Para obtener el valor lo podemos hacer por medio del *, `fmt.Println("something: ", *myPointer)`

    ejemplo:
     
    myValue := 12
    var myPointer *int = &myValue

    func (s *UserServce) Something() {....}
    func doMore(v *int) { ..... }
    doMore(&myPointer)


 - Los punteros tambien nos ayudan a pasar valores por referencia en las funciones ya que por defecto estas son majadas por valor. Esto es util cuando queremos modificar el valor de alguna propiedad de una structura, etc.

7 .- ¿Qué es “struct”?

  - Es un tipo de dato que nos ayuda a agrupar caracteristicas o propiedades que perteneces a una entidad, por ejemplo Usuario, podemos hacer uso de 
    una struct para definir sus caracteristicas.

   - Aparte de poderse usar para definir entidades, tambien puede ser usada para definir los valores de configuracion usados por un servicio, ejemplo:
   
          type Config struct {
              Host string
              Port string
              Envvar1 string
          }
     
8.- ¿Qué es “goroutine”?

   - Es una funcion que se ejecuta de forma paralela al proceso principal, esta puede ser usada para concurrencia y paralelismo en nuestars applicaciones.



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
-- Para este ejemplo estoy usado herramientas para mockear las interfaces usadas en los 
   diferentes files.

-- Hago el uso de mocks ya que he implementado dependency injection, esto nas facilita hacer test
   de una forma mas facil sin depender de la implementacion concreta.

-- Para definir los unittest hago uso de test tables en las que defino cada uno de los casos.

## Go to API documentation (swagger):
-- Para documentar el API hago uso de swaggo
   (swag init -d cmd,internal --parseDependency --parseInternal --parseDepth 2)

-- swagger esta disponible en esta URL: http://localhost:8080/swagger/index.html

## How to run it:

Para este ejemplo hago uso de `docker-compose` para levantar el servicio y la base de datos local
esto para no depender de un servicio externo de bases de datos.

levantar el service y la database:
`docker-compose up`

detener el service and database:
`docker-compose down`

