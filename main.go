package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"os"

	"github.com/indrora/feed-eater/config"
	"github.com/sunshineplan/limiter"
)

var (
	dividers = []string{
		"================================================",
		".:*~*:._.:*~*:._.:*~*:._.:*~*:._.:*~*:._.:*~*:._.:*~*:._.:*~*:.",
		"=^..^=   =^..^=   =^..^=    =^..^=    =^..^=    =^..^=    =^..^=",
		"_,.-'~'-.,__,.-'~'-.,__,.-'~'-.,__,.-'~'-.,__,.-'~'-.,_",
		".oOo.oOo.oOo.oOo.oOo.oOo.oOo.oOo.oOo.oOo.oOo.oOo.oOo.oOo.oOo.",
		"-=x=-=x=-=x=-=x=-=x=-=x=-=x=-=x=-=x=-=x=-=x=-=x=-=x=-=x=-=x=-",
		"pdbqpdbqpdbqpdbqpdbqpdbqpdbqpdbqpdbqpdbqpdbqpdbqpdbqpdbqpdbqp",
		`  .--.      .-'.      .--.      .--.      .--.      .--.      .--.      .--.
:::::.\::::::::.\::::::::.\::::::::.\::::::::.\::::::::.\::::::::.\::::::::.\
'      '--'      '.-'      '--'      '--'      '--'      '--'      '--'      '`,
	}
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
		idx := rand.UintN(uint(len(dividers)))
		fmt.Fprint(slowIO, "\n\n\n"+dividers[idx]+"\n\n\n")

	}

}
