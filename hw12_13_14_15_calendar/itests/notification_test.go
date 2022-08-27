package itests

import (
	"log"
	"os"
	"syscall"
	"testing"
	"time"

	integrationtests "github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/pkg/integration-tests"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

func TestNotification(t *testing.T) {
	const delay = 1 * time.Second
	opts := godog.Options{Output: colors.Colored(os.Stdout)}

	log.Printf("wait %s for service availability...", delay)
	time.Sleep(delay)

	var status int
	opts.Format = "pretty"
	opts.Randomize = 0
	opts.Paths = []string{"sender_features"}
	status = godog.TestSuite{
		Name:                "integration",
		ScenarioInitializer: integrationtests.NotificationFeatureContext,
		Options:             &opts,
	}.Run()

	syscall.Exit(status)
}
