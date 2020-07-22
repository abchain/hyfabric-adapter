package userstore

import (
	"bsn-sdk-go/pkg/common/errors"
	"bsn-sdk-go/pkg/core/entity/msp"
	"io/ioutil"
	"os"
	path2 "path"
	"strings"
)

func NewUserStore(path string) UserStore {
	us := &FileUserStore{
		FilePath: path,
	}

	return us
}

type FileUserStore struct {
	FilePath string
}

func (f *FileUserStore) Load(user *msp.UserData) error {
	key := storeKeyName(user)
	file := path2.Join(f.FilePath, key)

	if _, err1 := os.Stat(file); os.IsNotExist(err1) {
		return errors.New("user not found")
	}

	bytes, err := ioutil.ReadFile(file) // nolint: gas
	if err != nil {
		return err
	}
	if bytes == nil {
		return errors.New("user not found")
	}
	user.EnrollmentCertificate = bytes
	return nil

}

func (f *FileUserStore) Store(user *msp.UserData) error {
	key := storeKeyName(user)

	path := path2.Join(f.FilePath, key)

	valueBytes := user.EnrollmentCertificate

	err := os.MkdirAll(path2.Dir(path), 0700)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, valueBytes, 0600)

}

func (f *FileUserStore) LoadAll(appCode string) []*msp.UserData {

	var users []*msp.UserData

	//遍历文件夹下的文件
	files, err := ioutil.ReadDir(f.FilePath)
	if err != nil {
		return users
	}

	for _, file := range files {
		filePath := path2.Join(f.FilePath, file.Name())

		//判断文件名
		name := getPemName(file.Name(), appCode)
		if name != "" {
			//获取
			user := &msp.UserData{}
			bytes, err := ioutil.ReadFile(filePath) // nolint: gas
			if err == nil && bytes != nil {
				user.EnrollmentCertificate = bytes
				user.UserName = name
				user.AppCode = appCode
				users = append(users, user)
			}
		}
	}

	return users

}

func getPemName(name, appCode string) string {

	ext := "@" + appCode + "-cert.pem"

	i := strings.Index(name, ext)
	if i != -1 {
		return name[:i]
	} else {
		return ""
	}
}

func storeKeyName(user *msp.UserData) string {
	return user.UserName + "@" + user.AppCode + "-cert.pem"
}
