package main

type Driven interface {
	Drive()
}

type Car struct {
}

func (c *Car) Drive() {
	println("Car is being driven")
}

type Driver struct {
	Age int
}

type CarProxy struct {
	car    *Car
	driver *Driver
}

func (c *CarProxy) Drive() {
	if c.driver.Age >= 18 {
		c.car.Drive()
	} else {
		println("Driver too young")
	}
}

func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{&Car{}, driver}
}

func main() {
	car := NewCarProxy(&Driver{22})
	car.Drive()
}

// Proxy is a structural design pattern that provides an object that acts as a substitute for a real service object used by a client.
// A proxy receives client requests, does some work (access control, caching, etc.) and then passes the request to a service object.

// Protection Proxy is a structural design pattern that lets you provide access to some of the object's functionality.
