package main

import (
	"context"
	"flag"
	"log"

	"github.com/pkg/errors"
	"github.com/telecom-tower/sdk"
	"golang.org/x/image/colornames"
	"google.golang.org/grpc"
)

func check(err error, msg string) {
	if err != nil {
		err = errors.WithMessage(err, msg)
		log.Fatal(err)
	}
}

func main() {
	url := flag.String("url", "telecom-tower2.sofr.hefr.lan:10000", "grpc URL")
	next := flag.Bool("next", false, "mark as next message")
	color := flag.String("color", "white", "color")

	flag.Parse()

	c, ok := colornames.Map[*color]
	if !ok {
		c = colornames.White
	}

	var msg string
	if flag.NArg() != 1 {
		msg = "Bienvenue aux portes ouvertes de la Haute école d'ingénierie et d'architecture de Fribourg * "
	} else {
		msg = flag.Arg(0) + " * "
	}

	conn, err := grpc.Dial(*url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error dialing server: %v", err)
	}
	defer conn.Close() // nolint: errcheck
	client := sdk.NewClient(conn)

	check(client.StartDrawing(context.Background()), "Error getting frame")
	check(client.Init(), "Error initializing display")
	check(client.WriteText(msg, "8x8", 0, c, 0, sdk.PaintMode), "Error writing text")
	if *next {
		check(client.AutoRoll(0, sdk.RollingNext, 0, 0), "Error setting autoroll")
	} else {
		check(client.AutoRoll(0, sdk.RollingStart, 0, 0), "Error setting autoroll")
	}
	check(client.Render(), "Error rendering")
}
