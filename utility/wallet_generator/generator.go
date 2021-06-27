package main

import (
	"fmt"
	"github.com/overseven/try-network/utility"
)

func main() {
	//flagPrefix := flag.String("prefix", "", "generate public key with prefix")
	//flag.Parse()
	priv, pub, err := utility.Generate()
	if err != nil {
		fmt.Println("Error! ", err.Error())
		return
	}

	//if *flagPrefix != "" {
	//	_, err := strconv.ParseUint(*flagPrefix, 16, 64)
	//	if err != nil {
	//		fmt.Println("Prefix error!", err.Error())
	//		return
	//	}
	//
	//	for strings.ToLower(pubStr[:len(*flagPrefix)]) != strings.ToLower(*flagPrefix) {
	//		priv, err = cr.GenerateKey()
	//		if err != nil {
	//			fmt.Println("Error!", err.Error())
	//			return
	//		}
	//		pub = cr.CompressPubkey(&priv.PublicKey)
	//		pubStr = hex.EncodeToString(pub)
	//		fmt.Println("pub:", pubStr)
	//		fmt.Println(strings.ToLower(pubStr[:len(*flagPrefix)]), strings.ToLower(*flagPrefix))
	//	}
	//}

	fmt.Println("Private:", priv)
	fmt.Println("Public: ", pub)

	_, err = utility.ParseKeys(priv, pub)

	if err != nil {
		fmt.Println("Error is raised at the check of keys generation!", err.Error())
	}
}
