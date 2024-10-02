package main

import (
	"context"

	"go.viam.com/rdk/components/camera"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/robot/client"
	"go.viam.com/utils/rpc"
)

// type "go run main.go" to run this command
func main() {
	logger := logging.NewDebugLogger("client")
	machine, err := client.New(
		context.Background(),
		"test1-main.myj4cbg09b.viam.cloud",
		logger,
		client.WithDialOptions(rpc.WithEntityCredentials(
			"d004e62e-cc0a-4ea4-8e13-59a8bc028eb0",
			rpc.Credentials{
				Type:    rpc.CredentialsTypeAPIKey,
				Payload: "ur6joo0p8j87yryjs5e1c5x41h5mmi2q",
			})),
	)
	if err != nil {
		logger.Fatal(err)
	}

	defer machine.Close(context.Background())
	logger.Info("Resources:")
	logger.Info(machine.ResourceNames())

	// camera-1
	// camera1, err := camera.FromRobot(machine, "camera-1")
	// if err != nil {
	// 	logger.Error(err)
	// 	return
	// }
	// camera1ReturnValue, err := camera1.Properties(context.Background())
	// if err != nil {
	// 	logger.Error(err)
	// 	return
	// }
	// logger.Infof("camera-1 Properties return value: %+v", camera1ReturnValue)

	// // camera-2
	// camera2, err := camera.FromRobot(machine, "camera-2")
	// if err != nil {
	// 	logger.Error(err)
	// 	return
	// }
	// camera2ReturnValue, err := camera2.Properties(context.Background())
	// if err != nil {
	// 	logger.Error(err)
	// 	return
	// }
	// logger.Infof("camera-2 Properties return value: %+v", camera2ReturnValue)

	// // youtube-stream
	// youtubeStream, err := camera.FromRobot(machine, "youtube-stream")
	// if err != nil {
	// 	logger.Error(err)
	// 	return
	// }
	// youtubeStreamReturnValue, err := youtubeStream.Properties(context.Background())
	// if err != nil {
	// 	logger.Error(err)
	// 	return
	// }
	// logger.Infof("youtube-stream Properties return value: %+v", youtubeStreamReturnValue)

	// // sensor-1
	// sensor1, err := sensor.FromRobot(machine, "sensor-1")
	// if err != nil {
	// 	logger.Error(err)
	// 	return
	// }
	// sensor1ReturnValue, err := sensor1.Readings(context.Background(), map[string]interface{}{})
	// if err != nil {
	// 	logger.Error(err)
	// 	return
	// }
	// logger.Infof("sensor-1 Readings return value: %+v", sensor1ReturnValue)

	// camera-3
	camera3, err := camera.FromRobot(machine, "camera-3")
	if err != nil {
		logger.Error(err)
		return
	}
	camera3ReturnValue, err := camera3.Properties(context.Background())
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Infof("camera-3 Properties return value: %+v", camera3ReturnValue)

	_, _, err = camera3.Images(context.Background())

}
