/*
Copyright 2024.

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

package v1

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// nolint:unused
// log is for logging in this package.
var podlog = logf.Log.WithName("pod-resource")

// SetupPodWebhookWithManager registers the webhook for Pod in the manager.
func SetupPodWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&corev1.Pod{}).
		WithDefaulter(&PodCustomDefaulter{
			c: mgr.GetClient(),
		}).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate--v1-pod,mutating=true,failurePolicy=fail,sideEffects=None,groups="",resources=pods,verbs=create;update,versions=v1,name=mpod-v1.kb.io,admissionReviewVersions=v1

// PodCustomDefaulter struct is responsible for setting default values on the custom resource of the
// Kind Pod when those are created or updated.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as it is used only for temporary operations and does not need to be deeply copied.
type PodCustomDefaulter struct {
	c client.Client
}

// +kubebuilder:rbac:groups="",resources=namespaces,verbs=get;list;watch

var _ webhook.CustomDefaulter = &PodCustomDefaulter{}

// Default implements webhook.CustomDefaulter so a webhook will be registered for the Kind Pod.
func (d *PodCustomDefaulter) Default(ctx context.Context, obj runtime.Object) error {
	pod, ok := obj.(*corev1.Pod)

	if !ok {
		return fmt.Errorf("expected an Pod object but got %T", obj)
	}
	podlog.Info("Defaulting for Pod", "name", pod.GetName())

	// TODO(user): fill in your defaulting logic.
	ns := pod.GetNamespace()
	nsObj := &corev1.Namespace{}
	if err := d.c.Get(ctx, client.ObjectKey{Name: ns}, nsObj); err != nil {
		return fmt.Errorf("failed to get namespace: %w", err)
	}
	nsLabel, ok := nsObj.ObjectMeta.Labels["lxcfs.lcpu.dev/inject"]
	if !ok || nsLabel != "enabled" {
		if nsLabel != "conditional" {
			return nil
		}
		if an, ok := pod.ObjectMeta.Annotations["lxcfs.lcpu.dev/inject"]; !ok || an != "enabled" {
			return nil
		}
	}
	if an, ok := pod.ObjectMeta.Annotations["lxcfs.lcpu.dev/inject"]; ok && an == "disabled" {
		return nil
	}

	pod.Spec.Volumes = append(pod.Spec.Volumes, []corev1.Volume{
		{
			Name: "lxcfs-proc-cpuinfo",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/var/lib/lxcfs/proc/cpuinfo",
					Type: newHostPathType(corev1.HostPathFile),
				},
			},
		},
		{
			Name: "system-cpu-online",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/var/lib/lxcfs/sys/devices/system/cpu/online",
					Type: newHostPathType(corev1.HostPathFile),
				},
			},
		},
		{
			Name: "lxcfs-proc-diskstats",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/var/lib/lxcfs/proc/diskstats",
					Type: newHostPathType(corev1.HostPathFile),
				},
			},
		},
		{
			Name: "lxcfs-proc-meminfo",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/var/lib/lxcfs/proc/meminfo",
					Type: newHostPathType(corev1.HostPathFile),
				},
			},
		},
		{
			Name: "lxcfs-proc-stat",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/var/lib/lxcfs/proc/stat",
					Type: newHostPathType(corev1.HostPathFile),
				},
			},
		},
		{
			Name: "lxcfs-proc-swaps",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/var/lib/lxcfs/proc/swaps",
					Type: newHostPathType(corev1.HostPathFile),
				},
			},
		},
		{
			Name: "lxcfs-proc-uptime",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/var/lib/lxcfs/proc/uptime",
					Type: newHostPathType(corev1.HostPathFile),
				},
			},
		},
		{
			Name: "lxcfs-proc-slabinfo",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/var/lib/lxcfs/proc/slabinfo",
					Type: newHostPathType(corev1.HostPathFile),
				},
			},
		},
	}...)

	for i := range pod.Spec.Containers {
		pod.Spec.Containers[i].VolumeMounts = append(pod.Spec.Containers[i].VolumeMounts, []corev1.VolumeMount{
			{
				Name:      "lxcfs-proc-cpuinfo",
				MountPath: "/proc/cpuinfo",
			},
			{
				Name:      "system-cpu-online",
				MountPath: "/sys/devices/system/cpu/online",
			},
			{
				Name:      "lxcfs-proc-diskstats",
				MountPath: "/proc/diskstats",
			},
			{
				Name:      "lxcfs-proc-meminfo",
				MountPath: "/proc/meminfo",
			},
			{
				Name:      "lxcfs-proc-stat",
				MountPath: "/proc/stat",
			},
			{
				Name:      "lxcfs-proc-swaps",
				MountPath: "/proc/swaps",
			},
			{
				Name:      "lxcfs-proc-uptime",
				MountPath: "/proc/uptime",
			},
			{
				Name:      "lxcfs-proc-slabinfo",
				MountPath: "/proc/slabinfo",
			},
		}...)
	}

	return nil
}

func newHostPathType(hostPathType corev1.HostPathType) *corev1.HostPathType {
	return &hostPathType
}
