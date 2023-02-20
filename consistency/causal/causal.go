package causal

/*
	implementations of causal consistency
*/

type CausalEntity struct {
	// address for internal communication between nodes
	internalAddress string
	//
}

// this method must be called by KVServer instead of CausalEntity
// this method is used to execute the command from client
// func (ce *CausalEntity) Start(command interface{}, vcFromClientArg map[string]int32, vcFromKVServerArg map[string]int32) (map[string]int32, bool) {
// 	vcFromClient := util.BecomeSyncMap(vcFromClientArg)
// 	vcFromKVServer := util.BecomeSyncMap(vcFromKVServerArg)
// 	newLog := command.(config.Log)
// 	util.DPrintf("Log in Start(): %v ", newLog)
// 	util.DPrintf("vcFromClient in Start(): %v", vcFromClient)
// 	if newLog.Option == "Put" {
// 		isUpper := util.IsUpper(vcFromClient, vcFromKVServer)
// 		if isUpper {
// 			val, _ := vcFromKVServer.Load(ce.internalAddress)
// 			vcFromKVServer.Store(ce.internalAddress, val.(int32)+1)
// 			data, _ := json.Marshal(newLog)
// 			args := &causalrpc.AppendEntriesInCausalRequest{
// 				Log:         data,
// 				Vectorclock: util.BecomeMap(vcFromKVServer),
// 			}
// 			for i := 0; i < len(ce.members); i++ {
// 				if ce.members[i] != ce.address {
// 					go ce.sendAppendEntriesInCausal(ce.members[i], args)
// 				}
// 			}
// 			ce.log = append(ce.log, newLog)
// 			ce.persister.Put(newLog.Command.Key, newLog.Command.Value)
// 			ce.applyCh <- 1
// 			util.DPrintf("applyCh unread buffer: %v", len(ce.applyCh))
// 			return util.BecomeMap(ce.vectorClock), true
// 		} else {
// 			return nil, false
// 		}
// 	} else if newLog.Command.Option == "Get" {
// 		if util.Len(vcFromClient) == 0 {
// 			return nil, true
// 		} else {
// 			if ce.IsUpper(vcFromClient) {
// 				return nil, true
// 			}
// 			return nil, false
// 		}
// 	}
// 	util.DPrintf("here is Start() in Causal: log command option is false")
// 	return nil, false
// }

// func MakeCausalEntity(address string) *CausalEntity {
// 	ce := &CausalEntity{
// 		address: address,
// 	}
// 	return ce
// }
