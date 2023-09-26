package dpfm_api_output_formatter

import (
	"data-platform-api-postal-code-reads-rmq-kube/DPFM_API_Caller/requests"
	api_input_reader "data-platform-api-postal-code-reads-rmq-kube/DPFM_API_Input_Reader"
	"database/sql"
	"fmt"
)

func ConvertToPostalCode(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]PostalCode, error) {
	defer rows.Close()
	postalcodes := make([]PostalCode, 0, len(sdc.PostalCode))

	i := 0
	for rows.Next() {
		i++
		pm := &requests.PostalCode{}

		err := rows.Scan(
			&pm.PostalCode,
			&pm.Country,
			&pm.LocalSubRegion,
			&pm.LocalRegion,
			&pm.GlobalRegion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &postalcodes, err
		}

		data := pm

		postalcodes = append(postalcodes, PostalCode{
			PostalCode:     data.PostalCode,
			Country:        data.Country,
			LocalSubRegion: data.LocalSubRegion,
			LocalRegion:    data.LocalRegion,
			GlobalRegion:   data.GlobalRegion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &postalcodes, nil
	}

	return &postalcodes, nil
}
