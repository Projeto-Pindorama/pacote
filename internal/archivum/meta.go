/*
* meta.go - Get UNIX metadata for files and return in a struct
*
 */

package archivum

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

type Metadata struct {
	FType      rune
	Path       string
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

	// minor/major numbers are only used in device files */
	var deviceInfo *DeviceInfo = nil
	if (ftype == 'c') || (ftype == 'b') {
		deviceInfo = fi.DeviceInfo()
	}

	if ftype == 's' {
		return nil, fmt.Errorf("%w <%s>", errSLink, path)
	}

	Data := &Metadata{
		FType:      ftype,
		Path:       path,
		OctalMod:   "octalPermissions",
		Owner:      fi.Owner().Name,
		Group:      fi.Group().Name,
		DeviceInfo: deviceInfo,
	}

	return Data, nil
}

func MetadataToString(m *Metadata) string {
	if m.DeviceInfo == nil {
		return fmt.Sprintf(
			"%c %s %s %s %s",
			m.FType,
			m.Path,
			m.OctalMod,
			m.Owner,
			m.Group,
		)
	} else {
		return fmt.Sprintf(
			"%c %s %d %d %s %s %s",
			m.FType,
			m.Path,
			m.DeviceInfo.Major,
			m.DeviceInfo.Minor,
			m.OctalMod,
			m.Owner,
			m.Group,
		)
	}

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
