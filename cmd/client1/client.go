package main

import (
	"context"
	"image"
	"image/color"
	"log"
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
	// conn, err := grpc.Dial("telecom-tower.sofr.hefr.lan:10000", grpc.WithInsecure())
	conn, err := grpc.Dial("127.0.0.1:10000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error dialing server: %v", err)
	}
	defer conn.Close() // nolint: errcheck
	client := sdk.NewClient(conn)

	for {
		log.Print("Frame 1")
		check(client.StartDrawing(context.Background()), "Error getting frame")
		check(client.Init(), "Error during init")
		check(client.SetPixels(
			[]sdk.Pixel{
				sdk.Pixel{
					Point: image.Point{0, 0},
					Color: color.RGBA{255, 60, 60, 255},
				},
				sdk.Pixel{
					Point: image.Point{1, 1},
					Color: color.RGBA{60, 255, 255, 255},
				},
			}, 0, sdk.PaintMode), "Error setting pixels")

		check(client.DrawRectangle(
			image.Rect(2, 0, 6, 6), color.RGBA{255, 255, 255, 255},
			0, sdk.PaintMode), "Error setting pixels")

		check(client.DrawRectangle(
			image.Rect(4, 2, 8, 8), color.RGBA{0, 255, 0, 128},
			1, sdk.PaintMode), "Error setting pixels")

		check(client.WriteText(
			"Hello nice little wonderful world", "8x8", 10,
			color.RGBA{255, 255, 255, 255}, 0, sdk.PaintMode),
			"Error writing text")
		check(client.SetLayerOrigin(0, image.Point{0, 0}), "Error moving origin")
		check(client.WriteText(
			"World", "6x8", 20,
			color.RGBA{32, 32, 255, 180}, 1, sdk.OverMode),
			"Error writing text")
		check(client.Render(), "Error rendering")

		log.Print("Moving frame")
		for i := 1; i <= 20; i++ {
			check(client.StartDrawing(context.Background()), "Error getting stream")
			check(client.SetLayerOrigin(0, image.Point{i, 0}), "Error moving origin")
			check(client.Render(), "Error rendering")
		}
		for i := 1; i <= 8; i++ {
			check(client.StartDrawing(context.Background()), "Error getting stream")
			check(client.SetLayerOrigin(0, image.Point{20, i}), "Error moving origin")
			check(client.Render(), "Error rendering")
		}
		for i := 8; i >= -8; i-- {
			check(client.StartDrawing(context.Background()), "Error getting stream")
			check(client.SetLayerOrigin(0, image.Point{20, i}), "Error moving origin")
			check(client.Render(), "Error rendering")
		}
		for i := -8; i <= 0; i++ {
			check(client.StartDrawing(context.Background()), "Error getting stream")
			check(client.SetLayerOrigin(0, image.Point{20, i}), "Error moving origin")
			check(client.Render(), "Error rendering")
		}
		for i := 20; i >= 0; i-- {
			check(client.StartDrawing(context.Background()), "Error getting stream")
			check(client.SetLayerOrigin(0, image.Point{i, 0}), "Error moving origin")
			check(client.Render(), "Error rendering")
		}

		time.Sleep(500 * time.Millisecond)

		log.Print("Frame 2")
		check(client.StartDrawing(context.Background()), "Error getting stream")
		check(client.Clear(0, 1), "Error clearing display")
		check(client.SetPixels(
			[]sdk.Pixel{
				sdk.Pixel{
					Point: image.Point{1, 1},
					Color: color.RGBA{255, 60, 60, 255},
				},
				sdk.Pixel{
					Point: image.Point{0, 0},
					Color: color.RGBA{60, 255, 255, 255},
				},
			}, 0, sdk.PaintMode), "Error setting pixels")
		check(client.Render(), "Error rendering")

		time.Sleep(500 * time.Millisecond)
	}
}
