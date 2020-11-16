package pipline

import (
	"queuify/workers"
)

type Pipeline struct {
	workers []workers.Worker
}
