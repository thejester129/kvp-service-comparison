namespace Service;

using Amazon.DynamoDBv2;
using Amazon.DynamoDBv2.DocumentModel;
using global::Service.Extensions;
using global::Service.Persistence;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Routing;
using Newtonsoft.Json;
using Newtonsoft.Json.Serialization;

public class Service
{
    private readonly WebApplication app;
    private readonly Repository repository;
    public static JsonSerializerSettings jsonSettings = new()
    {
        ContractResolver = new CamelCasePropertyNamesContractResolver(),
        NullValueHandling = NullValueHandling.Ignore,
        DateTimeZoneHandling = DateTimeZoneHandling.Utc,
    };

    public Service(WebApplication app)
    {
        this.app = app;

        var config = new AmazonDynamoDBConfig();
        config.ServiceURL = "http://localhost:4566";
        var client = new AmazonDynamoDBClient(config);
        this.repository = new Repository(Table.LoadTable(client, "kvp-table"));

        this.MapRoutes();
    }

    public void Start()
    {
        app.Urls.Add("http://+:8080");
        app.Run();
    }

    private void MapRoutes()
    {
        app.MapGet("/{key}", Get);
        app.MapPut("/{key}", Put);
        app.MapPatch("/{key}", Patch);
        app.MapDelete("/{key}", Delete);
    }

    private IResult Get(HttpContext context)
    {
        var key = GetKey(context);

        var data = repository.Get(key).Result;

        if (data == null)
            return Results.NotFound();

        return Results.Ok(data);
    }

    private IResult Put(HttpContext context)
    {
        var body = context.ReadBodyAsDynamic().Result;

        if (body == null)
            return Results.BadRequest();

        var key = GetKey(context);
        body.key = key;

        var result = repository.Put(body).Result;

        if (result == null)
        {
            return Results.Problem();
        }

        return Results.Ok(result);
    }

    private IResult Patch(HttpContext context)
    {
        var body = context.ReadBodyAsDynamic().Result;

        if (body == null)
            return Results.BadRequest();

        var key = GetKey(context);
        body.key = key;

        var result = repository.Patch(body).Result;

        if (result == null)
            return Results.Problem();

        return Results.Ok(result);
    }

    private IResult Delete(HttpContext context)
    {
        var key = GetKey(context);

        var result = repository.Delete(key).Result;

        if (!result)
            return Results.Problem();

        return Results.Ok();
    }

    private string GetKey(HttpContext context) =>
        context.GetRouteValue("key")?.ToString() ?? throw new ArgumentException("Key is required");
}
