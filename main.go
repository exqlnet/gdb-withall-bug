package main

import (
	_ "gdb-withall-bug/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"gdb-withall-bug/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
