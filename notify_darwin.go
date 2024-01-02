package notify

import (
	"fmt"
	"os/exec"
)

var command = exec.Command

func (n *Notify) Send() error {
	// // if `terminal-notifier`` is not installed:
	// script := fmt.Sprintf(`display notification "%s" with title "(%s) %s"`, n.message, n.severity, n.title)
	// notifyCommand := command("osascript", "-e", script)
	// return notifyCommand.Run()
	notifyCmdName := "terminal-notifier"

	notifyCmd, err := exec.LookPath(notifyCmdName)
	if err != nil {
		return err
	}

	title := fmt.Sprintf("(%s) %s", n.severity, n.title)

	notifyCommand := command(notifyCmd, "-title", title, "-message", n.message)
	return notifyCommand.Run()
}
