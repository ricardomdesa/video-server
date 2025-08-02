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
	targetDir := "./assets/videos"
	// modulos := []string{
	// 	"mod1:1-Introdução",
	// }

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

	downloadVideos(sess, env.S3Bucket, targetDir, "mod1:1-Introdução:index.m3u8")

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

	// upload(sess, bucket, targetDir, []string{"mod1:1-Introdução:index.m3u8"})
}


func upload(sess *session.Session, bucketName string, targetDir string, modulos []string) {
	// Cria um uploader S3.
	uploader := s3manager.NewUploader(sess)
	for _, modulo := range modulos {
		err := ListarArquivosEDefinirRedis(targetDir, bucketName, modulo, uploader)
		if err != nil {
			log.Fatalf("Ocorreu um erro: %v", err)
		}
	}
}


func ListarArquivosEDefinirRedis(dirPath string, bucketName string,modulo string, uploader *s3manager.Uploader) error {
	// Obter a lista de itens no diretório
	entries, err := os.ReadDir(dirPath) // Use os.ReadDir para Go 1.16+
	if err != nil {
		return fmt.Errorf("erro ao ler o diretório %s: %w", dirPath, err)
	}

	for _, entry := range entries {
		// Ignorar subdiretórios, queremos apenas arquivos
		if entry.IsDir() {
			continue
		}

		filePath := filepath.Join(dirPath, entry.Name())

		redisKey := fmt.Sprintf("%s:%s", modulo, entry.Name())
		fmt.Println(redisKey)

		// Ler o conteúdo do arquivo
		fileContent, err := os.ReadFile(filePath) // Use os.ReadFile para Go 1.16+
		if err != nil {
			log.Printf("Erro ao ler o arquivo %s: %v", filePath, err)
			continue // Continuar para o próximo arquivo mesmo se este falhar
		}

		uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(redisKey),
			Body:   bytes.NewReader(fileContent),
		})

		fmt.Printf("Arquivo '%s' salvo no S3 com a chave '%s'\n", entry.Name(), redisKey)
	}
	return nil
}