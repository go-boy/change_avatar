package solarlunar

import (
	"fmt"
	"testing"
)

func TestSolarToChineseLuanr(t *testing.T) {
	solarDate := "2048-11-01"
	fmt.Println(SolarToChineseLuanr(solarDate))
}

func TestSolarToSimpleLunar(t *testing.T) {
	solarDate := "2019-11-12"
	fmt.Println(SolarToSimpleLuanr(solarDate))
}

func TestLunarToSolar(t *testing.T) {
	lunarDate := "2045-09-23"
	fmt.Println(LunarToSolar(lunarDate, false))
}
