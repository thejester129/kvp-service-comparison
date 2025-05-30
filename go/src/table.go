package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var client *dynamodb.Client
var tableName string = "kvp-table"

type TableParams struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func loadTable() {
	var config, err = config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("test", "test", "")),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	client = dynamodb.NewFromConfig(config, func(o *dynamodb.Options) {
		o.BaseEndpoint = aws.String("http://localhost:4566")
	})

	describeTable()
}

func describeTable() {
	_, err := client.DescribeTable(
		context.TODO(), &dynamodb.DescribeTableInput{TableName: aws.String(tableName)},
	)
	if err != nil {
		log.Fatalf("can't find table, %v", err)
		return
	}
	println("found table")
}

func getAllTableItems(ctx context.Context) ([]JsonItem, error) {
	var items []JsonItem = make([]JsonItem, 0)
	var err error
	var response *dynamodb.ScanOutput
	scanPaginator := dynamodb.NewScanPaginator(client, &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})
	for scanPaginator.HasMorePages() {
		response, err = scanPaginator.NextPage(ctx)
		if err != nil {
			break
		} else {
			var newItems []JsonItem
			err = attributevalue.UnmarshalListOfMaps(response.Items, &newItems)
			if err != nil {
				log.Printf("Couldn't unmarshal query response. Here's why: %v\n", err)
				break
			} else {
				items = append(items, newItems...)
			}
		}
	}
	return items, err
}

func getTableItem(ctx context.Context, key string) (JsonItem, error) {
	var item JsonItem

	response, err := client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"key": &types.AttributeValueMemberS{Value: key},
		},
	})

	if err != nil {
		log.Printf("Failed to get %v, :%v\n", key, err)
		return nil, err
	}

	err = attributevalue.UnmarshalMap(response.Item, &item)

	if err != nil {
		log.Printf("Failed to unmarshal: %v\n", err)
	}

	return item, err
}

func putTableItem(ctx context.Context, item JsonItem) error {
	dbItem, err := attributevalue.MarshalMap(item)
	if err != nil {
		panic(err)
	}
	_, err = client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(tableName), Item: dbItem,
	})
	if err != nil {
		log.Printf("Couldn't add item to table. Here's why: %v\n", err)
	}
	return err
}

func updateTableItem(ctx context.Context, item JsonItem) (JsonItem, error) {
	var err error
	var response *dynamodb.UpdateItemOutput
	var update expression.UpdateBuilder
	var updatedItem JsonItem

	for key, value := range item {
		if key != "key" {
			update = update.Set(expression.Name(key), expression.Value(value))
		}
	}

	expr, err := expression.NewBuilder().WithUpdate(update).Build()

	if err != nil {
		log.Printf("Couldn't build expression for update. Here's why: %v\n", err)
	} else {
		response, err = client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
			TableName:                 aws.String(tableName),
			Key:                       map[string]types.AttributeValue{"key": &types.AttributeValueMemberS{Value: item["key"].(string)}},
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			UpdateExpression:          expr.Update(),
			ReturnValues:              types.ReturnValueAllNew,
		})
		if err != nil {
			log.Printf("Couldn't update item: %v\n", err)
		} else {
			err = attributevalue.UnmarshalMap(response.Attributes, &updatedItem)
			if err != nil {
				log.Printf("Couldn't unmarshall update response. Here's why: %v\n", err)
			}
		}
	}

	return updatedItem, err
}

func deleteTableItem(ctx context.Context, key string) error {
	_, err := client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName), Key: map[string]types.AttributeValue{
			"key": &types.AttributeValueMemberS{Value: key},
		}})
	if err != nil {
		log.Printf("Couldn't delete %v from the table. Here's why: %v\n", key, err)
	}
	return err
}
