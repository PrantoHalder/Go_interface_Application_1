package main

import "fmt"

type person interface{
	showPerson() int
}
type personLeaves interface{
   showleave() int
}
type personAll interface {
	person
	personLeaves
}
type personDetails struct{
  name string
  age int
  salary int 
  pf int
  totalLeave int 
  leaveTaken int  
}

func (s personDetails) showPerson()int {
   return s.salary + s.pf 
}
func (p personDetails) showleave() int {
	return p.totalLeave -p.leaveTaken
}
func allShow(){
	s1 := personDetails{
		name:   "Prnto",
		age:    20,
		salary: 34000,
		pf:     3450,
	}
	s2 := personDetails{
		name:   "Shovon",
		age:    30,
		salary: 23450,
		pf:     3450,
	}
	s3 := personDetails{
		name:   "Halder",
		age:    23,
		salary: 21000,
		pf:     2345,
	}
	l1 := personDetails{
		totalLeave: 10,
		leaveTaken: 5,
	}
	l2 := personDetails{
		totalLeave: 10,
		leaveTaken: 4,
	}
	em1 := []personAll{s1,s2,s3,l1,l2}
	allTimeShow(em1)

}
func allTimeShow (r []personAll){
	for _,i := range r{
		fmt.Println("the salary ",i.showPerson())
		fmt.Println("Leave available",i.showleave())
	}
}