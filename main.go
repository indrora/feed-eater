package main

import (
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	"os"

	"github.com/indrora/feed-eater/config"
	fio "github.com/indrora/feed-eater/io"
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
	configPath := "config.toml"
	if len(os.Args) > 1 {
		configPath = os.Args[1]
	}

	limits := limiter.New(3000 / 8)
	limits.SetBurst(1)

	conf, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	var output io.Writer

	output, err = conf.GetOutput()

	if err != nil {
		panic(err)
	}

	if conf.General.Slow && conf.General.SpeedLimit > 0 {
		limits.SetLimit(limiter.Limit(conf.General.SpeedLimit))
		output = limits.Writer(output)
	}
	if conf.Output.FilterTTY {
		output = fio.NewTTYConverter(output)
	}

writeout:

	sources := conf.Sources

	if conf.General.Shuffle {
		// shuffle the set of sources
		rand.Shuffle(len(sources), func(i, j int) { sources[i], sources[j] = sources[j], sources[i] })
	}

	output.Write([]byte("\r\n\r\n\n"))

	for _, source := range conf.Sources {
		if source.Inhibit {
			continue
		}
		if source.Impl != nil {
			(*source.Impl).Print(output)
		}
		idx := rand.UintN(uint(len(dividers)))
		fmt.Fprint(output, "\r\n\r\n\n"+dividers[idx]+"\r\n\n\r\n")
	}

	if conf.General.Loop {
		goto writeout
	}

}
