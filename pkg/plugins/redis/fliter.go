package redis

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/klog"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
)

var _ framework.FilterPlugin = &RedisScheduler{}

func (s *RedisScheduler) Filter(pc *framework.PluginContext, pod *v1.Pod, nodeName string) *framework.Status {
	klog.V(3).Infof("filter pod: %v", pod.Name)
	return framework.NewStatus(framework.Success, "")
}
