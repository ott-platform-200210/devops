mport json
import boto3

# Initialize AWS Cognito client
cognito_client = boto3.client('cognito-idp')

# Replace with your Cognito User Pool ID
USER_POOL_ID = "us-east-2_5zEkiuvSQ"
DEFAULT_GROUP = "USER"  # Change this as needed

def lambda_handler(event, context):
    print("Event received: ", json.dumps(event))

    try:
        # Extract user attributes
        user_email_verified = event["request"]["userAttributes"].get("email_verified", "false")
        username = event["userName"]

        # Only add user to group if email is verified
        if user_email_verified.lower() == "true":
            response = cognito_client.admin_add_user_to_group(
                UserPoolId=USER_POOL_ID,
                Username=username,
                GroupName=DEFAULT_GROUP
            )
            print(f"User {username} added to group {DEFAULT_GROUP}")

        return event  # Cognito expects the event object to be returned

    except Exception as e:
        print(f"Error processing post confirmation: {str(e)}")
        raise

