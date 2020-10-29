package unsplash

import (
	"testing"
)

func TestSetUnsplashAcessKey(t *testing.T) {
	SetUnsplashAcessKey("gDiZdJTkd2JfA6W8dklxAUlPyGk1lSITxPlutej5eAA")
	_, total, _, err := SearchPhotosByWord("medical", 1, 9)
	if err != nil {
		t.Errorf("key invalid or api changed")
		return
	}
	if total <= 0 {
		t.Errorf("api did not work")
	}
}