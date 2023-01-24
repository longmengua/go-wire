package main

import wiregen "wire-demo/wire-gen"

func main() {
	e := wiregen.InitializeEvent()

	e.Start()
}
