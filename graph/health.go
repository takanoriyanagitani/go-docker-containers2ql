package graph

import (
	"time"

	tc "github.com/docker/docker/api/types/container"
	"github.com/takanoriyanagitani/go-docker-containers2ql/graph/model"
)

func ToHealthCheckResult(c *tc.HealthcheckResult) *model.HealthcheckResult {
	return &model.HealthcheckResult{
		Start:    c.Start.Format(time.RFC3339),
		End:      c.End.Format(time.RFC3339),
		ExitCode: int64(c.ExitCode),
		Output:   c.Output,
	}
}

func ToHealth(h *tc.Health) *model.Health {
	var m model.Health

	m.Status = model.HealthStatus(h.Status)
	m.FailingStreak = int64(h.FailingStreak)
	m.Log = make([]*model.HealthcheckResult, 0, len(h.Log))

	// Copy the log data
	for _, hr := range h.Log {
		m.Log = append(m.Log, ToHealthCheckResult(hr))
	}

	return &m
}

func StateToHealth(s *tc.State) *tc.Health { return s.Health }

func InspectResponseToState(i tc.InspectResponse) *tc.State { return i.State }

func InspectResponseToHealth(i tc.InspectResponse) *model.Health {
	var s *tc.State = InspectResponseToState(i)
	var h *tc.Health = StateToHealth(s)
	return ToHealth(h)
}
