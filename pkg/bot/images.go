package bot

import (
	"image"

	"github.com/ev3go/ev3dev/fb"
)

var gopherSquint = &fb.Monochrome{
	Pix: []uint8{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xe0, 0xff, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0, 0xff, 0xff, 0xff, 0x7f, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0xe0, 0xff, 0x01, 0x00, 0x00, 0xfc, 0x3f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x01, 0x00, 0x00, 0x00, 0x00, 0xfc, 0x03,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf8,
		0x07, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0,
		0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xe0, 0x03,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x0f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x3c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xe0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0xe0, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x80, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x78, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x1e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x3c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x07, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x70, 0x00, 0xfe, 0x1f, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0xc0, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0xc0, 0xc1, 0x07, 0x78, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x70, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xfe, 0x3f, 0x00, 0x80, 0xf3, 0x00, 0xc0, 0x01, 0x00, 0x00,
		0x00, 0x80, 0xff, 0x01, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc0, 0x1f, 0xfc,
		0x01, 0x00, 0x3e, 0x00, 0x80, 0x03, 0x00, 0x00, 0x00, 0xf0, 0xe7, 0x0f, 0x0c, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0xf0, 0x00, 0x80, 0x07, 0x00, 0x1c, 0x00, 0x00, 0x06, 0x00, 0x00,
		0x00, 0x3c, 0x00, 0x3c, 0x07, 0x00, 0xf8, 0xff, 0x01, 0x00, 0x00, 0x00, 0x00, 0x1c, 0x00, 0x00,
		0x1e, 0x00, 0x18, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x0e, 0x00, 0xf0, 0x03, 0x80, 0x1f, 0xc0,
		0x0f, 0x00, 0x00, 0x00, 0x00, 0x07, 0x00, 0x00, 0x38, 0x00, 0x30, 0x00, 0x00, 0x18, 0x00, 0x00,
		0x80, 0x03, 0x00, 0xe0, 0x00, 0xe0, 0x01, 0x00, 0x78, 0x00, 0x00, 0x00, 0xc0, 0x01, 0x00, 0x00,
		0xe0, 0x00, 0x60, 0x00, 0x00, 0x30, 0x00, 0x00, 0xc0, 0x01, 0x00, 0x60, 0x00, 0x38, 0x00, 0x00,
		0xe0, 0x00, 0x00, 0x00, 0xe0, 0x00, 0x00, 0x00, 0xc0, 0x01, 0xc0, 0x00, 0x00, 0x60, 0x00, 0x00,
		0x60, 0x00, 0x00, 0x30, 0x00, 0x0e, 0x00, 0x00, 0x80, 0x03, 0x00, 0x00, 0xf0, 0xff, 0xff, 0xff,
		0xff, 0x03, 0x80, 0x01, 0x00, 0x60, 0x00, 0x00, 0x30, 0x00, 0x00, 0x18, 0x00, 0x03, 0x00, 0x00,
		0x00, 0x06, 0x00, 0x00, 0xf8, 0xff, 0xff, 0xff, 0xff, 0x07, 0x00, 0x03, 0x00, 0xc0, 0x00, 0x00,
		0x30, 0x00, 0x00, 0x0c, 0x80, 0xff, 0xff, 0xff, 0xff, 0x0f, 0x00, 0x00, 0x8c, 0xff, 0x03, 0x00,
		0x00, 0x06, 0x00, 0xfe, 0x00, 0xc0, 0x00, 0x00, 0x18, 0x00, 0x00, 0x06, 0xc0, 0xff, 0xff, 0xff,
		0xff, 0x1f, 0x00, 0x00, 0xc6, 0xff, 0x07, 0x00, 0x00, 0x0c, 0x00, 0xfc, 0x01, 0x80, 0x01, 0x00,
		0x08, 0x00, 0x00, 0x03, 0x60, 0xf0, 0x3f, 0x00, 0x00, 0x30, 0x00, 0x00, 0xc6, 0xff, 0x07, 0x00,
		0x00, 0x18, 0x00, 0xfc, 0x03, 0x80, 0x01, 0x00, 0x0c, 0x00, 0x80, 0x03, 0x30, 0xf8, 0x7f, 0x00,
		0x00, 0x20, 0x00, 0x00, 0xc3, 0xff, 0x0f, 0x00, 0x00, 0x18, 0x00, 0xf8, 0x07, 0x00, 0x01, 0x00,
		0x04, 0xc0, 0x9f, 0x01, 0x18, 0xf8, 0xff, 0x00, 0x00, 0x60, 0x00, 0x80, 0xe1, 0xff, 0x0f, 0x00,
		0x00, 0x30, 0x00, 0xf0, 0x07, 0x00, 0x01, 0x00, 0x06, 0xe0, 0xff, 0x00, 0x08, 0xfc, 0xff, 0x00,
		0x00, 0xc0, 0x00, 0x80, 0xe1, 0xff, 0x0f, 0x00, 0x00, 0x30, 0x00, 0xf0, 0x07, 0x00, 0x03, 0x00,
		0x06, 0xf0, 0x7f, 0x00, 0x0c, 0xfc, 0xff, 0x00, 0x00, 0xc0, 0x00, 0xc0, 0xc0, 0xff, 0x0f, 0x00,
		0x00, 0x20, 0x00, 0xe0, 0x0f, 0x00, 0x03, 0x00, 0x02, 0xf0, 0x7f, 0x00, 0x04, 0xfc, 0xff, 0x01,
		0x00, 0x80, 0x00, 0xc0, 0xc0, 0xcf, 0x0f, 0x00, 0x00, 0x60, 0x00, 0xc0, 0x07, 0x00, 0x03, 0x00,
		0x02, 0xf8, 0x3f, 0x00, 0x06, 0xfc, 0xff, 0x01, 0x00, 0x80, 0x01, 0x40, 0xc0, 0xc7, 0x07, 0x00,
		0x00, 0x60, 0x00, 0xc0, 0x07, 0x00, 0x03, 0x00, 0x03, 0xf8, 0x1f, 0x00, 0x06, 0xfc, 0xff, 0x00,
		0x00, 0x00, 0x01, 0x60, 0x80, 0xc7, 0x07, 0x00, 0x00, 0x60, 0x00, 0x80, 0x03, 0x00, 0x03, 0x00,
		0x03, 0xf8, 0x1f, 0x00, 0x02, 0xfc, 0xfd, 0x00, 0x00, 0x00, 0x01, 0x60, 0x80, 0xcf, 0x03, 0x00,
		0x00, 0x40, 0x00, 0x80, 0x01, 0x00, 0x03, 0x00, 0x03, 0xf8, 0x0f, 0x00, 0x03, 0xf8, 0xf8, 0x00,
		0x00, 0x00, 0x03, 0x20, 0x00, 0xff, 0x01, 0x00, 0x00, 0x40, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x03, 0xf0, 0x0f, 0x00, 0x03, 0xf8, 0x78, 0x00, 0x00, 0x00, 0x03, 0x20, 0x00, 0xfc, 0x00, 0x00,
		0x00, 0x40, 0x00, 0x00, 0x03, 0x80, 0x01, 0x00, 0x02, 0xe0, 0x07, 0x00, 0x01, 0xf0, 0x38, 0x00,
		0x00, 0x00, 0x02, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc0, 0x00, 0x00, 0x03, 0x80, 0x01, 0x00,
		0x02, 0xc0, 0x07, 0x00, 0x01, 0xe0, 0x1f, 0x00, 0x00, 0x00, 0x02, 0x20, 0x00, 0x00, 0x00, 0x00,
		0x00, 0xc0, 0x00, 0x00, 0x06, 0xc0, 0x00, 0x00, 0x06, 0x80, 0x03, 0x00, 0x01, 0x80, 0x0f, 0x00,
		0x00, 0x00, 0x02, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc0, 0x00, 0x00, 0x06, 0xc0, 0x00, 0x00,
		0x06, 0x00, 0x03, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x20, 0x00, 0x00, 0x00, 0x00,
		0x00, 0xc0, 0x00, 0x00, 0x04, 0x60, 0x00, 0x00, 0x0c, 0x00, 0x03, 0x00, 0x01, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x02, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x0c, 0x30, 0x00, 0x00,
		0x0c, 0x80, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x20, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x40, 0x00, 0x00, 0x0c, 0x38, 0x00, 0x00, 0x18, 0x80, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x02, 0x60, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x08, 0x1c, 0x00, 0x00,
		0x30, 0x80, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x60, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x60, 0x00, 0x00, 0x18, 0x07, 0x00, 0x00, 0x60, 0xc0, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x03, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x60, 0x00, 0x00, 0xd8, 0x03, 0x00, 0x00,
		0xc0, 0xc0, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xc0, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x20, 0x00, 0x00, 0xf8, 0x00, 0x00, 0x00, 0x80, 0x43, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x01, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x30, 0x00, 0x00, 0x30, 0x00, 0x00, 0x00,
		0x00, 0x6f, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x80, 0x01, 0x80, 0x01, 0x00, 0x00, 0x00,
		0x00, 0x10, 0x00, 0x00, 0x30, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00,
		0x00, 0x80, 0x01, 0x80, 0x01, 0x00, 0x00, 0x00, 0x00, 0x18, 0x00, 0x00, 0x30, 0x00, 0x00, 0x00,
		0x00, 0x70, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0xc0, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00,
		0x00, 0x08, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00,
		0x00, 0xc0, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00,
		0x00, 0x30, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x60, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00,
		0x00, 0x06, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00, 0x00, 0x30, 0x00, 0x00, 0x18, 0x00, 0x00, 0x00,
		0x00, 0x20, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x60, 0x00, 0x00, 0x00,
		0x00, 0x30, 0x00, 0x00, 0x30, 0x00, 0x00, 0x00, 0x00, 0x30, 0x80, 0x3f, 0x18, 0x00, 0x00, 0x00,
		0x80, 0x01, 0x00, 0x00, 0x60, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00,
		0x00, 0x18, 0xf0, 0xff, 0x71, 0x00, 0x00, 0x00, 0xe0, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00,
		0x00, 0x10, 0x00, 0x00, 0x60, 0x00, 0x00, 0x00, 0x00, 0x0c, 0xfc, 0xff, 0xe3, 0x00, 0x00, 0x00,
		0x30, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0xc0, 0x01, 0x00, 0x00,
		0x00, 0x07, 0xfe, 0xff, 0x87, 0x03, 0x00, 0x00, 0x1c, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00,
		0x00, 0x18, 0x00, 0x00, 0x80, 0x03, 0x00, 0x00, 0x80, 0x03, 0xff, 0xff, 0x0f, 0x0f, 0x00, 0x00,
		0x07, 0x00, 0x00, 0x00, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x18, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00,
		0xe0, 0x00, 0xff, 0xff, 0x1f, 0x7c, 0x00, 0xf0, 0x01, 0x00, 0x00, 0x00, 0xc0, 0x00, 0x00, 0x00,
		0x00, 0x18, 0x00, 0x00, 0x00, 0x1c, 0x00, 0x00, 0x78, 0x80, 0xff, 0xff, 0x1f, 0xe0, 0xff, 0x3f,
		0x00, 0x00, 0x00, 0x00, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x18, 0x00, 0x00, 0x00, 0xf0, 0x00, 0x00,
		0x1f, 0x80, 0xff, 0xff, 0x1f, 0x00, 0xfc, 0x01, 0x00, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00,
		0x00, 0x08, 0x00, 0x00, 0x00, 0xc0, 0x1f, 0xf8, 0x03, 0x80, 0xff, 0xff, 0x1f, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0xfc, 0x3f,
		0x00, 0xf0, 0xff, 0xff, 0x7f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00,
		0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xb8, 0xff, 0xff, 0xe7, 0x01, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x80, 0x01, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x0c, 0xff, 0xff, 0x83, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x01, 0x00, 0x00,
		0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x06, 0xfc, 0xff, 0x00, 0x07, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x80, 0x01, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x03, 0xf0, 0x1f, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x01, 0x00, 0x00,
		0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x01, 0x00, 0x00, 0x00, 0x18, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x80, 0x01, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x80, 0x01, 0x00, 0x00, 0x00, 0x30, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
		0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00, 0x00, 0x30, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0xc0, 0x00, 0x00, 0x00, 0x00, 0x60, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
		0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x60, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0xc0, 0x00, 0x00, 0x00, 0x00, 0x60, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00,
		0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x60, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x80, 0x01, 0x00, 0x0e, 0x00, 0x30, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00,
		0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x01, 0xe0, 0xff, 0x00, 0x30, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x03, 0x7c, 0xc7, 0x07, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00,
		0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xfe, 0x0f, 0x06, 0x1e, 0x0e, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0xfc, 0x01, 0x06, 0xf8, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00,
		0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x60, 0x00, 0x03, 0xd8, 0x01, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x60, 0x00, 0x03, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00,
		0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x20, 0x00, 0x03, 0x18, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x20, 0x00, 0x03, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00,
		0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x20, 0x00, 0x03, 0x18, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x60, 0x00, 0x03, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00,
		0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x60, 0x00, 0x03, 0x18, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x60, 0x00, 0x03, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00,
		0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, 0x00, 0x03, 0x18, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0xc0, 0x00, 0x03, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00,
		0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc0, 0x00, 0x07, 0x18, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x80, 0xc1, 0x0f, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00,
		0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x1c, 0x0e, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x3e, 0xf8, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00,
		0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xe0, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00,
	},
	Stride: 24,
	Rect: image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: 178,
			Y: 128,
		},
	},
}

