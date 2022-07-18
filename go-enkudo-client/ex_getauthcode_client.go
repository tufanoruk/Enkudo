/** ex_getauthcode_client.go

% ln -s ../go_src openapiclinet
% go mod init go-enkudo-client
% go mod tidy
% go build

**/

package main

import (
	"context"
	"fmt"
	"os"

	openapiclient "enkudo-dsp-client/openapiclient"
)

func main() {
	authCodeReq := *openapiclient.NewAuthCodeReq("Target_example") // AuthCodeReq |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultEnkudoDSP.GetAuthCode(context.Background()).AuthCodeReq(authCodeReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultEnkudoDSP.GetAuthCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthCode`: AuthCodeRsp
	fmt.Fprintf(os.Stdout, "Response from `DefaultEnkudoDSP.GetAuthCode`: %v\n", resp)
}
