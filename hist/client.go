package dbn_hist

import (
	"context"
	"net/http"
	"time"

	"github.com/NimbleMarkets/dbn-go"
)

type Client interface {
	GetRange(ctx context.Context, jobParams SubmitJobParams) ([]byte, error)
	SymbologyResolve(ctx context.Context, params ResolveParams) (*Resolution, error)
	ListPublishers(ctx context.Context) ([]PublisherDetail, error)
	ListDatasets(ctx context.Context, dateRange DateRange) ([]string, error)
	ListSchemas(ctx context.Context, dataset string) ([]string, error)
	ListFields(ctx context.Context, encoding dbn.Encoding, schema dbn.Schema) ([]FieldDetail, error)
	ListUnitPrices(ctx context.Context, dataset string) ([]UnitPricesForMode, error)
	GetDatasetCondition(ctx context.Context, dataset string, dateRange DateRange) ([]ConditionDetail, error)
	GetDatasetRange(ctx context.Context, dataset string) (DateRange, error)
	GetRecordCount(ctx context.Context, metaParams MetadataQueryParams) (int, error)
	GetBillableSize(ctx context.Context, metaParams MetadataQueryParams) (int, error)
	GetCost(ctx context.Context, metaParams MetadataQueryParams) (float64, error)
	ListJobs(ctx context.Context, stateFilter string, sinceYMD time.Time) ([]BatchJob, error)
	ListFiles(ctx context.Context, jobID string) ([]BatchFileDesc, error)
	SubmitJob(ctx context.Context, jobParams SubmitJobParams) (*BatchJob, error)
}

type client struct {
	apiKey string

	httpClient *http.Client
}

func New(apiKey string, httpClient *http.Client) Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &client{
		apiKey:     apiKey,
		httpClient: httpClient,
	}
}
