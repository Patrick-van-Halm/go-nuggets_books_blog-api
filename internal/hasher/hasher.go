package hasher

import (
	"github.com/speps/go-hashids"
	"log"
	"os"
	"strconv"
)

var idHasher *hashids.HashID
func init() {
	length, err := strconv.Atoi(os.Getenv("ID_HASHER_MINLENGTH"))
	if err != nil {
		log.Fatalln(err)
		return
	}
	hd := hashids.NewData()
	hd.Salt = os.Getenv("ID_HASHER_SALT")
	hd.MinLength = length
	idHasher, err = hashids.NewWithData(hd)
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func HashID(id int) string {
	hash, err := idHasher.Encode([]int{id})
	if err != nil {
		log.Fatalln(err)
		return ""
	}
	return hash
}

func GetFromHashID(hash string) int {
	numbers, err := idHasher.DecodeWithError(hash)
	if err != nil {
		log.Fatalln(err)
		return 0
	}

	return numbers[0]
}