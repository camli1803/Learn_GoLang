package computer

type Address struct {
	City, State string
}
type Person struct {
	Name        string
	Age, Salary int
	Address
}
