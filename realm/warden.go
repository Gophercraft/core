package realm

// func (s *Session) InitWardenSession() {

// }

// func (s *Session) RequestHash() {
// 	p := etc.NewBuffer()
// 	p.WriteByte(packet.WARDEN_SMSG_HASH_REQUEST)
// 	p.Write(s.Warden.Module.Seed)
// 	d := p.Bytes()
// 	s.Warden.OutputCrypto.Encrypt(d)
// 	wp := packet.NewWorldPacket(packet.SMSG_WARDEN_DATA)
// 	wp.Write(d)
// 	s.SendPacket(wp)
// 	log.Println("Hash requested.")
// }

// func (s *Session) HandleHashResult(data *etc.Buffer) {
// 	data.SeekR(0)
// 	data.ReadByte()
// 	h := data.ReadBytes(20)
// 	if !bytes.Equal(h, s.Warden.Module.ClientKeySeedHash) {
// 		log.Println("Client failed Warden hash reply.")
// 		return
// 	}

// 	log.Println("Correct Warden module verified")

// 	s.Warden.InputCrypto = arc4.New(s.Warden.Module.ClientKeySeed)
// 	s.Warden.OutputCrypto = arc4.New(s.Warden.Module.ServerKeySeed)
// }

// func (s *Session) FailCheck() {

// }

// func (s *Session) HandleCheck(data *etc.Buffer) {
// 	data.SeekR(1)
// 	length := data.ReadUint16()
// 	checksum := data.ReadUint32()
// 	timingResult := data.ReadByte()
// 	if timingResult == 0x00 {
// 		log.Warnf("Failed timing check\n")
// 		s.FailCheck()
// 		return
// 	}
// 	newClientTicks := data.ReadUint32()
// 	log.Warnf("Packet length: %d Length: %d Checksum: 0x%x timingResult: %d newClientTicks: %d\n\n", data.Len(), length, checksum, timingResult, newClientTicks)
// 	for _, chk := range warden.GetChecks() {
// 		cid := chk.ID
// 		log.Println(chk.Type)
// 		switch chk.Type {
// 		case packet.MEM_CHECK:
// 			log.Println("Checking memory, ", cid)
// 			memresult := data.ReadByte()
// 			if memresult != 0 {
// 				log.Println("Client failed MEM_CHECK, not zero")
// 				s.FailCheck()
// 				return
// 			}
// 			res := data.ReadBytes(int(chk.Length))
// 			warden.CheckResults[cid] = res
// 			log.Println("MEM_CHECK", cid, "result", res)

// 		case packet.PAGE_CHECK_A, packet.PAGE_CHECK_B, packet.DRIVER_CHECK, packet.MODULE_CHECK:
// 			if data.ReadByte() != 0xE9 {
// 				log.Println("Client failed MODULE_CHECK")
// 				s.FailCheck()
// 				return
// 			}
// 		case packet.LUA_STR_CHECK:
// 			luaResult := data.ReadByte()
// 			if luaResult != 0 {
// 				log.Println("Client failed LUA_STR_CHECK, not zero")
// 				s.FailCheck()
// 				return
// 			}

// 			luaStrlen := data.ReadByte()
// 			luaStr := string(data.ReadBytes(int(luaStrlen)))
// 			log.Println("Lua string: ", luaStr)
// 			warden.CheckResults[cid] = []byte(luaStr)
// 		case packet.MPQ_CHECK:
// 			mpqResult := data.ReadByte()
// 			if mpqResult != 0 {
// 				log.Println("Client failed MPQ_CHECK, not zero")
// 				s.FailCheck()
// 				return
// 			}
// 			mpqHash := data.ReadBytes(20)
// 			log.Println("MPQ_CHECK", cid, "result", mpqHash)
// 			warden.CheckResults[cid] = mpqHash
// 		default:
// 			log.Println("Unhandled opcode", chk.Type)
// 		}

// 		log.Println("Check", cid, "passed.")
// 	}

