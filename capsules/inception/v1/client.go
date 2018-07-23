package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	tf_core_framework "tensorflow/core/framework"
	pb "tensorflow_serving/apis"

	google_protobuf "github.com/golang/protobuf/ptypes/wrappers"

	"google.golang.org/grpc"
)

const (
	VERSION = "v0.0.0"
)

func main() {
	servingAddress := flag.String("serving-address", "35.195.114.253:9000", "The tensorflow serving address")
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("Usage: " + os.Args[0] + " --serving-address localhost:10000 path/to/img.png")
		os.Exit(1)
	}

	imgPath, err := filepath.Abs(flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}

	imageBytes, err := ioutil.ReadFile(imgPath)
	if err != nil {
		log.Fatalln(err)
	}

	request := &pb.PredictRequest{
		ModelSpec: &pb.ModelSpec{
			Name:          "inception",
			SignatureName: "predict_images",
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
				StringVal: [][]byte{imageBytes},
			},
		},
	}

	conn, err := grpc.Dial(*servingAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot connect to the grpc server: %v\n", err)
	}
	defer conn.Close()

	client := pb.NewPredictionServiceClient(conn)

	resp, err := client.Predict(context.Background(), request)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp)
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
