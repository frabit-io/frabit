// Frabit - The next-generation database automatic operation platform
// Copyright © 2022-2023 Frabit Labs
//
// Licensed under the GNU General Public License, Version 3.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.gnu.org/licenses/gpl-3.0.txt
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"github.com/frabits/frabit/pkg/common/constant"
)

type WorkspaceLevel struct {
	Id        uint32
	Workspace string
	User      string
	Admin     bool
	Editor    bool
	Viewer    bool
}

type ProjectLevel struct {
	Id        uint32
	Workspace string
	Project   string
	User      string
	Admin     bool
	Editor    bool
	Viewer    bool
}

type DatabaseLevel struct {
	Id        uint32
	Workspace string
	Project   string
	Database  string
	User      string
	Admin     bool
	Editor    bool
	Viewer    bool
}

type GeneralLevel struct {
	Id        uint32
	Workspace string
	Project   string
	Database  string
	User      string
	Admin     bool
	Editor    bool
	Viewer    bool
}

// IsVisible determine the visible for a user based
func IsVisible(level constant.PrivilegeLevel, role constant.Role, user string) bool {
	switch level {
	case constant.GLOBAL:
		return true
	case constant.WORKSPACE:
		return true
	case constant.PROJECT:
		return true
	case constant.DATABASE:
		return true
	default:
		return false
	}
}
