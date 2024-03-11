package templating

import (
	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/Drafteame/taskrun/internal/models"
)

type JobTemplate struct {
	template      string
	jobModel      *models.Job
	data          map[string]string
	awsConfig     aws.Config
	finalEnvs     map[string]string
	dependantEnvs []string
}

func NewJobTemplate(job *models.Job) *JobTemplate {
	return &JobTemplate{
		template:      job.ToYAML(),
		jobModel:      job,
		data:          make(map[string]string),
		finalEnvs:     make(map[string]string),
		dependantEnvs: make([]string, 0),
	}
}

func (jt *JobTemplate) WithData(data map[string]string) *JobTemplate {
	jt.data = data
	return jt
}

func (jt *JobTemplate) WithAWSConfig(cfg aws.Config) *JobTemplate {
	jt.awsConfig = cfg
	return jt
}
