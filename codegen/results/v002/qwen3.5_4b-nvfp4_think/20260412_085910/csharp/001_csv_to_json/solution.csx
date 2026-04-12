var lines = File.ReadAllLines("input/data.csv");
var writer = new TextWriter(Console.Out);
var jsonOptions = new JsonSerializerOptions { WriteIndented = true };

for (int i = 1; i < lines.Length; i++) {
    var fields = lines[i].Split(',');
    if (fields.Length == 4) {
        for (var j = 0; j < 4; j++) {
            // Remove whitespace from each field
            var trimmed = fields[j].Trim();
            if (j == 1 && int.TryParse(trimmed, out int age)) {
                fields[j] = age.ToString();
            } else if (j == 3 && float.TryParse(trimmed, out float score)) {
                fields[j] = score.ToString();
            }
        }
    }
    
    var obj = new System.Text.Json.JsonElement();
    obj.Add("Name", fields[0]);
    obj.Add("Age", int.Parse(fields[1]));
    obj.Add("Email", fields[2]);
    obj.Add("Score", float.Parse(fields[3]));
    
    writer.Write($"@{obj}");
}

writer.WriteLine("");
writer.Flush();