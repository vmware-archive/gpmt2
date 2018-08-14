/*
Greenplum Magic Tool

Authored by Tyler Ramer, Ignacio Elizaga
Copyright 2018

Licensed under the Apache License, Version 2.0 (the "License")

*/
package cmd

import (
	"github.com/op/go-logging"
	"os"
	"fmt"
	"time"
)

// Logger name and logging format
var (
	log = logging.MustGetLogger(toolName)
	format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfile}(%{shortfunc}) ▶ %{level:.4s} ▶ %{color:reset} %{message}`,
	)
)

// Setup the logger.
func SetupLogger()  {

	// Check if the log level os parameter is set, to turn on debug logging
	var logLevel = os.Getenv("LOG_LEVEL")

	// New On Screen logger
	backendScreen := logging.NewLogBackend(os.Stderr, "", 0)
	screenFormatted := logging.NewBackendFormatter(backendScreen, format)
	screenLevelled := logging.AddModuleLevel(screenFormatted)

	// Setup logfile logger
	logFilename := fmt.Sprintf("/tmp/gpmt_log_%s", time.Now().Format("2006-01-02"))
	file, err := os.OpenFile(logFilename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	backendFile := logging.NewLogBackend(file, "", 0)
	fileFormatter := logging.NewBackendFormatter(backendFile, format)
	fileLeveled := logging.AddModuleLevel(fileFormatter)

	// Setup the logging level based on users request
	if logLevel == "DEBUG" {
		screenLevelled.SetLevel(logging.DEBUG, "")
		fileLeveled.SetLevel(logging.DEBUG, "")
	} else {
		screenLevelled.SetLevel(logging.INFO, "")
		fileLeveled.SetLevel(logging.INFO, "")
	}

	// Start the logger
	logging.SetBackend(screenLevelled, fileLeveled)

	// If the file creation failed, then place a warning message on the screen.
	// NOTE: the reason why place the err catcher here is simply because the logger has not
	// initialized yet, until we set the backend.
	if err != nil {
		log.Warningf("Cannot log onto logfile: %s, ERROR: %v", logFilename, err)
	} else {
		log.Debugf("All %s log messages are logged into: %s", toolName, logFilename)
	}

	log.Infof(logLevel)
}
