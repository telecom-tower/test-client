//go:generate protoc -I $GOPATH/src/github.com/telecom-tower/towerapi/v1 telecomtower.proto --go_out=plugins=grpc:$GOPATH/src/github.com/telecom-tower/towerapi/v1

package main

import (
	"context"
	"image"
	"image/color"
	"log"
	"time"

	"github.com/telecom-tower/sdk"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:10000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error dialing server: %v", err)
	}
	defer conn.Close()
	client := sdk.NewClient(conn)

	for {
		if err := client.StartDrawing(context.Background()); err != nil {
			log.Fatalf("Error getting stream: %v", err)
		}
		if err := client.Clear(0, 1); err != nil {
			log.Fatalf("Error clearing display: %v", err)
		}
		if err := client.SetPixels(
			[]sdk.Pixel{
				sdk.Pixel{
					Point: image.Point{0, 0},
					Color: color.RGBA{255, 60, 60, 255},
				},
				sdk.Pixel{
					Point: image.Point{1, 1},
					Color: color.RGBA{60, 255, 255, 255},
				},
			}, 0, sdk.PaintMode); err != nil {
			log.Fatalf("Error setting pixels: %v", err)
		}
		if err := client.DrawRectangle(
			image.Rect(2, 0, 6, 6),
			color.RGBA{255, 255, 255, 255},
			0, sdk.PaintMode); err != nil {
			log.Fatalf("Error setting pixels: %v", err)
		}
		if err := client.DrawRectangle(
			image.Rect(4, 2, 8, 8),
			color.RGBA{0, 255, 0, 128},
			1, sdk.PaintMode); err != nil {
			log.Fatalf("Error setting pixels: %v", err)
		}
		if err := client.WriteText(
			"Hello nice little wonderful world", "8x8", 10,
			color.RGBA{255, 255, 255, 255}, 0, sdk.PaintMode); err != nil {
			log.Fatalf("Error writing text: %v", err)
		}
		if err := client.WriteText(
			"World", "6x8", 20,
			color.RGBA{32, 32, 255, 180}, 0, sdk.OverMode); err != nil {
			log.Fatalf("Error writing text: %v", err)
		}
		if err := client.Render(); err != nil {
			log.Fatalf("Error rendering: %v", err)
		}

		time.Sleep(500 * time.Millisecond)

		if err := client.StartDrawing(context.Background()); err != nil {
			log.Fatalf("Error getting stream: %v", err)
		}
		if err := client.Clear(0, 1); err != nil {
			log.Fatalf("Error clearing display: %v", err)
		}
		if err := client.SetPixels([]sdk.Pixel{
			sdk.Pixel{
				Point: image.Point{1, 1},
				Color: color.RGBA{255, 60, 60, 255},
			},
			sdk.Pixel{
				Point: image.Point{0, 0},
				Color: color.RGBA{60, 255, 255, 255},
			},
		}, 0, sdk.PaintMode); err != nil {
			log.Fatalf("Error setting pixels: %v", err)
		}
		if err := client.Render(); err != nil {
			log.Fatalf("Error rendering: %v", err)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
