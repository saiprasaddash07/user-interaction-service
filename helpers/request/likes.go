package request

type Interaction struct {
	UserId    int64 `json:"userId" form:"userId" binding:"required"`
	ContentId int64 `json:"contentId" form:"contentId" binding:"required"`
}

// type ReadInteraction struct {
// 	ReadId    int64 `json:"readId,omitempty" form:"readId"`
// 	UserId    int64 `json:"userId" form:"userId" binding:"required"`
// 	ContentId int64 `json:"contentId" form:"contentId" binding:"required"`
// }
