// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package devrel_tutorial

import (
	// "fmt"
	"testing"

  // tutorial "github.com/googlecodelabs/tools/third_party/"
)

// Make into super package ~ although non-fatal/with "log" package? :D
func logIfError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

// TODO: Eventually reverse order to match proto definitions ~
// Test all Codelab Proto Element Definitions
func TestProtoDefinitionsHeading(t *testing.T) {
	element := &Heading{
		Text:  "hello",
		Level: 17,
	}
	t.Logf("%+v\n", element)
}

func TestProtoDefinitionsLink(t *testing.T) {
	element := &Link{
		Text:     "hello",
		Location: "https://codelabs.developers.google.com",
	}
	t.Logf("%+v\n", element)
}

// Left:
// Tutorial
// Step
// BlockContent
// InlineContent
// Button
// InfoBox
// CodeBlock
// Image
// List
// ListStep
// Survey
// Checklist
// FAQ