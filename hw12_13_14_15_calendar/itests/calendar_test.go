package itests

import (
	"os"
	"syscall"
	"testing"

	integrationtests "github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/pkg/integration-tests"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

func TestCalendar(t *testing.T) {
	opts := godog.Options{Output: colors.Colored(os.Stdout)}

	var status int
	opts.Format = "pretty"
	opts.Randomize = 0
	opts.Paths = []string{"calendar_event_features"}
	status = godog.TestSuite{
		Name:                "integration",
		ScenarioInitializer: integrationtests.EventCrudFeatureContext,
		Options:             &opts,
	}.Run()

	syscall.Exit(status)
}
