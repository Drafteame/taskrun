package config

import (
	"fmt"
	"github.com/Drafteame/taskrun/internal/models"
)

func GetJob(job, stage, file string) (*models.Job, error) {
	jobs, err := GetJobs(stage, file)
	if err != nil {
		return nil, err
	}

	var jobModel *models.Job

	for _, j := range jobs {
		aux := j

		if j.Name == job {
			jobModel = &aux
			break
		}
	}

	if jobModel == nil {
		return nil, fmt.Errorf("job %s not found on stage %s", job, stage)
	}

	return jobModel, nil
}
