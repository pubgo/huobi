package main

import (
	"github.com/huobirdcenter/huobi/cmd/accountclientexample"
	"github.com/huobirdcenter/huobi/cmd/accountwebsocketclientexample"
	"github.com/huobirdcenter/huobi/cmd/algoorderclientexample"
	"github.com/huobirdcenter/huobi/cmd/commonclientexample"
	"github.com/huobirdcenter/huobi/cmd/crossmarginclientexample"
	"github.com/huobirdcenter/huobi/cmd/etfclientexample"
	"github.com/huobirdcenter/huobi/cmd/isolatedmarginclientexample"
	"github.com/huobirdcenter/huobi/cmd/marketclientexample"
	"github.com/huobirdcenter/huobi/cmd/marketwebsocketclientexample"
	"github.com/huobirdcenter/huobi/cmd/orderclientexample"
	"github.com/huobirdcenter/huobi/cmd/orderwebsocketclientexample"
	"github.com/huobirdcenter/huobi/cmd/stablecoinclientexample"
	"github.com/huobirdcenter/huobi/cmd/subuserclientexample"
	"github.com/huobirdcenter/huobi/cmd/walletclientexample"
	"github.com/huobirdcenter/huobi/logging/perflogger"
)

func main() {
	runAll()
}

// Run all examples
func runAll() {
	commonclientexample.RunAllExamples()
	accountclientexample.RunAllExamples()
	orderclientexample.RunAllExamples()
	algoorderclientexample.RunAllExamples()
	marketclientexample.RunAllExamples()
	isolatedmarginclientexample.RunAllExamples()
	crossmarginclientexample.RunAllExamples()
	walletclientexample.RunAllExamples()
	subuserclientexample.RunAllExamples()
	stablecoinclientexample.RunAllExamples()
	etfclientexample.RunAllExamples()
	marketwebsocketclientexample.RunAllExamples()
	accountwebsocketclientexample.RunAllExamples()
	orderwebsocketclientexample.RunAllExamples()
}

// Run performance test
func runPerfTest() {
	perflogger.Enable(true)

	commonclientexample.RunAllExamples()
	accountclientexample.RunAllExamples()
	orderclientexample.RunAllExamples()
	marketclientexample.RunAllExamples()
	isolatedmarginclientexample.RunAllExamples()
	crossmarginclientexample.RunAllExamples()
	walletclientexample.RunAllExamples()
	etfclientexample.RunAllExamples()
}
