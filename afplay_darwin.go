package goafplay

import (
	"fmt"
	"os/exec"
)

type Cmd struct {
	Voice string
	Rate  float64
	Time  float64
}

func NewCmd() *Cmd {
	return &Cmd{}
}
func NewCmdV(voice string) *Cmd {
	return &Cmd{
		Voice: voice,
	}
}
func NewCmdT(t float64) *Cmd {
	return &Cmd{
		Time: t,
	}
}
func NewCmdVTR(voice string, t, rate float64) *Cmd {
	return &Cmd{
		Voice: voice,
		Rate:  rate,
		Time:  t,
	}
}

func (c *Cmd) SetVoice(voice string) {
	c.Voice = voice
}
func (c *Cmd) SetRate(rate float64) {
	c.Rate = rate
}
func (c *Cmd) SetTime(t float64) {
	c.Time = t
}

func (c Cmd) Do(message string) (err error) {
	args := make([]string, 0)
	if c.Voice != "" {
		args = append(args, "-v", c.Voice)
	}
	if c.Rate > 0.0 {
		args = append(args, "-r", fmt.Sprint(c.Rate))
	}
	if c.Time > 0.0 {
		args = append(args, "-t", fmt.Sprint(c.Time))
	}
	args = append(args, message)
	out, err := exec.Command("afplay", args...).CombinedOutput()
	fmt.Println(string(out))
	return
}

var AudioFileOpenErr = fmt.Errorf("Error: AudioFileOpen failed ('wht?')")

func Do(message string) (err error) {
	out, err := exec.Command("afplay", message).CombinedOutput()
	str := string(out)
	if str == "Error: AudioFileOpen failed ('wht?')\n" {
		return AudioFileOpenErr
	}
	return
}
