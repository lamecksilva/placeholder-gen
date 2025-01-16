package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"net/http"
	"net/url"
	"os"
	"placeholder-gen/pkg/utils"
	"strconv"
	"time"

	"github.com/fogleman/gg"
)

type PlaceholderImageParams struct {
	BackgroundColor color.Color
	LabelColor      color.Color
	Label           string
	Width           int
	Heigth          int
}

func generateImage(params *PlaceholderImageParams, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "image/png")

	// Create rectangle image
	img := image.NewRGBA(image.Rect(0, 0, params.Width, params.Heigth))
	// Add background color

	log.Printf("%v\n", params.BackgroundColor)

	draw.Draw(img, img.Bounds(), &image.Uniform{params.BackgroundColor}, image.Point{}, draw.Src)

	if len(params.Label) != 0 {
		dc := gg.NewContextForRGBA(img)

		err := dc.LoadFontFace("web/static/fonts/Switzal.ttf", float64(params.Heigth)/10)
		if err != nil {
			log.Printf("Error loading font: %v", err)
			http.Error(w, "Failure to generate image", http.StatusInternalServerError)
		}

		dc.SetColor(params.LabelColor)

		dc.DrawStringAnchored(params.Label, float64(params.Width)/2, float64(params.Heigth)/2, 0.5, 0.5)

		dc.Stroke()
	}
	if err := png.Encode(w, img); err != nil {
		http.Error(w, "Failure to generate image", http.StatusInternalServerError)
		return
	}
}

func ValidateAndGetParams(query url.Values) (*PlaceholderImageParams, error) {
	params := PlaceholderImageParams{}

	queryColor := query.Get("color")
	queryWidth := query.Get("width")
	queryHeigth := query.Get("heigth")
	queryLabel := query.Get("label")
	queryLabelColor := query.Get("label-color")

	if len(queryWidth) == 0 || len(queryHeigth) == 0 {
		return nil, errors.New("Width or Heigth not provided")
	}

	width, err := strconv.ParseInt(queryWidth, 10, 32)
	if err != nil {
		return nil, errors.New("Error with width value")
	}

	params.Width = int(width)

	heigth, err := strconv.ParseInt(queryHeigth, 10, 32)
	if err != nil {
		return nil, errors.New("Error with heigth value")
	}
	params.Heigth = int(heigth)

	if len(queryColor) != 0 {
		log.Printf("%v\n", queryColor)

		col, err := utils.ParseHexColor(queryColor)
		if err != nil {
			return nil, errors.New("Error with color value")
		}
		params.BackgroundColor = col
	} else {
		params.BackgroundColor = color.White
	}

	if len(queryLabel) != 0 {
		params.Label = queryLabel
	}

	if len(queryLabelColor) != 0 {
		log.Printf("%v\n", queryLabelColor)

		col, err := utils.ParseHexColor(queryLabelColor)
		if err != nil {
			return nil, errors.New("Error with color value")
		}
		params.LabelColor = col
	} else {
		params.LabelColor = utils.GetContrastColor(utils.ConvertToRGBA(params.BackgroundColor))
	}

	return &params, nil
}

func GenerateImageHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	params, err := ValidateAndGetParams(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	generateImage(params, w)

	fmt.Print("Image generated with success")
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/generate", GenerateImageHandler)

	serverPort := os.Getenv("PORT")
	if len(serverPort) == 0 {
		serverPort = "8080"
	}

	log.Printf("Server running on port: %v", serverPort)
	if err := http.ListenAndServe(":"+serverPort, nil); err != nil {
		log.Fatalf("Error when starting server: %v", err)
	}
}
