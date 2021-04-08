/*
Copyright IBM Corporation 2020

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package optimize

import (
	irtypes "github.com/konveyor/move2kube/internal/types"
	core "k8s.io/kubernetes/pkg/apis/core"
)

//envVarsOptimizer implements Optimizer interface
type envVarsOptimizer struct {
}

//Optimize uses data from ir containers to fill ir.services
func (opt *envVarsOptimizer) optimize(ir irtypes.IR) (irtypes.IR, error) {
	for serviceName, service := range ir.Services {
		for coreContainerIdx, coreContainer := range service.Containers {
			if irContainer, ok := ir.GetContainer(coreContainer.Image); ok {
				if ir.Services[serviceName].Containers[coreContainerIdx].Env == nil {
					ir.Services[serviceName].Containers[coreContainerIdx].Env = []core.EnvVar{}
				} else {
					for envVarName, envVarValue := range irContainer.EnvVars {
						found := false
						for _, pEnvVar := range coreContainer.Env {
							if pEnvVar.Name == envVarName {
								found = true
							}
						}
						if !found {
							ir.Services[serviceName].Containers[coreContainerIdx].Env = append(ir.Services[serviceName].Containers[coreContainerIdx].Env, core.EnvVar{Name: envVarName, Value: envVarValue})
						}
					}
				}
			}
		}
		ir.Services[serviceName] = service
	}

	return ir, nil
}
