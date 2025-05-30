import boto3

resource = boto3.resource('dynamodb', endpoint_url='http://localhost:4566', region_name='us-east-1')
table = resource.Table("kvp-table")

def get_item(key):
    response = table.get_item(
        Key={
            'key': key
        }
    )
    return response.get('Item', None)

def put_item(item):
    response = table.put_item(
                Item=item
    ) 
    if 'ResponseMetadata' in response and response['ResponseMetadata']['HTTPStatusCode'] == 200: 
        return item
    return None

def delete_item(key):
    response = table.delete_item(
        Key={
            'key': key
        }
    )
    if 'ResponseMetadata' in response and response['ResponseMetadata']['HTTPStatusCode'] == 200: 
        return True
    return False


