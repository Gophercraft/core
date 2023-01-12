package realm

// type TerrainMgr struct {
// 	Chunks *sync.Map
// }

// type MapChunkKey struct {
// 	MapID uint32
// 	Index terrain.TileChunkLookupIndex
// }

// type MapChunkValue struct {
// 	Loaded    time.Time
// 	ChunkData *terrain.ChunkData
// }

// func (ws *Server) LookupMapChunk(key MapChunkKey) (*terrain.ChunkData, error) {
// 	mapChunkValue, ok := ws.TerrainMgr.Chunks.Load(key)
// 	if !ok {
// 		// Load Map data from pack
// 		var cmap *dbdefs.Ent_Map
// 		ws.DB.Lookup(wdb.BucketKeyUint32ID, key.MapID, &cmap)
// 		if cmap == nil {
// 			return nil, fmt.Errorf("no map for %d", key.MapID)
// 		}

// 		mr, err := terrain.NewMapReader(ws.PackLoader, ws.Build(), cmap.Directory)
// 		if err != nil {
// 			return nil, err
// 		}

// 		cnkData, err := mr.GetChunkByIndex(&key.Index)
// 		if err != nil {
// 			return nil, err
// 		}

// 		ws.TerrainMgr.Chunks.Store(key, &MapChunkValue{
// 			Loaded:    time.Now(),
// 			ChunkData: cnkData,
// 		})
// 		return cnkData, nil
// 	}
// 	mcv := mapChunkValue.(*MapChunkValue)
// 	mcv.Loaded = time.Now()
// 	return mcv.ChunkData, nil
// }

// func (s *Session) UpdateArea() {
// 	// We need to see which chunk
// 	pos := s.Position()

// 	tcli := terrain.FindTileChunkIndex(pos.X, pos.Y)

// 	if s.CurrentChunkIndex != nil {
// 		// Player is in same chunk. Nothing to update
// 		if *s.CurrentChunkIndex == *tcli {
// 			return
// 		}
// 	}

// 	s.CurrentChunkIndex = tcli

// 	lookup := MapChunkKey{
// 		MapID: s.CurrentMap,
// 		Index: *tcli,
// 	}

// 	cnk, err := s.Server.LookupMapChunk(lookup)
// 	if err != nil {
// 		log.Warn(err)
// 		return
// 	}

// 	if cnk.AreaID != s.CurrentArea {
// 		s.CurrentArea = cnk.AreaID

// 		var area *dbdefs.Ent_AreaTable
// 		s.DB().Lookup(wdb.BucketKeyUint32ID, s.CurrentArea, &area)

// 		if area == nil {
// 			return
// 		}

// 		s.HandleZoneExperience(s.CurrentArea)

// 		for area.ParentAreaID != 0 {
// 			s.DB().Lookup(wdb.BucketKeyUint32ID, area.ParentAreaID, &area)
// 			if area == nil {
// 				break
// 			}

// 			s.HandleZoneExperience(uint32(area.ParentAreaID))
// 		}
// 	}

// 	s.CurrentArea = cnk.AreaID
// }

// const terrainMgrSweepInterval = 2 * time.Minute
// const unusedChunkLifetime = 2 * time.Minute

// func (ws *Server) InitTerrainMgr() {
// 	ws.TerrainMgr.Chunks = new(sync.Map)

// 	for {
// 		time.Sleep(terrainMgrSweepInterval)
// 		ws.TerrainMgr.Chunks.Range(func(k, v interface{}) bool {
// 			key := k.(MapChunkKey)
// 			value := v.(*MapChunkValue)
// 			if time.Since(value.Loaded) > unusedChunkLifetime {
// 				ws.TerrainMgr.Chunks.Delete(key)
// 			}
// 			return true
// 		})
// 	}
// }
