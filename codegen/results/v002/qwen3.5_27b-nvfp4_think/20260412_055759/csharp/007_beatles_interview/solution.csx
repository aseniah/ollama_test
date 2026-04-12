#r "System.Text.Json"

using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

string[] lines = File.ReadAllLines("input/input.csv");
var header = lines[0].Split(',');
List<object> output = new List<object>();

DateTime referenceDate = new DateTime(2025, 7, 1);

for (int i = 1; i < lines.Length; i++) {
    var values = lines[i].Split(',');
    
    // Parse full name into first and last
    string[] nameParts = values[0].Trim().Split(' ');
    string firstName = nameParts[0];
    string lastName = nameParts[1];
    
    // Parse birthday (M/D/YYYY format)
    string[] dateParts = values[1].Trim().Split('/');
    int month = int.Parse(dateParts[0]);
    int day = int.Parse(dateParts[1]);
    int year = int.Parse(dateParts[2]);
    DateTime birthday = new DateTime(year, month, day);
    string formattedBirthday = $"{year:D4}-{month:D2}-{day:D2}";
    
    // Calculate age as of July 1, 2025
    int age = referenceDate.Year - birthday.Year;
    if (referenceDate.DayOfYear < birthday.Date.DayOfYear) {
        age--;
    }
    
    // Build relatives list
    List<object> relatives = new List<object>();
    
    string fatherName = values[3].Trim();
    if (!string.IsNullOrEmpty(fatherName) && fatherName != "null") {
        string[] fParts = fatherName.Split(' ');
        relatives.Add(new JsonObject {
            {"FirstName", fParts[0]},
            {"LastName", fParts[1]},
            {"Relationship", "Father"}
        });
    }
    
    string motherName = values[4].Trim();
    if (!string.IsNullOrEmpty(motherName) && motherName != "null") {
        string[] mParts = motherName.Split(' ');
        relatives.Add(new JsonObject {
            {"FirstName", mParts[0]},
            {"LastName", mParts[1]},
            {"Relationship", "Mother"}
        });
    }
    
    string brotherName = values[5].Trim();
    if (!string.IsNullOrEmpty(brotherName) && brotherName != "null") {
        string[] bParts = brotherName.Split(' ');
        relatives.Add(new JsonObject {
            {"FirstName", bParts[0]},
            {"LastName", bParts[1]},
            {"Relationship", "Brother"}
        });
    }
    
    string sisterName = values[6].Trim();
    if (!string.IsNullOrEmpty(sisterName) && sisterName != "null") {
        string[] sParts = sisterName.Split(' ');
        relatives.Add(new JsonObject {
            {"FirstName", sParts[0]},
            {"LastName", sParts[1]},
            {"Relationship", "Sister"}
        });
    }
    
    // Create person object
    var person = new JsonObject {
        {"FirstName", firstName},
        {"LastName", lastName},
        {"Birthday", formattedBirthday},
        {"Age", age},
        {"Relatives", new JsonArray(relatives)}
    };
    
    output.Add(person);
}

// Write JSON array to stdout with formatting
var options = new JsonSerializerOptions { WriteIndented = true };
Console.WriteLine(JsonSerializer.Serialize(output, options));