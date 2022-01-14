package main

import "test/samplelib"

func main() {
	gmap := samplelib.GetMap("magic2")
	println("GetMap(id:magic2) => " + gmap)
}
