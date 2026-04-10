using System.Text.Json;
using System.Text.Json.Nodes;

var jsonString = File.ReadAllText("input/data.json");
var dataArray = JsonSerializer.Deserialize<JsonArray>(jsonString);

var filteredData = dataArray
    .Where(item => item["active"].GetBoolean() && item["age"].GetInt32() >= 30)
    .OrderBy(item => item["name"].GetString())
    .ToJsonArray();

Console.WriteLine(filteredData);