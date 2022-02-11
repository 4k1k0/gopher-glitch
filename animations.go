package main

import (
	"github.com/jpoz/glitch"
)

func glitchIt(gl *glitch.Glitch) {
	for _, style := range config.Styles {
		switch style {
		case GlitchStyleTransposeInput:
			glitchTransposeInput(gl)
		case GlitchStyleVerticalTransposeInput:
			glitchVerticalTransposeInput(gl)
		case GlitchStyleCompressionGhost:
			glitchCompressionGhost(gl)
		case GlitchStyleGhostStreach:
			glitchGhostStreach(gl)
		case GlitchStyleHalfLifeLeft:
			glitchHalfLifeLeft(gl)
		case GlitchStyleHalfLifeRight:
			glitchHalfLifeRight(gl)
		case GlitchStyleChannelShiftLeft:
			glitchChannelShiftLeft(gl)
		case GlitchStyleChannelShiftRight:
			glitchChannelShiftRight(gl)
		case GlitchStyleBlueBoost:
			glitchBlueBoost(gl)
		case GlitchStyleGreenBoost:
			glitchGreenBoost(gl)
		case GlitchStyleRedBoost:
			glitchRedBoost(gl)
		case GlitchStylePrismBurst:
			glitchPrismBurst(gl)
		case GlitchStyleNoise:
			glitchNoise(gl)
		}
	}
}

func glitchTransposeInput(gl *glitch.Glitch) {
	width, height := gl.Bounds.Max.X, gl.Bounds.Max.Y
	w, h := generateWidthAndHeight(width, height)
	gl.TransposeInput(w, h, config.TransposeInput)
}

func glitchVerticalTransposeInput(gl *glitch.Glitch) {
	width, height := gl.Bounds.Max.X, gl.Bounds.Max.Y
	w, h := generateWidthAndHeight(width, height)
	gl.VerticalTransposeInput(w, h, config.VertitalTransposeInput)
}

func glitchCompressionGhost(gl *glitch.Glitch) {
	gl.CompressionGhost()
}

func glitchGhostStreach(gl *glitch.Glitch) {
	gl.GhostStreach()
}

func glitchHalfLifeLeft(gl *glitch.Glitch) {
	width, height := gl.Bounds.Max.X, gl.Bounds.Max.Y
	w, h := generateWidthAndHeight(width, height)

	if width < height {
		gl.HalfLifeLeft(w, width)
	} else {
		gl.HalfLifeLeft(h, height)
	}

}

func glitchHalfLifeRight(gl *glitch.Glitch) {
	width, height := gl.Bounds.Max.X, gl.Bounds.Max.Y
	w, h := generateWidthAndHeight(width, height)
	if width < height {
		gl.HalfLifeRight(w, width)
	} else {
		gl.HalfLifeRight(h, height)
	}
}

func glitchChannelShiftLeft(gl *glitch.Glitch) {
	gl.ChannelShiftLeft()
}

func glitchChannelShiftRight(gl *glitch.Glitch) {
	gl.ChannelShiftRight()
}

func glitchBlueBoost(gl *glitch.Glitch) {
	gl.BlueBoost()
}

func glitchGreenBoost(gl *glitch.Glitch) {
	gl.GreenBoost()
}

func glitchRedBoost(gl *glitch.Glitch) {
	gl.RedBoost()
}

func glitchPrismBurst(gl *glitch.Glitch) {
	gl.PrismBurst()
}

func glitchNoise(gl *glitch.Glitch) {
	r, g, b := generateRGB()
	a := generateRandomNumber(1, 600)

	gl.Noise(r, g, b, float64(a))
}
