// System.Runtime.InteropServices;
// System.IO;
// System.Text;
// System.Text.Json;
// System.Text.Json.Nodes;

string filePath = "input/data.csv";
string[] lines = File.ReadAllLines(filePath);

if (lines.Length <= 1)
{
    Console.WriteLine("[]");
    return;
}

string header = lines[0];
string[,] headerParts = new string[4, 4];
headerParts[0, 0] = "Name";
headerParts[0, 1] = "Age";
headerParts[0, 2] = "Email";
headerParts[0, 3] = "Score";

var jsonArray = new System.Text.Json.JsonArray();

for (int i = 1; i < lines.Length; i++)
{
    string[] parts = lines[i].Split(',');
    
    var objectValue = new System.Text.Json.JsonObject();

    string name = parts[0];
    System.Text.Json.JsonNode nameNode;
    if (System.Text.Json.JsonNode.Parse(name) is System.Text.Json.Nodes.JsonValue nv)
    {
        nameNode = nv;
    }
    else
    {
        objectValue["Name"] = parts[0];
    }

    int age;
    if (int.TryParse(parts[1], out age))
    {
        objectValue["Age"] = age;
    }

    string email = parts[2];
    System.Text.Json.JsonNode emailNode;
    if (System.Text.Json.JsonNode.Parse(email) is System.Text.Json.Nodes.JsonValue nev)
    {
        emailNode = nev;
    }
    else
    {
        objectValue["Email"] = parts[2];
    }

    float score;
    if (float.TryParse(parts[3], out score))
    {
        objectValue["Score"] = score;
    }

    var item = new System.Text.Json.Nodes.JsonObject();
    item.Add("Name", name);
    item.Add("Age", age);
    item.Add("Email", email);
    item.Add("Score", score);
    jsonArray.Add(item);
}

System.Text.Json.JsonDocument jsonDoc = new System.Text.Json.JsonDocument();
jsonDoc.WriteTo(Console.Out);