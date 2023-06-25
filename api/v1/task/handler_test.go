package task

import (
	"github.com/sword/api-backend-challenge/model"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	testCases := []struct {
		Name        string
		Request     *model.Task
		Response    *model.Task
		ResponseErr error
	}{
		{
			Name: "Create Success",
			Request: &model.Task{
				ID:          "",
				Summary:     "",
				PerformedAt: time.Time{},
			},
		},
	}
}
