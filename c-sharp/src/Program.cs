using Autofac.Extensions.DependencyInjection;
using Microsoft.AspNetCore.Builder;
using Microsoft.Extensions.DependencyInjection;

namespace Service;

public class Program
{
    public static void Main(string[] args)
    {
        StartServer(args);
    }

    public static void StartServer(string[] args)
    {
        var builder = WebApplication.CreateBuilder(args);

        builder.Services.AddCors();

        using (var app = builder.Build())
        {
            app.UseCors(builder =>
                builder
                    .SetIsOriginAllowed(_ => true)
                    .AllowAnyMethod()
                    .AllowAnyHeader()
                    .AllowCredentials()
            );

            try
            {
                var service = new Service(app);
                service.Start();
            }
            catch (Exception e)
            {
                Console.WriteLine(e);
            }
        }
    }
}
