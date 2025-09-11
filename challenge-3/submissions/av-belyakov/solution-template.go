package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Age    int
	Salary float64
}

type Manager struct {
	Employees []Employee
}


// AddEmployee adds a new employee to the manager's list.
func (m *Manager) AddEmployee(e Employee) {
    (*m).Employees = append((*m).Employees, e)
}

// RemoveEmployee removes an employee by ID from the manager's list.
func (m *Manager) RemoveEmployee(id int) {
    num, isExist := m.findEmployeeByID(id)

    if !isExist {
        return
    }
    
    m.Employees = append((*m).Employees[:num], m.Employees[num+1:]...)
}

// GetAverageSalary calculates the average salary of all employees.
func (m *Manager) GetAverageSalary() float64 {
    var full float64
    var count float64 = float64(len((*m).Employees))
    
    if count == 0 {
        return 0
    }
    
    for _, v := range (*m).Employees {
        full += v.Salary
    }
    
    
    
	return full/count
}

// FindEmployeeByID finds and returns an employee by their ID.
func (m *Manager) FindEmployeeByID(id int) *Employee {
    num, isExist := m.findEmployeeByID(id)

    if isExist {
        return &m.Employees[num]
    }    

    return nil
}

func (m *Manager) findEmployeeByID(id int) (int, bool) {
    for k, v := range (*m).Employees {
        if v.ID == id {
            return k, true
        }        
    }
    
    return -1, false
}

func main() {
	manager := Manager{}
	manager.AddEmployee(Employee{ID: 1, Name: "Alice", Age: 30, Salary: 70000})
	manager.AddEmployee(Employee{ID: 2, Name: "Bob", Age: 25, Salary: 65000})
	manager.RemoveEmployee(1)
	averageSalary := manager.GetAverageSalary()
	employee := manager.FindEmployeeByID(2)

	fmt.Printf("Average Salary: %f\n", averageSalary)
	if employee != nil {
		fmt.Printf("Employee found: %+v\n", *employee)
	}
}
