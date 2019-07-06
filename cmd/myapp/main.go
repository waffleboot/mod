package main

import (
	"github.com/waffleboot/mod/internal/app/myapp"
	"github.com/waffleboot/mod/internal/pkg/myprivlib"
	lib "github.com/waffleboot/mod/pkg/mypubliclib"
	lib2 "github.com/waffleboot/mod/v2/pkg/mypubliclib"
)

func main() {
	myapp.Foo()
	myprivlib.Foo()
	lib.Foo()
	lib2.Foo(1)
}
