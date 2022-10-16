package decoder

import "encoding/json"

/*----------------------------------Decode JSON BLOCK--------------------------------------*/

func Worker(content string) (Addr string, targets []string) {
	type Json struct {
		Subnet struct {
			Address string   `json:"Address"`
			UHosts  []string `json:"UHosts"`
		} `json:"Subnet"`
	}
	var str Json
	err := json.Unmarshal([]byte(content), &str)
	if err != nil {
		panic(err)
	}
	return str.Subnet.Address, str.Subnet.UHosts
}
