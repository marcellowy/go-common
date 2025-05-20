package tools

import (
	"context"
	"fmt"
	"testing"
)

func TestRetryExecute(t *testing.T) {
	type args struct {
		ctx         context.Context
		tryMaxTimes int
		f           func(current int) error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				ctx:         context.Background(),
				tryMaxTimes: 2,
				f: func(current int) error {
					if current == 2 {
						return nil
					}
					return fmt.Errorf("wait next")
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RetryExecute(tt.args.ctx, tt.args.tryMaxTimes, tt.args.f); (err != nil) != tt.wantErr {
				t.Errorf("RetryExecute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
