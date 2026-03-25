var builder = WebApplication.CreateBuilder(args);

var portEnv = Environment.GetEnvironmentVariable("PORT");
var port = string.IsNullOrEmpty(portEnv) ? 8080 : int.Parse(portEnv);

builder.WebHost.ConfigureKestrel(options =>
{
    options.ListenAnyIP(port);
});

var app = builder.Build();
Console.WriteLine($"Server started in port {port}");
app.MapGet("/", () => "Hello world!");

app.Run();
