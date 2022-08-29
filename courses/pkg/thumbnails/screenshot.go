package thumbnails

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/chromedp/chromedp"
)

type ScreenShot struct {
	Ctx context.Context
	// Quality is the quality of the image to generate
	Quality int
	// LoadTime is the amount of time to wait before taking the screenshot
	LoadTime int
	// The name to save the file as
	Name string
	// Output directory
	OutputDir string
}

func NewScreenshot(ctx context.Context, name string) *ScreenShot {
	return &ScreenShot{
		Ctx:      ctx,
		Quality:  300,
		LoadTime: 5,
		Name:     name,
	}
}

func (s *ScreenShot) Generate(screenShotUrl string) error {
	//byte slice to hold captured image in bytes
	var buf []byte

	//setting image file extension to png but
	var ext string = "png"
	//if image quality is less than 100 file extension is jpeg
	if s.Quality < 100 {
		ext = "jpeg"
	}

	//setting options for headless chrome to execute with
	var options []chromedp.ExecAllocatorOption
	options = append(options, chromedp.WindowSize(1400, 900))
	options = append(options, chromedp.DefaultExecAllocatorOptions[:]...)

	//setup context with options
	actx, acancel := chromedp.NewExecAllocator(s.Ctx, options...)
	defer acancel()

	// create context
	ctx, cancel := chromedp.NewContext(actx)
	defer cancel()

	//configuring a set of tasks to be run
	tasks := chromedp.Tasks{
		//loads page of the URL
		chromedp.Navigate(screenShotUrl),
		chromedp.Sleep(time.Duration(s.LoadTime) * time.Second),
		//Captures Screenshot with current window size
		chromedp.CaptureScreenshot(&buf),
	}

	// running the tasks configured earlier and logging any errors
	if err := chromedp.Run(ctx, tasks); err != nil {
		return err
	}

	//write byte slice data of standard screenshot to file
	err := ioutil.WriteFile(
		fmt.Sprintf("%s/%s.%s",
			s.OutputDir,
			s.Name,
			ext,
		), buf, 0644)
	return err

}
