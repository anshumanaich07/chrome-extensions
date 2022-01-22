package audioservice

import (
	"fmt"
	"os/exec"
	"regexp"
)

func ConvertToAudio(ytURL string, audioTitle string) string {
	fmt.Println("youtube video URL: ", ytURL)

	// ytVideoIDLength := 11
	// ytVideoID := ytURL[len(ytURL)-ytVideoIDLength:]

	//command format:  youtube-dl --extract-audio --audio-format mp3 <link>
	cmdArgs := []string{}
	cmdArgs = append(cmdArgs, "--extract-audio")
	cmdArgs = append(cmdArgs, "--audio-format")
	cmdArgs = append(cmdArgs, "mp3")
	cmdArgs = append(cmdArgs, "--output")
	cmdArgs = append(cmdArgs, audioTitle)
	cmdArgs = append(cmdArgs, ytURL)

	cmd := exec.Command("youtube-dl", cmdArgs...)
	stdout, _ := cmd.StdoutPipe()
	oneByte := make([]byte, 100)
	// bar := progressbar.Default(100)

	cmd.Start()

	for {
		// bar.Add(1)
		_, err := stdout.Read(oneByte)
		if err != nil {
			break
		}
		r, _ := regexp.Compile("(100|(\\d{1,2}(\\.\\d+)*))%")

		downloadStatus := r.Find(oneByte)
		downloadStatusStr := string(downloadStatus)
		fmt.Println(downloadStatusStr)
	}

	cmd.Wait()
	return audioTitle
}
