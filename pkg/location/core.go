package location

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/kubernetes/pkg/scheduler/framework"
	"strings"
)

func (l Location) getNodeLocations(nodeName string) (string, string, string) {
	node, _ := l.nodes[nodeName]
	ls := node.Labels

	return ls[NodeCity], ls[NodeCountry], ls[NodeContinent]
}

func isAdmissibleNode(nodeInfo *framework.NodeInfo) bool {
	_, ok := nodeInfo.Node().Labels[Node]
	return ok
}

func getLocationLabelType(pod *v1.Pod) string {
	if pod.Labels[WorkloadRequiredLocation] != "" {
		return "required"
	}

	if pod.Labels[WorkloadPreferredLocation] != "" {
		return "preferred"
	}

	return ""
}

func verifyNodeHasRequiredLocation(pod *v1.Pod, node *v1.Node) *framework.Status {
	cities, countries, continents := getRequiredLocations(pod)

	if findAnyLocation(cities, node, NodeCity) ||
		findAnyLocation(countries, node, NodeCountry) ||
		findAnyLocation(continents, node, NodeContinent) {
		return framework.NewStatus(framework.Success)
	}

	return framework.NewStatus(framework.Unschedulable,
		fmt.Sprintf("Node '%s' does not have any of the required locations", node.Name))
}

func getLocations(pod *v1.Pod) ([]string, []string, []string) {
	if pod.Labels[WorkloadRequiredLocation] != "" {
		return parseLocations(pod, WorkloadRequiredLocation)
	}

	if pod.Labels[WorkloadPreferredLocation] != "" {
		return parseLocations(pod, WorkloadPreferredLocation)
	}

	return make([]string, 0), make([]string, 0), make([]string, 0)
}

func getRequiredLocations(pod *v1.Pod) ([]string, []string, []string) {
	return parseLocations(pod, WorkloadRequiredLocation)
}

func parseLocations(pod *v1.Pod, label string) ([]string, []string, []string) {
	locations := pod.Labels[label]

	divisions := strings.Split(locations, "-")
	cities := strings.Split(divisions[0], "_")
	countries := strings.Split(divisions[1], "_")
	continents := strings.Split(divisions[2], "_")

	return cities, countries, continents
}

func findAnyLocation(locations []string, node *v1.Node, label string) bool {
	for _, location := range locations {
		nodeLocation, _ := node.Labels[label]
		if nodeLocation == location {
			return true
		}
	}

	return false
}

func getPartialScore(value string, items []string, increment int64) int64 {
	score := int64(0)

	for _, item := range items {
		if value == item {
			score += increment
		}
	}

	return score
}
