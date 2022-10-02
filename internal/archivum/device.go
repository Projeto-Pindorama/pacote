package archivum

type DeviceInfo struct {
	Major, Minor uint32
}

func deviceInfoByDev(dev uint64) *DeviceInfo {
	major := uint32((dev & 0x00000000000fff00) >> 8)
	major |= uint32((dev & 0xfffff00000000000) >> 32)
	minor := uint32((dev & 0x00000000000000ff) >> 0)
	minor |= uint32((dev & 0x00000ffffff00000) >> 12)

	return &DeviceInfo{
		Major: major,
		Minor: minor,
	}

}
