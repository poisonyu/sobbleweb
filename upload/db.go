package upload

import "github.com/cyansobble/global"

func CreateFileInfo(f *FileInfo) error {
	return global.DB.Create(f).Error
}
