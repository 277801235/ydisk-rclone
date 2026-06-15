// Package notify provides a cross-platform desktop notification helper
// for rclone. It wraps beeep to send system-native notifications, and
// never returns an error to the caller — notification failures are
// silently ignored so they never interrupt the main operation.
package notify

import (
	"github.com/rclone/rclone/fs"

	"github.com/gen2brain/beeep"
)

// AppName is the application name used in desktop notifications.
const AppName = "rclone"

func init() {
	beeep.AppName = AppName
}

// Send sends a desktop notification with the given title and message.
// If the notification fails for any reason (unsupported platform,
// missing system service, etc.), the error is logged at DEBUG level
// and silently discarded.
func Send(title, message string) {
	if err := beeep.Notify(title, message, ""); err != nil {
		fs.Debugf(nil, "desktop notification failed (title=%q): %v", title, err)
	}
}

// SendAlert sends an urgent desktop notification with a beep sound.
// Use for critical errors like permission denied or quota exceeded.
func SendAlert(title, message string) {
	if err := beeep.Alert(title, message, ""); err != nil {
		fs.Debugf(nil, "desktop alert failed (title=%q): %v", title, err)
	}
}
