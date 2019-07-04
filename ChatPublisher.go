package WorldMapServer

type publisher_interface interface {
	BroadCast_Mode(int64,interface{})
	Unicast_Mode(int64,interface{})
	Trigger( interface{})
}

const (
	broadcast =iota+1
	unicast
)
const (
	publisher_EventType_Normal =iota+1

)


type publisher_event interface {
	GetInstance() interface{}
}

type publisher_event_normal struct {
	publisher_event
	detailType int64
	chatType int64
	broadcastType int64
	eventType uint32
	groupId int64
	user *User
}

func (self  *publisher_event_normal)GetInstance() interface{} {
	return DeepCopy(self)
}

type publisher struct {
	userIds []int64
}
//多播
func ( pub *publisher)BroadCast_Mode(eventType int64,event interface{})  {
	switch eventType {
	case publisher_EventType_Normal:
		tmp:=event.(publisher_event_normal)
		ChatManager.SystemGroupBroadcast(tmp.detailType,"",nil,tmp.chatType,tmp.groupId)
	}
}
//单播
func (pub *publisher)Unicast_Mode( eventType int64,event interface{} ) {
	switch eventType {
	case publisher_EventType_Normal:
		tmp:=event.(publisher_event_normal)
		ChatManager.SystemBroadcast(tmp.detailType,"",nil,tmp.chatType)
	}
}
//
func (pub *publisher)Trigger(event interface{}){
	switch event.(type) {
	case publisher_event_normal:
		tmp:=event.(publisher_event_normal)
		if tmp.broadcastType== broadcast {
			pub.BroadCast_Mode(publisher_EventType_Normal,tmp)
		}
		if tmp.broadcastType==unicast {
			pub.Unicast_Mode(publisher_EventType_Normal,tmp)
		}
	}
}
//
func DeepCopy(value interface{}) interface{} {
	if valueMap, ok := value.(map[string]interface{}); ok {
		newMap := make(map[string]interface{})
		for k, v := range valueMap {
			newMap[k] = DeepCopy(v)
		}
		return newMap
	} else if valueSlice, ok := value.([]interface{}); ok {
		newSlice := make([]interface{}, len(valueSlice))
		for k, v := range valueSlice {
			newSlice[k] = DeepCopy(v)
		}
		return newSlice
	}
	return value
}

//type singleton struct{}
//var ins *singleton
//var once sync.Once
//
//func GetIns() *singleton {
//     once.Do(func(){
//         ins = &singleton{}
//     })
//     return ins
// }
//
func (self *ChatManager_)InitSth()  {
	for key, value := range self.ChatBaseManager.MessageContentData {
		tmp:=&publisher_event_normal{
			detailType:value.Message_id,
		}
		if len(value.Messate_type)>1 {
			tmp.broadcastType=broadcast
		}else {
			tmp.broadcastType=unicast
		}
		self.detailChatType2Event[key]=tmp
	}
}
//detailChatType2Event  map[int64]publisher_event
//detailChatType2Event:make(map[int64]publisher_event),


