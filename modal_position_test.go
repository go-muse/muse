package muse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getModalPositionNameByWeight(t *testing.T) {
	weights := []Weight{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var expectedModalPositionName ModalPositionName
	for _, weight := range weights {
		switch {
		case weight < 0:
			expectedModalPositionName = ModalPositionNameHigh
		case weight == 0:
			expectedModalPositionName = ModalPositionNameNeutral
		case weight > 0:
			expectedModalPositionName = ModalPositionNameLow
		}
		assert.Equal(t, expectedModalPositionName, getModalPositionNameByWeight(weight))
	}
}

func TestNewModalPositionNameByWeight(t *testing.T) {
	weights := []Weight{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var expectedModalPositionName ModalPositionName
	var modalPosition *ModalPosition
	for _, weight := range weights {
		switch {
		case weight < 0:
			expectedModalPositionName = ModalPositionNameHigh
		case weight == 0:
			expectedModalPositionName = ModalPositionNameNeutral
		case weight > 0:
			expectedModalPositionName = ModalPositionNameLow
		}
		modalPosition = NewModalPositionByWeight(weight)
		assert.Equal(t, expectedModalPositionName, modalPosition.name)
		assert.Equal(t, weight, modalPosition.GetWeight())
	}
}