var gopher = &fb.Monochrome{
	Pix: []uint8{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xe0, 0xff, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0, 0xff, 0xff, 0xff, 0x7f, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0xe0, 0xff, 0x01, 0x00, 0x00, 0xfc, 0x3f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x01, 0x00, 0x00, 0x00, 0x00, 0xfc, 0x03,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf8,
		0x07, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0,
		0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xe0, 0x03,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x0f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x3c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xe0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0xe0, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x80, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x78, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x1e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x3c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x07, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x70, 0x00, 0xfe, 0x1f, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0xc0, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0xc0, 0xc1, 0x07, 0x78, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x70, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xfe, 0x3f, 0x00, 0x80, 0xf3, 0x00, 0xc0, 0x01, 0x00, 0x00,
		0x00, 0x80, 0xff, 0x01, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc0, 0x1f, 0xfc,
		0x01, 0x00, 0x3e, 0x00, 0x80, 0x03, 0x00, 0x00, 0x00, 0xf0, 0xe7, 0x0f, 0x0c, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0xf0, 0x00, 0x80, 0x07, 0x00, 0x1c, 0x00, 0x00, 0x06, 0x00, 0x00,
		0x00, 0x3c, 0x00, 0x3c, 0x07, 0x00, 0xf8, 0xff, 0x01, 0x00, 0x00, 0x00, 0x00, 0x1c, 0x00, 0x00,
		0x1e, 0x00, 0x18, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x0e, 0x00, 0xf0, 0x03, 0x80, 0x1f, 0xc0,
		0x0f, 0x00, 0x00, 0x00, 0x00, 0x07, 0x00, 0x00, 0x38, 0x00, 0x30, 0x00, 0x00, 0x18, 0x00, 0x00,
		0x80, 0x03, 0x00, 0xe0, 0x00, 0xe0, 0x01, 0x00, 0x78, 0x00, 0x00, 0x00, 0xc0, 0x01, 0x00, 0x00,
		0xe0, 0x00, 0x60, 0x00, 0x00, 0x30, 0x00, 0x00, 0xc0, 0x01, 0x00, 0x60, 0x00, 0x38, 0x00, 0x00,
		0xe0, 0x00, 0x00, 0x00, 0xe0, 0x00, 0x00, 0x00, 0xc0, 0x01, 0xc0, 0x00, 0x00, 0x60, 0x00, 0x00,
		0x60, 0x00, 0x00, 0x30, 0x00, 0x0e, 0x00, 0x00, 0x80, 0x03, 0x00, 0x00, 0x30, 0x00, 0xf8, 0x03,
		0x00, 0x03, 0x80, 0x01, 0x00, 0x60, 0x00, 0x00, 0x30, 0x00, 0x00, 0x18, 0x00, 0x03, 0x00, 0x00,
		0x00, 0x06, 0x00, 0x00, 0x18, 0x00, 0xfc, 0x0f, 0x00, 0x06, 0x00, 0x03, 0x00, 0xc0, 0x00, 0x00,
		0x30, 0x00, 0x00, 0x0c, 0x80, 0x01, 0x00, 0x1c, 0x00, 0x0c, 0x00, 0x00, 0x0c, 0x00, 0xfe, 0x0f,
		0x00, 0x06, 0x00, 0xfe, 0x00, 0xc0, 0x00, 0x00, 0x18, 0x00, 0x00, 0x06, 0xc0, 0x00, 0x80, 0x7f,
		0x00, 0x18, 0x00, 0x00, 0x06, 0x00, 0xff, 0x1f, 0x00, 0x0c, 0x00, 0xfc, 0x01, 0x80, 0x01, 0x00,
		0x08, 0x00, 0x00, 0x03, 0x60, 0x00, 0xc0, 0xff, 0x00, 0x30, 0x00, 0x00, 0x06, 0x00, 0xff, 0x1f,
		0x00, 0x18, 0x00, 0xfc, 0x03, 0x80, 0x01, 0x00, 0x0c, 0x00, 0x80, 0x03, 0x30, 0x00, 0xe0, 0xff,
		0x01, 0x20, 0x00, 0x00, 0x03, 0x00, 0xff, 0x3f, 0x00, 0x18, 0x00, 0xf8, 0x07, 0x00, 0x01, 0x00,
		0x04, 0xc0, 0x9f, 0x01, 0x18, 0x00, 0xe0, 0xff, 0x03, 0x60, 0x00, 0x80, 0x01, 0x80, 0xff, 0x3f,
		0x00, 0x30, 0x00, 0xf0, 0x07, 0x00, 0x01, 0x00, 0x06, 0xe0, 0xff, 0x00, 0x08, 0x00, 0xf0, 0xff,
		0x03, 0xc0, 0x00, 0x80, 0x01, 0x80, 0xff, 0x3f, 0x00, 0x30, 0x00, 0xf0, 0x07, 0x00, 0x03, 0x00,
		0x06, 0xf0, 0x7f, 0x00, 0x0c, 0x00, 0xf0, 0xff, 0x03, 0xc0, 0x00, 0xc0, 0x00, 0x00, 0xff, 0x3f,
		0x00, 0x20, 0x00, 0xe0, 0x0f, 0x00, 0x03, 0x00, 0x02, 0xf0, 0x7f, 0x00, 0x04, 0x00, 0xf0, 0xff,
		0x07, 0x80, 0x00, 0xc0, 0x00, 0x00, 0x3f, 0x3f, 0x00, 0x60, 0x00, 0xc0, 0x07, 0x00, 0x03, 0x00,
		0x02, 0xf8, 0x3f, 0x00, 0x06, 0x00, 0xf0, 0xff, 0x07, 0x80, 0x01, 0x40, 0x00, 0x00, 0x1f, 0x1f,
		0x00, 0x60, 0x00, 0xc0, 0x07, 0x00, 0x03, 0x00, 0x03, 0xf8, 0x1f, 0x00, 0x06, 0x00, 0xf0, 0xff,
		0x03, 0x00, 0x01, 0x60, 0x00, 0x00, 0x1e, 0x1f, 0x00, 0x60, 0x00, 0x80, 0x03, 0x00, 0x03, 0x00,
		0x03, 0xf8, 0x1f, 0x00, 0x02, 0x00, 0xf0, 0xf7, 0x03, 0x00, 0x01, 0x60, 0x00, 0x00, 0x3e, 0x0f,
		0x00, 0x40, 0x00, 0x80, 0x01, 0x00, 0x03, 0x00, 0x03, 0xf8, 0x0f, 0x00, 0x03, 0x00, 0xe0, 0xe3,
		0x03, 0x00, 0x03, 0x20, 0x00, 0x00, 0xfc, 0x07, 0x00, 0x40, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00,
		0x03, 0xf0, 0x0f, 0x00, 0x03, 0x00, 0xe0, 0xe3, 0x01, 0x00, 0x03, 0x20, 0x00, 0x00, 0xf0, 0x03,
		0x00, 0x40, 0x00, 0x00, 0x03, 0x80, 0x01, 0x00, 0x02, 0xe0, 0x07, 0x00, 0x01, 0x00, 0xc0, 0xe3,
		0x00, 0x00, 0x02, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc0, 0x00, 0x00, 0x03, 0x80, 0x01, 0x00,
		0x02, 0xc0, 0x07, 0x00, 0x01, 0x00, 0x80, 0x7f, 0x00, 0x00, 0x02, 0x20, 0x00, 0x00, 0x00, 0x00,
		0x00, 0xc0, 0x00, 0x00, 0x06, 0xc0, 0x00, 0x00, 0x06, 0x80, 0x03, 0x00, 0x01, 0x00, 0x00, 0x3e,
		0x00, 0x00, 0x02, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc0, 0x00, 0x00, 0x06, 0xc0, 0x00, 0x00,
		0x06, 0x00, 0x03, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x20, 0x00, 0x00, 0x00, 0x00,
		0x00, 0xc0, 0x00, 0x00, 0x04, 0x60, 0x00, 0x00, 0x0c, 0x00, 0x03, 0x00, 0x01, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x02, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x0c, 0x30, 0x00, 0x00,
		0x0c, 0x80, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x20, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x40, 0x00, 0x00, 0x0c, 0x38, 0x00, 0x00, 0x18, 0x80, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x02, 0x60, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x08, 0x1c, 0x00, 0x00,
		0x30, 0x80, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x60, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x60, 0x00, 0x00, 0x18, 0x07, 0x00, 0x00, 0x60, 0xc0, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x03, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x60, 0x00, 0x00, 0xd8, 0x03, 0x00, 0x00,
		0xc0, 0xc0, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xc0, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x20, 0x00, 0x00, 0xf8, 0x00, 0x00, 0x00, 0x80, 0x43, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x01, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x30, 0x00, 0x00, 0x30, 0x00, 0x00, 0x00,
		0x00, 0x6f, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x80, 0x01, 0x80, 0x01, 0x00, 0x00, 0x00,
		0x00, 0x10, 0x00, 0x00, 0x30, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00,
		0x00, 0x80, 0x01, 0x80, 0x01, 0x00, 0x00, 0x00, 0x00, 0x18, 0x00, 0x00, 0x30, 0x00, 0x00, 0x00,
		0x00, 0x70, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0xc0, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00,
		0x00, 0x08, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00,
		0x00, 0xc0, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00,
		0x00, 0x30, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x60, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00,
		0x00, 0x06, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00, 0x00, 0x30, 0x00, 0x00, 0x18, 0x00, 0x00, 0x00,
		0x00, 0x20, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x60, 0x00, 0x00, 0x00,
		0x00, 0x30, 0x00, 0x00, 0x30, 0x00, 0x00, 0x00, 0x00, 0x30, 0x80, 0x3f, 0x18, 0x00, 0x00, 0x00,
		0x80, 0x01, 0x00, 0x00, 0x60, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00,
		0x00, 0x18, 0xf0, 0xff, 0x71, 0x00, 0x00, 0x00, 0xe0, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00,
		0x00, 0x10, 0x00, 0x00, 0x60, 0x00, 0x00, 0x00, 0x00, 0x0c, 0xfc, 0xff, 0xe3, 0x00, 0x00, 0x00,
		0x30, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0xc0, 0x01, 0x00, 0x00,
		0x00, 0x07, 0xfe, 0xff, 0x87, 0x03, 0x00, 0x00, 0x1c, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00,
		0x00, 0x18, 0x00, 0x00, 0x80, 0x03, 0x00, 0x00, 0x80, 0x03, 0xff, 0xff, 0x0f, 0x0f, 0x00, 0x00,
		0x07, 0x00, 0x00, 0x00, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x18, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00,
		0xe0, 0x00, 0xff, 0xff, 0x1f, 0x7c, 0x00, 0xf0, 0x01, 0x00, 0x00, 0x00, 0xc0, 0x00, 0x00, 0x00,
		0x00, 0x18, 0x00, 0x00, 0x00, 0x1c, 0x00, 0x00, 0x78, 0x80, 0xff, 0xff, 0x1f, 0xe0, 0xff, 0x3f,
		0x00, 0x00, 0x00, 0x00, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x18, 0x00, 0x00, 0x00, 0xf0, 0x00, 0x00,
		0x1f, 0x80, 0xff, 0xff, 0x1f, 0x00, 0xfc, 0x01, 0x00, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00,
		0x00, 0x08, 0x00, 0x00, 0x00, 0xc0, 0x1f, 0xf8, 0x03, 0x80, 0xff, 0xff, 0x1f, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0xfc, 0x3f,
		0x00, 0xf0, 0xff, 0xff, 0x7f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00,
		0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xb8, 0xff, 0xff, 0xe7, 0x01, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x80, 0x01, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x0c, 0xff, 0xff, 0x83, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x01, 0x00, 0x00,
		0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x06, 0xfc, 0xff, 0x00, 0x07, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x80, 0x01, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x03, 0xf0, 0x1f, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x01, 0x00, 0x00,
		0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x01, 0x00, 0x00, 0x00, 0x18, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x80, 0x01, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x80, 0x01, 0x00, 0x00, 0x00, 0x30, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
		0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00, 0x00, 0x30, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0xc0, 0x00, 0x00, 0x00, 0x00, 0x60, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
		0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x60, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0xc0, 0x00, 0x00, 0x00, 0x00, 0x60, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00,
		0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x60, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x80, 0x01, 0x00, 0x0e, 0x00, 0x30, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00,
		0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x01, 0xe0, 0xff, 0x00, 0x30, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x03, 0x7c, 0xc7, 0x07, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00,
		0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xfe, 0x0f, 0x06, 0x1e, 0x0e, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0xfc, 0x01, 0x06, 0xf8, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00,
		0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x60, 0x00, 0x03, 0xd8, 0x01, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x60, 0x00, 0x03, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00,
		0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x20, 0x00, 0x03, 0x18, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x20, 0x00, 0x03, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00,
		0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x20, 0x00, 0x03, 0x18, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x60, 0x00, 0x03, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00,
		0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x60, 0x00, 0x03, 0x18, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x60, 0x00, 0x03, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00,
		0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, 0x00, 0x03, 0x18, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0xc0, 0x00, 0x03, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00,
		0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc0, 0x00, 0x07, 0x18, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x80, 0xc1, 0x0f, 0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00,
		0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x1c, 0x0e, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x3e, 0xf8, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00,
		0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xe0, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00,
	},
	Stride: 24,
	Rect: image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: 178,
			Y: 128,
		},
	},
}