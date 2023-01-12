package realm

const (
	ProcKilled = 1 << iota
)

type Process struct {
	Name          string
	CreationState SessionState
	ProcState     uint8
	Data          interface{}
	Cancel        chan<- bool
}

func (s *Session) FindProcess(name string) *Process {
	for _, proc := range s.Processes {
		if proc.Name == name {
			return proc
		}
	}
	return nil
}

// Creates a background function that will run until canceled, or until the session state changes to a lower value
func (s *Session) CreateProcess(name string, fn func(*Session, <-chan bool)) *Process {
	s.GuardSession.Lock()
	ch := make(chan bool)
	proc := &Process{
		Name:          name,
		Cancel:        ch,
		CreationState: s.state,
	}
	go func() {
		fn(s, ch)
		// If process returns without being killed,
		if proc.ProcState&ProcKilled == 0 {
			s.GuardSession.Lock()
			for i, allProc := range s.Processes {
				if proc == allProc {
					s.killProcess(i)
					break
				}
			}
			s.GuardSession.Unlock()
		}
	}()
	s.Processes = append(s.Processes, proc)
	s.GuardSession.Unlock()
	return proc
}

func (s *Session) killProcess(i int) {
	proc := s.Processes[i]
	proc.ProcState |= ProcKilled
	proc.Cancel <- true
	close(proc.Cancel)
	s.Processes = append(s.Processes[:i], s.Processes[i+1:]...)
}

func (s *Session) KillAllProcesses() {
	s.GuardSession.Lock()
	for i := range s.Processes {
		s.killProcess(i)
	}
	s.Processes = nil
	s.GuardSession.Unlock()
}

func (s *Session) KillProcessesWithTag(tag string) {
	s.GuardSession.Lock()

	for {
		var proc int = -1

		for i := range s.Processes {
			if s.Processes[i].Name == tag {
				proc = i
			}
		}

		if proc == -1 {
			break
		}

		s.killProcess(proc)
	}

	s.GuardSession.Unlock()
}
