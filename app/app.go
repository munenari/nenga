package app

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/munenari/nenga/app/domain"
	"github.com/munenari/nenga/app/query"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/munenari/nenga/env"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx      context.Context
	e        *echo.Echo
	launched bool
}

func NewApp(e *echo.Echo) *App {
	return &App{e: e}
}

func (a *App) Startup(ctx context.Context) {
	if err := a.startupInternal(ctx); err != nil {
		runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
			Type:          runtime.ErrorDialog,
			Message:       "起動準備に失敗しました\nアプリを終了します\n" + err.Error(),
			Buttons:       []string{"OK"},
			DefaultButton: "OK",
		})
		os.Exit(1)
	}
	a.launched = true
	runtime.EventsEmit(a.ctx, "launched")
	slog.Info("launched")
}

func (a *App) startupInternal(ctx context.Context) error {
	a.ctx = ctx
	d, err := os.UserConfigDir()
	if err != nil {
		d = os.TempDir()
	}
	d = filepath.Join(d, "nenga", env.GetEnvName())
	dbPath := filepath.Join(d, "data.ddb")
	if err := query.Init(dbPath); err != nil {
		return err
	}
	return nil
}

func (a *App) Shutdown(ctx context.Context) {
	defer query.Close()
}

func (a *App) IsLaunched() bool {
	return a.launched
}

func (a *App) AllNengaAtena(year int) ([]*domain.NengaAtena, error) {
	return query.Atena{}.List(query.DB(), year)
}

func (a *App) SaveNengaAtena(x *domain.NengaAtena) error {
	if x.ID == uuid.Nil {
		id, _ := uuid.NewV7()
		x.ID = id
	}
	return query.Atena{}.Save(query.DB(), x)
}

func (a *App) DeleteNengaAtena(id string) error {
	return query.Atena{}.Delete(query.DB(), id)
}
