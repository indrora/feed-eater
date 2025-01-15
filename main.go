package main

import (
	"fmt"
	"log"
	"os"

	"github.com/indrora/feed-eater/config"
	"github.com/sunshineplan/limiter"
)

func main() {

	limits := limiter.New(3000 / 8)
	limits.SetBurst(1)

	slowIO := limits.Writer(os.Stdout)

	conf, err := config.LoadConfig("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	for _, source := range conf.Sources {
		if source.Impl != nil {
			fmt.Println(source.Name)
			(*source.Impl).Print(slowIO)
		}
	}

}
