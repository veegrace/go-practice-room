//compute grades

package main

import (
	"fmt"
)

func main(){
	grd := Grades{[]float64 {97.0,85.9,55.6,78.6,66.7}}


	fmt.Println("Total Grade: ", grd.getTotalGrade())
	fmt.Println("Average Grade: ", grd.getAverageGrade())
	//fmt.Println("Highest Grade: ", grd.getHighestGrade())
}

type calculateGrades interface{
	getTotalGrade() float64
	getAverageGrade() float64
	//getHighestGrade() float64
}

type Grades struct{
	gradess []float64
}

func (g Grades) getTotalGrade() float64{
	var sum float64 = 0.0
	
	for i:=0; i < len(g.gradess); i++{
		sum = sum + g.gradess[i]
	}
	return sum
}

func (g Grades) getAverageGrade() float64{
	return g.getTotalGrade()/float64(len(g.gradess))  //len(g.gradess)
}

func (g Grades) getHighestGrade() float64{
	return  
}

