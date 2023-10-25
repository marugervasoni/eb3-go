package main

import	"fmt"


func main(){

}

//COMPOSICIÓN DE UNA FUNCION:----------------------------------------------------------------------------------------------------
//-pueden tener 0, 1 o muchas entradas; y puede retornar 0, 1 o muchos valores.

//palabraReservada + nombreFunc + parametrosEntrada + tipoDeRetorno
func suma(a int, b int) int{
//cuerpo
	sum := a+b
	return sum 
}


//SCOPES DE FUNCIONES:-------------------------------------------------------------------------------------------------------------
//- variables a nivel de PAQUETE - FUNCION - BLOQUE
//	LOCAL: dentro de función main
//	GLOBAL: fuera de función main. En parte superior del programa. Acceso desde cualquier función definida en el programa
//	PARAMETROS


//NOTACIÓN ELLIPSIS:----------------------------------------------------------------------------------------------------------------
//permite a las funciones recibir cantidad dinamica de parametros. Parámetros representados serán del mismo tipo de dato 

	//ejemplo 1
func sumar(valores...float64) float64{
	var resultado float64
	for _, valor := range valores {
		resultado += valor
	}
	return resultado
}
//invocamos la funcion y pasamos parametros en MAIN:
//sumar(2,3,4,5,6,7,8,9,4)

	//ejemplo 2: cuando tenemos parametros adicionales, ellipsis debe ir en ultimo lugar
func miFunc(valor1 string, valor2 string, valores...float64)


//MULTIRETORNO:----------------------------------------------------------------------------------------------------------------
//retornar más de 1 valor. Indicamos los tiposDeDatos de cada valor a retornar separado por coma y entre ()

//palabraReservada + nombreFunc + parametrosEntrada + tipoDatosDeRetorno
func operaciones(valor1, valor2 float64) (float64,float64,float64,float64) {
	suma := valor1 + valor2
	resta := valor1 - valor2
	multip := valor1 * valor2
	var divis float64

	if valor2 != 0 {
		divis = valor1 / valor2	
	}

	return suma, resta, multip, divis
}
//invocamos la funcion y pasamos parametros en MAIN:
//	s,r,m,d := operaciones(6,2)
//ejemplo de impresión por consola
//	fmt.Println("Suma:\t\t", s)


//RETORNO DE VALORES NOMBRADOS:-------------------------------------------------------------------------------------------------------
//definimos tiposDeDatos a retornar como nombreDeVariable, separado por coma y entre (). Dentro de la func 
//almacenamos en esas var el valor a retornar

//palabraReservada + nombreFunc + parametrosEntrada + nombreVarDeRetorno y tipoDatosDeRetorno 
func operaciones2(valor1, valor2 float64) (suma float64, resta float64, multip float64, divis float64) {
	suma = valor1 + valor2
	resta = valor1 - valor2
	multip = valor1 * valor2
	
	if valor2 != 0 {
		divis = valor1 / valor2	
	}

	return
}


//RETORNO DE FUNCIONES:----------------------------------------------------------------------------------------------------------------
//una funcion que devuelve otra funcion
 
//Ejemplo: funcion por operacion + funcion orquestadora 
func opSuma(valor1, valor2 float64) float64 {
	return valor1 + valor2
}
func opResta(valor1, valor2 float64) float64 {
	return valor1 - valor2
}
func opMultip(valor1, valor2 float64) float64 {
	return valor1 * valor2
}
func opDivis(valor1, valor2 float64) float64 {
	return valor1 / valor2
}

//	orquestador
//palabraReservada + nombreFunc + parametrosEntrada + funcDeRetorno
func operacionAritmetica(operador string) func(valor1, valor2 float64) float64 {
	switch operador {
	case "Suma":
		return opSuma
	case "Resta":
		return opResta
	case "Multip":
		return opMultip
	case "Divis":
		return opDivis
	}
	return nil
}
//invocamos la funcion y pasamos parametros en MAIN:
//indicamos operacion a realizar:
//	oper := operacionAritmetica("Suma")
//nos devuelve una función a la que passamos los valores con los cuales haremos la operación
//	r := oper(2,5)
//ejemplo de impresión por consola
//	fmt.Println(r)
