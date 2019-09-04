package main

import (
	"fmt"
	"sync"
)

type MyClass struct {
	attrib string
}

func (c *MyClass) SetAttrib(val string) {
	c.attrib = val
}

func (c *MyClass) GetAttrib() string {
	return c.attrib
}

var (
	once     sync.Once
	instance *MyClass
)

func GetMyClass() *MyClass {
	once.Do(func() {
		instance = &MyClass{"first"}
	})

	return instance
}

func main() {
	instance := GetMyClass()
	instance.SetAttrib("hola")
	instance = GetMyClass()
	fmt.Println(instance.GetAttrib())
}
