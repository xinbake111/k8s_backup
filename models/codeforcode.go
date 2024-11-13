package models

//加解密模块
type Encrypter struct {
	keysMap          map[string]string
	DictionaryEncMap map[string]string
}

func Constructor() Encrypter {
	var keys = []string{"a", "b", "c", "d","e","f","0","1","2","3","4","5","6","7","8","9"}
	var values = []string{"[", "]", "{", "}","(",")","-","*","&","^","!","@","#","$","%","+"}
	DictionaryEncMap := make(map[string]string)
	keysMap := make(map[string]string)
	for i := 0; i < len(keys); i++ {
		keysMap[keys[i]] = values[i]
		DictionaryEncMap[values[i]]=keys[i]
	}
	e := Encrypter{keysMap, DictionaryEncMap}
	return e
}

func (this *Encrypter) Encrypt(word1 string) string {
	var build string
	for i := 0; i < len(word1); i++ {
		build =build+this.keysMap[string(word1[i])]
	}
	return build
}

func (this *Encrypter) Decrypt(word2 string) string {

	var build string
	for i := 0; i < len(word2); i++ {
		build =build+this.DictionaryEncMap[string(word2[i])]
	}
	return build


}



