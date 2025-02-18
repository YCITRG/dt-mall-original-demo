package utility

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/dtm-labs/dtmcli/logger"
	"github.com/gin-gonic/gin/binding"
	"io"
	"net/http"
)

type UnifiedClient struct {
}

func (receiver *UnifiedClient) PostJson(url string, body map[string]interface{}) (*http.Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
}

type BizUnifiedClient struct {
	UnifiedClient
}

func (receiver *BizUnifiedClient) PostJSONAndBindJSON(url string, body map[string]any, obj any) error {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, binding.MIMEJSON, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		logger.Errorf("三方接口调用错误：%c", respBody)
		return errors.New("三方接口调用错误")
	}

	err = binding.JSON.BindBody(respBody, &obj)
	if err != nil {
		return err
	}
	return nil
}
