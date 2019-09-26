package ZUMServices

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

/*
ZSwrapper structure contains the
url, token and timeout info for
the ZUM Services
*/
type ZSwrapper struct {
	url     string
	Token   string
	Timeout int
}

func (service *ZSwrapper) check() error {
	service.url = "https://api.zum.services/v1"

	if service.Token == "" {
		return errors.New("All methods require an JWT access token. See https://zum.services/docs")
	}

	if service.Timeout == 0 {
		service.Timeout = 2000
	}

	return nil
}

// CreateAddress method creates a new ZUM address
func (service *ZSwrapper) CreateAddress() (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	data := url.Values{}

	response := service.makePostRequest("address", data)
	return response, nil
}

// DeleteAddress method deletes the specified address
func (service *ZSwrapper) DeleteAddress(address string) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeDeleteRequest("address/" + address)
	return response, nil
}

// GetAddress method gets the address details of the specified address
func (service *ZSwrapper) GetAddress(address string) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("address/" + address)

	return response, nil
}

// GetAddresses method views all addresses
// associated with the API token
func (service *ZSwrapper) GetAddresses() (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("address/all")
	return response, nil
}

// ScanAddress method scans an address for transactions between
// a 100 block range starting from the specified blockIndex.
func (service *ZSwrapper) ScanAddress(address string, blockIndex int) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("address/scan/" + address + "/" + strconv.Itoa(blockIndex))
	return response, nil
}

// GetAddressKeys method gets the public and
// secret spend keys of the specified address
func (service *ZSwrapper) GetAddressKeys(address string) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("address/keys/" + address)
	return response, nil
}

// IntegrateAddress method creates an integrated
// address with specified paymentID
func (service *ZSwrapper) IntegrateAddress(address string, paymentID string) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	data := url.Values{}
	data.Set("address", address)
	data.Set("paymentId", paymentID)

	response := service.makePostRequest("address/integrate", data)
	return response, nil
}

// GetIntegratedAddresses mthod returns all integrated
// address associated with the given normal address
func (service *ZSwrapper) GetIntegratedAddresses(address string) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("address/integrate/" + address)
	return response, nil
}

// GetFee method calculates the ZUM Services fee for
// an amount specified in ZUM with two decimal points.
func (service *ZSwrapper) GetFee(amount float64) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("transfer/fee/" + strconv.FormatFloat(amount, 'f', 2, 64))
	return response, nil
}

// CreateTransfer method sends a ZUM transaction with an
// address with the amount specified two decimal points.
func (service *ZSwrapper) CreateTransfer(
	fromAddress string,
	toAddress string,
	amount float64,
	fee float64,
	paymentID string,
	extra string) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	data := url.Values{}
	data.Set("from", fromAddress)
	data.Set("to", toAddress)
	data.Set("amount", strconv.FormatFloat(amount, 'f', 2, 64))
	data.Set("fee", strconv.FormatFloat(fee, 'f', 2, 64))
	if paymentID != "" {
		data.Set("paymentId", paymentID)
	}
	if extra != "" {
		data.Set("extra", extra)
	}

	response := service.makePostRequest("transfer", data)
	return response, nil
}

// GetTransfer method gets transaction details
// specified by transaction hash.
func (service *ZSwrapper) GetTransfer(transactionHash string) (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("transfer/" + transactionHash)
	return response, nil
}

// GetWallet method gets wallet container info and health check
func (service *ZSwrapper) GetWallet() (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("wallet")
	return response, nil
}

// GetStatus method gets the current status of the ZUM Services infrastructure
func (service *ZSwrapper) GetStatus() (*bytes.Buffer, error) {
	err := service.check()
	if err != nil {
		return nil, err
	}

	response := service.makeGetRequest("status")
	return response, nil
}

// Get Method
func (service *ZSwrapper) makeGetRequest(method string) *bytes.Buffer {
	url := service.url + "/" + method

	req, err := http.NewRequest("get", url, nil)
	if err != nil {
		println(err)
		return nil
	}

	req.Header.Add("Authorization", service.Token)
	return service.decodeResponse(req)
}

// Post Method
func (service *ZSwrapper) makePostRequest(method string, data url.Values) *bytes.Buffer {
	if method == "" {
		println("No method supplied.")
		return nil
	}

	url := service.url + "/" + method

	req, err := http.NewRequest("post", url, strings.NewReader(data.Encode()))

	if err != nil {
		println(err)
		return nil
	}

	req.Header.Add("Authorization", "Bearer "+service.Token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return service.decodeResponse(req)
}

// Delete Method
func (service *ZSwrapper) makeDeleteRequest(method string) *bytes.Buffer {
	if method == "" {
		println("No method supplied.")
		return nil
	}

	url := service.url + "/" + method

	req, err := http.NewRequest("delete", url, nil)
	if err != nil {
		println(err)
		return nil
	}

	req.Header.Add("Authorization", service.Token)
	return service.decodeResponse(req)
}

// Decode Response
func (service *ZSwrapper) decodeResponse(req *http.Request) *bytes.Buffer {
	client := &http.Client{}

	client.Timeout = time.Duration(service.Timeout) * time.Millisecond
	resp, err := client.Do(req)
	if err != nil {
		println(err)
		return nil
	}
	defer resp.Body.Close()

	responseBody, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		println(err1)
		return nil
	}

	return bytes.NewBuffer(responseBody)
}
