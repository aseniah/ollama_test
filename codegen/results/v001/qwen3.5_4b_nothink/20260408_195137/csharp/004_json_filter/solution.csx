using System.Text.Json;
using System.Text.Json.Nodes;

var inputPath = "input/data.json";
var fileContent = System.IO.File.ReadAllText(inputPath);
var root = JsonNode.Parse(fileContent);

var allData = root.Deserialize<JsonArray>();
var filteredData = allData.Where(item => 
    ((JsonNode)item)["active"] == true &&
    ((JsonNode)item)["age"] >= 30)
    .OrderBy(item => ((JsonNode)item)["name"])
    .ToArray();

var output = System.Text.Json.JsonSerializer.Serialize(filteredData, new System.Text.Json.JsonSerializerOptions { WriteIndented = false });
System.Console.WriteLine(output);