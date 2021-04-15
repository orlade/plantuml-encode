package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/term"

	"github.com/orlade/plantuml-encode/plantuml"

	"github.com/sirupsen/logrus"
)

func main() {
	if term.IsTerminal(0) {
		logrus.Error("no piped data")
		os.Exit(1)
	}

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		logrus.Error(err.Error())
		os.Exit(2)
	}

	s, err := plantuml.DeflateAndEncodeBytes(b)
	if err != nil {
		logrus.Error(err.Error())
		os.Exit(3)
	}

	fmt.Print(s)
}
