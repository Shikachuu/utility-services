package pkg

func (s *Server) routes() {
	s.router.HandleFunc("/health", s.newHealthHandler()).Methods("GET")
}
