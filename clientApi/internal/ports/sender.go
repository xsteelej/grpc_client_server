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
		s.Db.Write(context.Background(), portFromMap(id, values.(map[string]interface{})))
	}

	return len(p), nil
}

func portFromMap(id string, m map[string]interface{}) *grpc.Port {
	name, city, province, country, timezone := "", "", "", "", ""
	coordinates := make([]float64, 0)
	alias, regions, unlocs := make([]string, 0), make([]string, 0), make([]string, 0)

	if _, ok := m["name"]; ok {
		name = m["name"].(string)
	}

	if _, ok := m["city"]; ok {
		city = m["city"].(string)
	}

	if _, ok := m["province"]; ok {
		province = m["province"].(string)
	}

	if _, ok := m["country"]; ok {
		country = m["country"].(string)
	}

	if _, ok := m["alias"]; ok {
		alias = parseStringArray(m["alias"].([]interface{}))
	}

	if _, ok := m["regions"]; ok {
		regions = parseStringArray(m["regions"].([]interface{}))
	}
	if _, ok := m["coordinates"]; ok {
		coordinates = parseFloatArray(m["coordinates"].([]interface{}))
	}
	if _, ok := m["timezone"]; ok {
		timezone = m["timezone"].(string)
	}
	if _, ok := m["unlocs"]; ok {
		unlocs = parseStringArray(m["unlocs"].([]interface{}))
	}
	return &grpc.Port{
		Id:          id,
		Name:        name,
		City:        city,
		Province:    province,
		Country:     country,
		Alias:       alias,
		Regions:     regions,
		Coordinates: coordinates,
		Timezone:    timezone,
		Unlocs:      unlocs,
	}

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
