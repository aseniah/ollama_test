using System.Text.Json;
using System.Text.Json.Nodes;

var data = File.ReadAllText("input/data.json");
var array = JsonArray.Parse(data);

var filteredArray = array.Where(obj => obj["active"] is true && (int)obj["age"] >= 30)
                         .OrderBy(obj => obj["name"].ToString())
                         .ToJsonArray();

Console.WriteLine(filteredArray);