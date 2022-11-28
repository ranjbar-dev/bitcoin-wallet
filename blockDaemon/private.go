package blockDaemon

import (
	"encoding/json"
	"errors"
	"github.com/Amirilidan78/bitcoin-wallet/blockDaemon/response"
)

func (bd *blockDaemon) generateUrl(path string) string {

	return baseUrl + "/" + bd.conf.Protocol + "/" + bd.conf.Network + path
}

func (bd *blockDaemon) get(path string, result interface{}) error {

	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + bd.conf.Token

	httpResp, _, status, err := bd.hc.HttpGet(bd.generateUrl(path), headers)
	if err != nil {
		return err
	}

	if status == 400 {
		var errRes response.ErrorResponse
		err = json.Unmarshal(httpResp, &errRes)
		if err != nil {
			return err
		}
		return errors.New(errRes.Detail)
	}

	err = json.Unmarshal(httpResp, result)
	if err != nil {
		return err
	}

	return nil
}

func (bd *blockDaemon) post(path string, body interface{}, result interface{}) error {

	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + bd.conf.Token

	httpResp, _, status, err := bd.hc.HttpPost(bd.generateUrl(path), body, headers)
	if err != nil {
		return err
	}

	if status == 400 {
		var errRes response.ErrorResponse
		err = json.Unmarshal(httpResp, &errRes)
		if err != nil {
			return err
		}
		return errors.New(errRes.Detail)
	}

	err = json.Unmarshal(httpResp, result)
	if err != nil {
		return err
	}

	return nil
}
