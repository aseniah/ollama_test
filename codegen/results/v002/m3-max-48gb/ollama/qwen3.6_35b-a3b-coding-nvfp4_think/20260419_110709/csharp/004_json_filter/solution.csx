using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

string jsonContent = File.ReadAllText("input/data.json");
JsonNode rootNode = JsonNode.Parse(jsonContent);
JsonArray jsonArray = rootNode.AsArray();

var filteredList = new List<JsonObject>();
foreach (JsonObject obj in jsonArray)
{
    if (obj["active"] != null && obj["age"] != null)
    {
        bool isActive = obj["active"]!.AsValue().GetValue<bool>();
        int age = obj["age"]!.AsValue().GetValue<int>();

        if (isActive && age >= 30)
        {
            filteredList.Add(obj);
        }
    }
}

filteredList.Sort((a, b) => string.Compare(a["name"]?.ToString() ?? "", b["name"]?.ToString() ?? ""));

JsonArray resultArray = new JsonArray();
foreach (var item in filteredList)
{
    resultArray.Add(item.DeepClone());
}

Console.WriteLine(resultArray.ToJsonString());