package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"

import (
	"fmt"
)

type ResolutionType int

const (
	UndefinedResolution OrientationType = iota
	PixelsPerInchResolution
	PixelsPerCentimeterResolution
)

var resolutionTypeStrings = map[ResolutionType]string{
	UndefinedResolution:           "UndefinedResolution",
	PixelsPerInchResolution:       "PixelsPerInchResolution",
	PixelsPerCentimeterResolution: "PixelsPerCentimeterResolution",
}

func (ct *ResolutionType) String() string {
	if v, ok := resolutionTypeStrings[ResolutionType(*ct)]; ok {
		return v
	}
	return fmt.Sprintf("UnknownCompression[%d]", *ct)
}
