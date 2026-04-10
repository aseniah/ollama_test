using System;
using System.Text.Json;
using System.Text.Json.Nodes;

string json = File.ReadAllText("input/data.json");
JsonArray data = JsonNode.Parse(json)!.AsArray();

var filtered = from record in data
               where record["active"] is true && record["age"] is int && ((int)record["age"]) >= 30
               select record

              orderby record["name"]
              into sorted

              from sorted in sorted

JsonSerializer.Serialize(sorted.ToArray(), new JsonSerializerOptions { WriteIndented = true });