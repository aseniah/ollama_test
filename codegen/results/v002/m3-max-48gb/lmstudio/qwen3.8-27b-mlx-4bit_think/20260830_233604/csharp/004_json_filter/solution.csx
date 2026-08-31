using System.Text.Json;
using System.Text.Json.Nodes;

string json = File.ReadAllText("input/data.json");
JsonArray arr = JsonArray.Parse(json);

var filtered = new JsonArray();
foreach (JsonNode node in arr)
{
    var obj = node as JsonObject;
    bool active = obj["active"].GetValue<bool>();
    int age = obj["age"].GetValue<int>();
    if (active && age >= 30)
        filtered.Add(obj);
}

var sorted = filtered
    .OrderBy(x => ((x as JsonObject)["name"]).GetValue<string>(), StringComparer.Ordinal)
    .ToList();

JsonArray result = new JsonArray();
foreach (var item in sorted)
    result.Add(item);

var options = new JsonSerializerOptions { WriteIndented = true };
Console.Write(result.ToJsonString(options));