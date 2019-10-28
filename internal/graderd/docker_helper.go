package graderd

func ParseContainerState(status string) Status {
	// Container state can be one of "created", "running", "paused", "restarting", "removing", "exited", or "dead".
	switch status {
	case "created":
		return StatusPending
	case "running":
		return StatusStarted
	case "paused":

	case "restarting":
		return StatusStarted
	case "removing":
		return StatusComplete
	case "exited":
		return StatusComplete
	case "dead":
		return StatusFailed
	}
	return StatusFailed
}
