package main

import (
	"net/http"
)

type MyJsonName struct {
	Date  string `json:"date"`
	Rates struct {
		Usd float64 `json:"USD"`
		Aud float64 `json:"AUD"`
		Gbp float64 `json:"GBP"`
	} `json:"rates"`
	Base string `json:"base"`
}

func main() {

	chatString := ""

	http.HandleFunc("/getchat", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(chatString))
	})
	http.HandleFunc("/addchat", func(writer http.ResponseWriter, request *http.Request) {
		chatString = chatString + request.URL.Query()["msg"][0]
		//request.URL.Path[1:]
	})

	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/joke", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("joke of the day is: a man is funny haha"))
	})
	http.ListenAndServe(":8080", nil)

	//var theResponse MyJsonName
	//scanner := bufio.NewScanner(os.Stdin)
	//for scanner.Scan() {
	//	switch text := scanner.Text(); text {
	//	case "usd":
	//		fmt.Println(strconv.FormatFloat(theResponse.Rates.Usd, 'f', 4, 64))
	//	case "gbp":
	//		fmt.Println(strconv.FormatFloat(theResponse.Rates.Gbp, 'f', 4, 64))
	//	case "fetch":
	//		response, _ := http.Get("https://api.exchangeratesapi.io/latest?base=AUD")
	//		data, _ := ioutil.ReadAll(response.Body)
	//		json.Unmarshal([]byte(data), &theResponse)
	//		fmt.Println(theResponse.Base)
	//		fmt.Println(theResponse.Date)
	//	case "fetchraw":
	//		response, _ := http.Get("https://api.exchangeratesapi.io/latest?base=AUD")
	//		data, _ := ioutil.ReadAll(response.Body)
	//		fmt.Println(string(data))
	//	case "startserving":
	//
	//	default:
	//		fmt.Println("sorry not recognized")
	//	}
	//}

}
