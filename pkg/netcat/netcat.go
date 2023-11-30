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

package ssh

import (
	"os/exec"

	"github.com/frabits/frabit/pkg/common/constant"
)

type Netcat struct {
	BinPath    string
	DestAddr   string
	ListenPort uint32
}

func NewNetcat() *Netcat {
	nc := &Netcat{}
	nc.BinPath = exec.LookPath(constant.NETCAT)
	return nc
}
