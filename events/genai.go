package events

import (
	"strings"
	"time"

	"github.com/0x2142/frigate-notify/config"
	"github.com/0x2142/frigate-notify/models"
	"github.com/0x2142/frigate-notify/notifier"
	"github.com/rs/zerolog/log"
)

// processEvent handles preparing event for alerting
func processGenai(event models.Genai) {
	config.Internal.Status.LastEvent = time.Now()

	// Convert to human-readable timestamp
	genaiTime := time.Unix(int64(genai.StartTime), 0)
	log.Info().
		Str("genai_id", genai.ID).
		Str("description", genai.Description).
		Msg("Processing update...")
	log.Debug().
		Str("genai_id", genai.ID).
		Msgf("Event start time: %s", genaiTime)

	// Check that event passes configured filters
	if !checkFilters(event) {
		return
	}

	// Send alert with snapshot
	notifier.SendAlert(event)
}
