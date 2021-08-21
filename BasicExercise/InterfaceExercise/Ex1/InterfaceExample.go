package main

import "fmt"

type SalaryCaculator interface { //tạo 1 interface để tính lương cho từng kiểu đối tượng
	CaculateSalary() int // SalaryCaculator là kiểu đối tượng tổng quát
}

type Permanent struct { // nhân viên chính thức có lương cơ bản và lương thưởng
	emId     int
	basicpay int
	pf       int
}

type Contract struct { // nhân viên hợp đồng chỉ có lương cơ bản
	emId     int
	basicpay int
}

//Viết cách tính lương cho nhân viên chính thức.
func (p Permanent) CaculateSalary() int {
	return p.basicpay + p.pf
}

// Viết cách tính lương cho nhân viên hợp đồng.
func (c Contract) CaculateSalary() int {
	return c.basicpay
}

// Viết hàm tính lương tổng cho kiểu đối tượng tổng quát SalaryCaculator
func totalExpense(s []SalaryCaculator) int {
	expense := 0
	for _, v := range s {
		expense += v.CaculateSalary()
	}
	return expense
}

func main() {
	p1 := Permanent{1, 5000, 20}
	p2 := Permanent{2, 6000, 30}
	c3 := Contract{3, 3000}
	s := []SalaryCaculator{p1, p2, c3}
	fmt.Printf("Total Expense: $%d", totalExpense(s))
}
