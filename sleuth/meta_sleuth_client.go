package sleuth

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

type ChainID int64

const (
	ChainIDSolana          ChainID = -3
	ChainIDTron            ChainID = -2
	ChainIDBitcoin         ChainID = -1
	ChainIDEthereum        ChainID = 1
	ChainIDOptimism        ChainID = 10
	ChainIDCronos          ChainID = 25
	ChainIDBNBSmartChain   ChainID = 56
	ChainIDGnosis          ChainID = 100
	ChainIDPolygon         ChainID = 137
	ChainIDMantaPacific    ChainID = 169
	ChainIDBittorrent      ChainID = 199
	ChainIDFantomOpera     ChainID = 250
	ChainIDBoba            ChainID = 288
	ChainIDZksync          ChainID = 324
	ChainIDClvParachain    ChainID = 1024
	ChainIDPolygonZkevm    ChainID = 1101
	ChainIDWemix           ChainID = 1111
	ChainIDMoonbeam        ChainID = 1284
	ChainIDMoonriver       ChainID = 1285
	ChainIDMantle          ChainID = 5000
	ChainIDBase            ChainID = 8453
	ChainIDArbitrumOne     ChainID = 42161
	ChainIDCelo            ChainID = 42220
	ChainIDAvalancheCChain ChainID = 43114
	ChainIDLinea           ChainID = 59144
	ChainIDBlast           ChainID = 81457
	ChainIDAurora          ChainID = 1313161554
)

var ChainIdToName = map[ChainID]ChainName{
	ChainIDSolana:          ChainNameSolana,
	ChainIDTron:            ChainNameTron,
	ChainIDBitcoin:         ChainNameBitcoin,
	ChainIDEthereum:        ChainNameEthereum,
	ChainIDOptimism:        ChainNameOptimism,
	ChainIDCronos:          ChainNameCronos,
	ChainIDBNBSmartChain:   ChainNameBNBSmartChain,
	ChainIDGnosis:          ChainNameGnosis,
	ChainIDPolygon:         ChainNamePolygon,
	ChainIDMantaPacific:    ChainNameMantaPacific,
	ChainIDBittorrent:      ChainNameBittorrent,
	ChainIDFantomOpera:     ChainNameFantomOpera,
	ChainIDBoba:            ChainNameBoba,
	ChainIDZksync:          ChainNameZksync,
	ChainIDClvParachain:    ChainNameClvParachain,
	ChainIDPolygonZkevm:    ChainNamePolygonZkevm,
	ChainIDWemix:           ChainNameWemix,
	ChainIDMoonbeam:        ChainNameMoonbeam,
	ChainIDMoonriver:       ChainNameMoonriver,
	ChainIDMantle:          ChainNameMantle,
	ChainIDBase:            ChainNameBase,
	ChainIDArbitrumOne:     ChainNameArbitrumOne,
	ChainIDCelo:            ChainNameCelo,
	ChainIDAvalancheCChain: ChainNameAvalancheCChain,
	ChainIDLinea:           ChainNameLinea,
	ChainIDBlast:           ChainNameBlast,
	ChainIDAurora:          ChainNameAurora,
}

var ChainNameToId = map[ChainName]ChainID{
	ChainNameSolana:          ChainIDSolana,
	ChainNameTron:            ChainIDTron,
	ChainNameBitcoin:         ChainIDBitcoin,
	ChainNameEthereum:        ChainIDEthereum,
	ChainNameOptimism:        ChainIDOptimism,
	ChainNameCronos:          ChainIDCronos,
	ChainNameBNBSmartChain:   ChainIDBNBSmartChain,
	ChainNameGnosis:          ChainIDGnosis,
	ChainNamePolygon:         ChainIDPolygon,
	ChainNameMantaPacific:    ChainIDMantaPacific,
	ChainNameBittorrent:      ChainIDBittorrent,
	ChainNameFantomOpera:     ChainIDFantomOpera,
	ChainNameBoba:            ChainIDBoba,
	ChainNameZksync:          ChainIDZksync,
	ChainNameClvParachain:    ChainIDClvParachain,
	ChainNamePolygonZkevm:    ChainIDPolygonZkevm,
	ChainNameWemix:           ChainIDWemix,
	ChainNameMoonbeam:        ChainIDMoonbeam,
	ChainNameMoonriver:       ChainIDMoonriver,
	ChainNameMantle:          ChainIDMantle,
	ChainNameBase:            ChainIDBase,
	ChainNameArbitrumOne:     ChainIDArbitrumOne,
	ChainNameCelo:            ChainIDCelo,
	ChainNameAvalancheCChain: ChainIDAvalancheCChain,
	ChainNameLinea:           ChainIDLinea,
	ChainNameBlast:           ChainIDBlast,
	ChainNameAurora:          ChainIDAurora,
}

