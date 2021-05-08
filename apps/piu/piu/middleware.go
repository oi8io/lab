package piu

import (
	"fmt"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		fmt.Printf("[%d] %s in %v \n", c.StatusCode, c.Request.RequestURI, time.Since(t))
	}
}

func Mid1() HandlerFunc {
	name := "Mid1"
	return func(c *Context) {
		fmt.Println("start run", name)
		c.Next()
		fmt.Println("end   run", name)
	}
}

func Mid2() HandlerFunc {
	name := "Mid2"
	return func(c *Context) {
		// Start timer
		fmt.Println("start run", name)
		// Process request
		c.Next()
		fmt.Println("end   run", name)
		// Calculate resolution time
	}
}

func Mid3() HandlerFunc {
	name := "Mid3"
	return func(c *Context) {
		// Start timer
		fmt.Println("start run", name)
		// Process request
		c.Next()
		fmt.Println("end   run", name)
		// Calculate resolution time
	}
}
