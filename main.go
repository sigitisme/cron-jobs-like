package main

import (
	"time"
	// "github.com/astaxie/beego"
	"fmt"

	"github.com/astaxie/beego/toolbox"
)

func init() {
	perMinuteTask := toolbox.NewTask("perMinuteTask", "0 */1 * * * *", func() error {
		fmt.Printf("time: %v it is per minute task\n", time.Now().Format("2006-01-02 15:04:05"))

		return nil
	})

	perSecondTask := toolbox.NewTask("perSecondTask", "*/1 * * * * *", func() error {
		fmt.Printf("time: %v it is per second task\n", time.Now().Format("2006-01-02 15:04:05"))

		return nil
	})

	toolbox.AddTask("perMinuteTask", perMinuteTask)
	toolbox.AddTask("perSecondTask", perSecondTask)

	toolbox.StartTask()

}

func main() {
	// beego.Run()
	for {}
}
