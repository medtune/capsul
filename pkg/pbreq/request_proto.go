package pbreq

import (
	tf_core_framework "tensorflow/core/framework"
	pb "tensorflow_serving/apis"

	google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
)

// Meta .
type Meta struct {
	Name      string
	Signature string
	Version   int64
	UseDims   bool
}

// Predict .
func Predict(m *Meta, images ...[]byte) *pb.PredictRequest {
	r := &pb.PredictRequest{
		ModelSpec: &pb.ModelSpec{
			Name:          m.Name,
			SignatureName: m.Signature,
			Version: &google_protobuf.Int64Value{
				Value: m.Version,
			},
		},
		Inputs: map[string]*tf_core_framework.TensorProto{
			"images": &tf_core_framework.TensorProto{
				Dtype:     tf_core_framework.DataType_DT_STRING,
				StringVal: images,
			},
		},
	}

	if m.UseDims {
		r.Inputs["images"].TensorShape = &tf_core_framework.TensorShapeProto{
			Dim: []*tf_core_framework.TensorShapeProto_Dim{
				&tf_core_framework.TensorShapeProto_Dim{
					Size: int64(len(images)),
				},
			},
		}
	}

	return r
}

// PredictF32 .
func PredictF32(m *Meta, image []float32) *pb.PredictRequest {
	return &pb.PredictRequest{
		ModelSpec: &pb.ModelSpec{
			Name:          m.Name,
			SignatureName: m.Signature,
			Version: &google_protobuf.Int64Value{
				Value: m.Version,
			},
		},
		Inputs: map[string]*tf_core_framework.TensorProto{
			"images": &tf_core_framework.TensorProto{
				Dtype: tf_core_framework.DataType_DT_FLOAT,
				TensorShape: &tf_core_framework.TensorShapeProto{
					Dim: []*tf_core_framework.TensorShapeProto_Dim{
						&tf_core_framework.TensorShapeProto_Dim{
							Size: int64(1),
						},
						&tf_core_framework.TensorShapeProto_Dim{
							Size: int64(784),
						},
					},
				},
				FloatVal: image,
			},
		},
	}
}

// Status .
func Status(m *Meta) *pb.GetModelStatusRequest {
	return &pb.GetModelStatusRequest{
		ModelSpec: &pb.ModelSpec{
			Name: m.Name,
			Version: &google_protobuf.Int64Value{
				Value: m.Version,
			},
		},
	}
}

// Predict .
func PredictFTest(m *Meta, images []float32) *pb.PredictRequest {
	r := &pb.PredictRequest{
		ModelSpec: &pb.ModelSpec{
			Name:          m.Name,
			SignatureName: m.Signature,
			Version: &google_protobuf.Int64Value{
				Value: m.Version,
			},
		},
		Inputs: map[string]*tf_core_framework.TensorProto{
			"images": &tf_core_framework.TensorProto{
				Dtype:    tf_core_framework.DataType_DT_FLOAT,
				FloatVal: images,
			},
		},
	}

	if m.UseDims {
		r.Inputs["images"].TensorShape = &tf_core_framework.TensorShapeProto{
			Dim: []*tf_core_framework.TensorShapeProto_Dim{
				&tf_core_framework.TensorShapeProto_Dim{
					Size: int64(1),
				},
				&tf_core_framework.TensorShapeProto_Dim{
					Size: int64(224),
				},
				&tf_core_framework.TensorShapeProto_Dim{
					Size: int64(224),
				},
				&tf_core_framework.TensorShapeProto_Dim{
					Size: int64(3),
				},
			},
		}
	}

	return r
}
