/*
Licensed under the Apache License, Version 2.0 (the "License"); you may not use
this file except in compliance with the License. You may obtain a copy of the
License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed
under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
CONDITIONS OF ANY KIND, either express or implied. See the License for the
specific language governing permissions and limitations under the License.
*/

package fake

import (
	"errors"
)

type NodeGroup struct {
	WantErr  error
	Stable   bool
	Message  string
	Replicas *int32
}

func (n *NodeGroup) GetReplicas() (int32, error) {
	if n.WantErr != nil {
		return 0, n.WantErr
	}
	if n.Replicas != nil {
		return *n.Replicas, nil
	}
	return 0, errors.New("Replicas for fake.NodeGroup was nil, try setting fake.Factory.NodeReplicas.")
}

func (n *NodeGroup) SetReplicas(count int32) error {
	if n.WantErr != nil {
		return n.WantErr
	}
	*n.Replicas = count
	return nil
}

func (n *NodeGroup) Stabilized() (bool, string, error) {
	return n.Stable, n.Message, nil
}
