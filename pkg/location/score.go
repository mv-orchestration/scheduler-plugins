package location

import (
	"context"
	v1 "k8s.io/api/core/v1"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

func (l Location) Score(_ context.Context, _ *framework.CycleState, pod *v1.Pod, nodeName string) (int64, *framework.Status) {
	nodeCity, nodeCountry, nodeContinent := l.getNodeLocations(nodeName)
	cities, countries, continents := getLocations(pod)

	score := getPartialScore(nodeCity, cities, 30)
	score += getPartialScore(nodeCountry, countries, 20)
	score += getPartialScore(nodeContinent, continents, 10)

	return score, framework.NewStatus(framework.Success)
}

func (l Location) ScoreExtensions() framework.ScoreExtensions {
	return nil
}
