/*
 * helpers.go - General helper functions for file characteristics
 *
 * Copyright (C) 2022-2024: Pindorama
 *			Luiz Antônio Rangel (takusuman)
 *			Samuel Brederodes (callsamu)
 *			João Pedro Vieira (JoaoP-Vieira) 
 *
 * SPDX-Licence-Identifier: NCSA
 *
 */

package archivum

import (
	"errors"
	"io/fs"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

func groupByGID(gid uint32) (*user.Group, error) {
	return user.LookupGroupId(strconv.FormatUint(uint64(gid), 10))
}

func userByUID(uid uint32) (*user.User, error) {
	return user.LookupId(strconv.FormatUint(uint64(uid), 10))
}

func permsOctal(fm fs.FileMode) string {
	return "0" + strconv.FormatUint(uint64(fm), 8)
}

func IsModeHardlink(fi os.FileInfo) (uint64, error) {
	var nlinks uint64
	var ok bool
	var err error
	var fis *syscall.Stat_t	

	fis, ok = fi.Sys().(*syscall.Stat_t)
	if !ok {
		err = errors.New("could not convert stat value to syscall.Stat_t")
		return 0, err
	}

	nlinks = uint64(fis.Nlink);
	switch nlinks {
	/* Normally, s.Nlink being 1
	 * indicates that it is a
	 * regular file, so we will
	 * return zero, which is false. */
	case 1: 
		return 0, nil
	default:
		return nlinks, nil
	}
}
