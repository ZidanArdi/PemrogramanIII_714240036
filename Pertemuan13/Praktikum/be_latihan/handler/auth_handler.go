package handler

import (
	"be_latihan/config/middleware"
	"be_latihan/model"
	"be_latihan/pkg/password"
	"be_latihan/repository"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Register godoc
// @Summary Register user baru
// @Description Membuat akun user baru. Role dapat diisi admin atau user. Jika role kosong, backend akan memakai default admin.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body model.AuthRequest true "Payload register user"
// @Success 201 {object} model.CreatedResponse
// @Failure 400 {object} model.Response
// @Failure 409 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /register [post]
func Register(c *fiber.Ctx) error {
	var payload model.AuthRequest
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Message: "payload tidak valid",
			Error:   err.Error(),
		})
	}

	payload.Username = strings.TrimSpace(payload.Username)
	payload.Role = strings.TrimSpace(payload.Role)
	if payload.Role == "" {
		payload.Role = "admin"
	}

	if payload.Username == "" || payload.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Message: "username dan password wajib diisi",
		})
	}

	hashedPassword, err := password.HashPassword(payload.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "gagal membuat hash password",
			Error:   err.Error(),
		})
	}

	user := model.User{
		Username: payload.Username,
		Password: hashedPassword,
		Role:     payload.Role,
	}

	data, err := repository.InsertUser(&user)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(model.Response{
			Message: "username sudah digunakan atau data tidak valid",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(model.Response{
		Message: "register berhasil",
		Data: model.AuthUserResponse{
			ID:       data.ID,
			Username: data.Username,
			Role:     data.Role,
		},
	})
}

// Login godoc
// @Summary Login user
// @Description Melakukan login dan mengembalikan JWT jika username dan password valid.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body model.AuthRequest true "Payload login user"
// @Success 200 {object} model.SuccessResponse
// @Failure 400 {object} model.Response
// @Failure 401 {object} model.UnauthorizedResponse
// @Failure 500 {object} model.Response
// @Router /login [post]
func Login(c *fiber.Ctx) error {
	var payload model.AuthRequest
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Message: "payload tidak valid",
			Error:   err.Error(),
		})
	}

	user, err := repository.FindUserByUsername(strings.TrimSpace(payload.Username))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusUnauthorized).JSON(model.Response{
				Message: "username atau password salah",
				Error:   "unauthorized",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "gagal mencari user",
			Error:   err.Error(),
		})
	}

	if !password.CheckPasswordHash(payload.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Message: "username atau password salah",
			Error:   "unauthorized",
		})
	}

	token, err := middleware.GenerateJWT(user, 2*time.Hour)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "gagal membuat token",
			Error:   err.Error(),
		})
	}

	return c.JSON(model.Response{
		Message: "login berhasil",
		Data: model.LoginResponse{
			Token: token,
			User: model.AuthUserResponse{
				ID:       user.ID,
				Username: user.Username,
				Role:     user.Role,
			},
		},
	})
}

// ChangePassword godoc
// @Summary Change password user
// @Description Mengubah password user yang sedang login. Endpoint ini membutuhkan token JWT dan password lama yang sesuai.
// @Tags Auth
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body model.ChangePasswordRequest true "Payload change password"
// @Success 200 {object} model.SuccessResponse
// @Failure 400 {object} model.Response
// @Failure 401 {object} model.UnauthorizedResponse
// @Failure 500 {object} model.Response
// @Router /change-password [put]
func ChangePassword(c *fiber.Ctx) error {
	var payload model.ChangePasswordRequest
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Message: "payload tidak valid",
			Error:   err.Error(),
		})
	}

	payload.OldPassword = strings.TrimSpace(payload.OldPassword)
	payload.NewPassword = strings.TrimSpace(payload.NewPassword)
	if payload.OldPassword == "" || payload.NewPassword == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Message: "password lama dan password baru wajib diisi",
		})
	}

	username, ok := c.Locals("username").(string)
	if !ok || username == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Message: "token tidak valid atau belum dikirim",
			Error:   "unauthorized",
		})
	}

	user, err := repository.FindUserByUsername(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusUnauthorized).JSON(model.Response{
				Message: "user dari token tidak ditemukan",
				Error:   "unauthorized",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "gagal mencari user",
			Error:   err.Error(),
		})
	}

	if !password.CheckPasswordHash(payload.OldPassword, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Message: "password lama tidak sesuai",
			Error:   "unauthorized",
		})
	}

	hashedPassword, err := password.HashPassword(payload.NewPassword)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "gagal membuat hash password",
			Error:   err.Error(),
		})
	}

	if err := repository.UpdateUserPassword(user.Username, hashedPassword); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "gagal mengubah password",
			Error:   err.Error(),
		})
	}

	return c.JSON(model.Response{
		Message: "password berhasil diubah",
	})
}
