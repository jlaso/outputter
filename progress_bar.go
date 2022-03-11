package main

import (
	"fmt"
	"strings"
)

type Pbar struct {
	Percent int
	Current int64
	Total   int64
	Buff    string
	width   int
}

func NewPbar(total int64, width int) Pbar {
	pb := Pbar{
		Percent: 0,
		Current: 0,
		Total:   total,
		width:   width,
		Buff:    strings.Repeat("·", width),
	}
	return pb
}

func (pb *Pbar) New(total int64, width int) {
	pb.Percent = 0
	pb.Current = 0
	pb.Total = total
	pb.width = width
	pb.Buff = strings.Repeat("·", width)
}

func (pb *Pbar) Inc(i int64) {
	pb.Current += i
	pb.Percent = (int)(100 * pb.Current / pb.Total)
	var curr int
	curr = pb.Percent * pb.width / 100
	pb.Buff = strings.Repeat("=", curr) + strings.Repeat("·", pb.width-curr)
}

func (pb *Pbar) End() {
	pb.Current = pb.Total
	pb.Percent = 100
	pb.Buff = strings.Repeat("=", pb.width)
}

func (pb *Pbar) String() string {
	return fmt.Sprintf("\r[%s]%3d%% %8d/%d", pb.Buff, pb.Percent, pb.Current, pb.Total)
}
