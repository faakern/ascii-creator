package ascii

import (
	"image"
)

type Builder struct {
	generator *Generator
}

func (b *Builder) WithCharSet(charset CharSet) *Builder {
	b.generator.charset = charset
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

func NewBuilder() *Builder {
	return &Builder{generator: &Generator{}}
}
