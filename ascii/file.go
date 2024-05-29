package ascii

import (
	"hash/fnv"
	"os"
	"strings"
)

func ReadFile(ds string) [][]string {
	var banner []byte
	var err error
	// Чтение баннера по запросу
	if ds == "standard" {
		banner, err = os.ReadFile("./banners/standard.txt")
		if err != nil {
			return nil
		}
	} else if ds == "shadow" {
		banner, err = os.ReadFile("./banners/shadow.txt")
		if err != nil {
			return nil
		}
	} else if ds == "thinkertoy" {
		banner, err = os.ReadFile("./banners/thinkertoy.txt")
		banner = []byte(strings.ReplaceAll(string(banner), "\r", ""))
		if err != nil {
			return nil
		}
	} else {
		return nil
	}

	// Проверяет файл на изменения
	if !CheckFile(string(banner)) {
		return nil
	}

	// Разделение баннера по символам ASCII
	strBannerArr := strings.Split(string(banner), "\n\n")
	strBannerArr2 := make([][]string, len(strBannerArr))

	for i := range strBannerArr {
		strBannerArr2[i] = strings.Split(strBannerArr[i], "\n")
	}

	return strBannerArr2
}

func CheckFile(style string) bool {
	hasher := fnv.New32a()
	hasher.Write([]byte(style))
	hashValue := hasher.Sum32()
	if hashValue != 1766917683 && hashValue != 4281396044 && hashValue != 3075161722 {
		return false

	}
	return true
}
