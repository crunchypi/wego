// Copyright © 2020 wego authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lexvec

import (
	"fmt"
	"runtime"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func invalidRelationTypeError(typ RelationType) error {
	return errors.Errorf("invalid relation type: %s not in %s|%s|%s|%s", typ, PPMI, PMI, Collocation, LogCollocation)
}

type RelationType string

const (
	PPMI                RelationType = "ppmi"
	PMI                 RelationType = "pmi"
	Collocation         RelationType = "co"
	LogCollocation      RelationType = "logco"
	defaultRelationType              = PPMI
)

func (t *RelationType) String() string {
	if *t == RelationType("") {
		*t = defaultRelationType
	}
	return string(*t)
}

func (t *RelationType) Set(name string) error {
	typ := RelationType(name)
	if typ == PPMI || typ == PMI || typ == Collocation || typ == LogCollocation {
		*t = typ
		return nil
	}
	return invalidRelationTypeError(typ)
}

func (t *RelationType) Type() string {
	return t.String()
}

var (
	defaultBatchSize          = 100000
	defaultDim                = 10
	defaultDocInMemory        = false
	defaultGoroutines         = runtime.NumCPU()
	defaultInitlr             = 0.025
	defaultIter               = 15
	defaultMaxCount           = -1
	defaultMinCount           = 5
	defaultNegativeSampleSize = 5
	defaultSmooth             = 0.75
	defaultSubsampleThreshold = 1.0e-3
	defaultTheta              = 1.0e-4
	defaultToLower            = false
	defaultVerbose            = false

	defaultWindow = 5
)

type Options struct {
	BatchSize          int
	Dim                int
	DocInMemory        bool
	Goroutines         int
	Initlr             float64
	Iter               int
	MaxCount           int
	MinCount           int
	NegativeSampleSize int
	RelationType       RelationType
	Smooth             float64
	SubsampleThreshold float64
	Theta              float64
	ToLower            bool
	Verbose            bool

	Window int
}

func DefaultOptions() Options {
	return Options{
		BatchSize:          defaultBatchSize,
		Dim:                defaultDim,
		DocInMemory:        defaultDocInMemory,
		Goroutines:         defaultGoroutines,
		Initlr:             defaultInitlr,
		Iter:               defaultIter,
		MaxCount:           defaultMaxCount,
		MinCount:           defaultMinCount,
		NegativeSampleSize: defaultNegativeSampleSize,
		RelationType:       defaultRelationType,
		Smooth:             defaultSmooth,
		SubsampleThreshold: defaultSubsampleThreshold,
		Theta:              defaultTheta,
		ToLower:            defaultToLower,
		Verbose:            defaultVerbose,
		Window:             defaultWindow,
	}
}
func LoadForCmd(cmd *cobra.Command, opts *Options) {
	cmd.Flags().IntVar(&opts.BatchSize, "batch", defaultBatchSize, "batch size to train")
	cmd.Flags().IntVarP(&opts.Dim, "dim", "d", defaultDim, "dimension for word vector")
	cmd.Flags().IntVar(&opts.Goroutines, "goroutines", defaultGoroutines, "number of goroutine")
	cmd.Flags().BoolVar(&opts.DocInMemory, "in-memory", defaultDocInMemory, "whether to store the doc in memory")
	cmd.Flags().Float64Var(&opts.Initlr, "initlr", defaultInitlr, "initial learning rate")
	cmd.Flags().IntVar(&opts.Iter, "iter", defaultIter, "number of iteration")
	cmd.Flags().IntVar(&opts.MaxCount, "max-count", defaultMaxCount, "upper limit to filter words")
	cmd.Flags().IntVar(&opts.MinCount, "min-count", defaultMinCount, "lower limit to filter words")
	cmd.Flags().IntVar(&opts.NegativeSampleSize, "sample", defaultNegativeSampleSize, "negative sample size")
	cmd.Flags().Var(&opts.RelationType, "rel", fmt.Sprintf("relation type for co-occurrence words. One of %s|%s|%s|%s", PPMI, PMI, Collocation, LogCollocation))
	cmd.Flags().Float64Var(&opts.Smooth, "smooth", defaultSmooth, "smoothing value for co-occurence value")
	cmd.Flags().Float64Var(&opts.SubsampleThreshold, "threshold", defaultSubsampleThreshold, "threshold for subsampling")
	cmd.Flags().Float64Var(&opts.Theta, "theta", defaultTheta, "lower limit of learning rate (lr >= initlr * theta)")
	cmd.Flags().BoolVar(&opts.ToLower, "to-lower", defaultToLower, "whether the words on corpus convert to lowercase or not")
	cmd.Flags().BoolVar(&opts.Verbose, "verbose", defaultVerbose, "verbose mode")
	cmd.Flags().IntVarP(&opts.Window, "window", "w", defaultWindow, "context window size")

}

type ModelOption func(*Options)

func BatchSize(v int) ModelOption {
	return ModelOption(func(opts *Options) {
		opts.BatchSize = v
	})
}

func DocInMemory() ModelOption {
	return ModelOption(func(opts *Options) {
		opts.DocInMemory = true
	})
}

func Goroutines(v int) ModelOption {
	return ModelOption(func(opts *Options) {
		opts.Goroutines = v
	})
}

func Dim(v int) ModelOption {
	return ModelOption(func(opts *Options) {
		opts.Dim = v
	})
}

func Initlr(v float64) ModelOption {
	return ModelOption(func(opts *Options) {
		opts.Initlr = v
	})
}

func Iter(v int) ModelOption {
	return ModelOption(func(opts *Options) {
		opts.Iter = v
	})
}

func MaxCount(v int) ModelOption {
	return ModelOption(func(opts *Options) {
		opts.MaxCount = v
	})
}

func MinCount(v int) ModelOption {
	return ModelOption(func(opts *Options) {
		opts.MinCount = v
	})
}

func NegativeSampleSize(v int) ModelOption {
	return ModelOption(func(opts *Options) {
		opts.NegativeSampleSize = v
	})
}

func Relation(typ RelationType) ModelOption {
	return ModelOption(func(opts *Options) {
		opts.RelationType = typ
	})
}

func Smooth(v float64) ModelOption {
	return ModelOption(func(opts *Options) {
		opts.Smooth = v
	})
}

func SubsampleThreshold(v float64) ModelOption {
	return ModelOption(func(opts *Options) {
		opts.SubsampleThreshold = v
	})
}

func Theta(v float64) ModelOption {
	return ModelOption(func(opts *Options) {
		opts.Theta = v
	})
}

func ToLower() ModelOption {
	return ModelOption(func(opts *Options) {
		opts.ToLower = true
	})
}

func Verbose() ModelOption {
	return ModelOption(func(opts *Options) {
		opts.Verbose = true
	})
}

func Window(v int) ModelOption {
	return ModelOption(func(opts *Options) {
		opts.Window = v
	})
}
