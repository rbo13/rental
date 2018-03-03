package helpers

// ConvertStatusToString ...
func ConvertStatusToString(status bool) string {

	list := map[bool]string{
		false: "Vacant",
		true:  "Occupied",
	}

	return list[status]
}
