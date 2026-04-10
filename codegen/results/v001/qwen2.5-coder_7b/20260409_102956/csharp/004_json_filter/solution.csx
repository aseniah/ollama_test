using System.Text.Json;
using System.Text.Json.Nodes;

JsonArray data = JsonNode.Parse(File.ReadAllText("input/data.json")).AsArray();
var filteredData = data.Where(item => item["active"].GetValue<bool>() && item["age"].GetInt32() >= 30)
                       .OrderBy(item => item["name"].GetString());
JsonDocument resultDoc = JsonDocument.Create(new JsonArray(filteredData));
Console.WriteLine(resultDoc.RootElement.GetRawText());