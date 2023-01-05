package ascii

import (
	"image"
)

type Builder struct {
	generator *Generator
}

func (b *Builder) WithAlphaThreshold(threshold int) *Builder {
	b.generator.alphaThreshold = threshold
	return b
}

func (b *Builder) WithAlphaValue(value byte) *Builder {
	b.generator.alphaValue = value
	return b
}

func (b *Builder) WithCharSet(charset CharSet) *Builder {
	b.generator.charset = charset
	return b
}

func (b *Builder) WithGammaCorrection(correction float32) *Builder {
	b.generator.gammaCorrection = correction
	return b
}

func (b *Builder) WithInput() *InputBuilder {
	return &InputBuilder{*b}
}

func (b *Builder) Build() *Generator {
	return b.generator
}

type InputBuilder struct {
	Builder
}

func (b *InputBuilder) Image(image image.Image) *InputBuilder {
	b.generator.img = image
	return b
}

// NewBuilder will create a builder style construct,
// which provides you with a default Generator.
// The builder 'object' provides convenience functions for specifying generator properties,
// e.g. what character set to use for the ascii generation, and what input image to use.
func NewBuilder() *Builder {
	return &Builder{generator: &Generator{alphaValue: '@', alphaThreshold: 0, gammaCorrection: 1.0}}
}
