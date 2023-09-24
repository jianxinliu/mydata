package misc

import (
	"context"
	"fmt"

	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

func ShowMessage(ctx context.Context, message, title string) {
	_, err := rt.MessageDialog(ctx, rt.MessageDialogOptions{
		Title:   title,
		Message: message,
		Type:    rt.ErrorDialog,
	})
	if err != nil {
		fmt.Printf("dialog err: %v", err)
	}
}
