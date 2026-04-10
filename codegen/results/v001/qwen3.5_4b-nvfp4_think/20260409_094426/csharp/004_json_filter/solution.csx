using System.Text.Json;
using System.IO;

var data = File.ReadAllText("input/data.json");
var document = JsonDocument.Parse(data);

var records = new List<object>();
foreach (JsonElement element in document.Array!)
{
    var record = new Dictionary<string, object>
    {
        ["name"] = element.GetProperty("name").GetString(),
        ["age"] = element.GetProperty("age").GetInt32(),
        ["active"] = element.GetProperty("active").GetBool(),
        ["score"] = element.GetProperty("score").GetDouble()
    };
    records.Add(record);
}

var filtered = records.Where(r => (bool)r["active"] && r["age"] >= 30)
                     .OrderBy(r => r["name"])
                     .ToList();

Console.WriteLine(JsonSerializer.Serialize(filtered, new JsonSerializerOptions { WriteIndented = false }));