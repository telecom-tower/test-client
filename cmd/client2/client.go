package main

import (
	"context"
	"image"
	"image/color"
	"log"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/telecom-tower/sdk"
	"google.golang.org/grpc"
)

func check(err error, msg string) {
	if err != nil {
		err = errors.WithMessage(err, msg)
		log.Fatal(err)
	}
}

func main() {
	conn, err := grpc.Dial("127.0.0.1:10000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error dialing server: %v", err)
	}
	defer conn.Close() // nolint: errcheck
	client := sdk.NewClient(conn)

	for {
		check(client.StartDrawing(context.Background()), "Error getting frame")
		check(client.Clear(0, 1), "Error clearing display: %v")
		check(client.WriteText(
			"Hello Blue Masters * ", "8x8", 0,
			color.RGBA{128, 128, 255, 255}, 0, sdk.PaintMode),
			"Error writing text")
		check(client.AutoRoll(0, sdk.RollingStart, 0, 0), "Error setting autoroll")
		check(client.Render(), "Error rendering")

		for i := -10; i < 130; i++ {
			check(client.StartDrawing(context.Background()), "Error getting frame")
			check(client.Clear(1), "Error clearing display: %v")
			check(client.AutoRoll(0, sdk.RollingContinue, 0, 0), "Error setting autoroll")
			check(client.DrawRectangle(
				image.Rect(i, 0, i+8, 8), color.RGBA{0, 255, 0, 128},
				1, sdk.PaintMode), "Error setting pixels")
			check(client.Render(), "Error rendering")
			time.Sleep(100 * time.Millisecond)
		}
		time.Sleep(2 * time.Second)

		msg := strings.Split("Hello Telecom Tower, do you hear me ?", " ")
		txt := ""
		for _, w := range msg {
			if txt == "" {
				txt = w
			} else {
				txt = strings.Join([]string{txt, w}, " ")
			}
			check(client.StartDrawing(context.Background()), "Error getting frame")
			check(client.Clear(0, 1), "Error clearing display: %v")
			check(client.AutoRoll(0, 0, 0, 0), "Error stopping autoroll")
			check(client.WriteText(
				txt, "6x8", 0,
				color.RGBA{255, 0, 0, 255}, 0, sdk.PaintMode),
				"Error writing text")
			if len(txt)*6 > 128 {
				check(client.SetLayerOrigin(0, image.Point{len(txt)*6 - 128, 0}), "Error")
			}
			check(client.Render(), "Error rendering")
			time.Sleep(600 * time.Millisecond)
		}

		time.Sleep(1 * time.Second)

	}
}
