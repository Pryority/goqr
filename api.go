package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	qrcode "github.com/skip2/go-qrcode"
)

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {
	app := fiber.New()

	app.Get("/:chain/:txType/:value/:data", s.handleRouteParams)

	log.Println("JSON API server running on port:", s.listenAddr)

	app.Listen(s.listenAddr)
}

func (s *APIServer) handleRouteParams(c *fiber.Ctx) error {
	chain := c.Params("chain")
	txType := c.Params("txType")
	value := c.Params("value")
	data := c.Params("data")

	url := "http://localhost:3000" + c.OriginalURL()

	qr, err := generateQRCode(url)
	if err != nil {
		return err
	}

	// Save the QR code as a PNG image file
	err = saveQRCodeImage(qr, "qrcode.png")
	if err != nil {
		return err
	}

	response := fmt.Sprintf("Chain: %s\nTX Type: %s\nValue: %s\nData: %s\n\nQR Code:\n%s", chain, txType, value, data, qr)
	fmt.Println(c.SendString(response))
	return c.SendString(response)
}

func generateQRCode(content string) ([]byte, error) {
	qr, err := qrcode.New(content, qrcode.Low)
	if err != nil {
		return nil, err
	}
	return qr.PNG(256)
}

func saveQRCodeImage(qrCodeData []byte, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(qrCodeData)
	return err
}

type ApiError struct {
	Error string
}
