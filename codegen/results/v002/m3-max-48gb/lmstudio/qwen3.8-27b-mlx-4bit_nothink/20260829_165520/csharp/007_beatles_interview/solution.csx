using System;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Globalization;
using System.Linq;

string csvPath = "input/input.csv";
string[] lines = File.ReadAllLines(csvPath);

string header = lines[0];
string[] headers = header.Split(',');

// Find column indices
int nameIdx = Array.IndexOf(headers, "Name");
int birthdayIdx = Array.IndexOf(headers, "Birthday");
int diedIdx = Array.IndexOf(headers, "Died");
int fatherIdx = Array.IndexOf(headers, "Father");
int motherIdx = Array.IndexOf(headers, "Mother");
int brotherIdx = Array.IndexOf(headers, "Brother");
int sisterIdx = Array.IndexOf(headers, "Sister");

DateTime referenceDate = new DateTime(2025, 7, 1);

JsonArray result = new JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    if (string.IsNullOrWhiteSpace(lines[i])) continue;
    
    string[] fields = lines[i].Split(',');
    
    // Parse name
    string fullName = fields[nameIdx].Trim();
    string[] nameParts = fullName.Split(' ', 2);
    string firstName = nameParts[0];
    string lastName = nameParts.Length > 1 ? nameParts[1] : "";
    
    // Parse birthday
    string birthdayStr = fields[birthdayIdx].Trim();
    DateTime birthday = DateTime.Parse(birthdayStr, CultureInfo.InvariantCulture);
    
    // Parse died (if present)
    string diedStr = fields[diedIdx].Trim();
    DateTime? diedDate = null;
    if (!string.IsNullOrEmpty(diedStr) && diedStr != "null")
    {
        diedDate = DateTime.Parse(diedStr, CultureInfo.InvariantCulture);
    }
    
    // Calculate age
    int age;
    if (diedDate.HasValue)
    {
        age = diedDate.Value.Year - birthday.Year;
        if (diedDate.Value.Month < birthday.Month || 
            (diedDate.Value.Month == birthday.Month && diedDate.Value.Day < birthday.Day))
        {
            age--;
        }
    }
    else
    {
        age = referenceDate.Year - birthday.Year;
        if (referenceDate.Month < birthday.Month || 
            (referenceDate.Month == birthday.Month && referenceDate.Day < birthday.Day))
        {
            age--;
        }
    }
    
    // Format birthday as ISO
    string birthdayIso = birthday.ToString("yyyy-MM-dd");
    
    // Build relatives
    JsonArray relatives = new JsonArray();
    
    string father = fields[fatherIdx].Trim();
    if (!string.IsNullOrEmpty(father) && father != "null")
    {
        string[] fParts = father.Split(' ', 2);
        relatives.Add(new JsonObject
        {
            ["FirstName"] = fParts[0],
            ["LastName"] = fParts.Length > 1 ? fParts[1] : "",
            ["Relationship"] = "Father"
        });
    }
    
    string mother = fields[motherIdx].Trim();
    if (!string.IsNullOrEmpty(mother) && mother != "null")
    {
        string[] mParts = mother.Split(' ', 2);
        relatives.Add(new JsonObject
        {
            ["FirstName"] = mParts[0],
            ["LastName"] = mParts.Length > 1 ? mParts[1] : "",
            ["Relationship"] = "Mother"
        });
    }
    
    string brother = fields[brotherIdx].Trim();
    if (!string.IsNullOrEmpty(brother) && brother != "null")
    {
        string[] bParts = brother.Split(' ', 2);
        relatives.Add(new JsonObject
        {
            ["FirstName"] = bParts[0],
            ["LastName"] = bParts.Length > 1 ? bParts[1] : "",
            ["Relationship"] = "Brother"
        });
    }
    
    string sister = fields[sisterIdx].Trim();
    if (!string.IsNullOrEmpty(sister) && sister != "null")
    {
        string[] sParts = sister.Split(' ', 2);
        relatives.Add(new JsonObject
        {
            ["FirstName"] = sParts[0],
            ["LastName"] = sParts.Length > 1 ? sParts[1] : "",
            ["Relationship"] = "Sister"
        });
    }
    
    JsonObject person = new JsonObject
    {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthdayIso,
        ["Age"] = age,
        ["Relatives"] = relatives
    };
    
    result.Add(person);
}

string output = JsonSerializer.Serialize(result, new JsonSerializerOptions
{
    WriteIndented = true
});

Console.Write(output);