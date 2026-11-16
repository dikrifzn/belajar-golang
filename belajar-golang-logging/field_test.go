package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "dikrifzn").Info("Hello world")
	logger.WithField("username", "dikrifzn").WithField("username", "dikrifzn").Info("Hello world")

	logger.WithFields(logrus.Fields{
		"username": "dikrifzn",
		"email":    "dikrifzn@example.com",
	}).Info("Hello world")
}
