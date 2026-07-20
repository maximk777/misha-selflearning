package messaging

type Record struct {
	Topic, Key, Value string
	Offset            int64
}

func Partition(key string, partitions int) int {
	sum := 0
	for _, r := range key {
		sum += int(r)
	}
	if partitions < 1 {
		return 0
	}
	return sum % partitions
}
