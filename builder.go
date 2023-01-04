package ascii

import (
	"image"
)

type Builder struct {
	properties *Properties
}

func (b *Builder) WithCharSet(charset CharSet) *Builder {
	b.properties.charset = charset
	return b
}

func (b *Builder) WithInput() *InputBuilder {
	return &InputBuilder{*b}
}

func (b *Builder) Build() *Properties {
	return b.properties
}

type InputBuilder struct {
	Builder
}

func (b *InputBuilder) Image(image image.Image) *InputBuilder {
	b.properties.img = image
	return b
}

func NewBuilder() *Builder {
	return &Builder{properties: &Properties{}}
}
