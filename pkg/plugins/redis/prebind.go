package redis

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/klog"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
)

var _ framework.PreBindPlugin = &RedisScheduler{}

func (s *RedisScheduler) PreBind(pc *framework.PluginContext, p *v1.Pod, nodeName string) *framework.Status {
	nodeMap := s.handle.NodeInfoSnapshot().NodeInfoMap
	if node, ok := nodeMap[nodeName]; ok {
		klog.V(3).Infof("prebind node info: %+v", node.String())
		return framework.NewStatus(framework.Success, "")
	}
	return framework.NewStatus(framework.Error, "node is not exsit! ")
}
