import { DynamoDBClient } from "@aws-sdk/client-dynamodb";
import {
  DeleteCommand,
  DynamoDBDocument,
  DynamoDBDocumentClient,
  GetCommand,
  PutCommand,
} from "@aws-sdk/lib-dynamodb";

export default class Repository {
  private dbClient: DynamoDBDocumentClient;
  private tableName: string;

  constructor(tableName: string) {
    const client = new DynamoDBClient({
      region: "us-east-1",
      endpoint: "http://localhost:4566",
    });
    this.dbClient = DynamoDBDocumentClient.from(client);
    DynamoDBDocument;
    this.tableName = tableName;
  }

  async get(key: string) {
    const command = new GetCommand({
      TableName: this.tableName,
      Key: {
        key,
      },
    });

    const response = await this.dbClient.send(command);

    return response?.Item ?? null;
  }

  async put(item) {
    const command = new PutCommand({
      TableName: this.tableName,
      Item: {
        ...item,
      },
    });

    const response = await this.dbClient.send(command);

    if (response?.$metadata.httpStatusCode === 200) {
      return item;
    }

    return null;
  }

  async patch(item: any): Promise<boolean> {
    const current = await this.get(item.key);
    const patched = {
      ...current,
      ...item,
    };
    const command = new PutCommand({
      TableName: this.tableName,
      Item: {
        ...patched,
      },
    });

    const response = await this.dbClient.send(command);

    if (response?.$metadata.httpStatusCode === 200) {
      return true;
    }

    return false;
  }

  async delete(key: string): Promise<boolean> {
    const command = new DeleteCommand({
      TableName: this.tableName,
      Key: {
        key,
      },
    });

    const response = await this.dbClient.send(command);

    if (response?.$metadata.httpStatusCode === 200) {
      return true;
    }

    return false;
  }
}
