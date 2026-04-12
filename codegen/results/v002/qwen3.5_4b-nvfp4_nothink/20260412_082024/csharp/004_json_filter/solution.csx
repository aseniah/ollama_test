var content = File.ReadAllText("input/data.json");
using var jsonDocument = JsonDocument.Parse(content);
var allData = jsonDocument.RootElement.GetProperty("elements").EnumerateArray().Select(e => {
    return new { Name = e.GetNamedIntProperty("name"), Age = e["age"].GetInt32(), Active = e["active"].GetBoolean(), Score = e["score"].GetFloat() };
});
var result = allData.Where(r => r.Active && r.Age >= 30).OrderBy(r => r.Name);
var outputList = new System.Text.Json.Serialization.JavaScriptArray();
foreach (var item in result) {
    outputList.Add(new System.Text.Json.Serialization.JsonStringValue(item.Name));
}
Console.Write(System.Text.Json.JsonSerializer.Serialize(outputList, options: new System.Text.Json.JsonSerializerOptions()));