package goalb

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/dwdwow/golimiter"
)

const META_SLEUTH_REQ_PER_SECOND = 5
const META_SLEUTH_REQ_PER_DAY = 10000

type MetaSleuthRespErrCode int

const (
	MetaSleuthSuccess               MetaSleuthRespErrCode = 200000
	MetaSleuthUnauthorizedOperation MetaSleuthRespErrCode = 400001
	MetaSleuthRateLimitExceeded     MetaSleuthRespErrCode = 400002
	MetaSleuthInvalidParams         MetaSleuthRespErrCode = 400004
	MetaSleuthUserNotExist          MetaSleuthRespErrCode = 400005
	MetaSleuthServerBusy            MetaSleuthRespErrCode = 400006
	MetaSleuthInvalidAPIKey         MetaSleuthRespErrCode = 400007
	MetaSleuthInvalidAuthFormat     MetaSleuthRespErrCode = 400008
	MetaSleuthExpiredAPIKey         MetaSleuthRespErrCode = 400009
	MetaSleuthNotFound              MetaSleuthRespErrCode = 400010
	MetaSleuthInvalidAddress        MetaSleuthRespErrCode = 400011
	MetaSleuthDailyLimitExceeded    MetaSleuthRespErrCode = 400012
	MetaSleuthUnsupportedChain      MetaSleuthRespErrCode = 400013
	MetaSleuthInternalError         MetaSleuthRespErrCode = 500000
)

type MetaSleuthResp[D any] struct {
	RequestId string                `json:"request_id"`
	Code      MetaSleuthRespErrCode `json:"code"`
	Message   string                `json:"message"`
	Data      D                     `json:"data"`
}

type MetaSleuthSupportedChain struct {
	ChainId   int    `json:"chain_id"`
	ChainName string `json:"chain_name"`
}

type MetaSleuthEntityCategory struct {
	Name string `json:"name"`
	Code int    `json:"code"`
}

type MetaSleuthEntityAttribute struct {
	Name     string   `json:"name"`
	Code     int      `json:"code"`
	CompInfo []string `json:"comp_info"`
}

type MetaSleuthEntityDescription struct {
	Attributes []MetaSleuthEntityAttribute `json:"attributes"`
	Website    string                      `json:"website"`
	Twitter    string                      `json:"twitter"`
	Telegram   string                      `json:"telegram"`
	Discord    string                      `json:"discord"`
}

type MetaSleuthEntity struct {
	Entity      string                      `json:"entity"`
	Categories  []MetaSleuthEntityCategory  `json:"categories"`
	Attributes  []MetaSleuthEntityAttribute `json:"attributes"`
	Description MetaSleuthEntityDescription `json:"description"`
}

type MetaSleuthAddressLabel struct {
	ChainId        int                         `json:"chain_id"`
	Address        string                      `json:"address"`
	MainEntity     string                      `json:"main_entity"`
	MainEntityInfo MetaSleuthEntity            `json:"main_entity_info"`
	CompEntities   []string                    `json:"comp_entities"`
	Attributes     []MetaSleuthEntityAttribute `json:"attributes"`
	NameTag        string                      `json:"name_tag"`
}

type MetaSleuthOneChainSeveralAddressesLabel struct {
	ChainId   int                      `json:"chain_id"`
	Addresses []MetaSleuthAddressLabel `json:"addresses"`
}

