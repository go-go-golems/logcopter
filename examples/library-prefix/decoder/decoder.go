package decoder

func Decode(packet []byte) int {
	log.Trace().Int("bytes", len(packet)).Msg("decoding BLE packet")
	return len(packet)
}
