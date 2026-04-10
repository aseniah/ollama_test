#r "System.IO"
using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;

// Read expected format to understand output structure
string expectedFormatPath = Path.Combine("input", "expected_format.json");
if (File.Exists(expectedFormatPath))
{
    string expectedJson = File.ReadAllText(expectedFormatPath);
    JsonNode? expectedNode = JsonNode.Parse(expectedJson);
}

// Read CSV file
string csvPath = Path.Combine("input", "input.csv");
string[] lines = File.ReadAllLines(csvPath);
if (lines.Length == 0)
{
    Console.WriteLine("[]");
    return;
}

// Parse header and determine column indices
string[] headers = lines[0].Split(',').Select(h => h.Trim().Trim('"')).ToArray();

int nameIndex = Array.IndexOf(headers, "name");
int birthDateIndex = Array.IndexOf(headers, "birthdate");
if (birthDateIndex < 0) birthDateIndex = Array.IndexOf(headers, "date_of_birth");
if (birthDateIndex < 0) birthDateIndex = Array.IndexOf(headers, "dob");
if (birthDateIndex < 0) birthDateIndex = Array.IndexOf(headers, "birthday");

int emailIndex = Array.IndexOf(headers, "email");
if (emailIndex < 0) emailIndex = Array.IndexOf(headers, "email_address");

string[]? phoneIndexCheck = new string[] { "phone", "phone_number", "telephone" };
int phoneIndex = -1;
foreach (var p in phoneIndexCheck)
{
    if (Array.IndexOf(headers, p) >= 0)
    {
        phoneIndex = Array.IndexOf(headers, p);
        break;
    }
}

string[]? cityIndexCheck = new string[] { "city", "city_name" };
int cityIndex = -1;
foreach (var c in cityIndexCheck)
{
    if (Array.IndexOf(headers, c) >= 0)
    {
        cityIndex = Array.IndexOf(headers, c);
        break;
    }
}

string[]? countryIndexCheck = new string[] { "country", "country_name" };
int countryIndex = -1;
foreach (var c in countryIndexCheck)
{
    if (Array.IndexOf(headers, c) >= 0)
    {
        countryIndex = Array.IndexOf(headers, c);
        break;
    }
}

// Process data rows
var people = new List<JsonNode>();

DateTime cutoffDate = new DateTime(2025, 7, 1);

for (int i = 1; i < lines.Length; i++)
{
    string[] fields = lines[i].Split(',').Select(f => f.Trim().Trim('"')).ToArray();
    
    string name = nameIndex >= 0 && nameIndex < fields.Length ? fields[nameIndex] : "";
    if (string.IsNullOrEmpty(name)) continue;
    
    DateTime? birthDate = null;
    if (birthDateIndex >= 0 && birthDateIndex < fields.Length)
    {
        string dateStr = fields[birthDateIndex];
        if (!string.IsNullOrEmpty(dateStr))
        {
            birthDate = DateTime.ParseExact(dateStr, "yyyy-MM-dd", CultureInfo.InvariantCulture);
            if (birthDate == null)
            {
                birthDate = DateTime.Parse(dateStr, CultureInfo.InvariantCulture);
            }
        }
    }
    
    int? age = null;
    if (birthDate.HasValue)
    {
        int years = cutoffDate.Year - birthDate.Value.Year;
        bool hasHadBirthdayThisYear = cutoffDate.Month > birthDate.Value.Month ||
                                      (cutoffDate.Month == birthDate.Value.Month && 
                                       cutoffDate.Day >= birthDate.Value.Day);
        if (!hasHadBirthdayThisYear) years--;
        age = years;
    }
    
    string email = "";
    if (emailIndex >= 0 && emailIndex < fields.Length)
    {
        email = fields[emailIndex];
    }
    
    string phone = "";
    if (phoneIndex >= 0 && phoneIndex < fields.Length)
    {
        phone = fields[phoneIndex];
    }
    
    string city = "";
    if (cityIndex >= 0 && cityIndex < fields.Length)
    {
        city = fields[cityIndex];
    }
    
    string country = "";
    if (countryIndex >= 0 && countryIndex < fields.Length)
    {
        country = fields[countryIndex];
    }
    
    var personObj = new JsonObject();
    personObj["name"] = name;
    if (birthDate.HasValue)
    {
        personObj["birthdate"] = birthDate.Value.ToString("yyyy-MM-dd");
    }
    if (age.HasValue)
    {
        personObj["age"] = age.Value;
    }
    if (!string.IsNullOrEmpty(email))
    {
        personObj["email"] = email;
    }
    if (!string.IsNullOrEmpty(phone))
    {
        personObj["phone"] = phone;
    }
    if (!string.IsNullOrEmpty(city))
    {
        personObj["city"] = city;
    }
    if (!string.IsNullOrEmpty(country))
    {
        personObj["country"] = country;
    }
    
    people.Add(personObj);
}

// Output JSON array
var options = new JsonSerializerOptions
{
    WriteIndented = true,
    PropertyNamingPolicy = JsonNamingPolicy.CamelCase
};

Console.WriteLine(JsonSerializer.Serialize(people, options));