package routers

func (s *ServiceRouter) Wait() {
	s.wg.Wait()
}

func getFlushTimeout(timeoutMs int) int {
	if timeoutMs <= 0 {
		return 100
	}
	return timeoutMs
}