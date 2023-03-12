package service

import "sync"

type service struct {
	UserServices
	VideService
	CommentService
}

var ServiceApp *service
var syn sync.Once

func init() {
	syn.Do(func() {
		ServiceApp = new(service)
	})
}
