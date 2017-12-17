C/S模式
定义client和server共同遵守的协议于package中，

1.server端注册处理对象，对象的方法作为服务可以被远程调用
	方法必须满足如下严格格式：
	func (t *T)MethodName (argType T1,replyType *T2) error
2.server通过HTTP监听client端调用
