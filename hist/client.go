package dbn_hist

import (
	"net/http"
	"time"

	"github.com/NimbleMarkets/dbn-go"
)

type Client interface {
	GetRange(apiKey string, jobParams SubmitJobParams) ([]byte, error)
	SymbologyResolve(apiKey string, params ResolveParams) (*Resolution, error)
	ListPublishers(apiKey string) ([]PublisherDetail, error)
	ListDatasets(apiKey string, dateRange DateRange) ([]string, error)
	ListSchemas(apiKey string, dataset string) ([]string, error)
	ListFields(apiKey string, encoding dbn.Encoding, schema dbn.Schema) ([]FieldDetail, error)
	ListUnitPrices(apiKey string, dataset string) ([]UnitPricesForMode, error)
	GetDatasetCondition(apiKey string, dataset string, dateRange DateRange) ([]ConditionDetail, error)
	GetDatasetRange(apiKey string, dataset string) (DateRange, error)
	GetRecordCount(apiKey string, metaParams MetadataQueryParams) (int, error)
	GetBillableSize(apiKey string, metaParams MetadataQueryParams) (int, error)
	GetCost(apiKey string, metaParams MetadataQueryParams) (float64, error)
	ListJobs(apiKey string, stateFilter string, sinceYMD time.Time) ([]BatchJob, error)
	ListFiles(apiKey string, jobID string) ([]BatchFileDesc, error)
	SubmitJob(apiKey string, jobParams SubmitJobParams) (*BatchJob, error)
}

type client struct {
	apiKey string
	
	httpClient *http.Client
}
