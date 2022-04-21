package user

import (
	"net/http"
	user "user/model"
)

func SuccessInsert(data user.User) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "berhasil insert data user",
		"status":  true,
		"data":    data,
	}
}

func SuccessUpdate(data user.User) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "berhasil update data user",
		"status":  true,
		"data":    data,
	}
}
func BadRequest() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "terdapat kesalahan pada input data user",
		"status":  false,
		"data":    nil,
	}
}
