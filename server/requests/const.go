package requests

const (
	url           = "https://api.github.com"
	fetchEndpoint = "/repos/%s/%s/pulls/%d"
	sendEndpoint  = "/repos/%s/%s/issues/%d/comments"
)
