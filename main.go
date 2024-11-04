package main

import (
	"fmt"

	"github.com/ebitengine/purego"
	"github.com/mithrandie/csvq/lib/cli"
)

func main() {
	csvqlib, err := purego.Dlopen("./target/debug/libcsvq.so", purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	defer purego.Dlclose(csvqlib)
	var processData func(string) string
	purego.RegisterLibFunc(&processData, csvqlib, "process_data")
	result := processData("input")
	fmt.Println(result)

	cli.Run()
}
