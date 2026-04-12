using System.Text.Json;
using System.Text.Json.Nodes;

var jsonString = File.ReadAllText("input/data.json");
var jsonArray = JsonNode.Parse(jsonString).AsArray();

var filteredJsonArray = jsonArray
    .Where(item => item?["active"].GetValue<bool>() && item?["age"].GetValue<int>() >= 30)
    .OrderBy(item => item?["name"].GetValue<string>());

Console.WriteLine(filteredJsonArray.ToJsonString());