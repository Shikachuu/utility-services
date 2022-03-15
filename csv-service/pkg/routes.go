package pkg

func (s *Server) routes() {
	s.router.Get("/health", s.newHealthHandler())
	s.router.Post("/csv-service", s.newCSVServiceHandler())
}
