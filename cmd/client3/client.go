package main

import (
	"context"
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
	conn, err := grpc.Dial("telecom-tower.sofr.hefr.lan:10000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error dialing server: %v", err)
	}
	defer conn.Close() // nolint: errcheck
	client := sdk.NewClient(conn)

	check(client.StartDrawing(context.Background()), "Error getting frame")
	check(client.Clear(0, 1), "Error clearing display: %v")
	check(client.WriteText(
		"Hello Blue Masters * ", "8x8", 0,
		color.RGBA{128, 128, 255, 255}, 0, sdk.PaintMode),
		"Error writing text")
	check(client.AutoRoll(0, sdk.RollingStart, 0, 0), "Error setting autoroll")
	check(client.Render(), "Error rendering")
	time.Sleep(2 * time.Second)

	check(client.StartDrawing(context.Background()), "Error getting frame")
	check(client.Clear(0, 1), "Error clearing display: %v")
	check(client.WriteText(
		"Hello GDG Fribourg * ", "6x8", 0,
		color.RGBA{255, 128, 128, 255}, 0, sdk.PaintMode),
		"Error writing text")
	check(client.AutoRoll(0, sdk.RollingNext, 0, 0), "Error setting autoroll")
	check(client.Render(), "Error rendering")
	time.Sleep(5 * time.Second)

	check(client.StartDrawing(context.Background()), "Error getting frame")
	check(client.Clear(0, 1), "Error clearing display: %v")
	check(client.WriteText(
		"Portes ouvertes 2018 * ", "8x8", 0,
		color.RGBA{128, 128, 255, 255}, 0, sdk.PaintMode),
		"Error writing text")
	check(client.AutoRoll(0, sdk.RollingNext, 0, 0), "Error setting autoroll")
	check(client.Render(), "Error rendering")
	time.Sleep(2 * time.Second)

}
