package inception

import (
	"context"
	"io/ioutil"
	"log"
	tf_core_framework "tensorflow/core/framework"
	pb "tensorflow_serving/apis"

	google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
)

const (
	VERSION = "v0.0.1"
)

var (
	ModelName     = "inception"
	SignatureName = "predict_images"
)

func PredictRequest(b []byte) *pb.PredictRequest {
	request := &pb.PredictRequest{
		ModelSpec: &pb.ModelSpec{
			Name:          ModelName,
			SignatureName: SignatureName,
			Version: &google_protobuf.Int64Value{
				Value: int64(1),
			},
		},
		Inputs: map[string]*tf_core_framework.TensorProto{
			"images": &tf_core_framework.TensorProto{
				Dtype: tf_core_framework.DataType_DT_STRING,
				TensorShape: &tf_core_framework.TensorShapeProto{
					Dim: []*tf_core_framework.TensorShapeProto_Dim{
						&tf_core_framework.TensorShapeProto_Dim{
							Size: int64(1),
						},
					},
				},
				StringVal: [][]byte{b},
			},
		},
	}
	return request
}

func Predict(Client pb.PredictionServiceClient, ctx context.Context, request *pb.PredictRequest) (*pb.PredictResponse, error) {
	resp, err := Client.Predict(ctx, request)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func RunInference(ctx context.Context, request *pb.PredictRequest) (*pb.PredictResponse, error) {
	conn, err := grpc.Dial("", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := pb.NewPredictionServiceClient(conn)
	resp, err := client.Predict(ctx, request)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func RunInferenceOnImagePath(imgPath string) (*pb.PredictResponse, error) {
	ib, err := ioutil.ReadFile(imgPath)
	if err != nil {
		return nil, err
	}

	request := PredictRequest(ib)

	resp, err := RunInference(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func main() {
	resp, err := RunInferenceOnImagePath("ball.jpg")
	if err != nil {
		log.Panic(err)
	}
	log.Printf("resp: %v\n\n", resp)
	log.Println(resp.Outputs["classes"].StringVal[0])
}

/*

  outputs {
	key: "classes"
	value {
	  dtype: DT_STRING
	  tensor_shape {
	  dim {
		size: 1
	  }
	  dim {
		size: 5
	  }
	  }
	  string_val: "sports car, sport car"
	  string_val: "car wheel"
	  string_val: "racer, race car, racing car"
	  string_val: "grille, radiator grille"
	  string_val: "minivan"
	}
	}
	outputs {
	key: "scores"
	value {
	  dtype: DT_FLOAT
	  tensor_shape {
	  dim {
		size: 1
	  }
	  dim {
		size: 5
	  }
	  }
	  float_val: 9.68140888214
	  float_val: 7.5523018837
	  float_val: 7.47641181946
	  float_val: 6.99047279358
	  float_val: 6.82593536377
	}
	}
	model_spec {
	name: "inception"
	version {
	  value: 1
	}
	signature_name: "predict_images"
}*/
