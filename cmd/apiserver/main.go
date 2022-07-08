/*
Copyright 2022.

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

package main

import (
	"k8s.io/klog"
	"sigs.k8s.io/apiserver-runtime/pkg/builder"

	// +kubebuilder:scaffold:resource-imports
	cbtv1alpha1 "github.com/ihcsim/cbt-controller/pkg/apis/cbt/v1alpha1"
	"github.com/ihcsim/cbt-controller/pkg/storage"
)

func main() {
	err := builder.APIServer.
		// +kubebuilder:scaffold:resource-register
		WithResourceAndHandler(&cbtv1alpha1.VolumeSnapshotDelta{}, storage.NewStorageProvider(&cbtv1alpha1.VolumeSnapshotDelta{})).
		WithLocalDebugExtension().
		Execute()
	if err != nil {
		klog.Fatal(err)
	}
}
