package trustme

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"os/exec"
	"strings"
)

// (TODO) AssumeRole 하지 않고, Assume 해서 특정 커맨드를 실행하고, 프로그램 종료될 때 해당 TR policy를 지워줘도 좋겠다.
// ex) trustme run irsa--alpha-kr--ads-account -- aws sts get-caller-identity

func AssumeRole(ctx context.Context, roleName string, cmdStr string) (string, error) {
	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return "", err
	}

	// (TODO) AWS Client는 ctx에 넣어서 사용할까?
	client := sts.NewFromConfig(sdkConfig, func(o *sts.Options) {
		o.Region = "ap-northeast-2"
	})

	// (TODO) arn 형식이면 그대로 사용하고, 아니면 GetCallerIdentity 한 값에서 roleName 만 치환한다.
	roleARN := fmt.Sprintf("arn:aws:iam::314695318048:role/%s", roleName)
	roleSessionName := fmt.Sprintf("trustme-%s", roleName)

	output, err := client.AssumeRole(ctx, &sts.AssumeRoleInput{
		RoleArn:         &roleARN,
		RoleSessionName: &roleSessionName,
	})
	if err != nil {
		return "", err
	}

	//fmt.Println(*output.Credentials.AccessKeyId)
	//fmt.Println(*output.Credentials.SecretAccessKey)
	//fmt.Println(*output.Credentials.SessionToken)

	//os.Setenv("AWS_ACCESS_KEY_ID", *output.Credentials.AccessKeyId)
	//os.Setenv("AWS_SECRET_ACCESS_KEY", *output.Credentials.SecretAccessKey)
	//os.Setenv("AWS_SESSION_TOKEN", *output.Credentials.SessionToken)

	//cmd := exec.Command("aws", "sts", "get-caller-identity")
	args := strings.Split(cmdStr, " ")
	cmd := exec.Command(args[0], args[1:]...)

	a := fmt.Sprintf("%s=%s", "AWS_ACCESS_KEY_ID", *output.Credentials.AccessKeyId)
	b := fmt.Sprintf("%s=%s", "AWS_SECRET_ACCESS_KEY", *output.Credentials.SecretAccessKey)
	c := fmt.Sprintf("%s=%s", "AWS_SESSION_TOKEN", *output.Credentials.SessionToken)

	cmd.Env = []string{a, b, c}

	out, err := cmd.CombinedOutput()
	if err != nil {
		//return "", err
	}

	fmt.Println(strings.TrimSpace(string(out)))

	return *output.AssumedRoleUser.Arn, nil
}
