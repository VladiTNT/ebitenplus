package ebitenplus

import (
	"io"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// AssetLoader makes loading assets easier by centralizing error handling logic for io and decoding
// operations, which tend to require a lot of boilerplate.
//
// AssetLoader is not an asset manager. It doesn't cache assets, users should implement and handle their
// own caching logic if they wish to do that.
type AssetLoader struct {
	FileSystem fs.FS
	ErrFunc    func(error)
}

// NewAssetLoader is a constructor. It makes an AssetLoader with the given fs and error handling
// function.
func NewAssetLoader(fileSystem fs.FS, errFunc func(error)) *AssetLoader {
	return &AssetLoader{
		FileSystem: fileSystem,
		ErrFunc:    errFunc,
	}
}

// LoadImage opens the given file, decodes the image using DecodeImage and returns the ebiten.Image,
// if decoding or io fails, the function returns nil.
func (al *AssetLoader) LoadImage(name string) *ebiten.Image {
	f, err := al.FileSystem.Open(name)
	if err != nil {
		al.ErrFunc(err)
		return nil
	}
	defer f.Close()

	img, err := DecodeImage(f)
	if err != nil {
		al.ErrFunc(err)
		return nil
	}

	return ebiten.NewImageFromImage(img)
}

// LoadStream decodes the audio stream from the given file with the given encoding. It also returns
// an io.Closer so that you can close the file.
func (al *AssetLoader) LoadStream(name string, enc AudioEncoding) (AudioStream, io.Closer) {
	f, err := al.FileSystem.Open(name)
	if err != nil {
		al.ErrFunc(err)
		return nil, nil
	}

	stream, err := DecodeAudioStream(f, enc)
	if err != nil {
		al.ErrFunc(err)
		return nil, nil
	}

	return stream, f
}

// LoadRawAudio decodes the audio bytes from the given file with the given encoding.
func (al *AssetLoader) LoadRawAudio(name string, enc AudioEncoding) RawAudio {
	f, err := al.FileSystem.Open(name)
	if err != nil {
		al.ErrFunc(err)
		return nil
	}
	defer f.Close()

	stream, err := DecodeAudioStream(f, enc)
	if err != nil {
		al.ErrFunc(err)
		return nil
	}

	audioBytes, err := DecodeRawAudio(stream)
	if err != nil {
		al.ErrFunc(err)
		return nil
	}

	return audioBytes
}

// LoadFont loads a raw font file and returns a *text.GoTextFaceSource to use with text printing tools.
func (al *AssetLoader) LoadFont(name string) *text.GoTextFaceSource {
	f, err := al.FileSystem.Open(name)
	if err != nil {
		al.ErrFunc(err)
		return nil
	}
	defer f.Close()

	src, err := text.NewGoTextFaceSource(f)
	if err != nil {
		al.ErrFunc(err)
		return nil
	}

	return src
}

// LoadSpriteSheet loads a sprite sheet and returns it as an array of the sprites it contains, see more info with
// the ParseSpriteSheet function.
func (al *AssetLoader) LoadSpriteSheet(name string, spN, spLen, spHeight, sheetWidth int) []*ebiten.Image {
	return ParseSpriteSheet(al.LoadImage(name), spN, spLen, spHeight, sheetWidth)
}
