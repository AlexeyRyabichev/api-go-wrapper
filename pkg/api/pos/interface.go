package pos

import "context"

type (
	Manager interface {
		GetPointsOfSale(ctx context.Context, filters map[string]string) ([]PointOfSale, error)
		GetDayClosings(ctx context.Context, filters map[string]string) ([]DayClosing, error)
		GetCashIns(ctx context.Context, filters map[string]string) ([]CashIn, error)
		GetReasonCodes(ctx context.Context, filters map[string]string) ([]ReasonCode, error)
	}
)
