package notify

import (
	toast "github.com/jacobmarshall/go-toast"
	"time"
)

const (
	notificationBreakpoint = 5 * time.Minute
)

// SendNotification sends a platform specific desktop notification.
func SendNotification(title string, message string, elapsed time.Duration) error {

	if elapsed < notificationBreakpoint {
		return nil
	}

	notification := toast.Notification{
		AppID:   "WPDS",
		Title:   title,
		Message: message,
	}

	return notification.Push()

}