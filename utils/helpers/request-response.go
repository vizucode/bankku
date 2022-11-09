package helpers

func SuccessGetResponseData(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "Success",
		"message": "Success",
		"data":    data,
	}
}

func SuccessActionResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": msg,
	}
}

func FailedResponse(status string, message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  status,
		"message": message,
		"data":    map[string]interface{}{},
	}
}
