package internal

import (
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
