package jobs

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/ratelimit"
	"golang.org/x/sync/errgroup"

	"github.com/smartcontractkit/chainlink/deployment"
	"github.com/smartcontractkit/chainlink/deployment/environment/devenv"

	types "github.com/smartcontractkit/chainlink/system-tests/lib/cre/types"
)

func Create(offChainClient deployment.OffchainClient, don *devenv.DON, flags []string, jobSpecs types.DonJobs) error {
	if len(jobSpecs) == 0 {
		return nil
	}

	eg := &errgroup.Group{}
	jobRateLimit := ratelimit.New(5)

	for jobDesc, jobReqs := range jobSpecs {
		for _, jobReq := range jobReqs {
			eg.Go(func() error {
				jobRateLimit.Take()
				timeout := time.Second * 60
				ctx, cancel := context.WithTimeout(context.Background(), timeout)
				defer cancel()
				_, err := offChainClient.ProposeJob(ctx, jobReq)
				if err != nil {
					return errors.Wrapf(err, "failed to propose job %s for node %s", jobDesc.Flag, jobReq.NodeId)
				}
				if ctx.Err() != nil {
					return errors.Wrapf(err, "timed out after %s proposing job %s for node %s", timeout.String(), jobDesc.Flag, jobReq.NodeId)
				}

				return nil
			})
		}
	}

	if err := eg.Wait(); err != nil {
		return errors.Wrap(err, "failed to create at least one job for DON")
	}

	return nil
}

func calculateJobCount(jobSpecs types.DonJobs) int {
	count := 0
	for _, jobSpec := range jobSpecs {
		count += len(jobSpec)
	}

	return count
}
