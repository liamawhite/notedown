// Copyright 2024 Notedown Authors
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

package tasks

import (
	"github.com/notedownorg/notedown/pkg/ast"
	"github.com/notedownorg/notedown/pkg/workspace/documents/reader"
)

type TaskFilter func(ast.Task) bool

// Priorities are OR'd together
func FilterByPriority(priority ...int) TaskFilter {
	return func(task ast.Task) bool {
		for _, p := range priority {
			taskPriority := task.Priority()
			if taskPriority != nil && *taskPriority == p {
				return true
			}
		}
		return false
	}
}

func FilterByStatus(status ...ast.Status) TaskFilter {
	return func(task ast.Task) bool {
		for _, s := range status {
			if task.Status() == s {
				return true
			}
		}
		return false
	}
}

type DocumentFilter func(path string, document reader.Document) bool

func FilterByDocumentType(documentType string) DocumentFilter {
	return func(_ string, document reader.Document) bool {
		return document.Metadata["type"] == documentType
	}
}
