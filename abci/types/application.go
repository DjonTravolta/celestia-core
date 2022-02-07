package types

import (
	"context"
	fmt "fmt"
)

// Application is an interface that enables any finite, deterministic state machine
// to be driven by a blockchain-based replication engine via the ABCI.
// All methods take a RequestXxx argument and return a ResponseXxx argument,
// except CheckTx/DeliverTx, which take `tx []byte`, and `Commit`, which takes nothing.
// nolint:lll // ignore for interface
type Application interface {
	// Info/Query Connection
	Info(RequestInfo) ResponseInfo    // Return application info
	Query(RequestQuery) ResponseQuery // Query for state

	// Mempool Connection
	CheckTx(RequestCheckTx) ResponseCheckTx // Validate a tx for the mempool

	// Consensus Connection
	InitChain(RequestInitChain) ResponseInitChain // Initialize blockchain w validators/other info from TendermintCore
	PrepareProposal(RequestPrepareProposal) ResponsePrepareProposal
	// Signals the beginning of a block
	BeginBlock(RequestBeginBlock) ResponseBeginBlock
	// Deliver a tx for full processing
	DeliverTx(RequestDeliverTx) ResponseDeliverTx
	// Signals the end of a block, returns changes to the validator set
	EndBlock(RequestEndBlock) ResponseEndBlock
	// Commit the state and return the application Merkle root hash
	Commit() ResponseCommit
	// Create application specific vote extension
	ExtendVote(RequestExtendVote) ResponseExtendVote
	// Verify application's vote extension data
	VerifyVoteExtension(RequestVerifyVoteExtension) ResponseVerifyVoteExtension

	// State Sync Connection
	ListSnapshots(RequestListSnapshots) ResponseListSnapshots                // List available snapshots
	OfferSnapshot(RequestOfferSnapshot) ResponseOfferSnapshot                // Offer a snapshot to the application
	LoadSnapshotChunk(RequestLoadSnapshotChunk) ResponseLoadSnapshotChunk    // Load a snapshot chunk
	ApplySnapshotChunk(RequestApplySnapshotChunk) ResponseApplySnapshotChunk // Apply a shapshot chunk
}

//-------------------------------------------------------
// BaseApplication is a base form of Application

var _ Application = (*BaseApplication)(nil)

type BaseApplication struct {
}

func NewBaseApplication() *BaseApplication {
	return &BaseApplication{}
}

func (BaseApplication) Info(req RequestInfo) ResponseInfo {
	return ResponseInfo{}
}

func (BaseApplication) DeliverTx(req RequestDeliverTx) ResponseDeliverTx {
	return ResponseDeliverTx{Code: CodeTypeOK}
}

func (BaseApplication) CheckTx(req RequestCheckTx) ResponseCheckTx {
	return ResponseCheckTx{Code: CodeTypeOK}
}

func (BaseApplication) Commit() ResponseCommit {
	return ResponseCommit{}
}

func (BaseApplication) ExtendVote(req RequestExtendVote) ResponseExtendVote {
	return ResponseExtendVote{}
}

func (BaseApplication) VerifyVoteExtension(req RequestVerifyVoteExtension) ResponseVerifyVoteExtension {
	return ResponseVerifyVoteExtension{}
}

func (BaseApplication) Query(req RequestQuery) ResponseQuery {
	return ResponseQuery{Code: CodeTypeOK}
}

func (BaseApplication) InitChain(req RequestInitChain) ResponseInitChain {
	return ResponseInitChain{}
}

func (BaseApplication) BeginBlock(req RequestBeginBlock) ResponseBeginBlock {
	return ResponseBeginBlock{}
}

func (BaseApplication) EndBlock(req RequestEndBlock) ResponseEndBlock {
	return ResponseEndBlock{}
}

func (BaseApplication) ListSnapshots(req RequestListSnapshots) ResponseListSnapshots {
	return ResponseListSnapshots{}
}

func (BaseApplication) OfferSnapshot(req RequestOfferSnapshot) ResponseOfferSnapshot {
	return ResponseOfferSnapshot{}
}

func (BaseApplication) LoadSnapshotChunk(req RequestLoadSnapshotChunk) ResponseLoadSnapshotChunk {
	return ResponseLoadSnapshotChunk{}
}

