package tmplfunc

// SanitizeRequest converts boolean to string
func SanitizeRequest(status bool) string {

	list := map[bool]string{
		true:  "Accepted",
		false: "Pending",
	}

	return list[status]
}
