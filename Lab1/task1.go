package main

import "fmt"

func solveQuadratic(a, b, c int) {
	var discriminant = b*b - 4*a*c
	if discriminant < 0 {
		fmt.Println("No real roots")
	}
	if discriminant == 0 {
		var x = -b / (2 * a)
		fmt.Println("One root: ", x)
	}
	if discriminant > 0 {
		var x1 = (-b + discriminant) / (2 * a)
		var x2 = (-b - discriminant) / (2 * a)
		fmt.Println("Two roots: ", x1, x2)
	}
}

func main() {
	solveQuadratic(1, 2, 1)
}
