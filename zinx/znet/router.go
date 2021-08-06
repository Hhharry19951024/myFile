package znet

import "src/zinterface"

// 实现router时，先嵌入BaseRouter基类，然后根据需要重写
type BaseRouter struct{}

/* 这里之所以BaeRouter都为空，
是因为有的Router不需要PreHandel，PostH这两个业务
而Router全部继承后，可以省略
*/

// 在处理conn业务之前的钩子方法Hook
func (br *BaseRouter) PreHandle(req zinterface.IRequest) {}

// 在处理conn业务的主方法Hook
func (br *BaseRouter) Handle(req zinterface.IRequest) {}

// 在处理conn业务之后的钩子方法Hook
func (br *BaseRouter) PostHandle(req zinterface.IRequest) {}
