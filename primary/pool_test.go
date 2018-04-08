package primary

import (
	"sync"
	"io"
	"errors"
)

//一个安全的资源池，被管理的资源必须都实现io.Close接口

type Pool struct {
	m sync.Mutex
	res chan io.Closer
	factory func() (io.Closer,error)
	closed bool
}


//创建一个资源池
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size的值太小了。")
	}
	return &Pool{
		factory: fn,
		res:     make(chan io.Closer, size),
	}, nil
}