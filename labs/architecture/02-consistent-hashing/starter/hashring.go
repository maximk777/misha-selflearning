package hashring

import (
	"hash/fnv"
	"sort"
)

type Ring struct{ nodes []string }

func New(nodes []string) *Ring {
	copy := append([]string(nil), nodes...)
	sort.Strings(copy)
	return &Ring{copy}
}
func (r *Ring) Node(key string) string {
	if len(r.nodes) == 0 {
		return ""
	}
	h := fnv.New32a()
	_, _ = h.Write([]byte(key))
	return r.nodes[int(h.Sum32())%len(r.nodes)]
}
