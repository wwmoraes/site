package hugolite

import (
	"hash/fnv"

	"github.com/spf13/cast"
)

func FNV32a(v any) (int, error) {
	conv, err := cast.ToStringE(v)
	if err != nil {
		return 0, err
	}

	algorithm := fnv.New32a()
	algorithm.Write([]byte(conv))

	return int(algorithm.Sum32()), nil
}
