package main

import (
	"markdowndoc/cmd"
	"runtime"
)

func init()  {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main()  {
	cmd.Execute()
}
