using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

string jsonString = File.ReadAllText("input/data.json");
JsonArray jsonArray = JsonValue.Parse(jsonString).AsArray();

JsonArray filteredJsonArray = new JsonArray();
foreach (JsonObject jsonObject in jsonArray)
{
    if (jsonObject["active"].GetValue<bool>() && jsonObject["age"].GetInt32() >= 30)
    {
        filteredJsonArray.Add(jsonObject);
    }
}

JsonDocument sortedJsonDocument = JsonDocument.Parse(JsonSerializer.Serialize(filteredJsonArray));
using (JsonWriter writer = new Utf8JsonWriter(Console.OpenStandardOutput()))
{
    JsonSerializer.Serialize(writer, sortedJsonDocument.RootElement, new JsonSerializerOptions { WriteIndented = false });
}