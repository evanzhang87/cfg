package file

import (
	"github.com/fsnotify/fsnotify"
)

func (f *File) handleEvent(ev fsnotify.Event) (watched bool) {
	var err error
	if _, ok := f.ignoreEvents[ev.Op]; ok {
		return true
	}

	if ev.Op&fsnotify.Write == fsnotify.Write {
		if err = f.readfile(); err != nil {
			f.logger.Warnf("Write: readfile of %s failed: %v", f.filename, err)
			return
		}
	} else if ev.Op&fsnotify.Rename == fsnotify.Rename {
		if err = f.readfile(); err != nil {
			f.logger.Warnf("Rename: readfile of %s failed: %v", f.filename, err)
			return
		}
	} else if ev.Op&fsnotify.Remove == fsnotify.Remove {
		return
	}
	return true
}
