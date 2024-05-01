package test_contest_ozon

import (
	"bufio"
	"fmt"
	"os"
)

func Notifications() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var notificationID = 0
	var globalNotification = 0

	var usersCount, requestsCount int
	fmt.Fscan(in, &usersCount, &requestsCount)

	userNotificationMap := make([]int, usersCount+1)

	for i := 1; i <= usersCount; i++ {
		userNotificationMap[i] = 0
	}

	for i := 0; i < requestsCount; i++ {
		var requestType, userID int

		fmt.Fscan(in, &requestType, &userID)

		if requestType == 1 {
			notificationID++
			if userID == 0 {
				globalNotification = notificationID
			} else {
				userNotificationMap[userID] = notificationID
			}
		} else {
			if globalNotification > userNotificationMap[userID] {
				fmt.Fprintln(out, globalNotification)
			} else {
				fmt.Fprintln(out, userNotificationMap[userID])
			}
		}
	}
}
