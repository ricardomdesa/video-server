package video

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func ConvertSpeedup(inputDir, outputDir string) {
	files, err := os.ReadDir(inputDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		filename := file.Name()
		inputPath := filepath.Join(inputDir, filename)
		outputPath := filepath.Join(outputDir, fmt.Sprintf("conv_%s", filename))

		cmd := fmt.Sprintf(`ffmpeg -i %s -filter_complex "[0:v]setpts=0.6667*PTS[v];[0:a]atempo=1.5[a]" -map "[v]" -map "[a]" %s`, inputPath, outputPath)
		err := exec.Command(cmd).Run()
		if err != nil {
			log.Println(err)
		}
	}
}
