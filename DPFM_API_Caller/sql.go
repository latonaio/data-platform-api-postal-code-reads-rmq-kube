package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-postal-code-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-postal-code-reads-rmq-kube/DPFM_API_Output_Formatter"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var postalCode *[]dpfm_api_output_formatter.PostalCode
	var postalCodeAddress *[]dpfm_api_output_formatter.PostalCodeAddress
	for _, fn := range accepter {
		switch fn {
		case "PostalCode":
			func() {
				postalCode = c.PostalCode(mtx, input, output, errs, log)
			}()
		case "PostalCodeAddress":
			func() {
				postalCodeAddress = c.PostalCodeAddress(mtx, input, output, errs, log)
			}()
		}
	}

	data := &dpfm_api_output_formatter.Message{
		PostalCode:        postalCode,
		PostalCodeAddress: postalCodeAddress,
	}

	return data
}
