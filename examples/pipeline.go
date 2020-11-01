package main

import (
	"fmt"
)

var Pipeline = Example{}

type Example struct{}

func (ex Example) Consume(objs ...interface{}) (interface{}, error) {
	return objs, nil
}

func (ex Example) Init(objs ...interface{}) (interface{}, error) {
	fmt.Println("Example.Init()", objs)
	return nil, nil
}
