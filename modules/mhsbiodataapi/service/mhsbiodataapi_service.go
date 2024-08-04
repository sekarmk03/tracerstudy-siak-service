package service

import (
	"encoding/json"
	"log"
	"os"

	"tracerstudy-siak-service/common/config"
	"tracerstudy-siak-service/modules/mhsbiodataapi/entity"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	// "google.golang.org/grpc/codes"
	// "google.golang.org/grpc/status"
)

/*const (
	apiMaxRetries = 3
	sleepTime     = 500 * time.Millisecond
)*/

type MhsBiodataApiService struct {
	cfg config.Config
}

type MhsBiodataApiServiceUseCase interface {
	FetchMhsBiodataByNimFromSiakApi(nim string) (*entity.MhsBiodataApi, error)
}

func NewMhsBiodataApiService(cfg config.Config) *MhsBiodataApiService {
	return &MhsBiodataApiService{
		cfg: cfg,
	}
}

func (svc *MhsBiodataApiService) FetchMhsBiodataByNimFromSiakApi(nim string) (*entity.MhsBiodataApi, error) {
	data, err := os.ReadFile("/app/mock_response.json")
	if err != nil {
		log.Println("ERROR: [MhsBiodataService - FetchMhsBiodataByNimFromSiakApi] Failed to read mock response file:", err)
	}

	var mhssbiodata []entity.MhsBiodataApi
	if err := json.Unmarshal(data, &mhssbiodata); err != nil {
		log.Println("ERROR: [MhsBiodataService - FetchMhsBiodataByNimFromSiakApi] Error while unmarshalling mock response body:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	var foundMhs *entity.MhsBiodataApi
	for _, mhsbiodata := range mhssbiodata {
		if mhsbiodata.NIM == nim {
			foundMhs = &mhsbiodata
			break
		}
	}

	if foundMhs == nil {
		log.Println("WARNING: [MhsBiodataService - FetchMhsBiodataByNimFromSiakApi] Resource not found: nim", nim)
		return nil, status.Errorf(codes.NotFound, "mhs resource not found")
	}

	return foundMhs, nil
	/*payload := map[string]string{"nim": nim}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Println("ERROR: [MhsBiodataService - FetchMhsBiodataByNimFromSiakApi] Error while marshalling payload:", err)
		if _, isUnsupportedTypeError := err.(*json.UnsupportedTypeError); isUnsupportedTypeError {
			return nil, status.Errorf(codes.InvalidArgument, "invalid payload: unsupported data type")
		}
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	apiUrl := svc.cfg.SIAK_API.URL
	apiKey := svc.cfg.SIAK_API.KEY

	for attempt := 1; attempt <= apiMaxRetries; attempt++ {
		reqHttp, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(payloadBytes))
		if err != nil {
			log.Println("ERROR: [MhsBiodataService - FetchMhsBiodataByNimFromSiakApi] Error while creating HTTP request:", err)
			return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
		}

		reqHttp.Header.Set("Api-Key", apiKey)
		reqHttp.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(reqHttp)
		if err != nil {
			log.Println("ERROR: [MhsBiodataService - FetchMhsBiodataByNimFromSiakApi] Error while sending HTTP request:", err)

			if attempt == apiMaxRetries {
				log.Println("ERROR: [MhsBiodataService - FetchMhsBiodataByNimFromSiakApi] Maximum retries reached:", err)
				return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
			}

			time.Sleep(sleepTime)
			continue
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			if attempt == apiMaxRetries {
				log.Println("ERROR: [MhsBiodataService - FetchMhsBiodataByNimFromSiakApi] Maximum retries reached:", resp.StatusCode, resp.Body)
				return nil, status.Errorf(codes.Internal, "internal server error: HTTP request failed with status code: %d", resp.StatusCode)
			}

			time.Sleep(sleepTime)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("ERROR: [MhsBiodataService - FetchMhsBiodataByNimFromSiakApi] Error while reading HTTP response body:", err)
			return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
		}

		var apiResponse []entity.MhsBiodataApi
		if err := json.Unmarshal(body, &apiResponse); err != nil {
			log.Println("ERROR: [MhsBiodataService - FetchMhsBiodataByNimFromSiakApi] Error while unmarshalling HTTP response body:", err)
			return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
		}

		if len(apiResponse) == 0 {
			log.Println("WARNING: [MhsBiodataService - FetchMhsBiodataByNimFromSiakApi] Resource not found: nim", nim)
			return nil, status.Errorf(codes.NotFound, "mhs resource not found")
		}

		return &apiResponse[0], nil
	}

	log.Println("ERROR: [MhsBiodataService - FetchMhsBiodataByNimFromSiakApi] Maximum retries reached without success")
	return nil, status.Errorf(codes.Internal, "internal server error: maximum retries reached without success")
	*/
}
