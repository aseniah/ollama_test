using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

int main()
{
    var lines = File.ReadAllLines("input/input.csv");
    var firstLine = lines[0]; // headers
    var headerSplit = firstLine.Split(',');
    
    var persons = new List<Dictionary<string, string>>();
    
    for (int i = 1; i < lines.Length; i++)
    {
        var line = lines[i];
        var row = line.Split(',');
        
        var person = new Dictionary<string, string>();
        person["Name"] = row[0];
        person["BirthdayRaw"] = row[1];
        person["Died"] = row[2];
        person["Father"] = row[3];
        person["Mother"] = row[4];
        person["Brother"] = row[5];
        person["Sister"] = row[6];
        
        var firstName = person["Name"].Split(' ')[0];
        var lastName = person["Name"].Substring(person["Name"].LastIndexOf(' ') + 1);
        person["FirstName"] = firstName;
        person["LastName"] = lastName;
        
        var birthdayRaw = person["BirthdayRaw"];
        var parts = birthdayRaw.Split('/');
        var birthMonth = int.Parse(parts[0]);
        var birthDay = int.Parse(parts[1]);
        var birthYear = int.Parse(parts[2]);
        
        var birthday = $"{birthYear:D4}-{birthMonth:D2}-{birthDay:D2}";
        person["Birthday"] = birthday;
        
        var referenceDate = new DateTime(2025, 7, 1);
        var birthDate = new DateTime(birthYear, birthMonth, birthDay);
        
        var age = referenceDate.Year - birthDate.Year;
        if (referenceDate.Month > birthMonth || (referenceDate.Month == birthMonth && referenceDate.Day < birthDay))
        {
            age--;
        }
        
        person["Age"] = age;
        
        var relatives = new List<Dictionary<string, string>>();
        var father = person["Father"];
        if (!string.IsNullOrEmpty(father))
        {
            var fatherParts = father.Split(',');
            var relFatherFirstName = fatherParts[0];
            var relFatherLastName = fatherParts[1];
            
            relatives.Add(new Dictionary<string, string>
            {
                { "FirstName", relFatherFirstName },
                { "LastName", relFatherLastName },
                { "Relationship", "Father" }
            });
        }
        
        var mother = person["Mother"];
        if (!string.IsNullOrEmpty(mother))
        {
            var motherParts = mother.Split(',');
            var relMotherFirstName = motherParts[0];
            var relMotherLastName = motherParts[1];
            
            relatives.Add(new Dictionary<string, string>
            {
                { "FirstName", relMotherFirstName },
                { "LastName", relMotherLastName },
                { "Relationship", "Mother" }
            });
        }
        
        var brother = person["Brother"];
        if (!string.IsNullOrEmpty(brother))
        {
            var brotherParts = brother.Split(',');
            var relBrotherFirstName = brotherParts[0];
            var relBrotherLastName = brotherParts[1];
            
            relatives.Add(new Dictionary<string, string>
            {
                { "FirstName", relBrotherFirstName },
                { "LastName", relBrotherLastName },
                { "Relationship", "Brother" }
            });
        }
        
        var sister = person["Sister"];
        if (!string.IsNullOrEmpty(sister))
        {
            var sisterParts = sister.Split(',');
            var relSisterFirstName = sisterParts[0];
            var relSisterLastName = sisterParts[1];
            
            relatives.Add(new Dictionary<string, string>
            {
                { "FirstName", relSisterFirstName },
                { "LastName", relSisterLastName },
                { "Relationship", "Sister" }
            });
        }
        
        var personJson = new Dictionary<string, object>
        {
            { "FirstName", person["FirstName"] },
            { "LastName", person["LastName"] },
            { "Birthday", person["Birthday"] },
            { "Age", person["Age"] },
            { "Relatives", relatives }
        };
        
        persons.Add(personJson);
    }
    
    var output = System.Text.Json.JsonSerializer.Serialize(persons, new System.Text.Json.JsonSerializerOptions { WriteIndented = false });
    System.Console.WriteLine(output);
}