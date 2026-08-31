using System.Text.Json;
using System.Text.Json.Nodes;

string json = File.ReadAllText("input/data.json");
JsonNode root = JsonNode.Parse(json);
JsonArray array = root.AsArray();

List<JsonObject> filtered = new();
foreach (JsonObject obj in array.OfType<JsonObject>())
{
    bool active = (bool)obj["active"];
    int age = (int)obj["age"];
    if (active && age >= 30)
    {
        filtered.Add(obj);
    }
}

filtered.Sort((a, b) => string.Compare((string)a["name"], (string)b["name"], StringComparison.Ordinal));

JsonArray result = new();
foreach (JsonObject obj in filtered)
{
    result.Add(obj);
}

string output = result.ToJsonString(new JsonSerializerOptions { WriteIndented = false });
Console.Write(output);