package main

import "testing"

func TestGetChampions(t *testing.T) {
	tests := []struct {
		name       string
		statistics [][]Statistics
		userIds    []int
		steps      int
	}{
		{
			"check test 1",
			[][]Statistics{
				{{UserId: 1, Steps: 1000}, {UserId: 2, Steps: 1500}},
				{{UserId: 2, Steps: 1000}}},
			[]int{2}, 2500,
		},

		{
			"check test 2",
			[][]Statistics{
				{{UserId: 1, Steps: 2000}, {UserId: 2, Steps: 1500}},
				{{UserId: 2, Steps: 4000}, {UserId: 1, Steps: 3500}}},
			[]int{1, 2}, 5500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := getChampions(tt.statistics)

			// Check len arrs
			if len(tt.userIds) != len(res.UserIds) {
				t.Errorf("Need: %v | Real: %v\n", len(tt.userIds), len(res.UserIds))
			}

			// Check elements in arrs
			checkMap := make(map[int]struct{})
			for _, v := range res.UserIds {
				checkMap[v] = struct{}{}
			}
			for _, v := range tt.userIds {
				delete(checkMap, v)
			}
			if len(checkMap) != 0 {
				t.Errorf("Need: %v | Real: %v\n", tt.userIds, res.UserIds)
			}

			// Check steps
			if tt.steps != res.Steps {
				t.Errorf("Need: %v | Real: %v\n", tt.steps, res.Steps)
			}
		})
	}
}
