package main

import (
	"fmt"
	"math/rand"
	"sync"
	"syscall/js"
	"time"
)

var c chan bool
var mutex = sync.Mutex{}

var colors = map[string]string{
	"primary":   "",
	"secondary": "",
	"text":      "",
}

func init() {
	c = make(chan bool)
}

type RGBColor struct {
	Red   int
	Green int
	Blue  int
}

func GetRandomColorInRgb() RGBColor {
	rand.Seed(time.Now().UTC().UnixNano() * rand.Int63())
	Red := rand.Intn(255)
	rand.Seed(time.Now().UTC().UnixNano() * rand.Int63())
	Green := rand.Intn(255)
	rand.Seed(time.Now().UTC().UnixNano() * rand.Int63())
	blue := rand.Intn(255)
	c := RGBColor{Red, Green, blue}
	return c
}

func GetRandomColorInHex() string {
	color := GetRandomColorInRgb()
	hex := "#" + getHex(color.Red) + getHex(color.Green) + getHex(color.Blue)
	return hex
}

func getHex(num int) string {
	hex := fmt.Sprintf("%x", num)
	if len(hex) == 1 {
		hex = "0" + hex
	}
	return hex
}

func generateColors() {
	for {
		rand.Seed(time.Now().UTC().UnixNano() * rand.Int63())
		for color, _ := range colors {
			mutex.Lock()
			colors[color] = GetRandomColorInHex()
			mutex.Unlock()
		}
		time.Sleep(333 * time.Millisecond)
	}
}

func main() {
	go generateColors()

	js.Global().Set("GetColor", js.FuncOf(GetColor))
	<-c
}

func GetColor(this js.Value, args []js.Value) interface{} {
	var color string
	var exist bool

	if len(args) == 0 {
		args = append(args, js.ValueOf("primary"))
	}

	if color, exist = colors[args[0].String()]; !exist {
		color = colors["primary"]
	}

	return color
}
