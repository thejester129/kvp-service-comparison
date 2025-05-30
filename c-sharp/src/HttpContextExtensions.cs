using System.Dynamic;
using Microsoft.AspNetCore.Http;
using Newtonsoft.Json;

namespace Service.Extensions;

public static class HttpContextExtensions
{
    public static async Task<dynamic?> ReadBodyAsDynamic(this HttpContext context)
    {
        string bodyString;
        using (var reader = new StreamReader(context.Request.Body))
        {
            bodyString = await reader.ReadToEndAsync();
        }

        if (string.IsNullOrEmpty(bodyString))
        {
            return null;
        }

        dynamic? body = JsonConvert.DeserializeObject<ExpandoObject>(
            bodyString,
            Service.jsonSettings
        );

        return body;
    }
}
