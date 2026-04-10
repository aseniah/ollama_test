using System;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

var args = Args as IList<string> ?? new List<string>();
if (args.Count == 0 || !args[0].Equals("input/data.json", StringComparison.OrdinalIgnoreCase)) {
    Environment.Exit(1);
}

string json = File.ReadAllText(args[0]);
JsonElement root = JsonDocument.Parse(json).RootElement;

var result = root.EnumerateArray().Where(elem => {
    var name = (string)!(elem.GetProperty("name") ?? null);
    if (string.IsNullOrEmpty(name)) return false;
    int age = (int)(elem.GetProperty("age") ?? 0);
    bool active = (bool)(elem.GetProperty("active") ?? false);
    return active && age >= 30;
})
.OrderBy(elem => (string)!(elem.GetProperty("name") ?? null))
.ToArray();

if (result.Length == 0) {
    Console.WriteLine("[]");
} else {
    var jsonArray = new JsonArray();
    foreach (var item in result) {
        object data = new { 
            Name = (string)!(item.GetProperty("name") ?? null),
            Age = (int)(item.GetProperty("age") ?? 0),
            Active = (bool)(item.GetProperty("active") ?? false),
            Score = (double)(item.GetProperty("score") ?? 0f)
        };
        string jsonStr = JsonSerializer.Serialize(data);
        JsonArray AddItem(JsonArray array, string str) {
            if (array.IsList && array.Contains(System.Text.Json.JsonElement.ParseValue(jsonStr))) return array;
            object? newJsonNode = JsonDocument.Parse(jsonStr).RootElement;
            Array.Add(array, newJsonNode);
            return array;
        }
        System.Runtime.CompilerServices.Extension.IgnoreCompile();
    }
}