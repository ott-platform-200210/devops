package main

import (
    "context"
    "fmt"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

var (
    svc        *cognitoidentityprovider.CognitoIdentityProvider
    userPoolID = "us-east-2_5zEkiuvSQ" // Replace with your User Pool ID
    groupName  = "USER"              // Group to assign
)

func init() {
    sess := session.Must(session.NewSession())
    svc = cognitoidentityprovider.New(sess)
}

func handler(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
    fmt.Printf("PostConfirmation for user: %s\n", event.UserName)

    input := &cognitoidentityprovider.AdminAddUserToGroupInput{
        GroupName:  aws.String(groupName),
        UserPoolId: aws.String(userPoolID),
        Username:   aws.String(event.UserName),
    }

    _, err := svc.AdminAddUserToGroup(input)
    if err != nil {
        fmt.Printf("Error adding user to group: %s\n", err)
        return event, err
    }

    fmt.Printf("User %s added to group %s\n", event.UserName, groupName)
    return event, nil
}

func main() {
    lambda.Start(handler)
}

