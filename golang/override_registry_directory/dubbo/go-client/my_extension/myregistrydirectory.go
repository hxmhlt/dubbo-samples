/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package my_extension

import (
	"fmt"
)

import (
	"github.com/apache/dubbo-go/cluster"
	"github.com/apache/dubbo-go/common"
	"github.com/apache/dubbo-go/protocol"
	"github.com/apache/dubbo-go/registry"
)

type myRegistryDirectory struct {
	registry registry.Registry
}

func (myRegistryDirectory) GetUrl() common.URL {
	panic("implement me")
}

func (myRegistryDirectory) IsAvailable() bool {
	panic("implement me")
}

func (myRegistryDirectory) Destroy() {
	panic("implement me")
}

func (myRegistryDirectory) List(invocation protocol.Invocation) []protocol.Invoker {
	panic("implement me")
}

//subscribe from registry
func (dir *myRegistryDirectory) subscribe(url *common.URL) {
	dir.registry.Subscribe(url, dir)
}

func (dir *myRegistryDirectory) Notify(event *registry.ServiceEvent) {
	fmt.Println("get event success!!!!!!!!")
	fmt.Println(event)
}

func NewMyRegistryDirectory(url *common.URL, registry registry.Registry) (cluster.Directory, error) {
	dir := &myRegistryDirectory{
		registry: registry,
	}
	go dir.subscribe(url.SubURL)
	return dir, nil
}
