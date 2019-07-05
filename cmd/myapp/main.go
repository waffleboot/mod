package main

import (
	"github.com/waffleboot/mod/internal/app/myapp"
	"github.com/waffleboot/mod/internal/pkg/myprivlib"
	"github.com/waffleboot/mod/pkg/mypubliclib"
)

func main() {
	myapp.Foo()
	myprivlib.Foo()
	mypubliclib.Foo()
}
