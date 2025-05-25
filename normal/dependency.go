package main
import (
	"fmt"
)
/*
Problem Statement:
You have a list of projects and a list of dependency pairs. Each pair (X, Y) means that project X must be completed before project Y. Your task is to determine a valid build order for all the projects such that for every dependency (X, Y), X appears before Y in the order. If no valid build order exists because there is a circular dependency, your solution should indicate that it's not possible to build the projects.

Example:

Projects: a, b, c, d, e, f
Dependencies: (a, d), (f, b), (b, d), (f, a), (d, c)
A valid build order would be: f, e, a, b, d, c
*/
func main(){
	fmt.Println("Hello, World!")
}
