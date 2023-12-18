package trustme

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func GetCurrentIAMArn(ctx context.Context) (string, error) {
	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return "", err
	}

	client := sts.NewFromConfig(sdkConfig, func(o *sts.Options) {
		o.Region = "ap-northeast-2"
	})

	output, err := client.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return "", err
	}

	return *output.Arn, nil
}
