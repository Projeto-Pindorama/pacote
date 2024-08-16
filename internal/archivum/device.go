/*
 * device.go - Handle device information fetching on UNIX
 *
 * Copyright (C) 2022: Pindorama
 *		Samuel Brederodes (callsamu)
 *
 * SPDX-Licence-Identifier: NCSA
 *
 * Some of the alchemical numbers below were copied
 * from the Linux®-specific section of
 * golang.org/x/archive/tar/stat_unix.go. As per its
 * copyright header:
 * Copyright 2012 The Go Authors. All rights reserved.
 *
 * SPDX-Licence-Identifier: BSD-3-Clause
 *
 */

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
