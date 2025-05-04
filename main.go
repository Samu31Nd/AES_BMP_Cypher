package main

import (
	"labAES28_04/aesutils"
	"labAES28_04/ui"
	"log"
)

func main() {
	quitWithSelected, option := ui.GetOption("Operación a realizar:", []string{"Cifrado", "Decifrado"})
	if !quitWithSelected {
		return
	}
	var sufix string
	textHelpGetKey := []string{"Ingresa la llave de ", ""}
	if option == 0 {
		textHelpGetKey[0] += "cifrado"
		textHelpGetKey[1] += "cifrada"
		sufix = "e"
	} else {
		textHelpGetKey[0] += "decifrado"
		textHelpGetKey[1] += "decifrada"
		sufix = "d"
	}
	var mode int
	quitWithSelected, mode = ui.GetOption("Modo de operación:", []string{"ECB", "CBC", "CFB", "OFB", "CTR"})
	if !quitWithSelected {
		return
	}

	quitWithSelected, file := ui.GetFile()
	if !quitWithSelected {
		return
	}
	h, _, pixels := aesutils.ReadBmp(file)
	var key string
	quitWithSelected, key = ui.GetKey(16, "AES key (16 bytes)", textHelpGetKey[0])
	if !quitWithSelected {
		return
	}

	var c0 string
	if mode != 0 {
		quitWithSelected, c0 = ui.GetKey(16, "C0 (16 bytes)", "Ingresa el vector de inicialización:")
		if !quitWithSelected {
			return
		}
	}

	var err error
	var newName string
	var encryptedPixels []byte
	switch mode {
	case 0:
		if option == 0 {
			encryptedPixels, err = aesutils.CifrarAES_ECB([]byte(key), pixels)
		} else {
			encryptedPixels, err = aesutils.DecifrarAES_ECB([]byte(key), pixels)
		}
		newName = aesutils.GetNewBMPFilename(file, sufix+"ECB")
	case 1:
		if option == 0 {
			encryptedPixels, err = aesutils.CifrarAES_CBC([]byte(c0), []byte(key), pixels)
		} else {
			encryptedPixels, err = aesutils.DecifrarAES_CBC([]byte(c0), []byte(key), pixels)
		}
		newName = aesutils.GetNewBMPFilename(file, sufix+"CBC")
	case 2:
		if option == 0 {
			encryptedPixels, err = aesutils.CifrarAES_CFB([]byte(c0), []byte(key), pixels)
		} else {
			encryptedPixels, err = aesutils.DecifrarAES_CFB([]byte(c0), []byte(key), pixels)
		}
		newName = aesutils.GetNewBMPFilename(file, sufix+"CFB")
	case 3:
		if option == 0 {
			encryptedPixels, err = aesutils.CifrarAES_OFB([]byte(c0), []byte(key), pixels)
		} else {
			encryptedPixels, err = aesutils.DecifrarAES_OFB([]byte(c0), []byte(key), pixels)
		}
		newName = aesutils.GetNewBMPFilename(file, sufix+"OFB")
	case 4:
		if option == 0 {
			encryptedPixels, err = aesutils.CifrarAES_CTR([]byte(c0), []byte(key), pixels)
		} else {
			encryptedPixels, err = aesutils.DecifrarAES_CTR([]byte(c0), []byte(key), pixels)
		}
		newName = aesutils.GetNewBMPFilename(file, sufix+"CTR")
	default:
		panic("Index out of range!")
	}

	if err != nil {
		ui.ShowMsgDialog("La imagen tuvo errores al ser cifrada", true)
		log.Fatal(err)
	}

	err = aesutils.WriteBmpWithHeaderStruct(newName, h, encryptedPixels)
	if err != nil {
		ui.ShowMsgDialog("La imagen no pudo ser "+textHelpGetKey[1]+"!\n"+err.Error(), true)
	} else {
		ui.ShowMsgDialog("La imagen "+newName+" ha sido "+textHelpGetKey[1]+" con exito!", false)
	}

}
