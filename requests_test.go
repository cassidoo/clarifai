package clarifai

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	tokenResponse    = `{"access_token":"1234567890abcdefg","expires_in":36000,"scope": "api_access", "token_type": "Bearer"}`
	infoResponse     = `{"status_code":"OK","status_msg":"All images in request have completed successfully. ","results":{"max_image_size":100000,"default_language":"en","max_video_size":100000,"max_image_bytes":10485760,"min_image_size":1,"default_model":"default","max_video_bytes":104857600,"max_video_duration":1800,"max_batch_size":128,"max_video_batch_size":1,"min_video_size":1,"api_version":0.1}}`
	tagResponse      = `{"status_code":"OK","status_msg":"All images in request have completed successfully. ","meta":{"tag":{"timestamp":1443807051.1546,"model":"default","config":"0b2b7436987dd912e077ff576731f8b7"}},"results":[{"docid":273996447814733945748816681886883360608,"url":"http:\/\/www.clarifai.com\/img\/metro-north.jpg","status_code":"OK","status_msg":"OK","local_id":"","result":{"tag":{"classes":["train","railroad","station","rail","transportation","platform","railway","speed","departure","travel","business","traffic","modern","subway","waiting","locomotive","public","blur","night","urban"],"catids":["169","1342","740","836","98","1791","835","349","3469","9","37","650","93","1448","1309","1535","859","422","190","180"],"probs":[0.99939858913422,0.99803149700165,0.99704265594482,0.99504208564758,0.995012819767,0.99491918087006,0.99164134263992,0.97328984737396,0.97118103504181,0.9688708782196,0.96783900260925,0.96487069129944,0.95819878578186,0.95712018013,0.95362317562103,0.9521244764328,0.9489620923996,0.94855046272278,0.94031262397766,0.93990349769592]}},"docid_str":"31fdb2316ff87fb5d747554ba5267313"}]}`
	feedbackResponse = `{"status_code":"OK","status_msg":"Feedback successfully recorded."}`
	urlExample       = "http://www.clarifai.com/img/metro-north.jpg"
)

func TestInfo(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	client := NewClient(ClientID, ClientSecret)
	client.setAPIRoot(server.URL)

	defer server.Close()

	mux.HandleFunc("/v1/token", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, tokenResponse)
	})

	mux.HandleFunc("/v1/info", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, infoResponse)
	})

	_, err := client.Info()

	if err != nil {
		t.Errorf("requestAccessToken() should not return an err upon success: %v", err)
	}
}

func TestTagMultiple(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	client := NewClient(ClientID, ClientSecret)
	client.setAPIRoot(server.URL)

	defer server.Close()

	mux.HandleFunc("/v1/token", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, tokenResponse)
	})

	mux.HandleFunc("/v1/tag", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, tagResponse)
	})

	urls := []string{urlExample, urlExample}
	_, err := client.Tag(TagRequest{URLs: urls})

	if err != nil {
		t.Errorf("Tag() should not return error with valid request: %q\n", err)
	}
}

func TestTagEncoded(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	client := NewClient(ClientID, ClientSecret)
	client.setAPIRoot(server.URL)

	defer server.Close()

	mux.HandleFunc("/v1/token", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, tokenResponse)
	})

	mux.HandleFunc("/v1/tag", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, tagResponse)
	})

	data := "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAAMSURBVBhXY/j//z8ABf4C/qc1gYQAAAAASUVORK5CYII="
	_, err := client.TagEncoded(TagEncodedRequest{EncodedData: data})

	if err != nil {
		t.Errorf("TagEncoded() should not return error with valid request: %q\n", err)
	}
}

func TestFeedback(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	client := NewClient(ClientID, ClientSecret)
	client.setAPIRoot(server.URL)

	defer server.Close()

	mux.HandleFunc("/v1/feedback", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, feedbackResponse)
	})

	mux.HandleFunc("/v1/tag", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, tagResponse)
	})

	feedback := FeedbackForm{
		URLs:    []string{urlExample},
		AddTags: []string{"good", "work"},
	}
	_, err := client.Feedback(feedback)

	if err != nil {
		t.Errorf("Feedback() should not return error with valid request: %q\n", err)
	}
}
