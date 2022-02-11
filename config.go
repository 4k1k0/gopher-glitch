package main

import (
	"encoding/json"
	"os"
	"sync"
)

type GlitchStyle string

const (
	GlitchStyleTransposeInput         GlitchStyle = "TransposeInput"
	GlitchStyleVerticalTransposeInput GlitchStyle = "VerticalTransposeInput"
	GlitchStyleCompressionGhost       GlitchStyle = "CompressionGhost"
	GlitchStyleGhostStreach           GlitchStyle = "GhostStreach"
	GlitchStyleHalfLifeLeft           GlitchStyle = "HalfLifeLeft"
	GlitchStyleHalfLifeRight          GlitchStyle = "HalfLifeRight"
	GlitchStyleChannelShiftLeft       GlitchStyle = "ChannelShiftLeft"
	GlitchStyleChannelShiftRight      GlitchStyle = "ChannelShiftRight"
	GlitchStyleBlueBoost              GlitchStyle = "BlueBoost"
	GlitchStyleGreenBoost             GlitchStyle = "GreenBoost"
	GlitchStyleRedBoost               GlitchStyle = "RedBoost"
	GlitchStylePrismBurst             GlitchStyle = "PrismBurst"
	GlitchStyleNoise                  GlitchStyle = "Noise"
)

type Config struct {
	Path                   string        `json:"path"`
	Styles                 []GlitchStyle `json:"styles"`
	TransposeInput         bool          `json:"transpose_input"`
	VertitalTransposeInput bool          `json:"vertital_transpose_input"`
}

var config *Config
var once sync.Once

func init() {
	getInstancia()
}

func getInstancia() *Config {
	file, _ := os.Open("config.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	var conf Config
	err := decoder.Decode(&conf)
	if err != nil {
		conf.Path = "art"
	}

	once.Do(func() {
		config = &Config{
			Path:                   conf.Path,
			Styles:                 conf.Styles,
			TransposeInput:         conf.TransposeInput,
			VertitalTransposeInput: conf.VertitalTransposeInput,
		}
	})
	return config
}
