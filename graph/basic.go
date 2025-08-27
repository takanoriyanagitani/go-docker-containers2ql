package graph

import (
	"strconv"

	tc "github.com/docker/docker/api/types/container"
	tf "github.com/docker/docker/api/types/filters"
	"github.com/takanoriyanagitani/go-docker-containers2ql/graph/model"
)

func ToContainerState(cs tc.ContainerState) model.ContainerState {
	switch cs {
	case tc.StateCreated:
		return model.ContainerStateCreated
	case tc.StateRunning:
		return model.ContainerStateRunning
	case tc.StatePaused:
		return model.ContainerStatePaused
	case tc.StateRestarting:
		return model.ContainerStateRestarting
	case tc.StateRemoving:
		return model.ContainerStateRemoving
	case tc.StateExited:
		return model.ContainerStateExited
	case tc.StateDead:
		return model.ContainerStateDead
	default:
		return model.ContainerStateUnspecified
	}
}

func ToBasicInfoOutput(summary *tc.Summary) *model.BasicInfoOutput {
	return &model.BasicInfoOutput{
		ID:      summary.ID,
		Names:   summary.Names,
		Image:   summary.Image,
		ImageID: summary.ImageID,
		Command: summary.Command,
		Created: summary.Created,
		State:   ToContainerState(summary.State),
		Status:  summary.Status,
	}
}

type basicInfoInput struct{ *model.BasicInfoInput }

func (b basicInfoInput) AddExited(a tf.Args) {
	if nil == b.Exited {
		return
	}

	var es string = strconv.Itoa(int(*b.Exited))

	a.Add("exited", es)
}

func (b basicInfoInput) AddHealthStatus(a tf.Args) {
	if nil == b.Health {
		return
	}

	var hs string = b.Health.String()

	a.Add("health", hs)
}

func (b basicInfoInput) AddStatus(a tf.Args) {
	if nil == b.Status {
		return
	}

	var ss string = b.Status.String()

	a.Add("status", ss)
}

func (b basicInfoInput) AddString(a tf.Args, key string, val *string) {
	if nil == val {
		return
	}

	a.Add(key, *val)
}

func (b basicInfoInput) AddStrings(a tf.Args) {
	b.AddString(a, "id", b.ID)
	b.AddString(a, "name", b.Name)
	b.AddString(a, "network", b.Network)
}

func (b basicInfoInput) ToArgs() tf.Args {
	var a tf.Args = tf.NewArgs()

	b.AddExited(a)
	b.AddHealthStatus(a)
	b.AddStatus(a)
	b.AddStrings(a)

	return a
}

func ToListOptions(input *model.BasicInfoInput) tc.ListOptions {
	if nil == input {
		return tc.ListOptions{
			Size:    false,
			All:     false,
			Latest:  false,
			Since:   "",
			Before:  "",
			Limit:   0,
			Filters: tf.NewArgs(),
		}
	}

	var lmt int
	if nil != input.Limit {
		lmt = int(*input.Limit)
	}
	return tc.ListOptions{
		Size:    input.Size,
		All:     input.All,
		Latest:  false,
		Since:   "",
		Before:  "",
		Limit:   lmt,
		Filters: basicInfoInput{input}.ToArgs(),
	}
}
