// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

package launchpad

// Metadata describes the parameters for the API for a board and OS
type Metadata struct {
	DistroArchSeries string `json:"-"`
	LiveFS           string `json:"-"`
	Project          string `json:"project,omitempty"`
	SubArch          string `json:"subarch,omitempty"`
	ImageFormat      string `json:"image_format,omitempty"`
}

var boards = map[string]Metadata{
	"raspberrypi2-core16": {
		DistroArchSeries: "https://api.launchpad.net/1.0/ubuntu/xenial/armhf",
		LiveFS:           "+livefs/ubuntu/xenial/ubuntu-core",
		Project:          "ubuntu-core",
		SubArch:          "raspi2",
		ImageFormat:      "ubuntu-image",
	},
	"raspberrypi2-core18": {
		DistroArchSeries: "https://api.launchpad.net/1.0/ubuntu/bionic/armhf",
		LiveFS:           "+livefs/ubuntu/bionic/ubuntu-core",
		Project:          "ubuntu-core",
		SubArch:          "raspi2",
		ImageFormat:      "ubuntu-image",
	},
	"raspberrypi3-core16": {
		DistroArchSeries: "https://api.launchpad.net/1.0/ubuntu/xenial/armhf",
		LiveFS:           "+livefs/ubuntu/xenial/ubuntu-core",
		Project:          "ubuntu-core",
		SubArch:          "raspi3",
		ImageFormat:      "ubuntu-image",
	},
	"raspberrypi3-core18": {
		DistroArchSeries: "https://api.launchpad.net/1.0/ubuntu/bionic/armhf",
		LiveFS:           "+livefs/ubuntu/bionic/ubuntu-core",
		Project:          "ubuntu-core",
		SubArch:          "raspi3",
		ImageFormat:      "ubuntu-image",
	},
	"intelnuc-core16": {
		DistroArchSeries: "https://api.launchpad.net/1.0/ubuntu/xenial/amd64",
		LiveFS:           "+livefs/ubuntu/xenial/ubuntu-core",
		Project:          "ubuntu-core",
		SubArch:          "",
		ImageFormat:      "ubuntu-image",
	},
	"intelnuc-core18": {
		DistroArchSeries: "https://api.launchpad.net/1.0/ubuntu/bionic/amd64",
		LiveFS:           "+livefs/ubuntu/bionic/ubuntu-core",
		Project:          "ubuntu-core",
		SubArch:          "",
		ImageFormat:      "ubuntu-image",
	},
	"snapdragon-core16": {
		DistroArchSeries: "https://api.launchpad.net/1.0/ubuntu/xenial/arm64",
		LiveFS:           "+livefs/ubuntu/xenial/ubuntu-core",
		Project:          "ubuntu-core",
		SubArch:          "snapdragon",
		ImageFormat:      "ubuntu-image",
	},
	"snapdragon-core18": {
		DistroArchSeries: "https://api.launchpad.net/1.0/ubuntu/bionic/arm64",
		LiveFS:           "+livefs/ubuntu/bionic/ubuntu-core",
		Project:          "ubuntu-core",
		SubArch:          "snapdragon",
		ImageFormat:      "ubuntu-image",
	},
	"cm3-core16": {
		DistroArchSeries: "https://api.launchpad.net/1.0/ubuntu/xenial/armhf",
		LiveFS:           "+livefs/ubuntu/xenial/ubuntu-core",
		Project:          "ubuntu-core",
		SubArch:          "cm3",
		ImageFormat:      "ubuntu-image",
	},
	"cm3-core18": {
		DistroArchSeries: "https://api.launchpad.net/1.0/ubuntu/bionic/armhf",
		LiveFS:           "+livefs/ubuntu/bionic/ubuntu-core",
		Project:          "ubuntu-core",
		SubArch:          "cm3",
		ImageFormat:      "ubuntu-image",
	},
}
