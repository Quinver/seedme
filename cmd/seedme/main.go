package main

import (
	"fmt"
	"os"
	"os/exec"
	"seedme/internal/picker"
	"seedme/internal/search"
	"seedme/internal/stream"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env found")
	}
}

func main() {
	for _, dep := range []string{"webtorrent", "fzf", "mpv"} {
		if err := require(dep); err != nil {
			fmt.Println("error:", err)
			fmt.Println("please install it and try again")
			os.Exit(1)
		}
	}

	query := os.Args[1:]

	results, err := search.All(query)
	if err != nil {
		panic(err)
	}

	choice, err := picker.Pick(results)
	if err != nil {
		panic(err)
	}

	if err := stream.Play(choice.Magnet); err != nil {
		panic(err)
	}
}

func require(cmd string) error {
	_, err := exec.LookPath(cmd)
	if err != nil {
		return fmt.Errorf("%s not found in PATH", cmd)
	}
	return nil
}
