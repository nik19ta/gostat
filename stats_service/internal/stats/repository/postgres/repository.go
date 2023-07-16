package postgres

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/nik19ta/gostat/stats_service/internal/stats/model"
)

type StatsRepository struct {
	db *gorm.DB
}

func NewStatsRepository(db *gorm.DB) StatsRepository {
	return StatsRepository{db: db}
}

func (r StatsRepository) GetVisits() ([]model.Visits, error) {
	timeBoundary := time.Now().AddDate(0, 0, -30)

	var visits []model.Visits
	r.db.Where("time_entry > ?", timeBoundary).Find(&visits)

	return visits, nil
}

type Project struct {
	UUID string `json:"uuid"`
}

func (r StatsRepository) AddVisit(data model.Visits) error {
	result := r.db.Create(&data)

	return result.Error
}

func (r StatsRepository) VisitExtend(session string) error {
	var lastVisit model.Visits

	err := r.db.Model(&model.Visits{}).Where("session = ?", session).Order("time_entry desc").First(&lastVisit).Error
	if err != nil {
		return err
	}

	result := r.db.Model(&model.Visits{}).Where("uid = ?", lastVisit.UId).Update("time_leaving", time.Now())

	return result.Error
}