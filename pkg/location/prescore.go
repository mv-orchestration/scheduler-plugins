package location

import (
	"context"
	v1 "k8s.io/api/core/v1"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

func (l Location) PreScore(_ context.Context, _ *framework.CycleState, _ *v1.Pod, nodes []*v1.Node) *framework.Status {
	for _, node := range nodes {
		l.nodes[node.Name] = node
	}

	return framework.NewStatus(framework.Success)
}
