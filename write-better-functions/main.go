package main

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/google/uuid"
	"github.com/nfnt/resize"
)

type PhotoConfig struct {
	MaxWidth      int
	MaxHeight     int
	Quality       int
	Format        string
	WatermarkImg  image.Image
	ThumbnailSize uint
}

type ImageResult struct {
	Full      []byte
	Thumbnail []byte
}

type PhotoService struct {
	MaxFileSize  int
	WatermarkImg image.Image
	cfg          *PhotoConfig
	storage      Storage
}

type File struct {
	Size int
	Name string
}

func (ps *PhotoService) DecodeImage(file *File) (image.Image, error) {
	return nil, nil
}

func (ps *PhotoService) Upload(result *ImageResult) error {
	return nil
}

// Tip 1
// Bad
func BadUploadPhoto(ps *PhotoService, file *File) error {
	if ps.IsSupportedFormat(file) {
		if file.Size <= ps.MaxFileSize {
			img, err := ps.DecodeImage(file)
			if err == nil {
				result, err := ps.PrepareForUpload(img)
				if err == nil {
					return ps.Upload(result)
				} else {
					return err
				}
			} else {
				return err
			}
		} else {
			return errors.New("file too large")
		}
	} else {
		return errors.New("unsupported format")
	}
}

func GoodUploadPhoto(ps *PhotoService, file *File) error {
	if !ps.IsSupportedFormat(file) {
		return errors.New("unsupported format")
	}
	if file.Size > ps.MaxFileSize {
		return errors.New("file too large")
	}
	img, err := ps.DecodeImage(file)
	if err != nil {
		return err
	}

	result, err := ps.PrepareForUpload(img)
	if err != nil {
		return err
	}

	return ps.Upload(result)
}

// Tip 2
//
// func (s *PhotoService) PrepareForUpload(img image.Image) (*ImageResult, error) {
// 	// Resize
// 	if img.Bounds().Dx() > 1920 {
// 		img = resize.Resize(1920, 0, img, resize.Lanczos3)
// 	}

// 	// Compress
// 	var buf bytes.Buffer
// 	jpeg.Encode(&buf, img, &jpeg.Options{Quality: 85})

// 	// Generate thumbnail
// 	thumb := resize.Resize(200, 0, img, resize.Lanczos3)
// 	buf.Reset()
// 	jpeg.Encode(&buf, thumb, &jpeg.Options{Quality: 80})
// 	thumData := buf.Bytes()

// 	// Add watermark
// 	watermarked := imaging.Overlay(img, s.WatermarkImg, image.Pt(10, 10), 0.5)
// 	buf.Reset()
// 	jpeg.Encode(&buf, watermarked, &jpeg.Options{Quality: 90})

// 	return &ImageResult{
// 		Full:      buf.Bytes(),
// 		Thumbnail: thumData,
// 	}, nil
// }

func (s *PhotoService) PrepareForUpload(img image.Image) (*ImageResult, error) {
	img = s.Resize(img)
	thumb := s.CreateThumbnail(img)
	watermarked := s.AddWatermark(img)

	return &ImageResult{
		Full:      s.Compress(watermarked),
		Thumbnail: s.Compress(thumb),
	}, nil
}

func (s *PhotoService) Resize(img image.Image) image.Image {
	if img.Bounds().Dx() <= s.cfg.MaxWidth {
		return img
	}
	return resize.Resize(uint(s.cfg.MaxWidth), 0, img, resize.Lanczos3)
}

func (s *PhotoService) CreateThumbnail(img image.Image) image.Image {
	return resize.Resize(s.cfg.ThumbnailSize, 0, img, resize.Lanczos3)
}

func (s *PhotoService) AddWatermark(img image.Image) image.Image {
	return imaging.Overlay(img, s.cfg.WatermarkImg, image.Pt(10, 10), 0.5)
}

func (s *PhotoService) Compress(img image.Image) []byte {
	var buf bytes.Buffer
	jpeg.Encode(&buf, img, &jpeg.Options{Quality: 85})
	return buf.Bytes()
}

// Tip 3
// updateImage function name is not good, ResizeToMaxWidth
func (s *PhotoService) ResizeToMaxWidth(img image.Image) image.Image {
	if img.Bounds().Dx() <= s.cfg.MaxWidth {
		return img
	}
	return resize.Resize(uint(s.cfg.MaxWidth), 0, img, resize.Lanczos3)
}

// checkFile function name is not good, unclear,
func (s *PhotoService) IsSupportedFormat(f *File) bool {
	ext := strings.ToLower(filepath.Ext(f.Name))
	return ext == ".jpg" || ext == ".png" || ext == ".gif"
}

// Tip 4
//
//	func (s *PhotoService) PrepareForUpload(img image.Image, maxWidth int, maxHeight int, quality int, format string, addWatermark bool, watermarkPath string, generateThumb bool, thumbWidth int) (*ImageResult, error) {
//		// Implementation
//	}
//
// func (s *PhotoService) PrepareForUpload(img image.Image, cfg PhotoConfig) (*ImageResult, error) {
// 	// Implementation uses cfg.MaxWidth, cfg.Quality, etc.
// }
// call function
// result,err := s.PrepareForUpload(img,PhotoConfig{
// 	MaxWidth:      1920,
// 	MaxHeight:     1080,
// 	Quality:       85,
// 	Format:        "jpg",
// 	WatermarkImg:  watermarkImg,
// 	ThumbnailSize: 200,
// })

// Tip 5
// var storage *S3Client

// func SaveImage(userId string, data []byte) (string, error) {
// 	path := fmt.Sprintf("images/%s/%s.jpg", userId, uuid.New().String())
// 	err := storage.Upload(path, data)
// 	return path, err
// }

type Storage interface {
	Upload(path string, data []byte) error
}

func (s *PhotoService) SavePhoto(userID string, data []byte) (string, error) {
	path := fmt.Sprintf("images/%s/%s.jpg", userID, uuid.New().String())
	err := s.storage.Upload(path, data)
	return path, err
}
