package graderd

import "github.com/Capstone-auto-grader/grader-api-v2/internal/grader-task"

func ParseContainerState(status string) grader_task.Status {
	// Container state can be one of "created", "running", "paused", "restarting", "removing", "exited", or "dead".
	switch status {
	case "created":
		return grader_task.StatusPending
	case "running":
		return grader_task.StatusStarted
	case "paused":

	case "restarting":
		return grader_task.StatusStarted
	case "removing":
		return grader_task.StatusComplete
	case "exited":
		return grader_task.StatusComplete
	case "dead":
		return grader_task.StatusFailed
	}
	return grader_task.StatusFailed
}
