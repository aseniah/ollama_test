using System;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

string[] lines = File.ReadAllLines("input/input.csv");
JsonArray result = new JsonArray();

DateTime referenceDate = new DateTime(2025, 7, 1);

for (int i = 1; i < lines.Length; i++) {
    if (string.IsNullOrWhiteSpace(lines[i])) continue;
    
    string[] fields = lines[i].Split(',');
    string name = fields[0];
    string birthdayStr = fields[1];
    string fatherStr = fields[3];
    string motherStr = fields[4];
    string brotherStr = fields[5];
    string sisterStr = fields[6];
    
    string[] nameParts = name.Split(' ');
    string firstName = nameParts[0];
    string lastName = nameParts[nameParts.Length - 1];
    
    string[] birthdayParts = birthdayStr.Split('/');
    int bMonth = int.Parse(birthdayParts[0]);
    int bDay = int.Parse(birthdayParts[1]);
    int bYear = int.Parse(birthdayParts[2]);
    DateTime birthday = new DateTime(bYear, bMonth, bDay);
    
    int age = referenceDate.Year - birthday.Year;
    if (referenceDate < birthday.AddYears(age)) {
        age--;
    }
    
    var person = new JsonObject();
    person["FirstName"] = firstName;
    person["LastName"] = lastName;
    person["Birthday"] = birthday.ToString("yyyy-MM-dd");
    person["Age"] = age;
    
    JsonArray relatives = new JsonArray();
    
    string[] relativeNames = new[] { fatherStr, motherStr, brotherStr, sisterStr };
    string[] relationships = new[] { "Father", "Mother", "Brother", "Sister" };
    
    for (int j = 0; j < 4; j++) {
        if (relativeNames[j] != "null") {
            string[] relativeNameParts = relativeNames[j].Split(' ');
            var relative = new JsonObject();
            relative["FirstName"] = relativeNameParts[0];
            relative["LastName"] = relativeNameParts[relativeNameParts.Length - 1];
            relative["Relationship"] = relationships[j];
            relatives.Add(relative);
        }
    }
    
    person["Relatives"] = relatives;
    result.Add(person);
}

Console.WriteLine(JsonSerializer.Serialize(result, new JsonSerializerOptions { WriteIndented = true }));