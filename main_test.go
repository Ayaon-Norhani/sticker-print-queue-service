package main

import (
	"errors"
	"testing"
	"time"
)

func TestPrintJob_Validate(t *testing.T) {
	tests := []struct {
		name    string
		job     PrintJob
		wantErr error
	}{
		{
			name: "Valid Vinyl Sticker Job",
			job: PrintJob{
				ID:        "job_001",
				SKU:       "STK-CUSTOM-100",
				Quantity:  500,
				Material:  MaterialVinyl,
				CreatedAt: time.Now(),
			},
			wantErr: nil,
		},
		{
			name: "Valid Label Hologram Job",
			job: PrintJob{
				ID:        "job_002",
				SKU:       "LBL-LOGO-50",
				Quantity:  50,
				Material:  MaterialHologram,
				CreatedAt: time.Now(),
			},
			wantErr: nil,
		},
		{
			name: "Invalid SKU Prefix",
			job: PrintJob{
				ID:        "job_003",
				SKU:       "BAD-SKU-100",
				Quantity:  100,
				Material:  MaterialVinyl,
				CreatedAt: time.Now(),
			},
			wantErr: ErrInvalidSKU,
		},
		{
			name: "Quantity Below Minimum Limit",
			job: PrintJob{
				ID:        "job_004",
				SKU:       "STK-MINI-01",
				Quantity:  5,
				Material:  MaterialVinyl,
				CreatedAt: time.Now(),
			},
			wantErr: ErrQuantityOutOfRange,
		},
		{
			name: "Unsupported Material",
			job: PrintJob{
				ID:        "job_005",
				SKU:       "STK-GLOW-20",
				Quantity:  100,
				Material:  "glow_in_dark",
				CreatedAt: time.Now(),
			},
			wantErr: ErrInvalidMaterial,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.job.Validate()
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
