package helpers

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/lucsky/cuid"
	"golang.org/x/crypto/bcrypt"
)

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "../..")
}

func StringToHash(str string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), 10)
	return string(bytes), err
}

func CompareStringAndHash(str, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	return err == nil
}

func UUID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)

	if err != nil {
		fmt.Println("UUID Generation uuid has error: ", err)
		return "", err
	}

	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}

func SHA1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func MD5(str string) string {
	hashed := md5.New()
	hashed.Write([]byte(str))
	return hex.EncodeToString(hashed.Sum(nil))
}

func CUID() string {
	return cuid.New()
}
