package mnist

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	tf_core_framework "tensorflow/core/framework"
	pb "tensorflow_serving/apis"

	google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
	"gocv.io/x/gocv"

	"google.golang.org/grpc"
)

const (
	VERSION = "v0.0.0"
)

var (
	SERVER = "localhost:9000"
)

func PredictRequest(image []float32) *pb.PredictRequest {
	request := &pb.PredictRequest{
		ModelSpec: &pb.ModelSpec{
			Name:          "mnist",
			SignatureName: "predict_images",
			Version: &google_protobuf.Int64Value{
				Value: int64(1),
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
	return request
}

func PredictRequestFromBytes(ib []byte) (*pb.PredictRequest, error) {
	matb, err := gocv.IMDecode(ib, -1)
	if err != nil {
		return nil, fmt.Errorf("not an image %v", err)
	}

	//matb0 := gocv.NewMat()

	mat := gocv.NewMat()
	matb.ConvertTo(&mat, gocv.MatTypeCV32F)

	imgfloat := make([]float32, 0, 0)

	for i := 0; i < 28; i++ {
		for j := 0; j < 28; j++ {
			//fmt.Printf("value: %v %v: %v %v\n", i, j, matb.GetFloatAt(i, j), mat.GetFloatAt(i, j))
			imgfloat = append(imgfloat, mat.GetFloatAt(i, j))
		}
	}
	request := PredictRequest(imgfloat)
	return request, nil
}

func RunInference(ctx context.Context, request *pb.PredictRequest) (*pb.PredictResponse, error) {
	conn, err := grpc.Dial(SERVER, grpc.WithInsecure())
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

	request, err := PredictRequestFromBytes(ib)
	if err != nil {
		return nil, err
	}

	resp, err := RunInference(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func main() {
	resp, err := RunInferenceOnImagePath(os.Args[1])
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(resp)
}
