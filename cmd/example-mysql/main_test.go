// Copyright 2024 Google LLC. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"testing"
)

func TestParseTileLevelIndexWidth(t *testing.T) {
	for _, test := range []struct {
		pathLevel string
		pathIndex string
		wantLevel uint64
		wantIndex uint64
		wantWidth uint64
		wantErr   bool
	}{
		{
			pathLevel: "0",
			pathIndex: "x001/x234/067",
			wantLevel: 0,
			wantIndex: 1234067,
			wantWidth: 256,
		},
		{
			pathLevel: "0",
			pathIndex: "x001/x234/067.p/89",
			wantLevel: 0,
			wantIndex: 1234067,
			wantWidth: 89,
		},
		{
			pathLevel: "63",
			pathIndex: "x999/x999/x999/x999/x999/999.p/255",
			wantLevel: 63,
			wantIndex: 999999999999999999,
			wantWidth: 255,
		},
		{
			pathLevel: "0",
			pathIndex: "001",
			wantLevel: 0,
			wantIndex: 1,
			wantWidth: 256,
		},
		{
			pathLevel: "0",
			pathIndex: "x001/x234/067.p/",
			wantErr:   true,
		},
		{
			pathLevel: "0",
			pathIndex: "x001/x234/067.p",
			wantErr:   true,
		},
		{
			pathLevel: "0",
			pathIndex: "x001/x234/",
			wantErr:   true,
		},
		{
			pathLevel: "0",
			pathIndex: "x001/x234",
			wantErr:   true,
		},
		{
			pathLevel: "0",
			pathIndex: "x001/",
			wantErr:   true,
		},
		{
			pathLevel: "0",
			pathIndex: "x001",
			wantErr:   true,
		},
		{
			pathLevel: "1",
			pathIndex: "x001/.p/abc",
			wantErr:   true,
		},
		{
			pathLevel: "64",
			pathIndex: "x001/002",
			wantErr:   true,
		},
		{
			pathLevel: "-1",
			pathIndex: "x001/002",
			wantErr:   true,
		},
		{
			pathLevel: "abc",
			pathIndex: "x001/002",
			wantErr:   true,
		},
		{
			pathLevel: "8",
			pathIndex: "001/002",
			wantErr:   true,
		},
		{
			pathLevel: "8",
			pathIndex: "x001/0002",
			wantErr:   true,
		},
		{
			pathLevel: "8",
			pathIndex: "x001/-002",
			wantErr:   true,
		},
		{
			pathLevel: "8",
			pathIndex: "x001/002.p/256",
			wantErr:   true,
		},
	} {
		desc := fmt.Sprintf("pathLevel: %q, pathIndex: %q", test.pathLevel, test.pathIndex)
		t.Run(desc, func(t *testing.T) {
			gotLevel, gotIndex, gotWidth, err := parseTileLevelIndexWidth(test.pathLevel, test.pathIndex)
			if gotLevel != test.wantLevel {
				t.Errorf("got level %d want %d", gotLevel, test.wantLevel)
			}
			if gotIndex != test.wantIndex {
				t.Errorf("got index %d want %d", gotIndex, test.wantIndex)
			}
			if gotWidth != test.wantWidth {
				t.Errorf("got width %d want %d", gotWidth, test.wantWidth)
			}
			gotErr := err != nil
			if gotErr != test.wantErr {
				t.Errorf("got err %v want %v", gotErr, test.wantErr)
			}
		})
	}
}
