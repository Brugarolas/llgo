/*
 * Copyright (c) 2023 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package llgo

import (
	"github.com/goplus/llgo/x/gocmd"
	"github.com/goplus/mod/gopmod"
)

// -----------------------------------------------------------------------------

// NotFound returns if cause err is ErrNotFound or not
func NotFound(err error) bool {
	return gopmod.IsNotFound(err)
}

// -----------------------------------------------------------------------------

func BuildDir(dir string, conf *Config, build *gocmd.BuildConfig) (err error) {
	panic("todo")
}

func BuildPkgPath(workDir, pkgPath string, conf *Config, build *gocmd.BuildConfig) (err error) {
	panic("todo")
}

func BuildFiles(files []string, conf *Config, build *gocmd.BuildConfig) (err error) {
	panic("todo")
}

// -----------------------------------------------------------------------------
