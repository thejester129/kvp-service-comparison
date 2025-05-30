namespace Service.Persistence;

using System.Dynamic;
using Amazon.DynamoDBv2.DocumentModel;
using Newtonsoft.Json;

public class Repository
{
    private readonly Table table;

    public Repository(Table table)
    {
        this.table = table;
    }

    public async Task<dynamic?> Get(string key)
    {
        var doc = await this.table.GetItemAsync(new Primitive(key));

        var defaultModel = new ExpandoObject();
        defaultModel.TryAdd("key", key);

        if (doc == null)
            return defaultModel;

        var item = JsonConvert.DeserializeObject<ExpandoObject>(doc.ToJson(), Service.jsonSettings);

        if (item == null)
            return defaultModel;

        return item;
    }

    public async Task<dynamic?> Put(dynamic model)
    {
        var input = Document.FromJson(JsonConvert.SerializeObject(model, Service.jsonSettings));

        await this.table.PutItemAsync(input);

        return model;
    }

    public async Task<dynamic?> Patch(dynamic model)
    {
        var input = Document.FromJson(JsonConvert.SerializeObject(model));

        var result = await this.table.UpdateItemAsync(
            input,
            new UpdateItemOperationConfig { ReturnValues = ReturnValues.AllNewAttributes }
        );

        return JsonConvert.DeserializeObject<ExpandoObject>(result.ToJson(), Service.jsonSettings);
    }

    public async Task<bool> Delete(string key)
    {
        try
        {
            await table.DeleteItemAsync(new Primitive(key));
            return true;
        }
        catch (Exception ex)
        {
            Console.WriteLine($"Error deleting user with key {key}: {ex.Message}", ex);
            return false;
        }
    }
}
