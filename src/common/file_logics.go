package common

import (
	"io"
	"net/http"
	"os"

	"github.com/google/uuid"
)

type FileLogics interface {
	DownloadSaveImage(imageUrl string, saveFilePath string) (string, error)
}

type FileLogicsImp struct{}

/*
指定されたURL画像を別に指定した場所に保存する。
*/
func (fileLogicsImp *FileLogicsImp) DownloadSaveImage(imageUrl string, saveFilePath string) (string, error) {
	// 画像情報をダウンロードする。
	response, err := http.Get(imageUrl)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// 画像名用のUUIDを生成
	imageName, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	// 画像ファイルの作成
	outPutImage, err := os.Create(saveFilePath + imageName.String() + ".png")
	if err != nil {
		return "", err
	}
	defer outPutImage.Close()

	// 作成した画像ファイルにダウンロードした画像情報に上書きする。
	_, err = io.Copy(outPutImage, response.Body)
	if err != nil {
		return "", err
	}

	return outPutImage.Name(), nil
}
