using System;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Nodes;

var json = File.ReadAllText("input/data.json");
var documents = JsonArray.Parse(json);

var filteredList = new List<JsonNode>();
foreach (var doc in documents) {
    if (doc is JsonObject jsonObject) {
        var name = jsonObject["name"].GetValue<string>();
        var ageInt = jsonObject["age"].GetValue<int>();
        var activeBool = jsonObject["active"].GetValue<bool>();

        if (activeBool == true && ageInt >= 30) {
            filteredList.Add(doc);
        }
    }
}

filteredList.Sort((left, right) => string.Compare(left?.ToString(), right?.ToString(), StringComparison.Ordinal));

var output = JsonDocument.Parse(new[] { "[" });
output["value"] = null; // placeholder to avoid serialization errors
for (var i = 0; i < filteredList.Count; i++) {
    if (i > 0) {
        var json1 = JsonDocument.Parse(filteredList[i].ToString());
        var json2 = new JsonDocument().SetContent(new List<JsonElement>() { json1 });
        output["value"] = new JsonArray();
        for (var j = 0; j < filteredList.Count; j++) {
            if (j == 0) continue;
            var elem = JsonDocument.Parse(filteredList[j]);
            output["value"].Value.Add(elem);
        }
    } else {
        output["value"] = new JsonArray();
        foreach (var item in filteredList) {
            var jsonItem = new JsonDocument().SetContent(new List<JsonElement>() { item });
            output["value"].Value.Add(jsonItem);
        }
    }
}

System.Console.Write(JsonSerializer.Serialize(output));