// 	log.Println("Client passed all checks!")

// 	// log.Println("Client check data: ", data.Len(), spew.Sdump(data.ReadBytes(128)))
// 	printMap(warden.CheckResults)
// }

// func hexBytes(input []byte) string {
// 	var i []string
// 	for _, v := range input {
// 		i = append(i, fmt.Sprintf("0x%X", v))
// 	}
// 	return strings.Join(i, ", ")
// }
// func printMap(input map[uint16][]byte) {
// 	var i []int
// 	for k := range input {
// 		i = append(i, int(k))
// 	}

// 	sort.Ints(i)
// 	srcBuf := ""
// 	for _, v := range i {
// 		bs := input[uint16(v)]
// 		srcBuf += fmt.Sprintf("\t%d: []byte{ %s }, // %d\n", v, hexBytes(bs), len(bs))
// 	}
// 	log.Println(srcBuf)
// }

// func (s *Session) WardenResponse(data []byte) {
// 	if s.Warden == nil {
// 		log.Println("Received Warden data before Warden initialized")
// 		return
// 	}

// 	dat := data[6:]

// 	s.Warden.InputCrypto.Decrypt(dat)
// 	p := etc.FromBytes(dat)
// 	if p.Len() == 0 {
// 		log.Println("User sent too short packet")
// 	}
// 	log.Println("Leng??", p.Len())
// 	opcode := p.ReadByte()

// 	switch opcode {
// 	case packet.WARDEN_CMSG_CHEAT_CHECKS_RESULT:
// 		log.Println("Got cheat check.")
// 		s.HandleCheck(p)
// 	case packet.WARDEN_CMSG_MEM_CHECKS_RESULT:
// 		log.Println("Got mem checks")
// 	case packet.WARDEN_CMSG_MODULE_MISSING:
// 		log.Println("Client says Warden module is missing.")
// 		s.TransferWardenModule()
// 	case packet.WARDEN_CMSG_MODULE_OK:
// 		log.Println("Client says module is OK!")
// 		s.RequestHash()
// 	case packet.WARDEN_CMSG_HASH_RESULT:
// 		log.Println("Client sends hash result")
// 		s.HandleHashResult(p)
// 		s.InitializeModule()
// 	case packet.WARDEN_CMSG_MODULE_FAILED:
// 		log.Println("Fatal error: warden failed on target.")

// 	default:
// 		log.Println("Unrecognized packet:", opcode)
// 	}

// }

// func (s *Session) InitializeModule() {
// 	r := new(packet.WardenModuleInitRequest)
// 	r.Command1 = packet.WARDEN_SMSG_MODULE_INITIALIZE
// 	r.Size1 = 20
// 	r.Unk1 = 1
// 	r.Type = 1
// 	r.StringLibrary1 = 0
// 	r.Function1 = make([]uint32, 4)
// 	r.Function1[0] = 0x00024F80 // 0x00400000 + 0x00024F80 SFileOpenFile
// 	r.Function1[1] = 0x000218C0 // 0x00400000 + 0x000218C0 SFileGetFileSize
// 	r.Function1[2] = 0x00022530 // 0x00400000 + 0x00022530 SFileReadFile
// 	r.Function1[3] = 0x00022910 // 0x00400000 + 0x00022910 SFileCloseFile
// 	r.Checksum1 = packet.BuildChecksum(append([]byte{r.Unk1}, make([]byte, 19)...))

// 	r.Command2 = packet.WARDEN_SMSG_MODULE_INITIALIZE
// 	r.Size2 = 8
// 	r.Unk3 = 4
// 	r.Unk4 = 0
// 	r.StringLibrary2 = 0
// 	r.Function2 = 0x00419D40 // 0x00400000 + 0x00419D40 FrameScript::GetText
// 	r.Function2Set = 1
// 	r.Checksum2 = packet.BuildChecksum(append([]byte{r.Unk2}, make([]byte, 7)...))

