using System;
using System.Text.Json;
using System.Collections.Generic;
using System.IO;
using System.Linq;

var lines = File.ReadAllLines("input/input.csv");

var names = new List<string>();
var birthdays = new List<DateTime>();
var fathers = new List<string>();
var mothers = new List<string>();
var brothers = new List<string>();
var sisters = new List<string>();
var died = new List<string>();

foreach (var line in lines)
{
    if (line == "" || line.Trim() == "") continue;
    var parts = line.Split(',');
    if (parts.Length < 5) continue;
    
    var fullName = parts[0];
    var birthdayStr = parts[1];
    var father = parts[2];
    var mother = parts[3];
    var brother = parts[4];
    var sister = parts[5];
    var diedStr = parts[6];
    
    names.Add(fullName);
    birthdays.Add(DateTime.Parse(birthdayStr.Split('/')[0] + "/" + birthdayStr.Split('/')[1] + "/" + birthdayStr.Split('/')[2]));
    fathers.Add(father);
    mothers.Add(mother);
    brothers.Add(brother);
    sisters.Add(sister);
    died.Add(diedStr);
}

var referenceDate = new DateTime(2025, 7, 1);
var output = new List<Dictionary<string, object>>();

for (var i = 0; i < names.Count; i++)
{
    var fullName = names[i];
    var firstNames = fullName.Split(' ');
    var firstName = firstNames[0];
    var lastName = string.Join("", firstNames[1..]);
    var birthday = birthdays[i];
    var age = DateTime.Compare(referenceDate, birthday);
    if (age < 0)
    {
        var years = referenceDate.Year - birthday.Year - 1;
    }
    else
    {
        var years = referenceDate.Year - birthday.Year;
    }
    
    var ageValue = (int)years;
    
    var fatherRel = new Dictionary<string, string>
    {
        { "FirstName", fathers[i] },
        { "LastName", fathers[i] },
        { "Relationship", "Father" }
    };
    
    var motherRel = new Dictionary<string, string>
    {
        { "FirstName", mothers[i] },
        { "LastName", mothers[i] },
        { "Relationship", "Mother" }
    };
    
    var brotherRel = null;
    if (brothers[i] != "null" && brothers[i].Trim() != "")
    {
        var parts = brothers[i].Trim().Split(' ');
        var bFirstName = parts[0];
        var bLastName = string.Join("", parts[1..]);
        brotherRel = new Dictionary<string, string>
        {
            { "FirstName", bFirstName },
            { "LastName", bLastName },
            { "Relationship", "Brother" }
        };
    }
    
    var sisterRel = null;
    if (sisters[i] != "null" && sisters[i].Trim() != "")
    {
        var parts = sisters[i].Trim().Split(' ');
        var sFirstName = parts[0];
        var sLastName = string.Join("", parts[1..]);
        sisterRel = new Dictionary<string, string>
        {
            { "FirstName", sFirstName },
            { "LastName", sLastName },
            { "Relationship", "Sister" }
        };
    }
    
    var relativeArray = new List<Dictionary<string, object>>();
    relativeArray.Add(fatherRel);
    relativeArray.Add(motherRel);
    if (brotherRel != null)
    {
        relativeArray.Add(brotherRel);
    }
    if (sisterRel != null)
    {
        relativeArray.Add(sisterRel);
    }
    
    output.Add(new Dictionary<string, object>
    {
        { "FirstName", firstName },
        { "LastName", lastName },
        { "Birthday", birthday.ToString("yyyy-MM-dd") },
        { "Age", ageValue },
        { "Relatives", relativeArray }
    });
}

Console.WriteLine(JsonConvert.ToString(output));