var MetaSleuthSupportedChains = []MetaSleuthSupportedChain{
	{ChainId: -3, ChainName: "SOLANA"},
	{ChainId: -2, ChainName: "TRON"},
	{ChainId: -1, ChainName: "BITCOIN"},
	{ChainId: 1, ChainName: "ETHEREUM"},
	{ChainId: 10, ChainName: "OPTIMISM"},
	{ChainId: 25, ChainName: "CRONOS"},
	{ChainId: 56, ChainName: "BNB SMART CHAIN"},
	{ChainId: 100, ChainName: "GNOSIS"},
	{ChainId: 137, ChainName: "POLYGON"},
	{ChainId: 169, ChainName: "MANTA PACIFIC"},
	{ChainId: 199, ChainName: "BITTORRENT"},
	{ChainId: 250, ChainName: "FANTOM OPERA"},
	{ChainId: 288, ChainName: "BOBA"},
	{ChainId: 324, ChainName: "ZKSYNC"},
	{ChainId: 1024, ChainName: "CLV PARACHAIN"},
	{ChainId: 1101, ChainName: "POLYGON ZKEVM"},
	{ChainId: 1111, ChainName: "WEMIX3.0 MAINNET"},
	{ChainId: 1284, ChainName: "MOONBEAM"},
	{ChainId: 1285, ChainName: "MOONRIVER"},
	{ChainId: 5000, ChainName: "MANTLE"},
	{ChainId: 8453, ChainName: "BASE"},
	{ChainId: 42161, ChainName: "ARBITRUM ONE"},
	{ChainId: 42220, ChainName: "CELO"},
	{ChainId: 43114, ChainName: "AVALANCHE C-CHAIN"},
	{ChainId: 59144, ChainName: "LINEA"},
	{ChainId: 81457, ChainName: "BLAST"},
	{ChainId: 1313161554, ChainName: "AURORA"},
}

var MetaSleuthChainId = map[ChainName]int{
	"SOLANA":            -3,
	"TRON":              -2,
	"BITCOIN":           -1,
	"ETHEREUM":          1,
	"OPTIMISM":          10,
	"CRONOS":            25,
	"BNB SMART CHAIN":   56,
	"GNOSIS":            100,
	"POLYGON":           137,
	"MANTA PACIFIC":     169,
	"BITTORRENT":        199,
	"FANTOM OPERA":      250,
	"BOBA":              288,
	"ZKSYNC":            324,
	"CLV PARACHAIN":     1024,
	"POLYGON ZKEVM":     1101,
	"WEMIX3.0 MAINNET":  1111,
	"MOONBEAM":          1284,
	"MOONRIVER":         1285,
	"MANTLE":            5000,
	"BASE":              8453,
	"ARBITRUM ONE":      42161,
	"CELO":              42220,
	"AVALANCHE C-CHAIN": 43114,
	"LINEA":             59144,
	"BLAST":             81457,
	"AURORA":            1313161554,
}

type ChainName string

const (
	ChainNameSolana          ChainName = "SOLANA"
	ChainNameTron            ChainName = "TRON"
	ChainNameBitcoin         ChainName = "BITCOIN"
	ChainNameEthereum        ChainName = "ETHEREUM"
	ChainNameOptimism        ChainName = "OPTIMISM"
	ChainNameCronos          ChainName = "CRONOS"
	ChainNameBNBSmartChain   ChainName = "BNB SMART CHAIN"
	ChainNameGnosis          ChainName = "GNOSIS"
	ChainNamePolygon         ChainName = "POLYGON"
	ChainNameMantaPacific    ChainName = "MANTA PACIFIC"
	ChainNameBittorrent      ChainName = "BITTORRENT"
	ChainNameFantomOpera     ChainName = "FANTOM OPERA"
	ChainNameBoba            ChainName = "BOBA"
	ChainNameZksync          ChainName = "ZKSYNC"
	ChainNameClvParachain    ChainName = "CLV PARACHAIN"
	ChainNamePolygonZkevm    ChainName = "POLYGON ZKEVM"
	ChainNameWemix           ChainName = "WEMIX3.0 MAINNET"
	ChainNameMoonbeam        ChainName = "MOONBEAM"
	ChainNameMoonriver       ChainName = "MOONRIVER"
	ChainNameMantle          ChainName = "MANTLE"
	ChainNameBase            ChainName = "BASE"
	ChainNameArbitrumOne     ChainName = "ARBITRUM ONE"
	ChainNameCelo            ChainName = "CELO"
	ChainNameAvalancheCChain ChainName = "AVALANCHE C-CHAIN"
	ChainNameLinea           ChainName = "LINEA"
	ChainNameBlast           ChainName = "BLAST"
	ChainNameAurora          ChainName = "AURORA"
)

