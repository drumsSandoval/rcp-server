package main

import (
	"fmt"
	"net/rpc"
)

type Data struct {
	Alumno       string
	Materia      string
	Calificacion float64
}

func main() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	op := ""
	for {
		fmt.Printf( "\t\tMenu\n" +
			"\t1) Agregar calificación de una materia\n" +
			"\t2) Mostrar el promedio de un Alumno\n" +
			"\t3) Mostrar el promedio general\n" +
			"\t4) Mostrar el promedio de una materia\n" +
			"\t5) Ver info\n" +
			"\t6) Exit\n" +
			"\topc: ")
		fmt.Scanln(&op)

		switch op {
		case "1":
			var nombre, materia string
			var calificacion float64
			fmt.Println("Nombre del alumno: ")
			fmt.Scanln(&nombre)
			fmt.Println("Materia: ")
			fmt.Scanln(&materia)
			fmt.Println("Calificación: ")
			fmt.Scanln(&calificacion)

			data := Data{nombre, materia, calificacion}

			var result string
			err = c.Call("Server.AgregarCalificacionPorMateria", data, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Status: ", result)
			}
		case "2":
			var alumno string
			fmt.Print("Alumno: ")
			fmt.Scanln(&alumno)

			var result float64
			err = c.Call("Server.ObtenerPromedioAlumno", alumno, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Status: ", result)
			}
		case "3":
			var result float64
			err = c.Call("Server.ObtenerPromedioAlumnos", "", &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Status: ", result)
			}
		case "4":
			var materia string
			fmt.Print("Materia: ")
			fmt.Scanln(&materia)
			var result float64
			err = c.Call("Server.ObtenerPromedioPorMateria", materia, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Status: ", result)
			}
		case "5":
			var result string
			err = c.Call("Server.VerInfo", "", &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Status: ", result)
			}
		case "0":
			return
		}
	}
}
