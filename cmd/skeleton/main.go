package main

import (
	"skeleton/api"
)

func main() {
	skeletonPort := ":13131"

	api.New("conf.yaml").Start(skeletonPort)
}
