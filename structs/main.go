package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Part",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 90000,
		},
	}

	//&variable pointer возвращает ключ из памяти где расположено значение на которое указывает эта переменная
	//returns the momery adress that variable is stored at
	//*pointer  верни значение на которое указывает этот указатель

	// jimPointer := &jim
	jim.updateName("clark")
	jim.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
