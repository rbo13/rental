package tmplfunc

// ConvertStatusToString converts boolean to string
func ConvertStatusToString(status bool) string {

	list := map[bool]string{
		true:  "Vacant",
		false: "Occupied",
	}

	return list[status]
}
