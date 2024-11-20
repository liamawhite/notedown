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

package daily

import "github.com/notedownorg/notedown/pkg/providers/pkg/collections"

type Fetcher = collections.Fetcher[DailyClient, Daily]
type ListOption = collections.ListOption[Daily]

func (c *DailyClient) DailyNoteSummary() int {
	c.notesMutex.RLock()
	defer c.notesMutex.RUnlock()
	return len(c.notes)
}

// Opts are applied in order so filters should be applied before sorters
func (c *DailyClient) ListDailyNotes(fetcher Fetcher, opts ...ListOption) []Daily {
	return collections.List(c, fetcher, opts...)
}
