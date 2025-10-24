package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

func EncryptLargeFiles(infile, outfile string, key []byte) error {
	buf := make([]byte, 4096)  //numero mais adequado ao que faça sentido para vc
	in, err := os.Open(infile) // lendo o aquivo
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.OpenFile(outfile, os.O_RDWR|os.O_CREATE, 0777) //abre ou cria o arquivo de out
	if err != nil {
		return err
	}
	defer out.Close()

	block, err := aes.NewCipher(key) //cria o nosso bloc com o tamanho da key, ex: key de 16bits e etc.....
	if err != nil {
		return err
	}

	// marcado de marcação de autencidade da mensagem
	iv := make([]byte, block.BlockSize()) // inicialization vector (vertor de inciaalização) , estamos criando um slice de bites
	// no io.ReadFull estamos lendo um valor randomico do tamanho do nosso vetor de bites de entrada para uma saida
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	// bloco de sifer
	stream := cipher.NewCTR(block, iv)
	for {
		n, err := in.Read(buf) // lendo do arquivo de entrada
		if n > 0 {
			stream.XORKeyStream(buf, buf[:n]) //lendo de pedaço em pedaço ou seja de 4096 em 4096
			out.Write(buf[:n])
		}
		// verificando se o valor e igual a "endo file"
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}

	out.Write(iv) //devolvemos o inicialization value
	return nil
}
