using System.Text.Json;
using System.Text.Json.Nodes;

string json = File.ReadAllText("input/data.json");
JsonArray arr = JsonArray.Parse(json);

var filtered = new JsonArray();
foreach (JsonNode node in arr)
{
    var obj = (JsonObject)node;
    bool active = (bool)obj["active"];
    int age = (int)obj["age"];
    if (active && age >= 30)
    {
        filtered.Add(obj);
    }
}

var sorted = filtered.OrderBy(o => (string)((JsonObject)o)["name"]).ToList();

var output = new JsonArray();
foreach (var item in sorted)
{
    output.Add(item);
}

string result = output.ToJsonString(new JsonSerializerOptions { WriteIndented = true });
Console.WriteLine(result);