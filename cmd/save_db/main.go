package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/ricardomdesa/videostr/config"
	"github.com/ricardomdesa/videostr/internal/api/repositories/persistence"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func main() {
	log.Info("Starting video saver...")
	targetDir := "./assets/media/mod1/1-Introdução"
	modulos := []string{
		"mod1:1-Introdução",
	}

	env := config.NewEnv()
	redisConn := config.NewRedis(env, 0)
	defer redisConn.Close()

	persistenceRepo := persistence.NewRedisRepository(redisConn)
	if err := persistenceRepo.SaveJson(context.Background(), "./assets/media/mod.json", "classes_json"); err != nil {
		log.Fatalf("Failed to save classes JSON: %v", err)
		return
	}
	for _, modulo := range modulos {
		err := ListarArquivosEDefinirRedis(targetDir, modulo, redisConn)
		if err != nil {
			log.Fatalf("Ocorreu um erro: %v", err)
		}
	}
	log.Info("Video saver finished successfully")
	if err := redisConn.Close(); err != nil {
		log.Errorf("Failed to close Redis connection: %v", err)
	} else {
		log.Info("Redis connection closed successfully")
	}
	log.Info("Exiting video saver")
}

func ListarArquivosEDefinirRedis(dirPath string, modulo string, redisClient *redis.Client) error {
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

		// Salvar o conteúdo do arquivo (em bytes) no Redis
		// Usamos context.Background() para um contexto simples. Em aplicações reais, use um contexto apropriado.
		err = redisClient.Set(context.Background(), redisKey, fileContent, 0).Err()
		if err != nil {
			log.Printf("Erro ao salvar o arquivo %s no Redis com a chave %s: %v", filePath, redisKey, err)
			continue // Continuar para o próximo arquivo
		}

		fmt.Printf("Arquivo '%s' salvo no Redis com a chave '%s'\n", entry.Name(), redisKey)
	}
	return nil
}
