package znet

// IServer的接口实现，定义服务器模块
type Server struct {
	Name       string // 服务器名称
	IPVersrion string // 服务器IP版本
	IP         string // 服务器监听IP
	Port       int    // 服务器监听端口
}

func NewServer(name string) zinterface.IServer {
	s := &Server{
		Name:       name,
		IPVersrion: "tcp4",
		IP:         "0.0.0.0",
		Port:       8001,
	}
	return s
}

func (s *Server) Start() {

}

func (s *Server) Stop() {

}

func (s *Server) Run() {

}
