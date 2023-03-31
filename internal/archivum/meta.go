/*
* meta.go - Get UNIX metadata for files and return in a struct
*
 */

package archivum

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
)

type Metadata struct {
	FType      rune
	Path       string
	RealPath   string
	OctalMod   string 
	Owner      string
	Group      string
	DeviceInfo *DeviceInfo
}

var (
	errSLink = errors.New("symbolic links are unsupported")
)

func Scan(dir OperatingSystemFS, path string) (*Metadata, error) {
	fi, err := dir.Stat(path)
	if err != nil {
		return nil, err
	}

	ftype := determineFType(fi)

	// minor/major numbers are only used in device files
	var deviceInfo *DeviceInfo = nil
	var realPath  string = ""
	if (ftype == 'c') || (ftype == 'b') {
		deviceInfo = fi.DeviceInfo()
	}

	if ftype == 's' {
		realPath, _ = filepath.Abs(path)
		Data := &Metadata{
			FType:      ftype,
			Path:       path,
			RealPath:   realPath,
		}
		return Data, nil
	}

	Data := &Metadata{
		FType:      ftype,
		Path:       path,
		RealPath:   realPath,
		OctalMod:   "permissions",
		Owner:      fi.Owner().Name,
		Group:      fi.Group().Name,
		DeviceInfo: deviceInfo,
	}

	return Data, nil
}

func determineFType(fi os.FileInfo) rune {
	switch mode := fi.Mode(); {
	case mode.IsRegular():
		return 'f'
	case mode.IsDir():
		return 'd'
	case mode&fs.ModeSymlink != 0:
		return 's'
	case mode&fs.ModeDevice != 0:
		return 'b'
	case mode&fs.ModeCharDevice != 0:
		return 'c'
	case mode&fs.ModeNamedPipe != 0:
		return 'p'
	default:
		return '?'
	}
}
