package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	tests := []struct {
		name         string
		signal       Signal
		wantScore    int
		wantDecision string
	}{
		{name: "case_1", signal: Signal{Demand: 74, Capacity: 92, Latency: 19, Risk: 17, Weight: 13}, wantScore: 71, wantDecision: "review"},
		{name: "case_2", signal: Signal{Demand: 71, Capacity: 89, Latency: 24, Risk: 20, Weight: 13}, wantScore: 21, wantDecision: "review"},
		{name: "case_3", signal: Signal{Demand: 78, Capacity: 97, Latency: 22, Risk: 20, Weight: 5}, wantScore: 35, wantDecision: "review"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Score(tc.signal); got != tc.wantScore {
				t.Fatalf("score = %d, want %d", got, tc.wantScore)
			}
			if got := Classify(tc.signal); got != tc.wantDecision {
				t.Fatalf("decision = %s, want %s", got, tc.wantDecision)
			}
		})
	}
}
