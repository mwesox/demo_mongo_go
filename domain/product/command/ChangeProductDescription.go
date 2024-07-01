package command

type ChangeProductDescription struct {
	Description string `json:"description" binding:"required"`
}
