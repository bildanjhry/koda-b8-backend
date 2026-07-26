package libs

type QueryParams struct {
	SEARCH_PAR map[string]string `form:"search" binding:"omitempty"`
	ORDER_BY   string            `form:"order" binding:"omitempty,oneof=id username fullname created_at birth_date"`
	ORDER_TYPE string            `form:"sort" binding:"omitempty,oneof=asc desc"`
	LIMIT      int               `form:"limit"`
	OFFSET     int               `form:"page"`
}
