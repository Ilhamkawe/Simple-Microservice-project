package handler

import "transaction-service/logs"

type logHandler struct {
	logService logs.Service
}

func NewLogHandler(logService logs.Service) *logHandler {
	return &logHandler{logService}
}

