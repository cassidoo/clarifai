package clarifai

import (
	"encoding/json"
	"errors"
	"math/big"
)

// InfoResp represents the expected JSON response from /info/
type InfoResp struct {
	StatusCode    string `json:"status_code"`
	StatusMessage string `json:"status_msg"`
	Results       struct {
		MaxImageSize      int     `json:"max_image_size"`
		DefaultLanguage   string  `json:"default_language"`
		MaxVideoSize      int     `json:"max_video_size"`
		MaxImageBytes     int     `json:"max_image_bytes"`
		DefaultModel      string  `json:"default_model"`
		MaxVideoBytes     int     `json:"max_video_bytes"`
		MaxVideoDuration  int     `json:"max_video_duration"`
		MaxVideoBatchSize int     `json:"max_video_batch_size"`
		MinVideoSize      int     `json:"min_video_size"`
		MinImageSize      int     `json:"min_image_size"`
		MaxBatchSize      int     `json:"max_batch_size"`
		APIVersion        float32 `json:"api_version"`
	}
}

// TagRequest represents a JSON request for /tag/
type TagRequest struct {
	URLs     []string `json:"url"`
	LocalIDs []string `json:"local_ids,omitempty"`
	Model    string   `json:"model,omitempty"`
}

type BaseTagResp struct {
	StatusCode    string `json:"status_code"`
	StatusMessage string `json:"status_msg"`
	Meta          struct {
		Tag struct {
			Timestamp json.Number `json:"timestamp"`
			Model     string      `json:"model"`
			Config    string      `json:"config"`
		}
	}
}

// ImageTagResp represents the expected JSON response from /tag/ for images
type ImageTagResp struct {
	BaseTagResp
	Results []ImageTagResult
}

// ImageTagResp represents the expected JSON response from /tag/ for images
type VideoTagResp struct {
	BaseTagResp
	Results []VideoTagResult
}

// TagResult represents the expected data for a single tag result
type BaseTagResult struct {
	DocID         *big.Int `json:"docid"`
	URL           string   `json:"url"`
	StatusCode    string   `json:"status_code"`
	StatusMessage string   `json:"status_msg"`
	LocalID       string   `json:"local_id"`
	DocIDString   string   `json:"docid_str"`
}

// TagResult represents the expected data for a single tag result
type ImageTagResult struct {
	BaseTagResult
	Result struct {
		Tag struct {
			Classes []string  `json:"classes"`
			CatIDs  []string  `json:"catids"`
			Probs   []float32 `json:"probs"`
		}
	}
}

// TagResult represents the expected data for a single tag result
type VideoTagResult struct {
	BaseTagResult
	Result struct {
		Tag struct {
			Classes [][]string  `json:"classes"`
			CatIDs  []string    `json:"catids"`
			Probs   [][]float32 `json:"probs"`
		}
	}
}

// FeedbackForm is used to send feedback back to Clarifai
type FeedbackForm struct {
	DocIDs           []string `json:"docids,omitempty"`
	URLs             []string `json:"url,omitempty"`
	AddTags          []string `json:"add_tags,omitempty"`
	RemoveTags       []string `json:"remove_tags,omitempty"`
	DissimilarDocIDs []string `json:"dissimilar_docids,omitempty"`
	SimilarDocIDs    []string `json:"similar_docids,omitempty"`
	SearchClick      []string `json:"search_click,omitempty"`
}

// FeedbackResp is the expected response from /feedback/
type FeedbackResp struct {
	StatusCode    string `json:"status_code"`
	StatusMessage string `json:"status_msg"`
}

// Info will return the current status info for the given client
func (client *Client) Info() (*InfoResp, error) {
	res, err := client.commonHTTPRequest(nil, "info", "GET", false)

	if err != nil {
		return nil, err
	}

	info := new(InfoResp)
	err = json.Unmarshal(res, info)

	return info, err
}

// Tag allows the client to request tag data on a single, or multiple photos
func (client *Client) Tag(req TagRequest) (*ImageTagResp, error) {
	return client.ImageTag(req)
}

// ImageTag allows the client to request tag data on a single, or multiple photos
func (client *Client) ImageTag(req TagRequest) (*ImageTagResp, error) {
	if len(req.URLs) < 1 {
		return nil, errors.New("Requires at least one url")
	}

	res, err := client.commonHTTPRequest(req, "tag", "POST", false)

	if err != nil {
		return nil, err
	}

	tagres := new(ImageTagResp)
	err = json.Unmarshal(res, tagres)

	return tagres, err
}

// VideoTag allows the client to request tag data on a single, or multiple videos
func (client *Client) VideoTag(req TagRequest) (*VideoTagResp, error) {
	if len(req.URLs) < 1 {
		return nil, errors.New("Requires at least one url")
	}

	res, err := client.commonHTTPRequest(req, "tag", "POST", false)

	if err != nil {
		return nil, err
	}

	tagres := new(VideoTagResp)
	err = json.Unmarshal(res, tagres)

	return tagres, err
}

// Feedback allows the user to provide contextual feedback to Clarifai in order to improve their results
func (client *Client) Feedback(form FeedbackForm) (*FeedbackResp, error) {
	if form.DocIDs == nil && form.URLs == nil {
		return nil, errors.New("Requires at least one docid or url")
	}

	if form.DocIDs != nil && form.URLs != nil {
		return nil, errors.New("Request must provide exactly one of the following fields: {'DocIDs', 'URLs'}")
	}

	res, err := client.commonHTTPRequest(form, "feedback", "POST", false)

	feedbackres := new(FeedbackResp)
	err = json.Unmarshal(res, feedbackres)

	return feedbackres, err

}
