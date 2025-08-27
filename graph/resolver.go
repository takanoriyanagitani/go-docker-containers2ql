package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"context"

	tc "github.com/docker/docker/api/types/container"
	dc "github.com/docker/docker/client"
	"github.com/takanoriyanagitani/go-docker-containers2ql/graph/model"
)

type Resolver struct {
	*dc.Client
}

func (r *Resolver) List(
	ctx context.Context,
	opts tc.ListOptions,
) ([]tc.Summary, error) {
	return r.Client.ContainerList(ctx, opts)
}

func (r *Resolver) BasicInfo(
	ctx context.Context,
	i *model.BasicInfoInput,
) ([]*model.BasicInfoOutput, error) {
	var opts tc.ListOptions = ToListOptions(i)
	sa, e := r.List(ctx, opts)
	if nil != e {
		return nil, e
	}

	ret := make([]*model.BasicInfoOutput, 0, len(sa))

	for _, s := range sa {
		ret = append(ret, ToBasicInfoOutput(&s))
	}

	return ret, nil
}

func (r *Resolver) Inspect(
	ctx context.Context,
	id string,
) (tc.InspectResponse, error) {
	return r.Client.ContainerInspect(ctx, id)
}

func (r *Resolver) Health(
	ctx context.Context,
	basic *model.BasicInfoOutput,
) (*model.Health, error) {
	var id string = basic.ID
	ins, e := r.Inspect(ctx, id)
	if nil != e {
		return nil, e
	}
	return InspectResponseToHealth(ins), nil
}
