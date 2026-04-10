using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var file = "input/data.json";
if (!File.Exists(file)) {
    Console.WriteLine("[]");
} else {
    var content = File.ReadAllText(file);
    JsonArray array = Json.Parse(content) as JsonArray;
    
    if (array == null) {
        Console.WriteLine("[]");
    } else {
        List<JsonNode> filtered = new List<JsonNode>();
        
        foreach (var item in array) {
            var name = item["name"]?.GetString();
            var age = item["age"]?.GetInt32();
            var active = item["active"]?.GetBoolean() ?? false;
            var score = item["score"]?.GetFloat(); // We don't need this for filtering
            
            if (active == true && (age != null && age >= 30)) {
                filtered.Add(item);
            }
        }
        
        // Sort by name ascending
        filtered.Sort((a, b) => string.Compare(a["name"]?.GetString() ?? "", b["name"]?.GetString() ?? "", StringComparison.Ordinal));
        
        var json = JsonSerializer.Serialize(filtered, new JsonSerializerOptions { DiscriminatorName = null });
        Console.WriteLine(json);
    }
}