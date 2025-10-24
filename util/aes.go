package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
	"os"
)

func DecryptLargeFile(fileInput, fileOutput string, key []byte) error {
	infile, err := os.Open(fileInput)
	if err != nil {
		return err
	}
	defer infile.Close()
	block, err := aes.NewCipher(key) //chave e o principal item de segurança e deve ser passada por varivael de ambiente por exemplo
	if err != nil {
		return err
	}
	fi, err := infile.Stat() //verifica se arquivo exist, | dar o describe no arquivo
	if err != nil {
		return err
	}
	iv := make([]byte, block.BlockSize()) // inicialization vector (vertor de inciaalização) e o tamanho e determinado pela key
	// tamanho da mensagem, vai acontecer um calculo e retirar o iv , por a mensagem original não tem, pois ele foi insetido no final do arquivo
	msgLen := fi.Size() - int64(len(iv)) // mensagem - o ponto
	if err != nil {
		log.Fatal("erro lendo o iv: ", err)
	}

	// nosso aruivo escriptografado
	outfile, err := os.OpenFile(fileOutput, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer outfile.Close()
	buf := make([]byte, 4096)          //deve ser multiplo de 16bits
	stream := cipher.NewCTR(block, iv) //como se foce nosso fluxo de dados, lembrar de usar a mesma criptografia (mesmo algoritimo)
	for {
		n, err := infile.Read(buf) //o n se refere ao nosso blocos
		if n > 0 {
			// verificando se os ultimos bytes são o retorno da inicialização
			if n > int(msgLen) { //essa validação e para previnir de não tentarmos decriptografar o nosso array de inicialização
				n = int(msgLen)
			}
			msgLen -= int64(n)

			stream.XORKeyStream(buf, buf[:n]) //o xor para reverter(voltar para a configuração original)
			outfile.Write(buf[:n])            //gravação do buffer já convertido
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("tamanho de bytes lidos: %d:", n)
			break
		}
	}

	return nil
}

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
