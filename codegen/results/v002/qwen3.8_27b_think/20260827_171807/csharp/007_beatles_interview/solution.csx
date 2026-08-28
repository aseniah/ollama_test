using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

string csvContent = File.ReadAllText("input/input.csv");
string[] lines = csvContent.ReadAllLinesSafe();

string[] headers = lines[0].Split(',');
List<JsonObject> results = new List<JsonObject>();

DateTime asOfDate = new DateTime(2025, 7, 1);

foreach (string line in lines.Skip(1))
{
    if (string.IsNullOrWhiteSpace(line)) continue;
    string[] fields = line.Split(',');

    string fullName = fields[0].Trim();
    string[] nameParts = fullName.Split(' ');
    string firstName = nameParts[0];
    string lastName = nameParts[nameParts.Length - 1];

    string birthdayInput = fields[1].Trim();
    string[] bp = birthdayInput.Split('/');
    string birthdayFormatted = $"{bp[2]}-{bp[0].PadLeft(2, '0')}-{bp[1].PadLeft(2, '0')}";
    DateTime birthDate = new DateTime(int.Parse(bp[2]), int.Parse(bp[0]), int.Parse(bp[1]));

    string diedInput = fields[2].Trim();
    int age;
    if (diedInput == "null" || string.IsNullOrWhiteSpace(diedInput))
    {
        age = CalculateAge(birthDate, asOfDate);
    }
    else
    {
        string[] dp = diedInput.Split('/');
        DateTime deathDate = new DateTime(int.Parse(dp[2]), int.Parse(dp[0]), int.Parse(dp[1]));
        age = CalculateAge(birthDate, deathDate);
    }

    List<JsonObject> relatives = new List<JsonObject>();
    AddRelative(relatives, fields[3], "Father");
    AddRelative(relatives, fields[4], "Mother");
    AddRelative(relatives, fields[5], "Brother");
    AddRelative(relatives, fields[6], "Sister");

    JsonObject person = new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthdayFormatted,
        ["Age"] = age,
        ["Relatives"] = new JsonArray(relatives.ToArray())
    };

    results.Add(person);
}

JsonSerializerOptions options = new JsonSerializerOptions { WriteIndented = true };
string output = JsonSerializer.Serialize(results, options);
Console.WriteLine(output);

static int CalculateAge(DateTime birth, DateTime reference)
{
    int age = reference.Year - birth.Year;
    if (reference.Month < birth.Month || (reference.Month == birth.Month && reference.Day < birth.Day))
        age--;
    return age;
}

static void AddRelative(List<JsonObject> list, string value, string relationship)
{
    if (string.IsNullOrWhiteSpace(value) || value.Trim().ToLower() == "null") return;
    string[] parts = value.Trim().Split(' ');
    list.Add(new JsonObject
    {
        ["FirstName"] = parts[0],
        ["LastName"] = parts[parts.Length - 1],
        ["Relationship"] = relationship
    });
}

static class StringExt
{
    public static string[] ReadAllLinesSafe(this string s) => s.Split('\n').Select(l => l.TrimEnd('\r')).Where(l => !string.IsNullOrWhiteSpace(l)).ToArray();
}