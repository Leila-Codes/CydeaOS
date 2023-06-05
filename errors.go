package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func internalServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal Server Error"))
	logrus.WithField("Error", err.Error()).Error("Internal Server Error")
}
