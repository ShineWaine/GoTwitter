package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Creación de una variable tipo puntero "*", que recogerá todo
// lo que se entre por teclado
var reader *bufio.Reader

// Creación de la estructura User, que contiene los siguientes
// atributos.
type User struct {
	id       int
	username string
	email    string
	age      int
}

// definimos una variable del tipo mapa que almacenará los datos del
// usuario a través de su id y por supuesto será del tipo User.
var users map[int]User

// Variable entera para almacenar el ID y usarlo como indice de búsqueda
var id int

// Función de limpieza de terminal
func clearConsole() {
	// Genera una variable con la carga del comando
	cmd := exec.Command("clear")

	// Asigna un método de salida del comando
	cmd.Stdout = os.Stdout

	// Ejecuta el comando propiamente
	cmd.Run()
}

// Función de creación/actualización de datos
// Se le suministra un Id, ya sea por el contador de creación
// o por el seleccionado por usuario para borrar o actualizar
func insertaDatos(id int) {
	// Funcion de limpiar el terminal
	clearConsole()

	// Inserción de los datos de usuario a traves de la función
	// de entrada de datos 'readLine()'
	fmt.Println("Introduce un Username: ")
	username := readLine()
	fmt.Println("Introduce un e-mail: ")
	email := readLine()
	fmt.Println("Introduce la edad del usuario: ")
	age, err := strconv.Atoi(readLine())

	//Validación del casting string to int
	if err != nil {
		panic("Problemas de conversión de tipo en el campo edad.")
	}

	// Inserción de datos en el map, ya sea por creación como por
	// actualización
	user := User{id, username, email, age}
	users[id] = user
}

func crearUsuarios() {
	// Incremento de id para nueva inserción de usuario
	id++

	// Llamada al proceso de inserción de datos, suministrando
	// id para nuevo usuario
	insertaDatos(id)

	// Aviso de inserción completada
	fmt.Println("Creación de nuevo Usuario COMPLETADA.")
}

// Función que lista a todos los usuarios activos
func lista() {
	// Limpia pantalla terminal
	clearConsole()

	// Recorre el mapa para listar a todos los usuarios
	for id, user := range users {
		fmt.Println(id, " - ", user.username)
	}
}

// Función de Opción B del menu. Visualización de usuarios
func listarUsuarios() {
	lista()
	fmt.Println("\n Listado de Usuarios COMPLETO")
}

// Función de inserción de ID. Selección para Actualización/borrado
func insertaId() int {
	//Bucle de control para asegurar inserción correcta de ID
	for {
		id, err := strconv.Atoi(readLine())
		if err != nil {
			panic("Error en la conversión del numero ID.")
		}
		if _, ok := users[id]; ok {
			return id
		}
		fmt.Println("El ID proporcionado no es correcto. Introduzca ID correcto.")
	}
}

// Funcion de Opción C del menu. Actualización de datos de usuario
func actualizarUsuarios() {
	// Llama a la función que muestra todos los usuarios
	// existentes para poder escoger uno.
	lista()

	// Solicita la introducción del número de id a actualizar
	fmt.Println("\nIntroduce el ID del usuario a modificar:")
	id := insertaId()

	// Llama a la función de entrada de datos nuevamente.
	insertaDatos(id)
	fmt.Println("Inserción de ID realizada. Datos actualizados")
}

// Función de Opción D del menu. Borrado de usuario
func eliminarUsuarios() {
	// Llamada a lafunción de borrado de pantalla
	clearConsole()

	// Llama a la función para visualizar los usuarios existentes
	lista()

	// Solicita la inserción del ID de usuario a borrar
	fmt.Println("Introduce el número de ID de usuario a borrar: ")
	id := insertaId()

	//Verificada por la función la validez del ID, procede al borrado
	delete(users, id)
	fmt.Println("Proceso de borrado CONCLUIDO")
}

// Funcion de entrada de datos generica.
func readLine() string {
	//Entrada valores por teclado, indicando el último caracter que
	//finalizala entrada (tambien incluido)
	//opcion, err := reader.ReadString('\n') **opcion implementada en el if.
	if opcion, err := reader.ReadString('\n'); err != nil {
		panic("No es posible la entrada por teclado.")
	} else {
		//Esta linea elimina el "\n" final de la entrada por teclado
		return strings.TrimSuffix(opcion, "\n")
	}
}

// Función main. Lanzadora del script
func main() {

	//opcion: variable de elección de menu de acciones
	var opcion string

	//inicialización de la variable tipo User, usando el mapa creado
	// al inicio que contiene los datos de cada usuario.
	users = make(map[int]User)

	//Almacenamos en reader todo lo que se entre por teclado.
	reader = bufio.NewReader(os.Stdin)

	// Bucle infinito de muestra de Menu de Opciones
	for {
		//Menu de acciones posibles
		fmt.Println("\n")
		fmt.Println(("A) Crear. "))
		fmt.Println(("B) Listar. "))
		fmt.Println(("C) Actualizar. "))
		fmt.Println(("D) Eliminar. \n"))
		fmt.Print("...Elige una opción (X para salir): ")

		//Llamada a la función de selección de opcion
		opcion = readLine()

		// Filtrado para la salida del script
		if opcion == "x" || opcion == "X" {
			break
		}

		//Control de elecciones de acción
		switch opcion {
		case "a", "A", "Crear", "crear":
			crearUsuarios()
		case "b", "B", "Listar", "listar":
			listarUsuarios()
		case "c", "C", "Actualizar", "actualizar":
			actualizarUsuarios()
		case "d", "D", "Eliminar", "eliminar":
			eliminarUsuarios()
		default:
			fmt.Println("Opción elegida NO VÁLIDA.")
		}
	}
	fmt.Println("Script acabado")
}
