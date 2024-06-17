/*
Copyright 2022 The Katalyst Authors.

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

package cpu

import (
	"github.com/kubewharf/katalyst-core/cmd/katalyst-agent/app/agent/qrm"
	"github.com/kubewharf/katalyst-core/pkg/agent/qrm-plugins/cpu/consts"
	"github.com/kubewharf/katalyst-core/pkg/agent/qrm-plugins/cpu/dynamicpolicy"
	"github.com/kubewharf/katalyst-core/pkg/agent/qrm-plugins/cpu/nativepolicy"
	"github.com/kubewharf/katalyst-core/pkg/agent/qrm-plugins/cpu/vcipolicy"
)

func init() {
	qrm.RegisterCPUPolicyInitializer(consts.CPUResourcePluginPolicyNameDynamic, dynamicpolicy.NewDynamicPolicy)
	qrm.RegisterCPUPolicyInitializer(consts.CPUResourcePluginPolicyNameNative, nativepolicy.NewNativePolicy)
	qrm.RegisterCPUPolicyInitializer(consts.CPUResourcePluginPolicyNameStatic, vcipolicy.NewStaticPolicy)
}
