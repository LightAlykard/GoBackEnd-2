package main

import (
	"context"
	"fmt"
	"k8s-go-app/server"
	"k8s-go-app/version"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/LightAlykard/GoLibsTest/config"

	"github.com/LightAlykard/GoLibsTest/logic"
	"github.com/LightAlykard/GoLibsTest/telemetry"
	"go.uber.org/zap"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	l := zap.L().Sugar() //простой логер

	// // Слушаем сигналы ОС для завершения работы
	// ctx, cancel := newOSSignalContext()
	// defer cancel()

	// Настраиваем сборщик трейсов
	tp, err := telemetry.RunTracingCollection(ctx)
	if err != nil {
		l.Panic(err)
	}
	defer func() {
		if err = tp.Shutdown(context.Background()); err != nil {
			l.Errorf("failed to stop the traces collector: %v", err)
		}
	}()

	launchMode := config.LaunchMode(os.Getenv("LAUNCH_MODE"))
	if len(launchMode) == 0 {
		launchMode = config.LocalEnv
	}
	log.Printf("LAUNCH MODE: %v", launchMode)
	cfg, err := config.Load(launchMode, "./config")
	if err != nil {
		//l.Errorf("failed to stop the traces collector: %v", err)
		log.Fatal(err)
	}
	log.Printf("CONFIG: %+v", cfg)
	info := server.VersionInfo{
		Version: version.Version,
		Commit:  version.Commit,
		Build:   version.Build,
	}

	tr := tp.Tracer("server")
	// Запускаем сервер
	// s := &server.Server{
	// 	Tr: tr,
	// 	Logic: &logic.Logic{
	// 		Tr: tr,
	// 	},
	// }

	lg := logic.Logic{
		Tr: tr,
	}

	// go func() {
	// 	err := s.Start()
	// 	if err != nil && err != http.ErrServerClosed {
	// 		l.Panic(err)
	// 	}
	// }()

	srv := server.New(info, cfg.Port, tr, lg)

	go func() {
		err := srv.Serve(ctx)
		if err != nil {
			log.Println(fmt.Errorf("serve: %w", err))
			return
		}
	}()
	osSigChan := make(chan os.Signal, 1)
	signal.Notify(osSigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-osSigChan
	log.Println("OS interrupting signal has received")
	cancel()

	// <-ctx.Done()

	// err = s.Stop(context.Background())
	// if err != nil {
	// 	l.Error(err)
	// }
}

func newOSSignalContext() (context.Context, func()) {
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
	}()

	return ctx, func() {
		signal.Stop(c)
		cancel()
	}
}
