package internal

import (
	"net/http"
	"strconv"

	"github.com/skip2/go-qrcode"
)

type QRService struct {
}

func NewQRService() *QRService {
	return &QRService{}
}

func (s *QRService) Generate(text string, size int) ([]byte, error) {
	qr, err := qrcode.New(text, qrcode.Medium)
	if err != nil {
		return nil, err
	}

	png, err := qr.PNG(size)
	if err != nil {
		return nil, err
	}

	return png, nil
}

func (s *QRService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	if text == "" {
		http.Error(w, "missing 'text' parameter", http.StatusBadRequest)
		return
	}

	sizeText := r.URL.Query().Get("size")
	if sizeText == "" {
		sizeText = "256"
	}
	size, err := strconv.Atoi(sizeText)
	if err != nil {
		http.Error(w, "Invalid 'size' parameter", http.StatusBadRequest)
		return
	}

	qrCodeData, err := s.Generate(text, size)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(qrCodeData)
}
