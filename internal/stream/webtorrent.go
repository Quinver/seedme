package stream

import (
	"os"
	"os/exec"
)

func Play(magnet string) error {
	cmd := exec.Command("webtorrent", magnet, "--interactive-select", "--mpv", "--player-args=--ontop --no-border --fs")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
