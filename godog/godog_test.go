package godog

import (
	"fmt"
	"github.com/cucumber/godog"
)

var godogs int

func thereAreGodogs(arg1 int) error {
	godogs = arg1
	return nil
}

func iEat(arg1 int) error {
	godogs -= arg1
	return nil
}

func thereShouldBeRemaining(arg1 int) error {
	if godogs == arg1 {
		return nil
	}
	return fmt.Errorf("godog not the same should be %v instead of %v", arg1, godogs)
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^there are (\d+) godogs$`, thereAreGodogs)
	s.Step(`^I eat (\d+)$`, iEat)
	s.Step(`^there should be (\d+) remaining$`, thereShouldBeRemaining)
}

