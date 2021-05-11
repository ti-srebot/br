// Copyright 2019 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"go.uber.org/zap"

	"github.com/pingcap/br/pkg/lightning"
	"github.com/pingcap/br/pkg/lightning/config"
	"github.com/pingcap/br/pkg/lightning/log"
)

func main() {
	globalCfg := config.Must(config.LoadGlobalConfig(os.Args[1:], nil))
	fmt.Fprintf(os.Stdout, "Verbose debug logs will be written to %s\n\n", globalCfg.App.Config.File)

	app := lightning.New(globalCfg)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		sig := <-sc
		log.L().Info("got signal to exit", zap.Stringer("signal", sig))
		app.Stop()
	}()

	logger := log.L()

	// Lightning allocates too many transient objects and heap size is small,
	// so garbage collections happen too frequently and lots of time is spent in GC component.
	//
	// In a test of loading the table `order_line.csv` of 14k TPCC.
	// The time need of `encode kv data and write` step reduce from 52m4s to 37m30s when change
	// GOGC from 100 to 500, the total time needed reduce near 15m too.
	// The cost of this is the memory of lightnin at runtime grow from about 200M to 700M, but it's acceptable.
	// So we set the gc percentage as 500 default to reduce the GC frequency instead of 100.
	//
	// Local mode need much more memory than importer/tidb mode, if the gc percentage is too high,
	// lightning memory usage will also be high.
	if globalCfg.TikvImporter.Backend != config.BackendLocal {
		gogc := os.Getenv("GOGC")
		if gogc == "" {
			old := debug.SetGCPercent(500)
			log.L().Debug("set gc percentage", zap.Int("old", old), zap.Int("new", 500))
		}
	}

	err := app.GoServe()
	if err != nil {
		logger.Error("failed to start HTTP server", zap.Error(err))
		fmt.Fprintln(os.Stderr, "failed to start HTTP server:", err)
		return
	}

	err = func() error {
		if globalCfg.App.ServerMode {
			return app.RunServer()
		}
		cfg := config.NewConfig()
		if err := cfg.LoadFromGlobal(globalCfg); err != nil {
			return err
		}
		return app.RunOnce(context.Background(), cfg, nil)
	}()

	if err != nil {
		logger.Error("tidb lightning encountered error stack info", zap.Error(err))
		logger.Error("tidb lightning encountered error", log.ShortError(err))
		fmt.Fprintln(os.Stderr, "tidb lightning encountered error: ", err)
	} else {
		logger.Info("tidb lightning exit")
		fmt.Fprintln(os.Stdout, "tidb lightning exit")
	}

	// call Sync() with log to stdout may return error in some case, so just skip it
	if globalCfg.App.File != "" {
		syncErr := logger.Sync()
		if syncErr != nil {
			fmt.Fprintln(os.Stderr, "sync log failed", syncErr)
		}
	}

	if err != nil {
		exit(1)
	}
}

// main_test.go override exit to pass unit test.
var exit = os.Exit
