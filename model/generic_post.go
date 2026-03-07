package model

type GenericPost interface {
	// GetPost returns the base post
	GetPost() *Post
	// GetViewPost returns the view post with whether the data are converted from post
	GetViewPost() (*ViewPost, bool)
}

func (p *Post) GetPost() *Post {
	return p
}

func (p *Post) GetViewPost() (*ViewPost, bool) {
	return nil, true
}

func (vp *ViewPost) GetPost() *Post {
	return &vp.Post
}

func (vp *ViewPost) GetViewPost() (*ViewPost, bool) {
	return vp, false
}
