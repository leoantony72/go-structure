package middleware

import "test/internal/ports"

type TestMiddleware struct {
	tweetService ports.TestService
}

func NewTweetMiddleware(s ports.TestService) *TestMiddleware {
	return &TestMiddleware{
		tweetService: s,
	}
}
