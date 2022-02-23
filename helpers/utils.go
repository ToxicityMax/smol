package helpers

import (
	"crypto/rand"
	"fmt"
	"github.com/ToxicityMax/smol/logger"
	"github.com/ToxicityMax/smol/models"
	"github.com/go-co-op/gocron"
	"time"
)

func Create_unique_string(length int) string {
	bytestring := make([]byte, length)
	if _, err := rand.Read(bytestring); err != nil {
		panic(err)
	}
	// todo: query db to check if string exists already or not
	str := fmt.Sprintf("%X", bytestring)
	return str
}

func StartCron() {
	cron := gocron.NewScheduler(time.UTC)
	logger.Info("Starting Cron job...")
	cron.Every(1).
		Hours().Do(
		func() {
			var urls []models.URL
			var count int

			if err := models.DB.
				Find(&urls, "expiration_date >= ?", time.Now()). //query the db for expired links
				Count(&count).                                   // number of expired links
				Delete(&urls).Error; err != nil {                // deleting expired links
				logger.Error(fmt.Sprintf("Some error in query --> %s ", err))
				return
			} else if count > 0 {
				logger.Info(fmt.Sprintf("Removed %d expired links...", count))
				return
			}
		})
	cron.StartAsync()
}
