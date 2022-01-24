package main

import (
	"fmt"

	"github.com/speaud/demo-go-pkg/foobar"
)

func helloworld(b string) string {
	a := foobar.Method("asd")

	return "Hello World!!" + b + a
}

func main() {
	fmt.Println(helloworld("asd"))
}

/*

docker run --rm golang:1.5
docker run -it --rm golang

docker run --rm golang sh -c "go get github.com/speaud/demo-go-pkg && exec demo-go-pkg"

go get <path-to-repo>@<branch>


*/
