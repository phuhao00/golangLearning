



```go
//广播请求
func (self *GMManager_) GMBroadcastHandle(clientPacket *ClientPacket, realdata []byte) {
   sess := clientPacket.Session
   req := GM_SystemBroadcast{}
   err := json.Unmarshal(realdata, &req)
   if err != nil {
      fmt.Println("Unmarshal GM_SystemBroadcast failed", err)
      sess.SendGMJsonMsg(pbID.Message_GM_SystemBroadcast, ErrorCode.UnpackPkgErr, &GM_ActorId{clientPacket.Packet.ErrCode, nil})
      return
   }
//如果
   if req.BeginTime!=0&& req.EndTime!=0&& req.CapTime!=0&&req.Content!=""{

      if ChatManager.stopAllTicking {
         sess.SendGMJsonMsg(pbID.Message_GM_SystemTickerBroadcastStop, 1, &GM_ActorId{clientPacket.Packet.ErrCode, nil})
         return
      }
      sess.SendGMJsonMsg(pbID.Message_GM_SystemBroadcast, ErrorCode.None, &GM_ActorId{clientPacket.Packet.ErrCode, nil})
      ChatManager.tickBroadcastGroup.Add(1)
      go func(){
         now := time.Now().Unix() //这里需要now的操作
         nextTime := req.BeginTime
         if now>req.BeginTime {
            nextTime=now
         }
         for  {
            now = time.Now().Unix()//这里now如果不重新赋值是不会变化的
            if now > req.EndTime || ChatManager.stopAllTicking{
               break
            }
            time.Sleep(time.Second)
            if now >= nextTime {
               ChatManager.systemBroadcast(req.Content, nil, ChatChannel_notice)
               nextTime += req.CapTime
            }
         }
         ChatManager.tickBroadcastGroup.Done()
      }()
      return
   } else {
      ChatManager.systemBroadcast(req.Content, nil, ChatChannel_notice)
      sess.SendGMJsonMsg(pbID.Message_GM_SystemBroadcast, ErrorCode.None, &GM_ActorId{clientPacket.Packet.ErrCode, nil})
   }
}

//停止广播
func (self *GMManager_) GMForceStopChatTickBroadcast(clientPacket *ClientPacket, realdata []byte) {
   sess := clientPacket.Session
   req := GM_SystemBroadcast{}
   err := json.Unmarshal(realdata, &req)
   if err != nil {
      fmt.Println("Unmarshal GM_SystemBroadcast failed", err)
      sess.SendGMJsonMsg(pbID.Message_GM_SystemTickerBroadcastStop, ErrorCode.UnpackPkgErr, &GM_ActorId{clientPacket.Packet.ErrCode, nil})
      return
   }

   if ChatManager.stopAllTicking {
      sess.SendGMJsonMsg(pbID.Message_GM_SystemTickerBroadcastStop, 1, &GM_ActorId{clientPacket.Packet.ErrCode, nil})
      return
   }

   go func() {
      ChatManager.stopAllTicking = true
      ChatManager.tickBroadcastGroup.Wait()
      ChatManager.stopAllTicking = false
      sess.SendGMJsonMsg(pbID.Message_GM_SystemTickerBroadcastStop, ErrorCode.None, &GM_ActorId{clientPacket.Packet.ErrCode, nil})
   }()
}
```