package app

import (
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

var (
	config1 = &Config{
		Directories: []string{},
		LogLevel:    "info",
		Depth:       1,
		QuickMode:   false,
		Mode:        "fetch",
	}
	config2 = &Config{
		Directories: []string{string(os.PathSeparator) + "tmp"},
		LogLevel:    "error",
		Depth:       1,
		QuickMode:   true,
		Mode:        "pull",
	}
)

func TestSetup(t *testing.T) {
	mockApp := &App{Config: config1}
	var tests = []struct {
		input    *Config
		expected *App
	}{
		{config2, nil},
		{config1, mockApp},
	}
	for _, test := range tests {

		app, err := Setup(test.input)
		if err != nil {
			t.Errorf("Test Failed. error: %s", err.Error())
		}
		q := test.input.QuickMode
		if q && app != nil {
			t.Errorf("Test Failed.")
		} else if !q && app == nil {
			t.Errorf("Test Failed.")
		}

	}
}

func TestClose(t *testing.T) {
	mockApp := &App{}
	if err := mockApp.Close(); err != nil {
		t.Errorf("Test")
	}
}

func TestSetLogLevel(t *testing.T) {
	var tests = []struct {
		input string
	}{
		{"debug"},
		{"info"},
	}
	for _, test := range tests {
		setLogLevel(test.input)
		if test.input != log.GetLevel().String() {
			t.Errorf("Test Failed: %s inputted, actual: %s", test.input, log.GetLevel().String())
		}
	}
}

func TestOverrideConfig(t *testing.T) {
	var tests = []struct {
		inp1     *Config
		inp2     *Config
		expected *Config
	}{
		{config1, config2, config1},
	}
	for _, test := range tests {
		if output := overrideConfig(test.inp1, test.inp2); output != test.expected || test.inp2.Mode != output.Mode {
			t.Errorf("Test Failed: {%s, %s} inputted, output: %s, expected: %s", test.inp1.Directories, test.inp2.Directories, output.Directories, test.expected.Directories)
		}
	}
}

func TestExecQuickMode(t *testing.T) {
	var tests = []struct {
		inp1 []string
		inp2 *Config
	}{
		{[]string{""}, config1},
	}
	for _, test := range tests {
		if err := execQuickMode(test.inp1, test.inp2); err != nil {
			t.Errorf("Test Failed: %s", err.Error())
		}
	}
}
