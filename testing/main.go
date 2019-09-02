package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/damienfamed75/aseprite"
)

func main() {
	f := &aseprite.File{}

	file, _ := os.Open("player.json")

	bytes, _ := ioutil.ReadAll(file)

	json.Unmarshal(bytes, f)

	log.Printf("IsMap[%v]\n", f.Frames.IsMap)
	for k, f := range f.Frames.Frames() {
		log.Printf("[%s:%v]\n", k, f.Duration)
	}
	// log.Println(f.Frames.FrameMap)
}