const REQ_PER_SECOND = 5
const REQ_PER_DAY = 10000

type RespErrCode int

const (
	Success               RespErrCode = 200000
	UnauthorizedOperation RespErrCode = 400001
	RateLimitExceeded     RespErrCode = 400002
	InvalidParams         RespErrCode = 400004
	UserNotExist          RespErrCode = 400005
	ServerBusy            RespErrCode = 400006
	InvalidAPIKey         RespErrCode = 400007
	InvalidAuthFormat     RespErrCode = 400008
	ExpiredAPIKey         RespErrCode = 400009
	NotFound              RespErrCode = 400010
	InvalidAddress        RespErrCode = 400011
	DailyLimitExceeded    RespErrCode = 400012
	UnsupportedChain      RespErrCode = 400013
	InternalError         RespErrCode = 500000
)

type Resp[D any] struct {
	RequestId string      `json:"request_id" bson:"request_id"`
	Code      RespErrCode `json:"code" bson:"code"`
	Message   string      `json:"message" bson:"message"`
	Data      D           `json:"data" bson:"data"`
}

type SupportedChain struct {
	ChainId   ChainID   `json:"chain_id" bson:"chain_id"`
	ChainName ChainName `json:"chain_name" bson:"chain_name"`
}

type EntityCategory struct {
	Name string `json:"name" bson:"name"`
	Code int    `json:"code" bson:"code"`
}

type EntityAttribute struct {
	Name     string   `json:"name" bson:"name"`
	Code     int      `json:"code" bson:"code"`
	CompInfo []string `json:"comp_info" bson:"comp_info"`
}

type EntityDescription struct {
	Attributes []EntityAttribute `json:"attributes" bson:"attributes"`
	Website    string            `json:"website" bson:"website"`
	Twitter    string            `json:"twitter" bson:"twitter"`
	Telegram   string            `json:"telegram" bson:"telegram"`
	Discord    string            `json:"discord" bson:"discord"`
}

type Entity struct {
	Entity      string            `json:"entity" bson:"entity"`
	Categories  []EntityCategory  `json:"categories" bson:"categories"`
	Attributes  []EntityAttribute `json:"attributes" bson:"attributes"`
	Description EntityDescription `json:"description" bson:"description"`
}

type AddressLabel struct {
	ChainId        int               `json:"chain_id" bson:"chain_id"`
	Address        string            `json:"address" bson:"address"`
	MainEntity     string            `json:"main_entity" bson:"main_entity"`
	MainEntityInfo Entity            `json:"main_entity_info" bson:"main_entity_info"`
	CompEntities   []string          `json:"comp_entities" bson:"comp_entities"`
	Attributes     []EntityAttribute `json:"attributes" bson:"attributes"`
	NameTag        string            `json:"name_tag" bson:"name_tag"`
}

type OneChainSeveralAddressesLabel struct {
	ChainId   ChainID        `json:"chain_id" bson:"chain_id"`
	Addresses []AddressLabel `json:"addresses" bson:"addresses"`
}

