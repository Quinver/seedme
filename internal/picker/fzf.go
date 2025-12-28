package picker

import (
	"bytes"
	"fmt"
	"os/exec"
	"seedme/internal/model"
	"strings"
)

func Pick(torrents []model.Torrent) (model.Torrent, error) {
	var input bytes.Buffer

	for _, t := range torrents {
		fmt.Fprintf(
			&input,
			"[%s] [%4d seeds] %s | %s\n",
			t.Site,
			t.Seeds,
			t.Title,
			t.Magnet,
		)
	}

	cmd := exec.Command("fzf", "--prompt=Select torrent: ")
	cmd.Stdin = &input

	out, err := cmd.Output()
	if err != nil {
		return model.Torrent{}, err
	}

	line := strings.TrimSpace(string(out))
	parts := strings.SplitN(line, "|", 2)
	if len(parts) != 2 {
		return model.Torrent{}, fmt.Errorf("invalid selection")
	}
	magnet := strings.TrimSpace(parts[1])

	for _, t := range torrents {
		if t.Magnet == magnet {
			return t, nil
		}
	}

	return model.Torrent{}, fmt.Errorf("no torrent selected")
}
