package upload

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/cyansobble/global"
	"go.uber.org/zap"
)

func SaveUploadedFile(file *multipart.FileHeader) (string, error) {
	// mainProcess, _ := os.Executable()
	// mainPath := filepath.Dir(mainProcess)
	// dst := filepath.Join(mainPath, global.CONFIG.Local.Path, file.Filename)

	dst := filepath.Join(global.CONFIG.Local.Path, file.Filename)
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	if err = os.MkdirAll(filepath.Dir(dst), 0750); err != nil {
		global.LOGGER.Error("make directory dst failed", zap.Error(err))
		return "", err
	}

	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, src)

	return dst, err
}
