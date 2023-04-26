import json
import os

import boto3

sns = boto3.client("sns", region_name="us-west-2")

topic_arn = os.environ.get("SNS_TOPIC_ARN")
if not topic_arn:
    print("Missing SNS_TOPIC_ARN environment variable")
    exit(1)

message = {"default": "test message"}

try:
    response = sns.publish(
        TopicArn=topic_arn,
        Message=json.dumps({"default": json.dumps(message)}),
        MessageStructure="json",
    )
    print(f"Message published with ID: {response['MessageId']}")
except Exception as e:
    print(f"Failed to publish message: {e}")


print(response["ResponseMetadata"]["HTTPSStatusCode"])
