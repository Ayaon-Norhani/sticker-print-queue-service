package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
)

// Custom domain errors for validation and database layer
var (
	ErrInvalidSKU          = errors.New("SKU must start with 'STK-' or 'LBL-'")
	ErrQuantityOutOfRange  = errors.New("order quantity must be between 10 and 10,000 units")
	ErrInvalidMaterial     = errors.New("unsupported print material type")
	ErrDatabaseConnection  = errors.New("failed to acquire database connection")
)

const (
	MaterialVinyl    = "die_cut_vinyl"
	MaterialHologram = "hologram"
	MaterialClear    = "clear_vinyl"
)

// PrintJob represents an order line item destined for the manufacturing queue
type PrintJob struct {
	ID        string    `json:"id"`
	SKU       string    `json:"sku"`
	Quantity  int       `json:"quantity"`
	Material  string    `json:"material"`
	CreatedAt time.Time `json:"created_at"`
}

// Validate executes fast guard-clause validation for incoming order items
func (pj *PrintJob) Validate() error {
	if !strings.HasPrefix(pj.SKU, "STK-") && !strings.HasPrefix(pj.SKU, "LBL-") {
		return ErrInvalidSKU
	}

	if pj.Quantity < 10 || pj.Quantity > 10000 {
		return ErrQuantityOutOfRange
	}

	if pj.Material != MaterialVinyl && pj.Material != MaterialHologram && pj.Material != MaterialClear {
		return ErrInvalidMaterial
	}

	return nil
}

// OrderRepository manages PostgreSQL database operations for print queue jobs
type OrderRepository struct {
	db *sql.DB
}

// NewOrderRepository initializes a repository instance
func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

// InsertPrintJob stores validated jobs into PostgreSQL with context timeout handling
func (r *OrderRepository) InsertPrintJob(ctx context.Context, job *PrintJob) error {
	if r.db == nil {
		return ErrDatabaseConnection
	}

	query := `
		INSERT INTO print_queue (id, sku, quantity, material, created_at)
		VALUES ($1, $2, $3, $4, $5);
	`

	_, err := r.db.ExecContext(ctx, query, job.ID, job.SKU, job.Quantity, job.Material, job.CreatedAt)
	if err != nil {
		return fmt.Errorf("db insert failed: %w", err)
	}

	return nil
}

func main() {
	fmt.Println("Sticker Print Queue Service initialized.")
}
