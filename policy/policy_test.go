package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	signal := Signal{Demand: 74, Capacity: 92, Latency: 19, Risk: 17, Weight: 13}
	if got := Score(signal); got != 71 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 71, Capacity: 89, Latency: 24, Risk: 20, Weight: 13}
	if got := Score(signal); got != 21 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 78, Capacity: 97, Latency: 22, Risk: 20, Weight: 5}
	if got := Score(signal); got != 35 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
}
