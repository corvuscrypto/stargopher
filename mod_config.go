package stargopher

//AddPluginHandlers sets an array of functions to a map key
//to the appropriate handler map according to the behavior wanted.
//immediately invoked functions can also be set here.
func AddPluginHandlers() {
	/*
	   E.g. this would mean add actions before each ChatSent packet type is
	   handled:

	     beforeHandlers[ChatSent] = append(beforeHandlers[ChatSent], myCustomFunction)
	*/
}
