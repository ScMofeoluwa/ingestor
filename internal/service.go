package ingestor

import (
	"context"

	"golang.org/x/sync/errgroup"

	"github.com/ScMofeoluwa/ingestor/internal/config"
	"github.com/ScMofeoluwa/ingestor/internal/database"
	"github.com/ScMofeoluwa/ingestor/internal/utils"
)

type LogService struct {
	db *database.Queries
}

func NewLogService(cfg config.Config) *LogService {
	return &LogService{
		db: database.SetupDB(cfg),
	}
}

func (l *LogService) InsertLog(ctx context.Context, logs []utils.LogEntry) error {

	eg, egCtx := errgroup.WithContext(ctx)
	for _, log := range logs {
		log := log
		eg.Go(func() error {
			return l.db.CreateLog(egCtx, log)
		})
	}
	return eg.Wait()
}
