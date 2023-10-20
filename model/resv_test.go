package model

import (
	"testing"
	"time"
)

func TestResv_CalculateTimeBits(t *testing.T) {
	type args struct {
		startTime time.Time
		endTime   time.Time
	}
	tests := []struct {
		name             string
		args             args
		expectedStartBit int
		expectedEndBit   int
	}{
		{
			"00:00:00~00:01:00",
			args{
				startTime: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
				endTime:   time.Date(2020, time.January, 1, 0, 1, 0, 0, time.UTC),
			},
			0, 0,
		},
		{
			"00:00:00~00:05:00",
			args{
				startTime: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
				endTime:   time.Date(2020, time.January, 1, 0, 5, 0, 0, time.UTC),
			},
			0, 0,
		},
		{
			"00:00:00~00:05:01",
			args{
				startTime: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
				endTime:   time.Date(2020, time.January, 1, 0, 5, 1, 0, time.UTC),
			},
			0, 1,
		},
		{
			"00:01:00~00:06:00",
			args{
				startTime: time.Date(2020, time.January, 1, 0, 1, 0, 0, time.UTC),
				endTime:   time.Date(2020, time.January, 1, 0, 6, 0, 0, time.UTC),
			},
			0, 1,
		},
		{
			"00:06:00~00:07:00",
			args{
				startTime: time.Date(2020, time.January, 1, 0, 6, 0, 0, time.UTC),
				endTime:   time.Date(2020, time.January, 1, 0, 7, 0, 0, time.UTC),
			},
			1, 1,
		},
		{
			"00:06:00~00:30:00",
			args{
				startTime: time.Date(2020, time.January, 1, 0, 6, 0, 0, time.UTC),
				endTime:   time.Date(2020, time.January, 1, 0, 30, 0, 0, time.UTC),
			},
			1, 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Resv{
				StartTime: &tt.args.startTime,
				EndTime:   &tt.args.endTime,
			}
			got, got1 := r.CalculateTimeBits(tt.args.startTime, tt.args.endTime)
			if got != tt.expectedStartBit {
				t.Errorf("Resv.CalculateTimeBits() got StartBit = %v, want StartBit %v", got, tt.expectedStartBit)
			}
			if got1 != tt.expectedEndBit {
				t.Errorf("Resv.CalculateTimeBits() got EndBit = %v, want EndBit %v", got1, tt.expectedEndBit)
			}
		})
	}
}
