package redis

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
)

const (
	// Name is plugin name
	Name = "redis-scheduler"
)

// Args 调度器配置说明.
type Args struct {
	// 节点数量
	NodeNum int `json:"node_num,omitempty"`
}

type RedisScheduler struct {
	args   *Args
	handle framework.FrameworkHandle
}

func (s *RedisScheduler) Name() string {
	return Name
}

// New new 一个plugin.
func New(plArgs *runtime.Unknown, handle framework.FrameworkHandle) (framework.Plugin, error) {
	args := &Args{}
	if err := framework.DecodeInto(plArgs, args); err != nil {
		return nil, err
	}
	klog.V(3).Infof("--------> args: %+v", args)
	klog.V(3).Infof("-------------- handler %v", handle.NodeInfoSnapshot())
	return &RedisScheduler{
		args:   args,
		handle: handle,
	}, nil
}
