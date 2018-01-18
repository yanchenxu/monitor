package main

import (
	"fmt"
	"os"

	"github.com/bocheninc/base/log"
)

func main() {
	if len(os.Args) != 2 {
		log.Infoln("must ./monitor monitor.yaml")
	}

	fmt.Println(os.Args)
}
