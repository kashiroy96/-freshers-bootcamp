package main

import "fmt"

type salary interface {
	calSalary() int
}

type fullTimeEmp struct {
	basic    int
}

type contractor struct {
	basic    int
}

type freelancer struct {
	basic    int
}

func (emp fullTimeEmp) calSalary() int {
	return emp.basic * 30
}

func (emp contractor) calSalary() int {
	return emp.basic * 30
}

func (emp freelancer) calSalary() int {
	return emp.basic * 20
}

func main() {

	fte := fullTimeEmp{500}
	cont := contractor{100}
	freelan := freelancer{10}

	empsal := []salary{fte, cont, freelan}

	for _, emp := range empsal {
		fmt.Println(emp.calSalary())
	}

	fmt.Println()

}
