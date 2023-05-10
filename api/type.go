package api

type getPaginationRequest struct {
	Page  int32 `form:"page" binding:"required,min=1"`
	Limit int32 `form:"limit" binding:"required,min=5,max=100"`
}

type idURI struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}
