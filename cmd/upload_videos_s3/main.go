package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/ricardomdesa/videostr/config"
)


func main() {

	env := config.NewEnv()

	log.Print("Starting video saver...")
	targetDir := "../../assets/media/"

	// Cria uma sessão com a configuração do endpoint e credenciais.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"), // A região é necessária, mas não importa no LocalStack.
		// Configura o endpoint para o LocalStack.
		Endpoint: aws.String(env.S3Endpoint),
		// Permite o uso de credenciais fictícias no LocalStack.
		S3ForcePathStyle: aws.Bool(true),
		Credentials: credentials.NewStaticCredentials(env.AWSAccessKeyID, env.AWSSecretAccessKey, ""),
	})
	if err != nil {
		fmt.Println("Erro ao criar a sessão da AWS:", err)
		os.Exit(1)
	}

	fmt.Println("Bucket:", env.S3Bucket)
	if env.S3Bucket == "invalid-bucket" {
		fmt.Println("Bucket inválido, verifique a configuração do LocalStack.")
		os.Exit(1)
	}

	// upload(sess, env.S3Bucket, targetDir, ModulosMap)
	downloadVideos(sess, env.S3Bucket, targetDir, "mod18:1-Apresentando_caso.index.m3u8")

}

func upload(sess *session.Session, bucketName string, targetDir string, modulosMap map[string]string) {
	// Cria um uploader S3.
	uploader := s3manager.NewUploader(sess)
	for nome, modulo := range modulosMap {
		fmt.Printf("path do arquivo : %s e nome do arquivo : %s\n", filepath.Join(targetDir, modulo), nome)
		err := ListarArquivos(filepath.Join(targetDir, modulo), bucketName, nome, uploader)
		if err != nil {
			log.Fatalf("Ocorreu um erro: %v", err)
		}
	}
}

func downloadVideos(sess *session.Session, bucket string, targetDir string, name string) {
	downloader := s3manager.NewDownloader(sess)

	file, err := os.Create(filepath.Join(targetDir, name))
	if err != nil {
		log.Fatalf("Erro ao criar o arquivo %s: %v", targetDir, err)
	}
	defer file.Close()

	_, err = downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(name),
	})
	if err != nil {
		log.Fatalf("Erro ao fazer o download do vídeo: %v", err)
	}

	fmt.Printf("Vídeo baixado com sucesso: %s\n", targetDir)

}

func ListarArquivos(dirPath string, bucketName string, modulo string, uploader *s3manager.Uploader) error {
	// Obter a lista de itens no diretório
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("erro ao ler o diretório %s: %w", dirPath, err)
	}

	for _, entry := range entries {
		// Ignorar subdiretórios, queremos apenas arquivos
		if entry.IsDir() {
			continue
		}

		filePath := filepath.Join(dirPath, entry.Name())

		fileName := fmt.Sprintf("%s:%s", modulo, entry.Name())
		fmt.Println(fileName)

		// Verificar se o arquivo ja existe no S3
		_, err = uploader.S3.HeadObject(&s3.HeadObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(fileName),
		})
		if err == nil {
			fmt.Printf("Arquivo '%s' já existe no S3, pulando upload.\n", fileName)
			continue // Se o arquivo já existe, pular o upload
		}

		// Ler o conteúdo do arquivo
		fileContent, err := os.ReadFile(filePath) 
		if err != nil {
			log.Printf("Erro ao ler o arquivo %s: %v", filePath, err)
			continue // Continuar para o próximo arquivo mesmo se este falhar
		}

		uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(fileName),
			Body:   bytes.NewReader(fileContent),
		})

		fmt.Printf("Arquivo '%s' salvo no S3 com a chave '%s'\n", entry.Name(), fileName)
	}
	return nil
}