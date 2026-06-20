package model

type Response struct {
	Message string      `json:"message" example:"detail pesan"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty" example:"detail error"`
}

type SuccessResponse struct {
	Message string `json:"message" example:"request berhasil diproses"`
	Data    string `json:"data,omitempty" example:"contoh data response 200"`
}

type CreatedResponse struct {
	Message string `json:"message" example:"data berhasil dibuat"`
	Data    string `json:"data,omitempty" example:"contoh data response 201"`
}

type UnauthorizedResponse struct {
	Message string `json:"message" example:"token tidak valid atau belum dikirim"`
	Error   string `json:"error,omitempty" example:"unauthorized"`
}

type ForbiddenResponse struct {
	Message string `json:"message" example:"user tidak memiliki akses untuk fitur ini"`
	Error   string `json:"error,omitempty" example:"forbidden"`
}
