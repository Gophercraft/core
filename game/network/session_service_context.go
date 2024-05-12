package network

// TODO: consider using small array as fast lookup table for this.

// Associate a context object for a particular service ID
func (session *Session) SetServiceContext(service_id ServiceID, context any) error {
	// enforce mutex
	session.guard_service_contexts.Lock()
	defer session.guard_service_contexts.Unlock()

	// Associate context
	session.service_contexts[service_id] = context

	return nil
}

func (session *Session) ServiceContext(service_id ServiceID) any {
	session.guard_service_contexts.Lock()
	defer session.guard_service_contexts.Unlock()

	return session.service_contexts[service_id]
}
