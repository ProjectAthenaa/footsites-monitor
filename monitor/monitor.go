package monitor

import (
	"fmt"
	monitor_controller "github.com/ProjectAthenaa/sonic-core/protos/monitorController"
	"github.com/ProjectAthenaa/sonic-core/sonic/base"
	"github.com/ProjectAthenaa/sonic-core/sonic/face"
	"github.com/prometheus/common/log"
)

var _ face.MonitorCallback = (*Task)(nil)

type Task struct {
	*base.BMonitor

	stockendpoint string
	pid   string
	site string
}

func NewTask(data *monitor_controller.Task) (*Task, error) {
	task := &Task{
		BMonitor:    &base.BMonitor{Data: data},
		pid:         data.Metadata["LOOKUP_pid"],
		site: 		 data.Metadata["SITE"],
	}

	task.Callback = task

	return task, nil
}

func (tk *Task) TaskLoop() {
	var err error
	for {
		select {
		case <-tk.Ctx.Done():
			return
		default:
			if err = tk.iteration(); err != nil {
				log.Error("error completing iteration", err)
			}
		}
	}
}

func (tk *Task) OnStarting() {
	tk.stockendpoint = fmt.Sprintf("https://www.%s.com/zgw/product-core/v1/pdp/FL/sku/%s", tk.site, tk.pid)
}

func (tk *Task) OnStopping() {

}
