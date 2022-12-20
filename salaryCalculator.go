package main

import "fmt"

type caculateSalary interface {
	salaryCalculator() int
}
type emploeeInfo struct {
	Name   string
	age    int
	salary int
	pf     int
}

func implement() {
	e1 := emploeeInfo{
		Name:   "Pranto",
		age:    24,
		salary: 20000,
		pf:     15000,
	}
	e2 := emploeeInfo{
		Name:   "Shovon",
		age:    26,
		salary: 23000,
		pf:     13000,
	}
	employee := []caculateSalary {e1,e2}
	showEmployeeCalculator(employee)

}

func (p emploeeInfo ) salaryCalculator()int {
	return p.salary +p.pf
}
func showEmployeeCalculator (em []caculateSalary) {
  for _,v := range em{
	fmt.Println("salary",v.salaryCalculator())
  } 
}
	

