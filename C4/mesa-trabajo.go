package main

import "fmt"

// Algunas tiendas de e-commerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total. La empresa tiene
// tres tipos de productos: Pequeño, Mediano y Grande. Pero se espera que sean muchos más.
// Los costos adicionales para cada uno son:
// Pequeño: solo tiene el costo del producto.
// Mediano: el precio del producto + un 3% por mantenerlo en la tienda.
// Grande: el precio del producto + un 6% por mantenerlo en la tienda y un adicional de $2500 de costo de envío.
// El porcentaje de mantenerlo en la tienda se calcula sobre el precio del producto.
// Se requiere una función factory que reciba el tipo de producto y el precio, y retorne una interfaz Producto que tenga el método Precio.
// Se debe poder ejecutar el método Precio y que el método devuelva el precio total basándose en el costo del producto y los adicionales, en caso de que los tenga.

type product interface {
	price() float64
}

const (
	smallType = "SMALL"
	middleType = "MIDDLE"
	bigType = "BIG"
)

type small struct {
	cost float64
}
type middle struct {
	cost float64
}
type big struct {
	cost float64
}

func (s small) price() float64 {
	return s.cost
}
func (m middle) price() float64 {
	costMaintance := 1.03
	return m.cost * costMaintance
}
func (b big) price() float64 {
	costMaintance := 1.06
	aditional := 2500.0
	return b.cost * costMaintance + aditional
}

func productFactory(productType string, value float64) product {
	switch productType {
	case smallType:
		return small{cost: value}
	case middleType:
		return middle{cost: value}
	case bigType:
		return big{cost: value}
	}
	return nil
}

func main() {

	s := productFactory(smallType,1000)
	m := productFactory(middleType,2500)
	b := productFactory(bigType,5000)
	
	fmt.Println("Small product price: \t $",s.price())
	fmt.Println("Middle product price: \t $",m.price())
	fmt.Println("Big product price: \t $",b.price())

}