package handler

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/metadiv-go-tech/metagin/v2"
	"github.com/metadiv-go-tech/metagin/v2/base"
	"github.com/metadiv-go-tech/module_motor_garage/config"
)

var ApiCustomerBookInSheet = metagin.Get(
	"printCustomerBookInSheet",
	"Print Customer Book In Sheet",
	fmt.Sprintf("/api/%s/motor-garage/customer/book-in/sheet", config.SystemVersion),
	func(ctx metagin.Context[base.Empty, base.Empty]) {
		pdf, err := os.ReadFile("vortex-customerforms-merged.pdf")
		if err != nil {
			ctx.ErrWithStatus(http.StatusInternalServerError, errors.New(err.Error()))
			return
		}
		ctx.File(pdf, "customer-book-in-sheet.pdf")
	},
)
