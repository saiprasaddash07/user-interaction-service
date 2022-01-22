package request

type Interaction struct {
	LikeId    int64 `json:"likeId,omitempty" form:"likeId"`
	UserId    int64 `json:"userId" form:"userId" binding:"required"`
	ContentId int64 `json:"contentId" form:"contentId" binding:"required"`
	Status    int   `json:"status,omitempty" form:"status"`
}
