package main

import "fmt"

func main() {
	fmt.Println("Hello, Methods!")

	satvik := User{"Satvik", 24, "satvik@gmail.com", false, "2323133"}
	fmt.Println(satvik) 
	fmt.Printf("Satvik details are : %+v\n", satvik)  
	fmt.Printf("Name is %v\n", satvik.Name)


	// now i want a method to display the status of the user
	satvik.GetStatus()
	satvikNewEmail := satvik.NewMail()
	fmt.Println("Satvik new email is: ", satvikNewEmail)
	// there is no change in the original email as we are passing a copy of the original struct
	fmt.Printf("Old Email is %v\n", satvik.Email) 
	// before changing status
	satvik.GetStatus()
	// after changing status
	satvik.ChangeStatus()
	satvik.GetStatus()
}

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
	password string // unexported field to other files
}

// create a method
// instead of u User we can also pass a part of User struct like u.Status



func (u User) GetStatus(){ // the value passes here is a copy of the original value
	fmt.Println("Is user actice: ", u.Status)
}

func (u User) NewMail() string{
	u.Email = "test@go.dev"
	return u.Email
}

// in order to modify the original struct we need to use pointer receiver

func (u *User) ChangeStatus(){
	u.Status = !u.Status
	fmt.Println("Status is : ", u.Status)
}