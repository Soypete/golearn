/*
this is a theoretical implementation not working
*/
package main

import (
	"flag"
	"fmt"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/trees"
)

var (
	word *string
)

func init() {
	word = flag.String("name", "harold", "This is the method that the program in performing.")
}
func main() {
	flag.Parse()
	file := "../models/DecisionTress.h"
	jsonData := fmt.Sprintf("{name: %s}", *word)

	fmt.Println(jsonData)
	tree := trees.NewID3DecisionTree(0.6)
	tree.Load(file)
	data := base.NewDenseInstances()
	a := base.NewCategoricalAttribute()
	a.UnmarshalJSON([]byte(jsonData))
	fmt.Println(a)
	predictions, err := tree.Predict(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(predictions)
}
