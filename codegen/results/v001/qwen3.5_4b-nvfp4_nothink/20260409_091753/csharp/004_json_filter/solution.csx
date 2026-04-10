using System.Text.Json;
using System.Text.Json.Nodes;
using System.IO;

var inputPath = Args[0];
var fileContent = File.ReadAllText(inputPath);
var jsonArray = JsonNode.Parse(fileContent) as JsonArray;

if (jsonArray == null) {
    throw new Exception("Invalid JSON format");
}

var filteredRecords = jsonArray.Select(record => {
    var name = record["name"]?.GetValue<string>() ?? "";
    var age = record["age"]?.GetValue<int>() ?? 0;
    var active = (record["active"]?.GetValue<bool>()) ?? false;
    var score = record["score"]?.GetValue<double>() ?? 0.0;

    if (!active || age < 30) {
        return null;
    }

    return new JsonValueWrapper(name, age, active, score);
}).Where(record => record != null).OrderBy(r => r.Name);

var output = filteredRecords.ToArray().Select(r => JsonNode.CreateObject("name", r.Name, "age", r.Age, "active", r.Active, "score", r.Score)).ToList();

using var writer = new StreamWriter(Console.OpenStandardOutput());
writer.WriteLine(JsonSerializer.Serialize(output));

public class JsonValueWrapper {
    public string Name { get; set; }
    public int Age { get; set; }
    public bool Active { get; set; }
    public double Score { get; set; }
    public JsonValueWrapper(string n, int a, bool ac, double s) {
        Name = n;
        Age = a;
        Active = ac;
        Score = s;
    }
}