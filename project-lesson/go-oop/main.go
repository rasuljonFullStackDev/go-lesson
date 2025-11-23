package main

import (
	"fmt"
)

type Parent struct {
	FullName string
	Region   string
}

type User struct {
	Parent
	Name string
	Age  int
}

func newUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}
func (u *User) read() {
	fmt.Println("Name and user", u.Name, u.Age, u.Parent.FullName, u.Parent.Region)
}

func (u User) greate() {
	fmt.Println("Hello " + u.Name)
}

// interface
type Payment interface {
	Pay(amout float64)
}
type Payme struct {
}

func (Payme) Pay(amout float64) {
	fmt.Println("payme Paid:", amout)
}

type Click struct{}

func (Click) Pay(amout float64) {
	fmt.Println("click Paid:", amout)
}
func MakePayment(p Payment, amount float64) {
	p.Pay(amount)
}

type userService interface {
	GetUserName() string
}
type RealUser struct {
	Name string
}

func (u RealUser) GetUserName() string {
	return u.Name
}

type Counter struct {
	Value int
}

func (c *Counter) increment() {
	c.Value++
}

type Logger interface {
	Log(message string)
}
type ConsoleLogger struct{}

func (ConsoleLogger) Log(message string) {
	fmt.Println("Log", message)
}

type App struct {
	Logger Logger
}

// {{ .User}}
func main() {
	user := User{Name: "rasuljon"}
	user.greate()
	user2 := newUser("rasuljon", 25)
	user2.Parent.FullName = "Rasuljon"
	user2.Parent.Region = "Toshkent"
	user2.read()
	MakePayment(Payme{}, 1000)
	MakePayment(Click{}, 1000)
	// servis ishlatish
	var service userService = RealUser{Name: "Rasuljon"}
	fmt.Println("foydalanuvchi,", service.GetUserName())
	// counter ishlatish
	counter := Counter{}
	counter.increment()
	fmt.Println("counter:", counter.Value)
	// interface
	app := App{Logger: ConsoleLogger{}}
	app.Logger.Log("salom server")

}
