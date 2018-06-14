// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sum

import "testing"

func TestRecursive(t *testing.T) {
	tt := []struct {
		name    string
		numbers []int
		sum     int
	}{
		{"one to five", []int{1, 2, 3, 4, 5}, 15},
		{"empty slice", []int{}, 0},
		{"nil slice", nil, 0},
		{"one and minus one", []int{-1, 1}, 0},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := recursive(tc.numbers)
			if s != tc.sum {
				t.Errorf("expected sum to be %d; got %d", tc.sum, s)
			}
		})
	}

	// 	expected := 6
	// 	actual := All(1, 2, 3)
	// 	if expected != actual {
	// 		t.Errorf("expected %d to eq %d", actual, expected)
	// 	}
}
