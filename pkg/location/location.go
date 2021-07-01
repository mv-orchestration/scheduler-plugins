package location

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

type Location struct {
	frameworkHandler framework.Handle
	nodes            map[string]*v1.Node
}

var _ framework.FilterPlugin = &Location{}
var _ framework.PreScorePlugin = &Location{}
var _ framework.ScorePlugin = &Location{}

const (
	Name = "Location"
)

func New(_ runtime.Object, handle framework.Handle) (framework.Plugin, error) {
	plugin := &Location{
		frameworkHandler: handle,
	}

	return plugin, nil
}

func (l Location) Name() string {
	return Name
}
