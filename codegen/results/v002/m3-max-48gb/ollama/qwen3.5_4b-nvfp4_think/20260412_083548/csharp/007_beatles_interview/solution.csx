System.Console.WriteLine($"[");
System.Text.Json.JsonDocument json;
using System.IO;
using System.Collections.Generic;
using System.Text.Json.Nodes;

var lines = System.IO.File.ReadAllLines("input/input.csv");
var header = lines[0];

for (int i = 1, end = lines.Length - 1, j = 0, l; i < end; i++) {
    var parts = lines[i].Split(',', StringSplitOptions.RemoveEmptyEntries);
    if (parts[j + 3] == "null") continue;
    
    string name = parts[0];
    string[] nameParts = name.Split(new char[] {' ' }, 2, StringSplitOptions.RemoveEmptyEntries);
    string firstName = nameParts.Length > 0 ? nameParts[0] : "";
    string lastName = nameParts.Length > 1 ? (nameParts[^1] == nameParts[nameParts.Length - 2] ? 
                                              System.Linq.Enumerable.Default("") + " " + nameParts[^1] 
                                             : nameParts[^1]) : "";
    
    if (String.IsNullOrEmpty(firstName)) { firstName = "Unknown"; }
    if (String.IsNullOrEmpty(lastName)) { lastName = "Unknown"; }
    
    var ageYears = 0, dayIndex = -1;
    for (var k = 0, n = parts[1].Split('/').Length, l = 0; k < n; k++) {
        if (System.DateTime.TryParse(parts[1], System.Globalization.CultureInfo.InvariantCulture, DateTimeStyles.AllowInputFormatting, out System.DateTime) || 
            System.DateTime.TryParse(parts[1], '0', 3, System.DateTimeFormatStyle.DMyDate)) { break; }
    }
    
    var[] parts = lines[i].Split(',', StringSplitOptions.RemoveEmptyEntries);
    string firstNamePart = parts[0] ?? "";
    string lastNamePart = parts[^1] ?? "";
    string firstYear = parts[j + 3] ?? "";
    string firstMonth = parts[1] ?? "";
    int.TryParse(firstYear, out System.DateTime birthDate);
    var dateParts = firstYear.Split('-');
    ageYears = (dateParts.Length > 2 && System.DateTime.Parse(new string[] {'', '0', '0', firstYear}, null)) ? 
              40 : (dateParts.Length >= 3) ? int.Parse(firstYear[0]) : 0;
    
    var jsonDoc = new System.Text.Json.JsonDocument { value = firstName };
    var jsonArray = new System.Text.Json.JSONArray();
    for (int l = 1, end2 = parts.Length, j = 0; l < end2; l++) {
        string relativeName = parts[0 + l] ?? "";
        if (String.IsNullOrEmpty(relativeName)) { continue; }
        
        var[] relParts = relativeName.Split(new char[] {' ' }, 1, StringSplitOptions.RemoveEmptyEntries);
        var firstNameRel = relParts.Length > 0 ? 
                            System.Linq.Enumerable.Default("") : "Unknown";
        var lastNameRel = relParts.Length > 1 ? 
                         (relParts[^1] == "" ? "Unknown" : relParts[^1]) : "Unknown";
        
    }
}