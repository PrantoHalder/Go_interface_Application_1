package main 

import "fmt"

type leaveCalculator interface {
	calculateLeave()int
}
type leave struct{
	name string
	employeeID int
	leaveTaken int
	totalLeave int 
}
func (p leave)calculateLeave()int{
  return p.totalLeave -p.leaveTaken
}
func show(){
	a1 := leave{
		name:       "Pranto",
		employeeID: 1234,
		leaveTaken: 2,
		totalLeave: 10,
	}
	a2 :=leave{
		name:       "Shovon",
		employeeID: 1223,
		leaveTaken: 4,
		totalLeave: 10,
	}
	
	employShowen := []leaveCalculator {a1,a2}
	showLeaveAvailable(employShowen)
}
func showLeaveAvailable (e []leaveCalculator){
  for _,r := range e{
	fmt.Println("the leave available",r.calculateLeave())
  }
}