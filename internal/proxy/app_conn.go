package proxy

import (
	"context"
	"time"

	"github.com/go-kit/kit/metrics"
	abciclient "github.com/tendermint/tendermint/abci/client"
	"github.com/tendermint/tendermint/abci/types"
)

//go:generate ../../scripts/mockery_generate.sh AppConnConsensus|AppConnMempool|AppConnQuery|AppConnSnapshot

//----------------------------------------------------------------------------------------
// Enforce which abci msgs can be sent on a connection at the type level

type AppConnConsensus interface {
	SetResponseCallback(abciclient.Callback)
	Error() error

	InitChain(context.Context, types.RequestInitChain) (*types.ResponseInitChain, error)

	BeginBlock(context.Context, types.RequestBeginBlock) (*types.ResponseBeginBlock, error)
	DeliverTxAsync(context.Context, types.RequestDeliverTx) (*abciclient.ReqRes, error)
	EndBlock(context.Context, types.RequestEndBlock) (*types.ResponseEndBlock, error)
	Commit(context.Context) (*types.ResponseCommit, error)
}

type AppConnMempool interface {
	SetResponseCallback(abciclient.Callback)
	Error() error

	CheckTxAsync(context.Context, types.RequestCheckTx) (*abciclient.ReqRes, error)
	CheckTx(context.Context, types.RequestCheckTx) (*types.ResponseCheckTx, error)

	FlushAsync(context.Context) (*abciclient.ReqRes, error)
	Flush(context.Context) error
}

type AppConnQuery interface {
	Error() error

	Echo(context.Context, string) (*types.ResponseEcho, error)
	Info(context.Context, types.RequestInfo) (*types.ResponseInfo, error)
	Query(context.Context, types.RequestQuery) (*types.ResponseQuery, error)
}

type AppConnSnapshot interface {
	Error() error

	ListSnapshots(context.Context, types.RequestListSnapshots) (*types.ResponseListSnapshots, error)
	OfferSnapshot(context.Context, types.RequestOfferSnapshot) (*types.ResponseOfferSnapshot, error)
	LoadSnapshotChunk(context.Context, types.RequestLoadSnapshotChunk) (*types.ResponseLoadSnapshotChunk, error)
	ApplySnapshotChunk(context.Context, types.RequestApplySnapshotChunk) (*types.ResponseApplySnapshotChunk, error)
}

//-----------------------------------------------------------------------------------------
// Implements AppConnConsensus (subset of abciclient.Client)

type appConnConsensus struct {
	metrics *Metrics
	appConn abciclient.Client
}

func NewAppConnConsensus(appConn abciclient.Client, metrics *Metrics) AppConnConsensus {
	return &appConnConsensus{
		metrics: metrics,
		appConn: appConn,
	}
}

func (app *appConnConsensus) SetResponseCallback(cb abciclient.Callback) {
	app.appConn.SetResponseCallback(cb)
}

func (app *appConnConsensus) Error() error {
	return app.appConn.Error()
}

func (app *appConnConsensus) InitChain(
	ctx context.Context,
	req types.RequestInitChain,
) (*types.ResponseInitChain, error) {
	defer addTimeSample(app.metrics.MethodTiming.With("method", "init_chain", "type", "sync"))()
	return app.appConn.InitChain(ctx, req)
}

func (app *appConnConsensus) BeginBlock(
	ctx context.Context,
	req types.RequestBeginBlock,
) (*types.ResponseBeginBlock, error) {
	defer addTimeSample(app.metrics.MethodTiming.With("method", "begin_block", "type", "sync"))()
	return app.appConn.BeginBlock(ctx, req)
}

func (app *appConnConsensus) DeliverTxAsync(
	ctx context.Context,
	req types.RequestDeliverTx,
) (*abciclient.ReqRes, error) {
	defer addTimeSample(app.metrics.MethodTiming.With("method", "deliver_tx", "type", "async"))()
	return app.appConn.DeliverTxAsync(ctx, req)
}

func (app *appConnConsensus) EndBlock(
	ctx context.Context,
	req types.RequestEndBlock,
) (*types.ResponseEndBlock, error) {
	defer addTimeSample(app.metrics.MethodTiming.With("method", "deliver_tx", "type", "sync"))()
	return app.appConn.EndBlock(ctx, req)
}

func (app *appConnConsensus) Commit(ctx context.Context) (*types.ResponseCommit, error) {
	defer addTimeSample(app.metrics.MethodTiming.With("method", "commit", "type", "sync"))()
	return app.appConn.Commit(ctx)
}

func (app *appConnConsensus) PreprocessTxsSync(
	ctx context.Context,
	req types.RequestPreprocessTxs,
) (*types.ResponsePreprocessTxs, error) {
	return app.appConn.PreprocessTxsSync(ctx, req)
}

//------------------------------------------------
// Implements AppConnMempool (subset of abciclient.Client)

type appConnMempool struct {
	metrics *Metrics
	appConn abciclient.Client
}

