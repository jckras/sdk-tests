// package main

// import (
// 	"context"

// 	"go.viam.com/rdk/components/camera"
// 	"go.viam.com/rdk/logging"
// 	"go.viam.com/rdk/robot/client"
// 	"go.viam.com/utils/rpc"
// )

// func main() {
// 	logger := logging.NewDebugLogger("client")
// 	machine, err := client.New(
// 		context.Background(),
// 		"ada-main.ikzfbcf5i3.viam.cloud",
// 		logger,
// 		client.WithDialOptions(rpc.WithEntityCredentials(
// 			"783426e6-7f5c-4821-918e-0e9c9e00b99a",
// 			rpc.Credentials{
// 				Type:    rpc.CredentialsTypeAPIKey,
// 				Payload: "sbqyy0lvrxdtwrjaiml3f740kbq1qdij",
// 			})),
// 	)
// 	if err != nil {
// 		logger.Fatal(err)
// 	}

// 	defer machine.Close(context.Background())
// 	logger.Info("Resources:")
// 	logger.Info(machine.ResourceNames())

// 	// Note that the pin supplied is a placeholder. Please change this to a valid pin.
// 	// board-1
// 	// board1, err := board.FromRobot(machine, "board-1")
// 	// if err != nil {
// 	//   logger.Error(err)
// 	//   return
// 	// }
// 	// board1ReturnValue, err := board1.GPIOPinByName("16")
// 	// if err != nil {
// 	//   logger.Error(err)
// 	//   return
// 	// }
// 	// logger.Infof("board-1 GPIOPinByName return value: %+v", board1ReturnValue)

// 	// camera-3
// 	camera3, err := camera.FromRobot(machine, "camera-3")
// 	if err != nil {
// 		logger.Error(err)
// 		return
// 	}
// 	camera3ReturnValue, err := camera3.Properties(context.Background())
// 	if err != nil {
// 		logger.Error(err)
// 		return
// 	}
// 	logger.Infof("camera-3 Properties return value: %+v", camera3ReturnValue)

// 	_, _, err = camera3.Images(context.Background())

// 	// camera2, err := camera.FromRobot(machine, "camera-2")
// 	// if err != nil {
// 	// 	logger.Error(err)
// 	// 	return
// 	// }
// 	// camera2ReturnValue, err := camera2.Properties(context.Background())
// 	// if err != nil {
// 	// 	logger.Error(err)
// 	// 	return
// 	// }
// 	// logger.Infof("camera-2 Properties return value: %+v", camera2ReturnValue)

// 	// _, _, err = camera2.Images(context.Background())

// 	// camera1, err := camera.FromRobot(machine, "camera-1")
// 	// if err != nil {
// 	// 	logger.Error(err)
// 	// 	return
// 	// }
// 	// camera1ReturnValue, err := camera1.Properties(context.Background())
// 	// if err != nil {
// 	// 	logger.Error(err)
// 	// 	return
// 	// }
// 	// logger.Infof("camera-1 Properties return value: %+v", camera1ReturnValue)

// 	// _, _, err = camera1.Images(context.Background())

// }
package main

import (
	"context"

	"go.viam.com/rdk/components/board"
	"go.viam.com/rdk/components/camera"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/robot/client"
	"go.viam.com/utils/rpc"
)

func main() {
	logger := logging.NewDebugLogger("client")
	machine, err := client.New(
		context.Background(),
		"ada-main.ikzfbcf5i3.viam.cloud",
		logger,
		client.WithDialOptions(rpc.WithEntityCredentials(
			"783426e6-7f5c-4821-918e-0e9c9e00b99a",
			rpc.Credentials{
				Type:    rpc.CredentialsTypeAPIKey,
				Payload: "sbqyy0lvrxdtwrjaiml3f740kbq1qdij",
			})),
	)
	if err != nil {
		logger.Fatal(err)
	}

	defer machine.Close(context.Background())
	logger.Info("Resources:")
	logger.Info(machine.ResourceNames())

	// Note that the pin supplied is a placeholder. Please change this to a valid pin.
	// board-1
	board1, err := board.FromRobot(machine, "board-1")
	if err != nil {
		logger.Error(err)
		return
	}
	board1ReturnValue, err := board1.GPIOPinByName("16")
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Infof("board-1 GPIOPinByName return value: %+v", board1ReturnValue)

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

	// // camera-4
	// camera4, err := camera.FromRobot(machine, "camera-4")
	// if err != nil {
	// 	logger.Error(err)
	// 	return
	// }
	// camera4ReturnValue, err := camera4.Properties(context.Background())
	// if err != nil {
	// 	logger.Error(err)
	// 	return
	// }
	// logger.Infof("camera-4 Properties return value: %+v", camera4ReturnValue)

	// logitech-cam
	logitechCam, err := camera.FromRobot(machine, "logitech-cam")
	if err != nil {
		logger.Error(err)
		return
	}
	logitechCamReturnValue, err := logitechCam.Properties(context.Background())
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Infof("logitech-cam Properties return value: %+v", logitechCamReturnValue)
	logitechCam.Images(context.Background())

}
