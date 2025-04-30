package main

import (
	"fmt"
	"labAES28_04/ui"
	//"labAES28_04/aesutils"
)

func main() {
	//aqui va algo asi como menu.seleccionar()
	//donde sera cifrar o decifrar
	textHelpGetKey := "Ingresa la llave de "
	quitWithSelected, option := ui.GetOption("Operación a realizar:", []string{"Cifrado", "Decifrado"})
	if !quitWithSelected {
		return
	}

	if option == 0 {
		textHelpGetKey += "cifrado"
	} else {
		textHelpGetKey += "decifrado"
	}
	//var mode int
	quitWithSelected, _ = ui.GetOption("Modo de operación:", []string{"ECB", "CBC", "CFB", "OFB", "CTR"})
	if !quitWithSelected {
		return
	}

	file := ui.GetFile()
	quitWithSelected, _ = ui.GetKey(16, "AES key (16 bytes)", textHelpGetKey)
	if !quitWithSelected {
		return
	}

	quitWithSelected, _ = ui.GetKey(16, "C0 (16 bytes)", "Ingresa el vector de inicialización:")
	if !quitWithSelected {
		return
	}
	fmt.Println(file)

	/* UNICAMENTE CBC

	if option == cifrar {
		*luego sera menu.elegirArchivo(), regresa path*
		*key := []byte("llave_ingresada")*

		inputFile := "nombre_archivo.bmp"
		outputEncrypted := "nombre_archivo_e_cbc.bmp"

		data, err := os.ReadFile(inputFile)
		if err != nil {
			fmt.Println("Error leyendo archivo")
			return
		}

		encryptedData, err := aesutils.encryptAES(data, key)
		if err != nil {
			fmt.Println("Error cifrando")
			return
		}
		os.WriteFile(outputEncrypted,encryptedData, 0644 )
	}

	if option == descifrar {

	}
	*/

	//cifrarArchivoCBC(archivo)
	//cifrarArchivoCFB(archivo)
	//cifrarArchivoOFB(archivo)
	//else
}
