package model

import (
	"fmt"
)

type JobNumber uint32

const NullJobNumber JobNumber = 0

func (n JobNumber) String() string {
	return fmt.Sprintf("%d", n)
}