// 	r.Command3 = packet.WARDEN_SMSG_MODULE_INITIALIZE
// 	r.Size3 = 8
// 	r.Unk5 = 1
// 	r.Unk6 = 1
// 	r.StringLibrary3 = 0
// 	r.Function3 = 0x0046AE20 // 0x00400000 + 0x0046AE20 PerformanceCounter
// 	r.Function3Set = 1
// 	r.Checksum3 = packet.BuildChecksum(append([]byte{r.Unk5}, make([]byte, 7)...))

// 	p := r.Encode()
// 	s.Warden.OutputCrypto.Encrypt(p)

// 	pkt := packet.NewWorldPacket(packet.SMSG_WARDEN_DATA)
// 	pkt.Write(p)
// 	s.SendPacket(pkt)
// 	log.Println("Module initialized.")
// 	s.RequestData()
// }

// func (s *Session) UpdateWarden() {
// 	if s.Warden == nil {
// 		return
// 	}

// 	s.Warden.CurrentTimestamp = timeMS()
// 	diff := s.Warden.CurrentTimestamp - s.Warden.PreviousTimestamp
// 	s.Warden.PreviousTimestamp = s.Warden.CurrentTimestamp

// 	log.Println("Time since last update: ", diff)
// 	s.RequestData()
// }

// // Naive warden server
// func (s *Session) RequestData() {
// 	pkt := etc.FromBytes(nil)
// 	pkt.WriteByte(packet.WARDEN_SMSG_CHEAT_CHECKS_REQUEST)

// 	xorByte := s.Warden.Module.ClientKeySeed[0]
// 	gc := warden.GetChecks()

// 	for _, chk := range gc {
// 		if chk.Type == packet.DRIVER_CHECK {
// 			pkt.WriteByte(uint8(len(chk.Str)))
// 			log.Println("Appending driver name", chk.Str)
// 			pkt.Write(append([]byte(chk.Str), 0))
// 		}
// 	}
// 	pkt.WriteByte(0x00)
// 	pkt.WriteByte(packet.TIMING_CHECK ^ xorByte)
// 	var index uint8 = 1
// 	log.Println("Got checks, ", len(gc))

// 	ix := 0
// 	for _, chk := range gc {
// 		pkt.WriteByte(chk.Type ^ xorByte)

// 		ix++
// 		dat, err := hex.DecodeString(chk.Data)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		switch chk.Type {
// 		case packet.MEM_CHECK:
// 			pkt.WriteByte(0x00)
// 			pkt.WriteUint32(uint32(chk.Address))
// 			pkt.WriteByte(chk.Length)
// 		case packet.PAGE_CHECK_A:
// 		case packet.PAGE_CHECK_B:
// 			pkt.Write(dat)
// 			pkt.WriteUint32(uint32(chk.Address))
// 			pkt.WriteByte(chk.Length)
// 		case packet.MPQ_CHECK:
// 		case packet.LUA_STR_CHECK:
// 			pkt.WriteByte(index)
// 			index++
// 		case packet.DRIVER_CHECK:
// 			pkt.Write(dat)
// 			pkt.WriteByte(index)
// 			index++
// 		case packet.MODULE_CHECK:
// 			seed := make([]byte, 4)
// 			rand.Read(seed)
// 			pkt.Write(seed)
// 			h := hmac.New(sha1.New, seed)
// 			h.Write([]byte(chk.Str))
// 			hm := h.Sum(nil)
// 			pkt.Write(hm)
// 		}
// 	}

// 	pkt.WriteByte(xorByte)
// 	d := pkt.Bytes()
// 	log.Println("Encoded check packet", len(d))
// 	s.Warden.OutputCrypto.Encrypt(d)
// 	wp := packet.NewWorldPacket(packet.SMSG_WARDEN_DATA)
// 	wp.Write(d)
// 	s.SendPacket(wp)
// }

// func timeMS() int64 {
// 	return time.Now().UnixNano() / int64(time.Millisecond)
// }
