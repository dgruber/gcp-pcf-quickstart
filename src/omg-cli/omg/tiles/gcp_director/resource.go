/*
 * Copyright 2017 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package gcp_director

import (
	"omg-cli/config"
)

var tile = config.Tile{
	Product: config.OpsManagerMetadata{
		Name: "BOSH Director",
	},
}

type Tile struct{}

func (*Tile) Definition(*config.EnvConfig) config.Tile {
	return tile
}
func (*Tile) BuiltIn() bool {
	return true
}

func (*Tile) NoConfig() bool {
	return false
}
