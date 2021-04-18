package location

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

func (l Location) Filter(_ context.Context, _ *framework.CycleState, pod *v1.Pod, nodeInfo *framework.NodeInfo,
) *framework.Status {

	if nodeInfo.Node() == nil {
		return framework.NewStatus(framework.Error, fmt.Sprintf("Node is nil '%s'", nodeInfo.Node().Name))
	}

	if !isAdmissibleNode(nodeInfo) {
		return framework.NewStatus(framework.Error,
			fmt.Sprintf("Node '%s' does not have the '%s' label", nodeInfo.Node().Name, Node))
	}

	if queryType := getLocationLabelType(pod); queryType == "preferred" {
		return verifyNodeHasRequiredLocation(pod, nodeInfo.Node())
	}

	return framework.NewStatus(framework.Success)
}
