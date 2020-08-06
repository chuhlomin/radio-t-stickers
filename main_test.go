package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var splitMessageTests = []struct {
	in  string
	out []string
}{
	{"Красота нечеловеческая", []string{"Красота", "нечеловеческая"}},
	{"Цена конская", []string{"Цена конская"}},
	{"Это позор какой-то", []string{"Это позор какой-то"}},
	{"Это сделано чужими для хищников", []string{"Это сделано чужими", "для хищников"}},
	{"В действительности всё совсем не так, как на самом деле", []string{"В действительности", "всё совсем не так,", "как на самом деле"}},
	{"Рефакторинг этот не_хочешь ты", []string{"Рефакторинг этот", "не хочешь ты"}},
}

func TestSplitMessage(t *testing.T) {
	for _, c := range splitMessageTests {
		assert.Equal(t, c.out, splitMessage(c.in))
	}
}

var calculateSizesTests = []struct {
	in  int
	out Sizes
}{
	{
		1,
		Sizes{
			CanvasHeight:        154,
			BackgroundHeight:    134,
			BalloonHeight:       99,
			AvatarPositionY:     52,
			AvatarCircleCenterY: 93,
			Lines:               []int{107},
			BackgroundShape:     backgroundShapes[0],
			BalloonShape:        balloonShapes[0],
		},
	},
	{
		2,
		Sizes{
			CanvasHeight:        192,
			BackgroundHeight:    172,
			BalloonHeight:       137,
			AvatarPositionY:     90,
			AvatarCircleCenterY: 131,
			Lines:               []int{107, 145},
			BackgroundShape:     backgroundShapes[1],
			BalloonShape:        balloonShapes[1],
		},
	},
	{
		3,
		Sizes{
			CanvasHeight:        230,
			BackgroundHeight:    210,
			BalloonHeight:       175,
			AvatarPositionY:     128,
			AvatarCircleCenterY: 169,
			Lines:               []int{107, 145, 183},
			BackgroundShape:     backgroundShapes[2],
			BalloonShape:        balloonShapes[2],
		},
	},
}

func TestCalculateSizes(t *testing.T) {
	for _, c := range calculateSizesTests {
		assert.Equal(t, c.out, calculateSizes(c.in))
	}
}
