// comments API.
package main

import (
	"produx/infrastructure"
	"produx/infrastructure/container"
)

func main() {
	containerInstance, err := container.GetInstance()
	if err != nil {
		panic(err.Error())
	}

	err = infrastructure.Start(containerInstance)
	if err != nil {
		panic(err.Error())
	}
}
