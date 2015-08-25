/***
Copyright 2014 Cisco Systems Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/contiv/objmodel/tools/modelgen/generators"
)

// Simple test to parse json schema
func TestParseJsonSchema(t *testing.T) {
	if err := generators.ParseTemplates(); err != nil {
		t.Fatal(err)
	}

	dir, err := os.Open("testdata")
	if err != nil {
		t.Fatal(err)
	}

	dirnames, err := dir.Readdirnames(0)
	if err != nil {
		t.Fatal(err)
	}

	for _, name := range dirnames {
		t.Logf("Parsing suite %s", name)
		basepath := filepath.Join("testdata", name)
		input, err := ioutil.ReadFile(filepath.Join(basepath, "input.json"))
		if err != nil {
			t.Fatal(err)
		}

		// Parse the input json string
		schema, err := ParseSchema(input)
		if err != nil {
			t.Fatalf("Error parsing json schema. Err: %v", err)
		}

		// Generate the code
		goStr, err := schema.GenerateGo()
		if err != nil {
			t.Fatalf("Error generating go code. Err: %v", err)
		}

		output, err := ioutil.ReadFile(filepath.Join(basepath, "output.go"))
		if err != nil {
			t.Fatal(err)
		}

		if goStr != string(output) {
			fmt.Printf("Generated string:\n%s\n", goStr)
			t.Fatalf("Generated string from input was not equal to output string")
		}
	}
}
