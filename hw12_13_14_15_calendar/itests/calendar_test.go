package itests

import (
	"log"
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/pkg/integration-tests"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

const delay = 1 * time.Second

var opts = godog.Options{Output: colors.Colored(os.Stdout)}

func TestCalendar(t *testing.T) {
	log.Printf("wait %s for service availability...", delay)
	time.Sleep(delay)

	opts.Format = "pretty"
	opts.Randomize = 0
	opts.Paths = []string{"calendar_event_features"}
	status := godog.TestSuite{
		Name:                "integration",
		ScenarioInitializer: integration_tests.EventCrudFeatureContext,
		Options:             &opts,
	}.Run()

	opts.Paths = []string{"sender_features"}
	status = godog.TestSuite{
		Name:                "integration",
		ScenarioInitializer: integration_tests.NotificationFeatureContext,
		Options:             &opts,
	}.Run()

	syscall.Exit(status)
}
