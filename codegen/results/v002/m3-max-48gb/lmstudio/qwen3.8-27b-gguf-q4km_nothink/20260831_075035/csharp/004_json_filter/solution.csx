using System.Text.Json;
using System.Text.Json.Nodes;

string json = File.ReadAllText("input/data.json");
JsonArray arr = JsonArray.Parse(json);

var filtered = new List<JsonObject>();
foreach (var item in arr)
{
    var obj = (JsonObject)item;
    bool active = (bool)obj["active"]!;
    int age = (int)obj["age"]!;
    if (active && age >= 30)
    {
        filtered.Add(obj);
    }
}

// Sort by name ascending
filtered.Sort((a, b) => string.Compare((string)a["name"]!, (string)b["name"]!, StringComparison.Ordinal));

JsonArray result = new JsonArray();
foreach (var obj in filtered)
{
    result.Add(obj.DeepClone());
}

Console.WriteLine(result.ToJsonString(new JsonSerializerOptions { WriteIndented = true }));