package ports

import (
	"context"
	"encoding/json"
	grpc "github.com/xsteelej/grpc_client_server/grpc"
)

type Sender struct {
	Db grpc.PortsDatabaseClient
}

func (s *Sender) Write(p []byte) (n int, err error) {
	var ports map[string]interface{}
	parseErr := json.Unmarshal(p, &ports)
	if err != nil {
		return 0, parseErr
	}

	for id, values := range ports {
		var portMap = values.(map[string]interface{})
		port := &grpc.Port{
			Id:          id,
			City:        portMap["city"].(string),
			Province:    portMap["province"].(string),
			Country:     portMap["country"].(string),
			Alias:       parseStringArray(portMap["alias"].([]interface{})),
			Regions:     parseStringArray(portMap["regions"].([]interface{})),
			Coordinates: parseFloatArray(portMap["coordinates"].([]interface{})),
			Timezone:    portMap["timezone"].(string),
			Unlocs:      parseStringArray(portMap["unlocs"].([]interface{})),
		}
		s.Db.Write(context.Background(), port, nil)
	}

	return len(p), nil
}

func parseStringArray(interfaces []interface{}) []string {
	strings := make([]string, len(interfaces))
	for i, value := range interfaces {
		strings[i] = value.(string)
	}
	return strings
}

func parseFloatArray(interfaces []interface{}) []float64 {
	floats := make([]float64, len(interfaces))
	for i, value := range interfaces {
		floats[i] = value.(float64)
	}
	return floats
}

func parseIntArray(interfaces []interface{}) []int32 {
	ints := make([]int32, len(interfaces))
	for i, value := range interfaces {
		ints[i] = value.(int32)
	}
	return ints
}
