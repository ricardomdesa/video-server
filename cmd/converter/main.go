package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	inputDir = "./"
)

func main() {
	convertSpeedup(inputDir)
}


func convertSpeedup(inputDir string) {
	files, err := os.ReadDir(inputDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		filename := file.Name()
		if fileExtension := filepath.Ext(filename); fileExtension == ".mp4" {
			inputPath := filepath.Join(inputDir, filename)
			outputPath := fmt.Sprintf("conv_%s", filename)

			cmd := fmt.Sprintf(`/opt/homebrew/bin/ffmpeg -i %s -filter_complex "[0:v]setpts=0.6667*PTS[v];[0:a]atempo=1.5[a]" -map "[v]" -map "[a]" %s`, inputPath, outputPath)
			err := exec.Command(cmd).Run()
			if err != nil {
				log.Println("Error aqui")
				log.Println(err)
			}
		}
	}
}
