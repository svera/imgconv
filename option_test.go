package imgconv

import (
	"image"
	"io/ioutil"
	"testing"
)

func TestOption(t *testing.T) {
	mark := &image.NRGBA{
		Rect:   image.Rect(0, 0, 4, 4),
		Stride: 4 * 4,
		Pix: []uint8{
			0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x40, 0xff, 0x00, 0x00, 0xbf, 0xff, 0x00, 0x00, 0xff,
			0x00, 0xff, 0x00, 0x40, 0x6e, 0x6d, 0x25, 0x70, 0xb0, 0x14, 0x3b, 0xcf, 0xbf, 0x00, 0x40, 0xff,
			0x00, 0xff, 0x00, 0xbf, 0x14, 0xb0, 0x3b, 0xcf, 0x33, 0x33, 0x99, 0xef, 0x40, 0x00, 0xbf, 0xff,
			0x00, 0xff, 0x00, 0xff, 0x00, 0xbf, 0x40, 0xff, 0x00, 0x40, 0xbf, 0xff, 0x00, 0x00, 0xff, 0xff,
		},
	}

	o := New()
	if o.Format.Format != JPEG {
		t.Error("Format is not expect one.")
	}
	o.SetWatermark(mark, 100)
	if mark != o.Watermark.Mark || o.Watermark.Opacity != 100 {
		t.Error("SetWatermark result is not expect one.")
	}
	o.SetResize(0, 0, 33)
	if o.Resize.Width != 0 || o.Resize.Height != 0 || o.Resize.Percent != 33 {
		t.Error("SetResize result is not expect one.")
	}
	base, err := Open("testdata/video-001.png")
	if err != nil {
		t.Error(err)
		return
	}
	if err := o.Convert(base, ioutil.Discard); err != nil {
		t.Error("Failed to Convert.")
	}
}

func TestConvertExt(t *testing.T) {
	o := New()
	if err := o.SetFormat("tif"); err != nil {
		t.Error("Failed to SetFormat.")
	}
	if o.ConvertExt("testdata/video-001.png") != "testdata/video-001.tif" {
		t.Error("ConvertExt result is not expect one.")
	}
}
