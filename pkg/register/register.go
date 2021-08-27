package register

import (
	"github.com/humorliang/scheduler-framework/pkg/plugins/redis"
	"github.com/spf13/cobra"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
)

// Register 注册一个插件.
func Register() *cobra.Command {
	return app.NewSchedulerCommand(
		app.WithPlugin(redis.Name, redis.New),
	)
}
