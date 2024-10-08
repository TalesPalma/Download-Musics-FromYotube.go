package converters

import (
	"log"
	"os"
	"path/filepath"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func ConvertMp4ToMp3(fileName string) {
	path := "musics"
	inputFile := filepath.Join(path, fileName)
	outputFile := filepath.Join(path, fileName[0:len(fileName)-4]) + ".mp3" // Remove .mp4 add mp3 extension

	//Convert mp4 to mp3
	err := ffmpeg_go.
		Input(inputFile).
		Output(outputFile, ffmpeg_go.KwArgs{"q:a": 0, "map": "a"}).
		Run()

	if err != nil {
		log.Fatalf("Erro ao converter: %v", err)
	}

	// //Remove file mp4
	os.Remove(inputFile)

	log.Println("Conversão concluída com sucesso!")
}