type MetaSleuthClient struct {
	Header  http.Header
	Limiter *golimiter.ReqLimiter
	logger  *slog.Logger
}

func NewMetaSleuthClient(apiKey string, limiter *golimiter.ReqLimiter, logger *slog.Logger) *MetaSleuthClient {
	if apiKey == "" {
		apiKey = os.Getenv("META_SLEUTH_API_KEY")
		if apiKey == "" {
			panic("META_SLEUTH_API_KEY is not set")
		}
	}

	if logger == nil {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}

	logger = logger.With("logger", "meta_sleuth")

	if limiter == nil {
		limiter = golimiter.NewReqLimiter(time.Second, META_SLEUTH_REQ_PER_SECOND)
	}

	header := http.Header{}
	header.Add("API-KEY", apiKey)
	header.Add("Content-Type", "application/json")

	return &MetaSleuthClient{
		Header:  header,
		Limiter: limiter,
		logger:  logger,
	}
}

func NewDefaultMetaSleuthClient() *MetaSleuthClient {
	return NewMetaSleuthClient("", nil, nil)
}

func metaSleuthClientReq[D any](c *MetaSleuthClient, method string, path string, body map[string]any) (D, error) {
	c.Limiter.Wait(context.Background())
	ul := "https://aml.blocksec.com/address-label/api/v3" + path
	var bodyBytes []byte
	var err error
	if body != nil {
		bodyBytes, err = json.Marshal(body)
		if err != nil {
			return *new(D), fmt.Errorf("metasleuth: marshal body failed: %w", err)
		}
	}
	req, err := http.NewRequest(method, ul, bytes.NewReader(bodyBytes))
	if err != nil {
		return *new(D), fmt.Errorf("metasleuth: create request failed: %w", err)
	}
	req.Header = c.Header
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return *new(D), fmt.Errorf("metasleuth: request failed: %w", err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var respData MetaSleuthResp[D]
	if err := decoder.Decode(&respData); err != nil {
		return *new(D), fmt.Errorf("metasleuth: decode response failed: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return *new(D), fmt.Errorf("metasleuth: status code: %d, error code: %d, message: %s", resp.StatusCode, respData.Code, respData.Message)
	}
	return respData.Data, nil
}

func (c *MetaSleuthClient) GetSupportedChains() ([]MetaSleuthSupportedChain, error) {
	return metaSleuthClientReq[[]MetaSleuthSupportedChain](c, http.MethodGet, "/chain-list", nil)
}

func (c *MetaSleuthClient) GetAddressLabels(chainId int, address string) (MetaSleuthAddressLabel, error) {
	return metaSleuthClientReq[MetaSleuthAddressLabel](c, http.MethodPost, "/labels", map[string]any{
		"chain_id": chainId,
		"address":  address,
	})
}

func (c *MetaSleuthClient) GetBatchLabels(chainId int, addresses ...string) ([]MetaSleuthAddressLabel, error) {
	return metaSleuthClientReq[[]MetaSleuthAddressLabel](c, http.MethodPost, "/batch-labels", map[string]any{
		"chain_id":  chainId,
		"addresses": addresses,
	})
}

func (c *MetaSleuthClient) GetBatchLabelsMultiChain(chainIds []int, addresses []string) ([]MetaSleuthOneChainSeveralAddressesLabel, error) {
	return metaSleuthClientReq[[]MetaSleuthOneChainSeveralAddressesLabel](c, http.MethodPost, "/multi-chains-labels", map[string]any{
		"chain_ids": chainIds,
		"addresses": addresses,
	})
}

func (c *MetaSleuthClient) GetEntity(entityName string) (MetaSleuthEntity, error) {
	return metaSleuthClientReq[MetaSleuthEntity](c, http.MethodPost, "/entity", map[string]any{
		"entity": entityName,
	})
}
