package config

import (
	"fmt"

	"github.com/Drafteame/taskrun/internal/models"
)

func GetJobs(stage, file string) ([]models.Job, error) {
	cfg, err := LoadConfigFromPath(file)
	if err != nil {
		return nil, err
	}

	if stage == "" {
		stage = cfg.DefaultStage
	}

	if stage == "" {
		for stageJob := range cfg.Jobs {
			stage = stageJob
			break
		}
	}

	if stage == "" {
		return nil, errNoStageDefined
	}

	stageJobs, ok := cfg.Jobs[stage]
	if !ok {
		return []models.Job{}, fmt.Errorf("stage %s not found", stage)
	}

	return stageJobs, nil
}
