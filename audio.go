package ebitenplus

import (
	"io"

	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

// RawAudio is an alias for a byte slice.
type RawAudio []byte

// AudioStream is an interface that represents a format independent audio stream.
type AudioStream interface {
	io.ReadSeeker
	SampleRate() int
	Length() int64
}

// AudioEncoding is an enum type for the supported audio encodings: wav, mp3 and vorbis (.ogg files).
type AudioEncoding int

const (
	WAV AudioEncoding = iota
	MP3
	VORBIS
)

// DecodeAudioStream returns the right F32 audio stream from src depending on the given enc.
func DecodeAudioStream(src io.Reader, enc AudioEncoding) (AudioStream, error) {
	switch enc {
	case WAV:
		return wav.DecodeF32(src)
	case MP3:
		return mp3.DecodeF32(src)
	case VORBIS:
		return vorbis.DecodeF32(src)
	default:
		return nil, ErrInvalidFormat
	}
}

// DecodeRawAudio consumes the stream and gives back a RawAudio that can be used with many players.
func DecodeRawAudio(stream AudioStream) (RawAudio, error) {
	return io.ReadAll(stream)
}
