package container

import (
	"context"
	"io"

	"github.com/cpuguy83/go-docker"

	"github.com/docker/docker/api/types"
)

type Container interface {
	Attach(context.Context, ...AttachOption) (AttachIO, error)
	Inspect(context.Context) (types.ContainerJSON, error)
	Logs(context.Context, ...LogsReadOption) (io.ReadCloser, error)
	Start(context.Context, ...StartOption) error
	Kill(context.Context, ...KillOption) error
	Stop(context.Context, ...StopOption) error
	ID() string
}

type container struct {
	id     string
	client *docker.Client
}

func (c *container) ID() string {
	return c.id
}