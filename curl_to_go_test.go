package curl_to_go

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	for _, tt := range tests {
		actual := Parse(tt.in)
		if !almostEqual(actual, tt.expect) {
			t.Errorf("[%s] want <%s>, got <%s>\n", tt.in, tt.expect, actual)
		}
	}
}

func almostEqual(s1, s2 string) bool {
	a1 := toSlice(s1)
	a2 := toSlice(s2)
	return reflect.DeepEqual(a1, a2)
}

func toSlice(s string) []string {
	var resp []string
	arr := strings.Split(s, "\n")
	for i := range arr {
		if strings.TrimSpace(arr[i]) != "" {
			resp = append(resp, arr[i])
		}
	}
	sort.Slice(resp, func(i, j int) bool {
		return resp[i] < resp[j]
	})
	return resp
}

var (
	tests = []struct {
		in     string
		expect string
	}{
		{
			`curl canhazip.com`,
			`// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
resp, err := http.Get("canhazip.com")
if err != nil {
	// handle err
}
defer resp.Body.Close()`,
		},
		{
			`curl https://api.example.com/surprise \
     -u banana:coconuts \
     -d "sample data"`,
			`// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

body := strings.NewReader(` + "`sample data`" + `)
req, err := http.NewRequest("POST", "https://api.example.com/surprise", body)
if err != nil {
	// handle err
}
req.SetBasicAuth("banana", "coconuts")
req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

resp, err := http.DefaultClient.Do(req)
if err != nil {
	// handle err
}
defer resp.Body.Close()`,
		},
		{
			`curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer b7d03a6947b217efb6f3ec3bd3504582" -d '{"type":"A","name":"www","data":"162.10.66.0","priority":null,"port":null,"weight":null}' "https://api.digitalocean.com/v2/domains/example.com/records"`,
			`// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

type Payload struct {
	Type     string      ` + "`json:\"type\"`" + `
	Name     string      ` + "`json:\"name\"`" + `
	Data     string      ` + "`json:\"data\"`" + `
	Priority interface{} ` + "`json:\"priority\"`" + `
	Port     interface{} ` + "`json:\"port\"`" + `
	Weight   interface{} ` + "`json:\"weight\"`" + `
}

data := Payload {
	// fill struct
}
payloadBytes, err := json.Marshal(data)
if err != nil {
	// handle err
}
body := bytes.NewReader(payloadBytes)

req, err := http.NewRequest("POST", "https://api.digitalocean.com/v2/domains/example.com/records", body)
if err != nil {
	// handle err
}
req.Header.Set("Content-Type", "application/json")
req.Header.Set("Authorization", "Bearer b7d03a6947b217efb6f3ec3bd3504582")

resp, err := http.DefaultClient.Do(req)
if err != nil {
	// handle err
}
defer resp.Body.Close()`,
		},
		{
			`curl -u "demo" -X POST -d @file1.txt -d @file2.txt https://example.com/upload`,
			`// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

f1, err := os.Open("file1.txt")
if err != nil {
	// handle err
}
defer f1.Close()
f2, err := os.Open("file2.txt")
if err != nil {
	// handle err
}
defer f2.Close()
payload := io.MultiReader(f1, f2)
req, err := http.NewRequest("POST", "https://example.com/upload", payload)
if err != nil {
	// handle err
}
req.SetBasicAuth("demo", "<PASSWORD>")
req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

resp, err := http.DefaultClient.Do(req)
if err != nil {
	// handle err
}
defer resp.Body.Close()`,
		},
		{
			`curl -X POST https://api.easypost.com/v2/shipments \
     -u API_KEY: \
     -d 'shipment[to_address][id]=adr_HrBKVA85' \
     -d 'shipment[from_address][id]=adr_VtuTOj7o' \
     -d 'shipment[parcel][id]=prcl_WDv2VzHp' \
     -d 'shipment[is_return]=true' \
     -d 'shipment[customs_info][id]=cstinfo_bl5sE20Y'`,
			`// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

body := strings.NewReader(` + "`shipment[to_address][id]=adr_HrBKVA85&shipment[from_address][id]=adr_VtuTOj7o&shipment[parcel][id]=prcl_WDv2VzHp&shipment[is_return]=true&shipment[customs_info][id]=cstinfo_bl5sE20Y`)" + `
req, err := http.NewRequest("POST", "https://api.easypost.com/v2/shipments", body)
if err != nil {
	// handle err
}
req.SetBasicAuth("API_KEY", "")
req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

resp, err := http.DefaultClient.Do(req)
if err != nil {
	// handle err
}
defer resp.Body.Close()`,
		},
	}
)
