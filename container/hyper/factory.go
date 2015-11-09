// Copyright 2014 Google Inc. All Rights Reserved.
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

package hyper

import (
	"github.com/golang/glog"
	"github.com/google/cadvisor/container"
	"github.com/google/cadvisor/fs"
	info "github.com/google/cadvisor/info/v1"
)

const (
	// HyperNamespace is namespace under which Hyper aliases are unique.
	HyperNamespace = "hyper"
)

type hyperFactory struct {
	client             *HyperClient
	machineInfoFactory info.MachineInfoFactory
	fsInfo             fs.FsInfo
}

func (self *hyperFactory) String() string {
	return HyperNamespace
}

func (self *hyperFactory) NewContainerHandler(name string, inHostNamespace bool) (handler container.ContainerHandler, err error) {
	handler, err = newHyperContainerHandler(
		self.client,
		name,
		self.machineInfoFactory,
		self.fsInfo,
	)
	return
}

func (self *hyperFactory) CanHandleAndAccept(name string) (bool, bool, error) {
	_, err := isHyperVirtualMachine(name)
	if err != nil {
		return false, false, nil
	}

	return true, true, nil
}

func (self *hyperFactory) DebugInfo() map[string][]string {
	return map[string][]string{}
}

// Register root container before running this function!
func Register(factory info.MachineInfoFactory, fsInfo fs.FsInfo) error {
	glog.Infof("Registering Hyper factory")
	f := &hyperFactory{
		client:             NewHyperClient(),
		machineInfoFactory: factory,
		fsInfo:             fsInfo,
	}
	container.RegisterContainerHandlerFactory(f)
	return nil
}
