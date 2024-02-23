package main

import (
	"bytes"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

func TestShouldLogDebugMessagesAndAbove(t *testing.T) {
	println("TestShouldLogDebugMessagesAndAbove")
	t.Setenv("LOG_LEVEL", "debug")

	buffer := new(bytes.Buffer)
	logger := createLogger(buffer)

	logger.Debugf("debug message: %v", "40e72733-2c05-40d4-a6bd-357ebb37dd75")
	logger.Infof("info message: %v", "9f35397c-65c1-4d59-bf83-4d578e0ab104")
	logger.Warnf("warn message: %v", "f13077ec-b2d8-44e6-a237-c84b369c7a1e")
	logger.Errorf("error message: %v", "129cd79c-e568-4a8e-8bd9-82eec11f42a4")

	snaps.MatchSnapshot(t, buffer.String())
}

func TestShouldLogInfoMessagesAndAbove(t *testing.T) {
	println("TestShouldLogInfoMessagesAndAbove")
	t.Setenv("LOG_LEVEL", "info")

	buffer := new(bytes.Buffer)
	logger := createLogger(buffer)

	logger.Debugf("debug message: %v", "40e72733-2c05-40d4-a6bd-357ebb37dd75")
	logger.Infof("info message: %v", "9f35397c-65c1-4d59-bf83-4d578e0ab104")
	logger.Warnf("warn message: %v", "f13077ec-b2d8-44e6-a237-c84b369c7a1e")
	logger.Errorf("error message: %v", "129cd79c-e568-4a8e-8bd9-82eec11f42a4")

	snaps.MatchSnapshot(t, buffer.String())
}

func TestShouldLogWarnMessagesAndAbove(t *testing.T) {
	println("ShouldLogWarnMessagesAndAbove")
	t.Setenv("LOG_LEVEL", "warn")

	buffer := new(bytes.Buffer)
	logger := createLogger(buffer)

	logger.Debugf("debug message: %v", "40e72733-2c05-40d4-a6bd-357ebb37dd75")
	logger.Infof("info message: %v", "9f35397c-65c1-4d59-bf83-4d578e0ab104")
	logger.Warnf("warn message: %v", "f13077ec-b2d8-44e6-a237-c84b369c7a1e")
	logger.Errorf("error message: %v", "129cd79c-e568-4a8e-8bd9-82eec11f42a4")

	snaps.MatchSnapshot(t, buffer.String())
}

func TestShouldLogErrorMessagesAndAbove(t *testing.T) {
	println("TestShouldLogErrorMessagesAndAbove")
	t.Setenv("LOG_LEVEL", "error")

	buffer := new(bytes.Buffer)
	logger := createLogger(buffer)

	logger.Debugf("debug message: %v", "40e72733-2c05-40d4-a6bd-357ebb37dd75")
	logger.Infof("info message: %v", "9f35397c-65c1-4d59-bf83-4d578e0ab104")
	logger.Warnf("warn message: %v", "f13077ec-b2d8-44e6-a237-c84b369c7a1e")
	logger.Errorf("error message: %v", "129cd79c-e568-4a8e-8bd9-82eec11f42a4")

	snaps.MatchSnapshot(t, buffer.String())
}
