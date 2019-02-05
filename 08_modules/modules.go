package main

import (
	"fmt"

	// $GOPATH/src/
	// import a module
	// if this is a github repo it would work on the same way
	"github.com/mstruebing/go-examples/08_modules/pkg"
)

func main() {
	// call DoStuff of the subpackage
	pkg.DoStuff()

	// prints the constant declared in the subpackage
	fmt.Println(pkg.MyConst)

	x := pkg.Add(1, 2)
	fmt.Println(x)
}
