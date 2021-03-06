package stdimpl

import "github.com/medtune/capsul/pkg/pbreq"

var (
	// Mnist .
	Mnist = &pbreq.Meta{
		Name:      "mnist",
		Signature: "predict_images",
		Version:   1,
	}

	// Inception .
	Inception = &pbreq.Meta{
		Name:      "inception",
		Signature: "predict_images",
		Version:   1,
		UseDims:   true,
	}

	//MuraIRNV2 .
	MuraIRNV2 = &pbreq.Meta{
		Name:      "mura_inception_resnet_v2",
		Signature: "predict_images",
		Version:   1,
	}

	// MuraMNV2 .
	MuraMNV2 = &pbreq.Meta{
		Name:      "mura_mobilenet_v2",
		Signature: "predict_images",
		Version:   1,
	}

	// Chexray .
	ChexrayMNV2 = &pbreq.Meta{
		Name:      "chexray",
		Signature: "predict_images",
		Version:   1,
	}

	// Chexray .
	ChexrayDN121 = &pbreq.Meta{
		Name:      "chexray_densenet_121",
		Signature: "predict",
		Version:   1,
		UseDims:   true,
	}

	// Map .
	Map = map[string]*pbreq.Meta{
		"mnist":          Mnist,
		"inception":      Inception,
		"mura-irn-v2":    MuraIRNV2,
		"mura-mn-v2":     MuraMNV2,
		"chexray-mn-v2":  ChexrayMNV2,
		"chexray-dn-121": ChexrayDN121,
	}
)
