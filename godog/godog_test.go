package main

import (
	"fmt"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/cucumber/messages-go/v10"
	"os"
	"testing"
)

func thereAreGodogs(available int) error {
	godogs = available
	return nil
}

func iEat(arg1 int) error {
	godogs -= arg1
	return nil
}

func thereShouldBeRemaining(arg1 int) error {
	if godogs != arg1 {
		return fmt.Errorf("godog not the same should be %v instead of %v", arg1, godogs)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^there are (\d+) godogs$`, thereAreGodogs)
	s.Step(`^I eat (\d+)$`, iEat)
	s.Step(`^there should be (\d+) remaining$`, thereShouldBeRemaining)

	s.BeforeScenario(func(pickle *messages.Pickle) {
		godogs = 0 //reset
	})
}

var options = godog.Options{
	Format:              "progress",
	Output:              colors.Colored(os.Stdout),
}

func TestMain(m *testing.M) {
	status := godog.RunWithOptions("godogs", func(s *godog.Suite) {
		FeatureContext(s)
	}, options)

	statusMRun := m.Run()
	if statusMRun > status {
		status = statusMRun
	}
	os.Exit(status)
}