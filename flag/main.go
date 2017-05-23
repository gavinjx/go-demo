package main

// go run flag/main.go -level 2 -conf ./online.conf
import (
	"flag"
	"fmt"
)

func main() {

	level := flag.Int64("level", 1, "level num")
	confPath := flag.String("conf", "/data/conf/online.conf", "conf path")

	flag.Parse()
	fmt.Println(*level, *confPath)
}
