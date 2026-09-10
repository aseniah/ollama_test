using System.Text.Json;
using System.Text.Json.Nodes;

var lines = File.ReadAllLines("input/input.csv");
var people = new JsonArray();

var referenceDate = new DateTime(2025, 7, 1);

for (int i = 1; i < lines.Length; i++)
{
    var line = lines[i].Trim();
    if (string.IsNullOrEmpty(line)) continue;

    var parts = line.Split(',');
    if (parts.Length < 7) continue;

    var fullName = parts[0].Trim();
    var nameParts = fullName.Split(' ');
    var firstName = nameParts[0];
    var lastName = nameParts[nameParts.Length - 1];

    var birthday = parts[1].Trim();
    var birthdayParsed = DateTime.Parse(birthday);
    var birthISO = birthdayParsed.ToString("yyyy-MM-dd");

    // Calculate age as of July 1, 2025
    int age = referenceDate.Year - birthdayParsed.Year;
    if (referenceDate < birthdayParsed.AddYears(age))
        age--;

    var relatives = new JsonArray();

    // Father
    var father = parts[3].Trim();
    if (father != "null" && !string.IsNullOrEmpty(father))
    {
        var fatherParts = father.Split(' ');
        relatives.Add(new JsonObject {
            ["FirstName"] = fatherParts[0],
            ["LastName"] = fatherParts[fatherParts.Length - 1],
            ["Relationship"] = "Father"
        });
    }

    // Mother
    var mother = parts[4].Trim();
    if (mother != "null" && !string.IsNullOrEmpty(mother))
    {
        var motherParts = mother.Split(' ');
        relatives.Add(new JsonObject {
            ["FirstName"] = motherParts[0],
            ["LastName"] = motherParts[motherParts.Length - 1],
            ["Relationship"] = "Mother"
        });
    }

    // Brother
    var brother = parts[5].Trim();
    if (brother != "null" && !string.IsNullOrEmpty(brother))
    {
        var brotherParts = brother.Split(' ');
        relatives.Add(new JsonObject {
            ["FirstName"] = brotherParts[0],
            ["LastName"] = brotherParts[brotherParts.Length - 1],
            ["Relationship"] = "Brother"
        });
    }

    // Sister
    var sister = parts[6].Trim();
    if (sister != "null" && !string.IsNullOrEmpty(sister))
    {
        var sisterParts = sister.Split(' ');
        relatives.Add(new JsonObject {
            ["FirstName"] = sisterParts[0],
            ["LastName"] = sisterParts[sisterParts.Length - 1],
            ["Relationship"] = "Sister"
        });
    }

    var person = new JsonObject {
        ["FirstName"] = firstName,
        ["LastName"] = lastName,
        ["Birthday"] = birthISO,
        ["Age"] = age,
        ["Relatives"] = relatives
    };

    people.Add(person);
}

Console.WriteLine(people.ToJsonString(new JsonWriterOptions { Indented = true }));