func (BaseApplication) ApplySnapshotChunk(req RequestApplySnapshotChunk) ResponseApplySnapshotChunk {
	return ResponseApplySnapshotChunk{}
}

func (BaseApplication) PrepareProposal(req RequestPrepareProposal) ResponsePrepareProposal {
	return ResponsePrepareProposal{}
}

//-------------------------------------------------------

// GRPCApplication is a GRPC wrapper for Application
type GRPCApplication struct {
	app Application
}

func NewGRPCApplication(app Application) *GRPCApplication {
	return &GRPCApplication{app}
}

func (app *GRPCApplication) Echo(ctx context.Context, req *RequestEcho) (*ResponseEcho, error) {
	return &ResponseEcho{Message: req.Message}, nil
}

func (app *GRPCApplication) Flush(ctx context.Context, req *RequestFlush) (*ResponseFlush, error) {
	return &ResponseFlush{}, nil
}

func (app *GRPCApplication) Info(ctx context.Context, req *RequestInfo) (*ResponseInfo, error) {
	res := app.app.Info(*req)
	return &res, nil
}

func (app *GRPCApplication) DeliverTx(ctx context.Context, req *RequestDeliverTx) (*ResponseDeliverTx, error) {
	res := app.app.DeliverTx(*req)
	return &res, nil
}

func (app *GRPCApplication) CheckTx(ctx context.Context, req *RequestCheckTx) (*ResponseCheckTx, error) {
	res := app.app.CheckTx(*req)
	return &res, nil
}

func (app *GRPCApplication) Query(ctx context.Context, req *RequestQuery) (*ResponseQuery, error) {
	res := app.app.Query(*req)
	return &res, nil
}

func (app *GRPCApplication) Commit(ctx context.Context, req *RequestCommit) (*ResponseCommit, error) {
	res := app.app.Commit()
	return &res, nil
}

func (app *GRPCApplication) InitChain(ctx context.Context, req *RequestInitChain) (*ResponseInitChain, error) {
	res := app.app.InitChain(*req)
	return &res, nil
}

func (app *GRPCApplication) BeginBlock(ctx context.Context, req *RequestBeginBlock) (*ResponseBeginBlock, error) {
	res := app.app.BeginBlock(*req)
	return &res, nil
}

func (app *GRPCApplication) EndBlock(ctx context.Context, req *RequestEndBlock) (*ResponseEndBlock, error) {
	res := app.app.EndBlock(*req)
	return &res, nil
}

func (app *GRPCApplication) ListSnapshots(
	ctx context.Context, req *RequestListSnapshots) (*ResponseListSnapshots, error) {
	res := app.app.ListSnapshots(*req)
	return &res, nil
}

func (app *GRPCApplication) OfferSnapshot(
	ctx context.Context, req *RequestOfferSnapshot) (*ResponseOfferSnapshot, error) {
	res := app.app.OfferSnapshot(*req)
	return &res, nil
}

func (app *GRPCApplication) LoadSnapshotChunk(
	ctx context.Context, req *RequestLoadSnapshotChunk) (*ResponseLoadSnapshotChunk, error) {
	res := app.app.LoadSnapshotChunk(*req)
	return &res, nil
}

func (app *GRPCApplication) ApplySnapshotChunk(
	ctx context.Context, req *RequestApplySnapshotChunk) (*ResponseApplySnapshotChunk, error) {
	res := app.app.ApplySnapshotChunk(*req)
	return &res, nil
}

func (app *GRPCApplication) ExtendVote(
	ctx context.Context, req *RequestExtendVote) (*ResponseExtendVote, error) {
	res := app.app.ExtendVote(*req)
	return &res, nil
}

func (app *GRPCApplication) VerifyVoteExtension(
	ctx context.Context, req *RequestVerifyVoteExtension) (*ResponseVerifyVoteExtension, error) {
	res := app.app.VerifyVoteExtension(*req)
	return &res, nil
}

func (app *GRPCApplication) PrepareProposal(
	ctx context.Context, req *RequestPrepareProposal) (*ResponsePrepareProposal, error) {
	res := app.app.PrepareProposal(*req)
	return &res, nil
}

func (app *GRPCApplication) ProcessProposal(
	ctx context.Context, req *RequestProcessProposal) (*ResponseProcessProposal, error) {
	fmt.Println("ProcessProposal not implemented yet.")
	// res := app.app.ProcessProposal(*req)
	return nil, nil
}
