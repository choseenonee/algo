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

	var usersCount, requestsCount int
	fmt.Fscan(in, &usersCount, &requestsCount)

	userNotificationMap := map[int][]int{}

	for i := 1; i <= usersCount; i++ {
		userNotificationMap[i] = []int{0}
	}

	for i := 0; i < requestsCount; i++ {
		var requestType, userID int

		fmt.Fscan(in, &requestType, &userID)

		if requestType == 1 {
			notificationID++
			if userID == 0 {
				for key, val := range userNotificationMap {
					userNotificationMap[key] = append(val, notificationID)
				}
			} else {
				userNotificationMap[userID] = append(userNotificationMap[userID], notificationID)
			}
		} else {
			userNotifications := userNotificationMap[userID]
			fmt.Fprintln(out, userNotifications[len(userNotifications)-1])
		}
	}
}
