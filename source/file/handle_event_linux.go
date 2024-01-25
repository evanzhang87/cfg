package file

import (
	"github.com/fsnotify/fsnotify"
)

func (f *File) handleEvent(ev fsnotify.Event) (watched bool) {
	var err error
	if _, ok := f.ignoreEvents[ev.Op]; ok {
		return true
	}

	if ev.Op&fsnotify.Write == fsnotify.Write || ev.Op&fsnotify.Create == fsnotify.Create {
		if err = f.readfile(); err != nil {
			f.logger.Warnf("Write|Create: readfile of %s failed: %v", f.filename, err)
			return
		}
	} else if ev.Op&fsnotify.Rename == fsnotify.Rename {
		if err = f.readfile(); err != nil {
			f.logger.Warnf("Rename: readfile of %s failed: %v", f.filename, err)
			return
		}
		if err = f.watcher.Remove(ev.Name); err != nil {
			f.logger.Infof("Rename: remove event %s failed: %v", ev.Name, err)
			return
		}
	} else if ev.Op&fsnotify.Remove == fsnotify.Remove {
		return
	} else if ev.Op&fsnotify.Chmod == fsnotify.Chmod {
		//can't distinguish chmod and mv(override)
		if err = f.readfile(); err != nil {
			f.logger.Warnf("Chmod: readfile of %s failed: %v", f.filename, err)
			return
		}
	}
	return true
}
