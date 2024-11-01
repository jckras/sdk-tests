package main

import (
	"context"
	"fmt"
	"net/url"

	"go.mongodb.org/mongo-driver/bson"
	pbData "go.viam.com/api/app/data/v1"
	"go.viam.com/rdk/logging"
	"go.viam.com/utils/rpc"
)

var (
	viamBaseURL  = "https://app.viam.com"
	viamAPIKey   = "fc3aoh6xx5kmfzb135guvxjcf12l02k8"
	viamAPIKeyID = "b5cf5cb8-5dcc-41d4-8b1f-ca79e28cef72"
)

func getDataClient(ctx context.Context, viamBaseURL, viamAPIKeyID, viamAPIKey string, logger logging.Logger) (pbData.DataServiceClient, error) {
	u, err := url.Parse(viamBaseURL + ":443")
	if err != nil {
		return nil, err
	}

	opts := rpc.WithEntityCredentials(
		viamAPIKeyID,
		rpc.Credentials{
			Type:    rpc.CredentialsTypeAPIKey,
			Payload: viamAPIKey,
		})
	conn, err := rpc.DialDirectGRPC(ctx, u.Host, logger.AsZap(), opts)
	if err != nil {
		return nil, err
	}

	return pbData.NewDataServiceClient(conn), nil
}

//another way to do it!
// func getDataClient(ctx context.Context) (pbData.DataServiceClient, error) {
// 	accessToken := os.Getenv("ACCESS_TOKEN")
// 	clientConn, err := rpc.DialDirectGRPC(
// 		ctx,
// 		"app.viam.com:443",
// 		nil,
// 		rpc.WithStaticAuthenticationMaterial(accessToken),
// 	)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "failed to establish gRPC connection")
// 	}

// 	return pbData.NewDataServiceClient(clientConn), nil
// }

func main() {
	dataClient, err := getDataClient(context.Background(), viamBaseURL, viamAPIKeyID, viamAPIKey, logging.NewLogger("viam"))
	// dataClient, err := getDataClient(context.Background())
	if err != nil {
		fmt.Println("Error getting data client:", err)
		return
	}
	print("i got dataClient", dataClient)

	// Create BSON documents for MongoDB queries
	matchStage := bson.M{"$match": bson.M{"organization_id": "e76d1b3b-0468-4efd-bb7f-fb1d2b352fcb"}}
	limitStage := bson.M{"$limit": 1}
	// Convert to BSON byte arrays
	matchBytes, err := bson.Marshal(matchStage)
	if err != nil {
		fmt.Println("Error marshalling match stage:", err)
		return
	}
	limitBytes, err := bson.Marshal(limitStage)
	if err != nil {
		fmt.Println("Error marshalling limit stage:", err)
		return
	}

	// Create the TabularDataByMQL request
	mqlRequest := &pbData.TabularDataByMQLRequest{
		OrganizationId: "e76d1b3b-0468-4efd-bb7f-fb1d2b352fcb",
		MqlBinary:      [][]byte{matchBytes, limitBytes}, // Pass both stages as MqlBinary
	}

	resp, err := dataClient.TabularDataByMQL(context.Background(), mqlRequest)
	if err != nil {
		fmt.Println("Error calling TabularDataByMQL:", err)
		return
	}

	// Print the response to verify
	if len(resp.GetData()) > 0 {
		fmt.Println("Received data:", resp.GetData())
	} else {
		fmt.Println("No data received.")
	}
}
