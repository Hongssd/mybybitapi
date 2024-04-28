package mybybitapi

import (
	"bytes"
	"compress/gzip"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"sync"
	"time"

	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

const (
	BIT_BASE_10 = 10
	BIT_SIZE_64 = 64
	BIT_SIZE_32 = 32
)

type RequestType string

const (
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
	PUT    = "PUT"
)

var NIL_REQBODY = []byte{}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var log = logrus.New()

func SetLogger(logger *logrus.Logger) {
	log = logger
}

func GetPointer[T any](v T) *T {
	return &v
}

func HmacSha256(secret, data string) []byte {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return h.Sum(nil)
}

// Request 发送请求
func Request(url string, reqBody []byte, method string, isGzip bool) ([]byte, error) {
	return RequestWithHeader(url, reqBody, method, map[string]string{}, isGzip)
}

func RequestWithHeader(url string, reqBody []byte, method string, headerMap map[string]string, isGzip bool) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range headerMap {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	if isGzip { // 请求 header 添加 gzip
		req.Header.Add("Content-Encoding", "gzip")
		req.Header.Add("Accept-Encoding", "gzip")
	}
	req.Close = true
	req.Body = io.NopCloser(bytes.NewBuffer(reqBody))

	log.Debug("reqURL: ", req.URL.String())
	if reqBody != nil && len(reqBody) > 0 {
		log.Debug("reqBody: ", string(reqBody))
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body := resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		body, err = gzip.NewReader(resp.Body)
		if err != nil {
			log.Error(err)
			return nil, err
		}
	}
	data, err := io.ReadAll(body)
	// log.Debug(string(data))
	return data, err
}

type MyBybit struct {
}

const (
	BYBIT_API_HTTP           = "api.bybit.com"
	BYBIT_API_WEBSOCKET      = "stream.bybit.com"
	BYBIT_API_HTTP_AWS       = "api.bytick.com"
	BYBIT_API_WEBSOCKET_AWS  = "stream.bytick.com"
	BYBIT_API_HTTP_TEST      = "api-testnet.bybit.com"
	BYBIT_API_WEBSOCKET_TEST = "stream-testnet.bybit.com"
	IS_GZIP                  = true
)

type ServerType int

const (
	BASE ServerType = iota
	AWS
	TEST
)

var SERVER_TYPE = BASE

func SetServerType(serverType ServerType) {
	SERVER_TYPE = serverType
}

type APIType int

const (
	REST APIType = iota
	WS_PUBLIC
	WS_PRIVATE
	WS_BUSINESS
)

// spot 現貨
// linear USDT永續, USDC永續, USDC交割
// inverse 反向合約，包含反向永續, 反向交割
// option 期權
type Category string

const (
	SPOT    Category = "spot"
	LINEAR  Category = "linear"
	INVERSE Category = "inverse"
	OPTION  Category = "option"
)

func (c Category) String() string {
	return string(c)
}

type Client struct {
	APIKey     string
	SecretKey  string
	Referer    string
	RecvWindow string
}

type RestClient struct {
	c *Client
}

type PublicRestClient RestClient

type PrivateRestClient RestClient

type RestClientOption func(c *Client)

func WithRefer(referer string) RestClientOption {
	return func(c *Client) {
		c.Referer = referer
	}
}

func WithRecvWindow(recvWindow int64) RestClientOption {
	return func(c *Client) {
		c.RecvWindow = strconv.FormatInt(recvWindow, BIT_BASE_10)
	}
}

func (*MyBybit) NewRestClient(APIKey, SecretKey string, options ...RestClientOption) *RestClient {
	client := &RestClient{
		c: &Client{
			APIKey:    APIKey,
			SecretKey: SecretKey,
		},
	}
	for _, option := range options {
		option(client.c)
	}
	if client.c.RecvWindow == "" {
		client.c.RecvWindow = "5000"
	}
	return client
}

func (c *RestClient) PublicRestClient() *PublicRestClient {
	return &PublicRestClient{
		c: c.c,
	}
}

func (c *RestClient) PrivateRestClient() *PrivateRestClient {
	return &PrivateRestClient{
		c: c.c,
	}
}

// 通用接口调用
func bybitCallAPI[T any](client *Client, url url.URL, reqBody []byte, method string) (*BybitRestRes[T], error) {
	body, err := Request(url.String(), reqBody, method, IS_GZIP)
	if err != nil {
		return nil, err
	}
	res, err := handlerCommonRest[T](body)
	if err != nil {
		return nil, err
	}
	return res, res.handlerError()
}

// 通用鉴权接口调用
func bybitCallAPIWithSecret[T any](client *Client, url url.URL, reqBody []byte, method string) (*BybitRestRes[T], error) {

	timestamp := strconv.FormatInt(time.Now().UnixMilli(), BIT_BASE_10)
	//requestPath := url.Path
	query := url.RawQuery

	hmacSha256Data := timestamp + client.APIKey + client.RecvWindow
	if query != "" {
		hmacSha256Data += query
	}
	if len(reqBody) != 0 {
		hmacSha256Data += string(reqBody)
	}
	sign := base64.StdEncoding.EncodeToString(HmacSha256(client.SecretKey, hmacSha256Data))

	log.Warn("bybit timestamp: ", timestamp)
	log.Warn("bybit method: ", method)
	log.Warn("bybit query: ", query)
	log.Warn("bybit reqBody: ", string(reqBody))
	log.Warn("bybit hmacSha256Data: ", hmacSha256Data)
	log.Warn("bybit sign: ", sign)

	body, err := RequestWithHeader(url.String(), reqBody, method,
		map[string]string{
			"X-BAPI-API-KEY":     client.APIKey,
			"X-BAPI-SIGN":        sign,
			"X-BAPI-TIMESTAMP":   timestamp,
			"X-Referer":          client.Referer,
			"X-BAPI-RECV-WINDOW": client.RecvWindow,
		}, IS_GZIP)
	if err != nil {
		return nil, err
	}
	res, err := handlerCommonRest[T](body)
	if err != nil {
		return nil, err
	}
	return res, res.handlerError()
}

// URL标准封装 带路径参数
func bybitHandlerRequestAPIWithPathQueryParam[T any](apiType APIType, request *T, name string) url.URL {
	query := bybitHandlerReq(request)
	u := url.URL{
		Scheme:   "https",
		Host:     BybitGetRestHostByAPIType(apiType),
		Path:     name,
		RawQuery: query,
	}
	return u
}

// URL标准封装 不带路径参数
func bybitHandlerRequestAPIWithoutPathQueryParam(apiType APIType, name string) url.URL {
	// query := bybitHandlerReq(request)
	u := url.URL{
		Scheme:   "https",
		Host:     BybitGetRestHostByAPIType(apiType),
		Path:     name,
		RawQuery: "",
	}
	return u
}

func bybitHandlerReq[T any](req *T) string {
	var argBuffer bytes.Buffer

	t := reflect.TypeOf(req)
	v := reflect.ValueOf(req)
	if v.IsNil() {
		return ""
	}
	t = t.Elem()
	v = v.Elem()
	count := v.NumField()
	for i := 0; i < count; i++ {
		argName := t.Field(i).Tag.Get("json")
		switch v.Field(i).Elem().Kind() {
		case reflect.String:
			argBuffer.WriteString(argName + "=" + v.Field(i).Elem().String() + "&")
		case reflect.Int, reflect.Int64:
			argBuffer.WriteString(argName + "=" + strconv.FormatInt(v.Field(i).Elem().Int(), BIT_BASE_10) + "&")
		case reflect.Float32, reflect.Float64:
			argBuffer.WriteString(argName + "=" + decimal.NewFromFloat(v.Field(i).Elem().Float()).String() + "&")
		case reflect.Bool:
			argBuffer.WriteString(argName + "=" + strconv.FormatBool(v.Field(i).Elem().Bool()) + "&")
		case reflect.Struct:
			sv := reflect.ValueOf(v.Field(i).Interface())
			ToStringMethod := sv.MethodByName("String")
			args := make([]reflect.Value, 0)
			result := ToStringMethod.Call(args)
			argBuffer.WriteString(argName + "=" + result[0].String() + "&")
		case reflect.Slice:
			s := v.Field(i).Interface()
			d, _ := json.Marshal(s)
			argBuffer.WriteString(argName + "=" + url.QueryEscape(string(d)) + "&")
		case reflect.Invalid:
		default:
			log.Errorf("req type error %s:%s", argName, v.Field(i).Elem().Kind())
		}
	}
	return strings.TrimRight(argBuffer.String(), "&")
}

func BybitGetRestHostByAPIType(apiType APIType) string {

	switch apiType {
	case REST:
		switch SERVER_TYPE {
		case BASE:
			return BYBIT_API_HTTP
		case AWS:
			return BYBIT_API_HTTP_AWS
		case TEST:
			return BYBIT_API_HTTP_TEST
		}
	default:
	}
	return ""
}

func interfaceStringToFloat64(inter interface{}) float64 {
	return stringToFloat64(inter.(string))
}

func interfaceStringToInt64(inter interface{}) int64 {
	return int64(inter.(float64))
}

func stringToFloat64(str string) float64 {
	f, _ := strconv.ParseFloat(str, BIT_SIZE_64)
	return f
}

type MySyncMap[K any, V any] struct {
	smap sync.Map
}

func NewMySyncMap[K any, V any]() MySyncMap[K, V] {
	return MySyncMap[K, V]{
		smap: sync.Map{},
	}
}
func (m *MySyncMap[K, V]) Load(k K) (V, bool) {
	v, ok := m.smap.Load(k)

	if ok {
		return v.(V), true
	}
	var resv V
	return resv, false
}
func (m *MySyncMap[K, V]) Store(k K, v V) {
	m.smap.Store(k, v)
}

func (m *MySyncMap[K, V]) Delete(k K) {
	m.smap.Delete(k)
}
func (m *MySyncMap[K, V]) Range(f func(k K, v V) bool) {
	m.smap.Range(func(k, v any) bool {
		return f(k.(K), v.(V))
	})
}

func (m *MySyncMap[K, V]) Length() int {
	length := 0
	m.Range(func(k K, v V) bool {
		length += 1
		return true
	})
	return length
}

func (m *MySyncMap[K, V]) MapValues(f func(k K, v V) V) *MySyncMap[K, V] {
	var res = NewMySyncMap[K, V]()
	m.Range(func(k K, v V) bool {
		res.Store(k, f(k, v))
		return true
	})
	return &res
}
