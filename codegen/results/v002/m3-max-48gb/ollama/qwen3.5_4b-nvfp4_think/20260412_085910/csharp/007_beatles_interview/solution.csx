System.Console.Write("Processing...");
var csvData = System.File.ReadAllText("input/input.csv");
var lines = csvData.Split('\n');
var header = lines[0].Split(',').Select(c => c.Trim()).ToArray();

var rows = new List<Dictionary<string, string>>();
for (int i = 1; i < lines.Length; i++) {
    var line = lines[i];
    if (string.IsNullOrWhiteSpace(line)) continue;
    var cols = line.Split(',');
    var row = new Dictionary<string, string>(cols.Length);
    for (int j = 0; j < cols.Length && j < header.Length; j++) {
        row[header[j]] = cols[j];
    }
    rows.Add(row);
}

var jsonConverter = System.Text.Json.JsonDocumentSerializer?.ToJson ?? (x => JsonConvert(x));
var relativeOutput = System.Text.Json.Nodes.JsonNode.FromValue(System.Object, System.Collections.Generic.List<JsonProperty>, System.String, System.Int32, System.DateTime).ToList();

foreach (var row in rows) {
    var parts = row["Name"] ?? string.Empty;
    var nameParts = parts.Split(' ').ToArray();
    var firstName = (nameParts.Length >= 1) ? nameParts[0] : "";
    var lastName = (nameParts.Length >= 2) ? nameParts[^1].Trim() : "";

    var birthdayPart = row["Birthday"] ?? string.Empty;
    var ageDays = System.DateTime.ParseExact(birthdayPart, "d/M/yyyy", System.Globalization.CultureInfo.InvariantCulture).ToGregorianCalendar();
    var referenceDate = new DateTime(2025, 7, 1);
    var birthYear = birthdayPart.Split('-')[0] ?? string.Empty;
    var birthDay = birthdayPart.Split('-')[1].Trim() ?? string.Empty;
    var birthMonth = birthdayPart.Split('-')[2] ?? string.Empty;
    var age = (int)Math.Floor(System.Math.Ceiling(ReferenceDate.Year - birthYear) - (System.DateTime.Parse(birthdayPart, "d/M/yyyy", System.Globalization.CultureInfo.InvariantCulture) >= ReferenceDate)) + 1850;
}

var expectedJson = System.IO.File.ReadAllText("input/expected_format.json");