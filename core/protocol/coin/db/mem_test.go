// Copyright © 2017-2018  Stratumn SAS
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

package db

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelDB(t *testing.T) {
	var tmpDirs []string

	defer func() {
		for _, dir := range tmpDirs {
			os.RemoveAll(dir)
		}
	}()

	testImplementation(t, func(*testing.T) DB {
		filename, err := ioutil.TempDir("", "")
		assert.NoError(t, err, "ioutil.TempDir()")
		tmpDirs = append(tmpDirs, filename)

		db, err := OpenLevelDBFile(filename, nil)
		assert.NoError(t, err, "OpenLevelDBFile()")

		return db
	})
}
