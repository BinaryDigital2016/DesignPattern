package main

import "fmt"

type Company interface {
	Add(c Company)
	Show()
}

//type DefaultCompany struct {
//}
//
//func (d *DefaultCompany) Add(c Company) {
//
//}
//
//func (d *DefaultCompany) Show() {}

type CompanyA struct {
	Name        string
	Departments []Company
}

func (ca *CompanyA) Add(c Company) {
	if ca.Departments == nil {
		ca.Departments = make([]Company, 0)
	}
	ca.Departments = append(ca.Departments, c)
}

func (ca *CompanyA) Show() {
	fmt.Printf("Company %s\n", ca.Name)
	for _, v := range ca.Departments {
		v.Show()
	}
}

type FinanceDepartment struct {
	Name string
}

func (f *FinanceDepartment) Add(c Company) {

}

func (f *FinanceDepartment) Show() {
	fmt.Printf("department %s\n", f.Name)
}

type HRDepartment struct {
	Name string
}

func (h *HRDepartment) Add(c Company) {

}

func (h *HRDepartment) Show() {
	fmt.Printf("department %s\n", h.Name)
}

func main() {
	c := CompanyA{Name: "总公司", Departments: make([]Company, 0)}
	d1 := FinanceDepartment{Name: "Finance"}
	d2 := HRDepartment{Name: "Human Resource"}
	c.Add(&d1)
	c.Add(&d2)

	c1 := CompanyA{Name: "分公司A", Departments: make([]Company, 0)}
	d3 := FinanceDepartment{Name: "Finance"}
	d4 := HRDepartment{Name: "Human Resource"}
	c1.Add(&d3)
	c1.Add(&d4)
	c.Add(&c1)

	c2 := CompanyA{Name: "分公司B", Departments: make([]Company, 0)}
	d5 := FinanceDepartment{Name: "Finance"}
	d6 := HRDepartment{Name: "Human Resource"}
	c2.Add(&d5)
	c2.Add(&d6)
	c.Add(&c2)

	c.Show()
}