func NewAppConnMempool(appConn abciclient.Client, metrics *Metrics) AppConnMempool {
	return &appConnMempool{
		metrics: metrics,
		appConn: appConn,
	}
}

func (app *appConnMempool) SetResponseCallback(cb abciclient.Callback) {
	app.appConn.SetResponseCallback(cb)
}

func (app *appConnMempool) Error() error {
	return app.appConn.Error()
}

func (app *appConnMempool) FlushAsync(ctx context.Context) (*abciclient.ReqRes, error) {
	defer addTimeSample(app.metrics.MethodTiming.With("method", "flush", "type", "async"))()
	return app.appConn.FlushAsync(ctx)
}

func (app *appConnMempool) Flush(ctx context.Context) error {
	defer addTimeSample(app.metrics.MethodTiming.With("method", "flush", "type", "sync"))()
	return app.appConn.Flush(ctx)
}

func (app *appConnMempool) CheckTxAsync(ctx context.Context, req types.RequestCheckTx) (*abciclient.ReqRes, error) {
	defer addTimeSample(app.metrics.MethodTiming.With("method", "check_tx", "type", "async"))()
	return app.appConn.CheckTxAsync(ctx, req)
}

func (app *appConnMempool) CheckTx(ctx context.Context, req types.RequestCheckTx) (*types.ResponseCheckTx, error) {
	defer addTimeSample(app.metrics.MethodTiming.With("method", "check_tx", "type", "sync"))()
	return app.appConn.CheckTx(ctx, req)
}

//------------------------------------------------
// Implements AppConnQuery (subset of abciclient.Client)

type appConnQuery struct {
	metrics *Metrics
	appConn abciclient.Client
}

func NewAppConnQuery(appConn abciclient.Client, metrics *Metrics) AppConnQuery {
	return &appConnQuery{
		metrics: metrics,
		appConn: appConn,
	}
}

func (app *appConnQuery) Error() error {
	return app.appConn.Error()
}

func (app *appConnQuery) Echo(ctx context.Context, msg string) (*types.ResponseEcho, error) {
	defer addTimeSample(app.metrics.MethodTiming.With("method", "echo", "type", "sync"))()
	return app.appConn.Echo(ctx, msg)
}

func (app *appConnQuery) Info(ctx context.Context, req types.RequestInfo) (*types.ResponseInfo, error) {
	defer addTimeSample(app.metrics.MethodTiming.With("method", "info", "type", "sync"))()
	return app.appConn.Info(ctx, req)
}

func (app *appConnQuery) Query(ctx context.Context, reqQuery types.RequestQuery) (*types.ResponseQuery, error) {
	defer addTimeSample(app.metrics.MethodTiming.With("method", "query", "type", "sync"))()
	return app.appConn.Query(ctx, reqQuery)
}

//------------------------------------------------
// Implements AppConnSnapshot (subset of abciclient.Client)

type appConnSnapshot struct {
	metrics *Metrics
	appConn abciclient.Client
}

func NewAppConnSnapshot(appConn abciclient.Client, metrics *Metrics) AppConnSnapshot {
	return &appConnSnapshot{
		metrics: metrics,
		appConn: appConn,
	}
}

func (app *appConnSnapshot) Error() error {
	return app.appConn.Error()
}

func (app *appConnSnapshot) ListSnapshots(
	ctx context.Context,
	req types.RequestListSnapshots,
) (*types.ResponseListSnapshots, error) {
	defer addTimeSample(app.metrics.MethodTiming.With("method", "list_snapshots", "type", "sync"))()
	return app.appConn.ListSnapshots(ctx, req)
}

func (app *appConnSnapshot) OfferSnapshot(
	ctx context.Context,
	req types.RequestOfferSnapshot,
) (*types.ResponseOfferSnapshot, error) {
	defer addTimeSample(app.metrics.MethodTiming.With("method", "offer_snapshot", "type", "sync"))()
	return app.appConn.OfferSnapshot(ctx, req)
}

func (app *appConnSnapshot) LoadSnapshotChunk(
	ctx context.Context,
	req types.RequestLoadSnapshotChunk) (*types.ResponseLoadSnapshotChunk, error) {
	defer addTimeSample(app.metrics.MethodTiming.With("method", "load_snapshot_chunk", "type", "sync"))()
	return app.appConn.LoadSnapshotChunk(ctx, req)
}

func (app *appConnSnapshot) ApplySnapshotChunk(
	ctx context.Context,
	req types.RequestApplySnapshotChunk) (*types.ResponseApplySnapshotChunk, error) {
	defer addTimeSample(app.metrics.MethodTiming.With("method", "apply_snapshot_chunk", "type", "sync"))()
	return app.appConn.ApplySnapshotChunk(ctx, req)
}

// addTimeSample returns a function that, when called, adds an observation to m.
// The observation added to m is the number of seconds ellapsed since addTimeSample
// was initially called. addTimeSample is meant to be called in a defer to calculate
// the amount of time a function takes to complete.
func addTimeSample(m metrics.Histogram) func() {
	start := time.Now()
	return func() { m.Observe(time.Since(start).Seconds()) }
}
