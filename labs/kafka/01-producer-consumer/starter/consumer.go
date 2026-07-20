package messaging

// CommitAfterSuccess records the at-least-once rule: only a handled record advances offset.
func CommitAfterSuccess(handled bool, offset int64) int64 {
	if handled {
		return offset
	}
	return -1
}
