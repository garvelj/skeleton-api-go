package main

import (
	"skeleton/api"
)

func main() {
	skeletonPort := ":13131"

	api.New("").Start(skeletonPort)
}
