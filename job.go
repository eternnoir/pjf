package pjf

import "context"

type Job interface {
	Init(config map[string]interface{}) error
	Process(ctx context.Context) error
}
