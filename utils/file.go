package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

const path = "./assets/images"

func StoreFile(file *multipart.FileHeader, name string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			return err
		}
	}

	src, err := file.Open()
	if err != nil {
		return err
	}

	d, err := os.Create(fmt.Sprintf("./assets/images/%s", name))
	if err != nil {
		return err
	}

	_, err = io.Copy(d, src)
	return err
}

func StoreFiles(files []*multipart.FileHeader, name string) (err error) {
	if _, err = os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
		return
	}

	for index, file := range files {
		src, err := file.Open()
		if err != nil {
			fmt.Println(err)
		}

		d, err := os.Create(fmt.Sprintf("./assets/images/%s-%d", name, index+1))
		if err != nil {
			fmt.Println(err)
		}

		_, err = io.Copy(d, src)
		if err != nil {
			fmt.Println(err)
		}
	}

	return
}
