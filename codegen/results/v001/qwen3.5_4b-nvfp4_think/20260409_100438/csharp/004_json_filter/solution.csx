using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json.Nodes;

var json = File.ReadAllText("input/data.json");
var data = JsonArray.Parse(json);

var results = data
    .Where(item => 
        // Convert active to bool if it's a boolean node
        (bool)item["active"] && 
        // age is 30 or older, handle both int and possible float conversion
        (int)item["age"] >= 30
    )
    .OrderBy(item => item["name"])
    .ToList();

Console.WriteLine(JsonArray.Serialize(results));