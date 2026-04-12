#r "System.Text.Json"

using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;

string[] lines = File.ReadAllLines("input/input.csv");
DateTime referenceDate = new DateTime(2025, 7, 1);

var people = new List<object>();

for (int i = 1; i < lines.Length; i++) {
    string line = lines[i];
    string[] fields = line.Split(',');
    
    string fullName = fields[0];
    string birthdayStr = fields[1];
    string father = fields[3];
    string mother = fields[4];
    string brother = fields[5];
    string sister = fields[6];
    
    var nameParts = fullName.Split(' ');
    string firstName = nameParts[0];
    string lastName = nameParts[nameParts.Length - 1];
    
    DateTime birthday = DateTime.Parse(birthdayStr);
    
    int age = referenceDate.Year - birthday.Year;
    if (referenceDate.Year != birthday.Year || 
        (referenceDate.Month > birthday.Month || 
         (referenceDate.Month == birthday.Month && referenceDate.Day >= birthday.Day)) == false) {
        if (referenceDate < new DateTime(referenceDate.Year, birthday.Month, birthday.Day)) {
            age--;
        }
    }
    
    string birthdayIso = birthday.ToString("yyyy-MM-dd");
    
    var relatives = new List<object>();
    
    if (!string.IsNullOrWhiteSpace(father) && father != "null") {
        var fatherParts = father.Split(' ');
        relatives.Add(new { 
            FirstName = fatherParts[0], 
            LastName = fatherParts[fatherParts.Length - 1], 
            Relationship = "Father" 
        });
    }
    
    if (!string.IsNullOrWhiteSpace(mother) && mother != "null") {
        var motherParts = mother.Split(' ');
        relatives.Add(new { 
            FirstName = motherParts[0], 
            LastName = motherParts[motherParts.Length - 1], 
            Relationship = "Mother" 
        });
    }
    
    if (!string.IsNullOrWhiteSpace(brother) && brother != "null") {
        var brotherParts = brother.Split(' ');
        relatives.Add(new { 
            FirstName = brotherParts[0], 
            LastName = brotherParts[brotherParts.Length - 1], 
            Relationship = "Brother" 
        });
    }
    
    if (!string.IsNullOrWhiteSpace(sister) && sister != "null") {
        var sisterParts = sister.Split(' ');
        relatives.Add(new { 
            FirstName = sisterParts[0], 
            LastName = sisterParts[sisterParts.Length - 1], 
            Relationship = "Sister" 
        });
    }
    
    people.Add(new {
        FirstName = firstName,
        LastName = lastName,
        Birthday = birthdayIso,
        Age = age,
        Relatives = relatives
    });
}

var options = new JsonSerializerOptions {
    WriteIndented = true
};

Console.WriteLine(JsonSerializer.Serialize(people, options));