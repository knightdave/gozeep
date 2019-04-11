package gozeep

// Plugin implemetns base plugin
type Plugin struct {
}

// Override to update the envelope or http headers when receiving a
// message.
func (p Plugin) ingress(envelope string, httpHeaders map[string]string, operation string) (string, map[string]string) {
	return envelope, httpHeaders
}

func (p Plugin) egress(envelope string, httpHeaders map[string]string, operation string, bindingOptions string) (string, map[string]string) {
	return envelope, httpHeaders
}

// HistoryPlugin implements base history plugin
type HistoryPlugin struct {
	buffer []map[string]interface{}
}

// GetLastSent gets last sent message
func (h HistoryPlugin) GetLastSent() map[string]interface{} {
	lastTx := h.buffer[len(h.buffer)-1]
	if lastTx != nil && lastTx["sent"] != nil {
		return lastTx["sent"].(map[string]interface{})
	}
	return nil
}

// GetLastReceived gets last received message
func (h HistoryPlugin) GetLastReceived() map[string]interface{} {
	lastTx := h.buffer[len(h.buffer)-1]
	if lastTx != nil && lastTx["received"] != nil {
		return lastTx["received"].(map[string]interface{})
	}
	return nil

}

// Ingress method
// TODO: implement Ingress method

// Egress method
// TODO: implement Egress method
