package aesutils

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strings"
)

type HeaderBMP struct {
	Signature       [2]byte
	FileSize        uint32
	Reserved1       uint16
	Reserved2       uint16
	PixelOffset     uint32
	DIBHeaderSize   uint32
	Width           int32
	Height          int32
	Planes          uint16
	BitsPerPixel    uint16
	Compression     uint32
	ImageSize       uint32
	XPixelsPerM     int32
	YPixelsPerM     int32
	ColorsUsed      uint32
	ColorsImportant uint32
}

func ReadBmp(fileName string) (headerStruct HeaderBMP, headerBytes, bytesBMP []byte) {
	file, errOpen := os.Open(fileName)
	if errOpen != nil {
		panic(errOpen)
	}
	defer file.Close()

	var header HeaderBMP

	// Leer el header
	if err := binary.Read(file, binary.LittleEndian, &header); err != nil {
		panic(err)
	}
	headerStruct = header

	// Leer todo el archivo (incluso los bytes extra del padding)
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	totalSize := stat.Size()

	// Leer los datos de píxeles (todo después del header)
	pixelDataSize := totalSize - int64(header.PixelOffset)
	bytesBMP = make([]byte, pixelDataSize)

	_, err = file.Seek(int64(header.PixelOffset), io.SeekStart)
	if err != nil {
		panic(err)
	}
	_, err = io.ReadFull(file, bytesBMP)
	if err != nil {
		panic(err)
	}

	// Leer header en bytes
	_, _ = file.Seek(0, io.SeekStart)
	headerBytes = make([]byte, header.PixelOffset)
	_, _ = io.ReadFull(file, headerBytes)

	return
}

func WriteBmp(fileName string, header, pixelContent []byte) {
	out, errCreate := os.Create(fileName)
	if errCreate != nil {
		panic(errCreate)
	}
	defer func(out *os.File) {
		errClose := out.Close()
		if errClose != nil {
			panic(errClose)
		}
	}(out)

	_, err := out.Write(header)
	if err != nil {
		panic(err)
	}
	_, err = out.Write(pixelContent)
	if err != nil {
		panic(err)
	}
}

func WriteBmpWithHeaderStruct(fileName string, header HeaderBMP, pixelContent []byte) error {
	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer out.Close()

	// Ajustamos los tamaños del header antes de escribir
	header.ImageSize = uint32(len(pixelContent))
	header.FileSize = header.PixelOffset + header.ImageSize

	// Escribimos el header serializado
	err = binary.Write(out, binary.LittleEndian, header)
	if err != nil {
		return err
	}

	// En caso de que PixelOffset sea mayor que el tamaño del header (por ejemplo si hay paleta), rellenamos con ceros
	headerSize := binary.Size(header)
	padding := int(header.PixelOffset) - headerSize
	if padding > 0 {
		_, err = out.Write(make([]byte, padding))
		if err != nil {
			return err
		}
	}

	// Escribimos el contenido (cifrado)
	_, err = out.Write(pixelContent)
	if err != nil {
		return err
	}

	return nil
}

func InvertImage(h HeaderBMP, pixels []byte) []byte {
	rowSize := ((int(h.Width)*3 + 3) / 4) * 4
	for y := 0; y < int(h.Height); y++ {
		rowStart := y * rowSize
		for x := 0; x < int(h.Width); x++ {
			i := rowStart + x*3
			pixels[i] = 255 - pixels[i]
			pixels[i+1] = 255 - pixels[i+1]
			pixels[i+2] = 255 - pixels[i+2]
		}
	}
	return pixels
}

func GetNewBMPFilename(prevName, method string) string {
	if splitted, ok := strings.CutSuffix(prevName, ".bmp"); ok {
		return splitted + "_" + method + ".bmp"
	}
	panic("Not a BMP file!")
}

func PrintBMPRGBValues(pixels []byte, width, height int) {
	rowSize := ((width*3 + 3) / 4) * 4

	fmt.Println("===== BMP RGB Pixel Data =====")
	for y := height - 1; y >= 0; y-- { // BMP guarda desde la fila inferior
		rowStart := y * rowSize
		for x := 0; x < width; x++ {
			i := rowStart + x*3
			if i+2 >= len(pixels) {
				fmt.Print("(?, ?, ?) ")
				continue
			}
			b, g, r := pixels[i], pixels[i+1], pixels[i+2]
			fmt.Printf("(%3d,%3d,%3d) ", r, g, b)
		}
		fmt.Println()
	}
	fmt.Println("==============================")
}
