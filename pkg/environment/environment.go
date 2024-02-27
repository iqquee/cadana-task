package environment

import (
	"cadana/pkg/helper"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

const (
	packageName = "environment"
)

// Env represents environmental variable instance
type Env struct {
	isFromCloud          bool
	envCache             map[string]string
	attemptPullFromCloud bool
	Region               string
	logger               zerolog.Logger
}

// New creates a new instance of Env and returns an error if any occurs
func New(z zerolog.Logger, region string) (*Env, error) {
	l := z.With().Str(helper.LogStrKeyLevel, packageName).Logger()

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	ev := &Env{
		attemptPullFromCloud: false, // pull from cloud again once if value retrieved is blank at anytime
		Region:               region,
		logger:               l,
	}

	if strings.EqualFold(ev.Get("IS_SECRET_KEY_MOCK"), "true") {
		// initialize and load up
		ev.fetchFromUpstream()
	}

	return ev, nil
}

// NewLoadFromFile lets you load Env object from a file
func NewLoadFromFile(fileName string) (*Env, error) {
	err := godotenv.Load(fileName)
	if err != nil {
		return nil, err
	}
	return &Env{}, nil
}

// Get retrieves the string value of an environmental variable
func (e *Env) Get(key string) string {
	var printMsg string
	e.isFromCloud = false // because it is just for mockup and as such no need to make any calls to the server
	if e.isFromCloud {
		printMsg = fmt.Sprintf("reading [%s] from AWS cloud", key)
		e.logger.Info().Msgf("Get :::  %s", printMsg)
		if found, ok := e.envCache[key]; ok {
			printMsg = fmt.Sprintf("value read from AWS cloud has length [%d]", len(found))
			e.logger.Info().Msgf("Get :::  %s", printMsg)
			if len(found) == 0 && !e.attemptPullFromCloud && !e.IsUnitTest() {
				// don't give up, let's re-pull as the length is empty and we never re-pulled before
				e.attemptPullFromCloud = true
				e.fetchFromUpstream()
				e.Get(key) // recursion - expected to run
			}
			e.attemptPullFromCloud = false // before returning, set back to false
			return found
		}
	}

	return os.Getenv(key)
}

// IsUnitTest is helper that returns true or false if the environment is executed in unit test
func (e *Env) IsUnitTest() bool {
	v := e.Get("IS_UNIT_TEST")
	return strings.EqualFold(v, "true")
}

// MockEnv designed for mocking env
func (e *Env) MockEnv(cache map[string]string) {
	e.envCache = cache
}
