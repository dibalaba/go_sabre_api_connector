"io/ioutil"
	"net/http"
	"bytes"


request, _ := http.NewRequest("POST", "https://api-crt.cert.havail.sabre.com/v1/offers/shop", bytes.NewBuffer(jsonValue))
request.Header.Set("Content-Type", "application/json")
request.Header.Set("Accept", "application/json")
request.Header.Set("Authorization", "Bearer  T1RLAQK41r/+HidhhS8WxTUsa5GeCHPuSxCuwhQRjVwMtVVC0ku2KTZEAADA647YQGgUqzYqbe9txvOw+iCIDzL7YUezPvhqwBYJkC/cWGZDLY1U/X16YFKXiuMPltsLZIlGQB7dq9saGe7aAsLqLnJV633Oaq+304LRInqbFjjYVE7SVppLDoWsRlItNiBKTTVB0DVBooz3l4HKZe8Yi6st9d3+5G/3IpvAPETKSA2WjqSeW6fWquohlC6Q0aO0vz8DmwojFVwu1GTn/+2pyL00tFrdwtOPemnSFOpjW+NI8psr9wsye+ZTil1O")
client := &http.Client{}
response, err := client.Do(request)

if err != nil {
	fmt.Printf("The HTTP request failed with error %s\n", err)
} else {
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
}