package template

import (
	"github.com/awslabs/goformation/cloudformation"
	"github.com/awslabs/goformation/cloudformation/resources"
)

func ValidateAWSCloudWatchAlarm(
	projectName, configName, resourceName string,
	template *cloudformation.Template,
	res *resources.AWSCloudWatchAlarm,
) error {
	if res.AlarmName != "" {
		return resourceError(res, resourceName, "Names are overwritten")
	}

	res.AlarmName = normalizeName("fenrir", projectName, configName, resourceName, 255)

	return nil
}
