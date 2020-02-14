package plivo

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type MediaService struct {
	client *Client
	Media
}
type Media struct {
	ContentType string `json:"content_type,omitempty" url:"content_type,omitempty"`
	FileName    string `json:"file_name,omitempty" url:"file_name,omitempty"`
	MediaID     string `json:"media_id,omitempty" url:"media_id,omitempty"`
	Size        int    `json:"size,omitempty" url:"size,omitempty"`
	UploadTime  string `json:"upload_time,omitempty" url:"upload_time,omitempty"`
	URL         string `json:"url,omitempty" url:"url,omitempty"`
	Status      string `json:"status,omitempty" url:"status,omitempty"`
	StatusCode  string `json:"status_code,omitempty" url:"status_code,omitempty"`
}

type MediaMeta struct {
	Previous   *string
	Next       *string
	TotalCount int `json:"total_count" url:"api_id"`
	Offset     int `json:"offset,omitempty" url:"offset,omitempty"`
	Limit      int `json:"limit,omitempty" url:"limit,omitempty"`
}
type MediaResponseBody struct {
	Media []string `json:"objects" url:"objcts"`
	ApiID string   `json:"api_id" url:"api_id"`
}

type BaseListMediaResponse struct {
	ApiID string    `json:"api_id" url:"api_id"`
	Meta  MediaMeta `json:"meta" url:"meta"`
	Media []string  `json:"objects" url:"objcts"`
}

type MediaUpload struct {
	FileName []string
}

type MediaListParams struct {
	Limit  int `url:"limit,omitempty"`
	Offset int `url:"offset,omitempty"`
}

func (service *MediaService) Upload(params MediaUpload) (response *MediaResponseBody, err error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	for i := 0; i < len(params.FileName); i++ {
		file, errFile1 := os.Open(params.FileName[i])
		defer file.Close()
		part1, errFile1 := writer.CreateFormFile("file", filepath.Base(params.FileName[i]))
		_, errFile1 = io.Copy(part1, file)
		if errFile1 != nil {
			return nil, errFile1
		}
		filerror := writer.Close()
		if filerror != nil {
			return nil, filerror
		}
	}
	requestUrl := service.client.BaseUrl
	requestUrl.Path = fmt.Sprintf(baseRequestString, fmt.Sprintf(service.client.AuthId+"/Media"))
	request, err := http.NewRequest("POST", requestUrl.String(), payload)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.SetBasicAuth(service.client.AuthId, service.client.AuthToken)
	err = service.client.ExecuteRequest(request, response)
	return
}

func mustOpen(f string) *os.File {
	r, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	return r
}

func (service *MediaService) Get(media_id string) (response *Media, err error) {
	req, err := service.client.NewRequest("GET", nil, "Media/%s", media_id)
	if err != nil {
		return
	}
	resp := &Media{}
	err = service.client.ExecuteRequest(req, resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	return resp, nil
}

func (service *MediaService) List(param MediaListParams) (response *BaseListMediaResponse, err error) {
	req, err := service.client.NewRequest("GET", param, "Media")
	if err != nil {
		return
	}
	resp := &BaseListMediaResponse{}
	err = service.client.ExecuteRequest(req, resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	return resp, nil
}
