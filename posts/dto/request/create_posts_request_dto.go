package request

type CreatePostsRequest struct {
	Title string `validate:"required,min=4,max=6" json:"title" errormsg:"Title is required and must be between 4 and 200 characters"`
}
