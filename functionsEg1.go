package main
import "fmt"
func rect(l,b int) int {
	return l*b
}
func squ(a int) int {
	return a*a
}
func circlePerimeter(r float64) float64{
	return 2*3.14*r
}
func circleArea(R float64) float64{
	return 3.14*R*R
}
func main() {
fmt.Println("the area of the rectangle is",rect(13,12))
fmt.Println("The area of the square is ",squ(4))
fmt.Println("the perimeter of the circle is",circlePerimeter(6.8))
fmt.Println("the area of the circle is ",circleArea(12.4))
}