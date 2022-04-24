package user

import (
	"net/http"
	mUser "user/model"
	"user/view"
	user "user/view/User"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type UserController struct {
	Repo  mUser.UserModel
	Valid *validator.Validate
}

func (uc *UserController) InsertNewUser(c echo.Context) error {
	var User user.InsertUserRequest

	if err := c.Bind(&User); err != nil {
		log.Warn("salah input")
		return c.JSON(http.StatusBadRequest, user.BadRequest())
	}

	if err := uc.Valid.Struct(User); err != nil {
		log.Warn(err.Error())
		return c.JSON(http.StatusBadRequest, user.BadRequest())
	}

	newUser := mUser.User{Nama: User.Nama, Email: User.Email, Password: User.Password}
	res, err := uc.Repo.Insert(newUser)

	if err != nil {
		log.Warn("masalah pada server")
		return c.JSON(http.StatusInternalServerError, view.InternalServerError())
	}
	log.Info("berhasil insert")
	return c.JSON(http.StatusCreated, user.SuccessInsert(res))
}

func (uc *UserController) GetAllUser(c echo.Context) error {

	res, err := uc.Repo.GetAll()

	if err != nil {
		log.Warn("masalah pada server")
		return c.JSON(http.StatusInternalServerError, view.InternalServerError())
	}
	log.Info("berhasil get all data")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "berhasil get all data",
		"status":  true,
		"data":    res,
	})
}

func (uc *UserController) GetUserbyid(c echo.Context) error {

	res, err := uc.Repo.GetUser()

	if err != nil {
		log.Warn("masalah pada server")
		return c.JSON(http.StatusInternalServerError, view.InternalServerError())
	}
	log.Info("berhasil get all data")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "berhasil get all data",
		"status":  true,
		"data":    res,
	})
}
func (uc *UserController) UpdateUser(c echo.Context) error {
	var tmpUser user.UpdateUserRequest

	if err := c.Bind(&tmpUser); err != nil {
		log.Warn("salah input")
		return c.JSON(http.StatusBadRequest, user.BadRequest())
	}

	if err := uc.Valid.Struct(tmpUser); err != nil {
		log.Warn(err.Error())
		return c.JSON(http.StatusBadRequest, user.BadRequest())
	}

	updateUser := mUser.User{Nama: tmpUser.Nama, Email: tmpUser.Email, Password: tmpUser.Password}
	res, err := uc.Repo.Update(updateUser)

	if err != nil {
		log.Warn("masalah pada server")
		return c.JSON(http.StatusInternalServerError, view.InternalServerError())
	}
	log.Info("berhasil Update")
	return c.JSON(http.StatusCreated, user.SuccessUpdate(res))
}

func (uc *UserController) DeleteUser(c echo.Context) error {
	res, err := uc.Repo.Delete()

	if err != nil {
		log.Warn("masalah pada server")
		return c.JSON(http.StatusInternalServerError, view.InternalServerError())
	}
	log.Info("berhasil Hapus data")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "berhasil Hapus data",
		"status":  true,
		"data":    res,
	})
}