var SupportedChains = []SupportedChain{
	{ChainId: ChainIDSolana, ChainName: ChainNameSolana},
	{ChainId: ChainIDTron, ChainName: ChainNameTron},
	{ChainId: ChainIDBitcoin, ChainName: ChainNameBitcoin},
	{ChainId: ChainIDEthereum, ChainName: ChainNameEthereum},
	{ChainId: ChainIDOptimism, ChainName: ChainNameOptimism},
	{ChainId: ChainIDCronos, ChainName: ChainNameCronos},
	{ChainId: ChainIDBNBSmartChain, ChainName: ChainNameBNBSmartChain},
	{ChainId: ChainIDGnosis, ChainName: ChainNameGnosis},
	{ChainId: ChainIDPolygon, ChainName: ChainNamePolygon},
	{ChainId: ChainIDMantaPacific, ChainName: ChainNameMantaPacific},
	{ChainId: ChainIDBittorrent, ChainName: ChainNameBittorrent},
	{ChainId: ChainIDFantomOpera, ChainName: ChainNameFantomOpera},
	{ChainId: ChainIDBoba, ChainName: ChainNameBoba},
	{ChainId: ChainIDZksync, ChainName: ChainNameZksync},
	{ChainId: ChainIDClvParachain, ChainName: ChainNameClvParachain},
	{ChainId: ChainIDPolygonZkevm, ChainName: ChainNamePolygonZkevm},
	{ChainId: ChainIDWemix, ChainName: ChainNameWemix},
	{ChainId: ChainIDMoonbeam, ChainName: ChainNameMoonbeam},
	{ChainId: ChainIDMoonriver, ChainName: ChainNameMoonriver},
	{ChainId: ChainIDMantle, ChainName: ChainNameMantle},
	{ChainId: ChainIDBase, ChainName: ChainNameBase},
	{ChainId: ChainIDArbitrumOne, ChainName: ChainNameArbitrumOne},
	{ChainId: ChainIDCelo, ChainName: ChainNameCelo},
	{ChainId: ChainIDAvalancheCChain, ChainName: ChainNameAvalancheCChain},
	{ChainId: ChainIDLinea, ChainName: ChainNameLinea},
	{ChainId: ChainIDBlast, ChainName: ChainNameBlast},
	{ChainId: ChainIDAurora, ChainName: ChainNameAurora},
}

type Client struct {
	Header  http.Header
	Limiter *golimiter.ReqLimiter
	logger  *slog.Logger
}

func New(apiKey string, limiter *golimiter.ReqLimiter, logger *slog.Logger) *Client {
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
		limiter = golimiter.NewReqLimiter(time.Second, REQ_PER_SECOND)
	}

	header := http.Header{}
	header.Add("API-KEY", apiKey)
	header.Add("Content-Type", "application/json")

	return &Client{
		Header:  header,
		Limiter: limiter,
		logger:  logger,
	}
}

var defaultClient *Client = New("", nil, nil)

func DefaultClient() *Client {
	return defaultClient
}

func req[D any](c *Client, method string, path string, body map[string]any) (D, error) {
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
	var respData Resp[D]
	if err := decoder.Decode(&respData); err != nil {
		return *new(D), fmt.Errorf("metasleuth: decode response failed: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return *new(D), fmt.Errorf("metasleuth: status code: %d, error code: %d, message: %s", resp.StatusCode, respData.Code, respData.Message)
	}
	return respData.Data, nil
}

func (c *Client) GetSupportedChains() ([]SupportedChain, error) {
	return req[[]SupportedChain](c, http.MethodGet, "/chain-list", nil)
}

func (c *Client) GetAddressLabels(chainId ChainID, address string) (AddressLabel, error) {
	return req[AddressLabel](c, http.MethodPost, "/labels", map[string]any{
		"chain_id": chainId,
		"address":  address,
	})
}

func (c *Client) GetBatchLabels(chainId ChainID, addresses ...string) ([]AddressLabel, error) {
	return req[[]AddressLabel](c, http.MethodPost, "/batch-labels", map[string]any{
		"chain_id":  chainId,
		"addresses": addresses,
	})
}

func (c *Client) GetBatchLabelsMultiChain(chainIds []ChainID, addresses []string) ([]OneChainSeveralAddressesLabel, error) {
	return req[[]OneChainSeveralAddressesLabel](c, http.MethodPost, "/multi-chains-labels", map[string]any{
		"chain_ids": chainIds,
		"addresses": addresses,
	})
}

func (c *Client) GetEntity(entityName string) (Entity, error) {
	return req[Entity](c, http.MethodPost, "/entity", map[string]any{
		"entity": entityName,
	})
}
