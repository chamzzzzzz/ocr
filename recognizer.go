package ocr

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type Image struct {
	File   string
	Width  int
	Height int
}

type Size struct {
	Width  int
	Height int
}

type Point struct {
	X int
	Y int
}

type BoudingBox struct {
	Origin Point
	Size   Size
}

type Observation struct {
	Confidence int
	Text       string
	BoudingBox BoudingBox
}

type Result struct {
	Error        error
	Code         string
	Message      string
	Image        Image
	Observations []*Observation
	Cost         time.Duration
	TotalCost    time.Duration
}

type Recognizer interface {
	Recognize(file string) *Result
}

type MacRecognizer struct {
}

func (recognizer *MacRecognizer) Recognize(file string) *Result {
	result := &Result{}
	cmd := exec.Command("mac-ocr-cli", file)
	var out bytes.Buffer
	cmd.Stdout = &out

	start := time.Now()
	err := cmd.Run()
	elapsed := time.Now().Sub(start)
	result.Code = fmt.Sprintf("%d", cmd.ProcessState.ExitCode())
	result.Cost = elapsed
	result.TotalCost = elapsed
	output := strings.TrimSuffix(out.String(), "\n")
	if err != nil {
		result.Error = err
		result.Message = output
		return result
	}

	for i, line := range strings.Split(output, "\n") {
		if i == 0 {
			fields := strings.SplitN(line, " ", 2)
			if len(fields) == 2 {
				result.Image.File = fields[1]
				resolution := strings.Split(fields[0], "x")
				if len(resolution) == 2 {
					result.Image.Width, _ = strconv.Atoi(resolution[0])
					result.Image.Height, _ = strconv.Atoi(resolution[1])
				}
			}

		} else {
			fields := strings.SplitN(line, " ", 3)
			if len(fields) == 3 {
				observation := &Observation{}
				if normalizeConfidence, err := strconv.ParseFloat(fields[0], 64); err == nil {
					observation.Confidence = int(normalizeConfidence * 100)
				}
				observation.Text = fields[2]

				boudingBox := strings.Split(strings.Trim(fields[1], "[]"), ",")
				if len(boudingBox) == 4 {
					observation.BoudingBox.Origin.X, _ = strconv.Atoi(boudingBox[0])
					observation.BoudingBox.Origin.Y, _ = strconv.Atoi(boudingBox[1])
					observation.BoudingBox.Size.Width, _ = strconv.Atoi(boudingBox[2])
					observation.BoudingBox.Size.Height, _ = strconv.Atoi(boudingBox[3])
				}
				result.Observations = append(result.Observations, observation)
			}
		}
	}
	return result
}
