using System.Text.Json;
using System.Text.Json.Nodes;

string inputPath = "input/data.json";
string jsonContent = File.ReadAllText(inputPath);

JsonArray records = JsonArray.Parse(jsonContent);

var filtered = new List<JsonObject>();

foreach (JsonNode node in records)
{
    var obj = node.AsObject();
    bool active = (bool)obj["active"];
    int age = (int)obj["age"];

    if (active && age >= 30)
    {
        filtered.Add(obj);
    }
}

// Sort by name ascending
filtered.Sort((a, b) => string.Compare((string)a["name"], (string)b["name"], StringComparison.Ordinal));

// Write output
var resultArray = new JsonArray();
foreach (var obj in filtered)
{
    resultArray.Add(obj.DeepClone());
}

Console.WriteLine(resultArray.ToJsonString(new JsonSerializerOptions { WriteIndented = false }));