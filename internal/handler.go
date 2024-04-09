package ingestor

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ScMofeoluwa/ingestor/internal/utils"
)

type LogHandler struct {
	service *LogService
	ctx     context.Context
}

func NewLogHandler(service *LogService) *LogHandler {
	return &LogHandler{
		service: service,
		ctx:     context.Background(),
	}
}

func (l *LogHandler) InsertLog(w http.ResponseWriter, r *http.Request) {
	var req []utils.LogEntry
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorResponse(w, &ErrorPayload{Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}

	if err := l.service.InsertLog(l.ctx, req); err != nil {
		ErrorResponse(w, &ErrorPayload{Message: err.Error(), StatusCode: http.StatusInternalServerError})
		return
	}
	SuccessResponse(w, &SuccessPayload{Message: "log successfully ingested", StatusCode: http.StatusOK})
}
