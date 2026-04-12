var csv = File.ReadAllLines(@"input/input.csv");
var header = csv[0].Split(',');
var firstNames = new Dictionary<string, int>(); // Store birth years for quick lookup by name (optional optimization)
var relatives = new Dictionary<string, string[]>(); // Map full names to their relative lists

foreach (var line in csv.Skip(1)) {
    var fields = line.Split(',');
    var fullName = fields[0];
    var birthday = fields[1].Split('/')[0] + '/' + fields[1].Split('/')[1] + '/' + fields[1].Split('/')[2];
    // Parse Father, Mother, Brother, Sister (columns 4, 5, 6, 7). Index 3 is Died.
    var father = fields[3];
    var mother = fields[4];
    var brother = fields[6];
    var sister = fields[7];

    string[] rels;
    if (string.IsNullOrEmpty(father)) rels = new[] { "Mother", brother, sister };
    else {
        rels = new[] { father, mother, brother, sister };
    }
    foreach (var r in rels) {
        if (string.IsNullOrEmpty(r)) continue;
        relatives[fullName] = [..relatives[fullName]];
        // Add to firstNames list for quick lookup later
        var parts = r.Split(' ');
        if (parts.Length >= 2) firstNames[r + " " + fields[3]] += 1; 
    }
}

var result = new List<object>();
var targetDate = new DateTime(2025, 7, 1);
foreach (var entry in input) {
    var parts = entry.Split(',');
    var firstName = parts[0];
    var lastName = parts[1];
    var birthdayStr = parts[2].Split('/')[0] + '/' + parts[2].Split('/')[1] + '/' + parts[2].Split('/')[2];
    var fullBirthDate = DateTime.ParseExact(birthdayStr, "M/d/yyyy", null);
    var ageYears = targetDate - fullBirthDate; // Calculate difference in years, not months
    var age = (int)ageYears.TotalDays / 365.2425; // More accurate age calculation

    result.Add([
        ["FirstName", firstName],
        ["LastName", lastName],
        ["Birthday", birthdayStr],
        ["Age", (int)age]
    ]);
}

Console.Output($"[{string.Join(",\n  ", result)}]");