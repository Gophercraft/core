package client

// func (cl *Client) InitWarden() {
// 	gen := crypto.NewSessionKeyGenerator(sha1.New, cl.SessionKey[:40])
// 	outputKey := make([]byte, 16)
// 	inputKey := make([]byte, 16)

// 	gen.Read(inputKey)
// 	gen.Read(outputKey)

// 	cl.Warden = &warden.Warden{
// 		InputCrypto:  arc4.New(inputKey),
// 		OutputCrypto: arc4.New(outputKey),
// 		Module:       warden.Module_79C0768D657977D697E10BAD956CCED1,
// 	}
// }

// func (cl *Client) HandleWarden(data []byte) {
// 	if cl.Warden == nil {
// 		cl.InitWarden()
// 	}

// 	request := data[4:]
// 	cl.Warden.OutputCrypto.Decrypt(request)
// 	op := request[0]
// 	yo.Println("Got warden opt", op)
// 	switch op {
// 	case packet.WARDEN_SMSG_MODULE_USE:
// 		yo.Println("Warden initialized.")

// 		d, _ := packet.UnmarshalWardenModuleUse(request)
// 		if !bytes.Equal(d.ModuleID, warden.Module_79C0768D657977D697E10BAD956CCED1.Hash) {
// 			log.Fatal("Invalid module ID: ", hex.EncodeToString(d.ModuleID))
// 		}
// 		pkt := packet.NewEtcBuffer(nil)
// 		pkt.WriteByte(packet.WARDEN_CMSG_MODULE_OK)
// 		da := pkt.Encode()
// 		cl.Warden.InputCrypto.Encrypt(da)
// 		gp := packet.NewGamePacket(packet.CMSG_WARDEN_DATA)
// 		gp.Write(da)
// 		cl.Send(gp)
// 	case packet.WARDEN_SMSG_MODULE_INITIALIZE:
// 		d, _ := packet.UnmarshalWardenModuleInitRequest(request)
// 		yo.Println(spew.Sdump(d))
// 	case packet.WARDEN_SMSG_CHEAT_CHECKS_REQUEST:
// 		chks, err := warden.UnmarshalWardenServerChecks(request, cl.Warden.Module.ClientKeySeed[0])
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		log.Warn("Checks requested.")
// 		rspBuffer := packet.NewEtcBuffer(nil)
// 		lng := 0
// 		for _, v := range chks {
// 			log.Warn("CHECK TYPE", v.Type)
// 			switch v.Type {
// 			case packet.MEM_CHECK:
// 				lng++
// 				rspBuffer.WriteByte(0)
// 				res := warden.CheckResults[v.ID]
// 				if res == nil {
// 					log.Fatal("No result found for ", v.ID)
// 				}
// 				rspBuffer.WriteBytes(res)
// 			case packet.MODULE_CHECK:
// 				lng++
// 				rspBuffer.WriteByte(0xE9)
// 			case packet.MPQ_CHECK:
// 				lng++
// 				rspBuffer.WriteByte(0)
// 				rspBuffer.WriteBytes(warden.CheckResults[v.ID])
// 			default:
// 				yo.Println("UNHANDLED MODULE TYPE", v.Type)
// 			}
// 		}
// 		data := rspBuffer.Encode()
// 		pkt := packet.NewEtcBuffer(nil)
// 		pkt.WriteUint16(uint16(lng))
// 		pkt.WriteUint32(packet.BuildChecksum(data))
// 		pkt.WriteByte(0x01)
// 		pkt.WriteUint32(1136948141)
// 		pkt.WriteBytes(data)
// 		dd := pkt.Encode()
// 		cl.Warden.InputCrypto.Encrypt(dd)
// 		wp := packet.NewGamePacket(packet.CMSG_WARDEN_DATA)
// 		wp.Write(dd)
// 		cl.Send(wp)
// 	case packet.WARDEN_SMSG_HASH_REQUEST:
// 		pkt := packet.NewEtcBuffer(nil)
// 		pkt.WriteByte(packet.WARDEN_CMSG_HASH_RESULT)
// 		pkt.WriteBytes(warden.Module_79C0768D657977D697E10BAD956CCED1.ClientKeySeedHash)
// 		da := pkt.Encode()
// 		cl.Warden.InputCrypto.Encrypt(da)
// 		gp := packet.NewGamePacket(packet.CMSG_WARDEN_DATA)
// 		gp.Write(da)
// 		cl.Send(gp)
// 		cl.Warden.InputCrypto = arc4.New(cl.Warden.Module.ClientKeySeed)
// 		cl.Warden.OutputCrypto = arc4.New(cl.Warden.Module.ServerKeySeed)
// 	default:
// 		log.Fatal("Unknown op", op)
// 	}
// }
