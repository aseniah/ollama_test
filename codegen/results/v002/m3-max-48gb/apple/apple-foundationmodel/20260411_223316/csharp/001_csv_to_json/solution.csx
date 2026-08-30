using System;
using System.IO;
using System.Text.Json;

class Program
{
    static void Main(string[] args)
    {
        var csvLines = File.ReadAllLines("input/data.csv");
        var jsonArray = ProcessCsvToJsonArray(csvLines);
        JsonSerializer.Serialize(jsonArray, Console.Out);
    }

    static JsonArray ProcessCsvToJsonArray(IList<string> csvLines)
    {
        var jsonArray = new JsonArray();

        // Skip header
        if (csvLines.Count > 1)
        {
            for (int i = 1; i < csvLines.Count; i++)
            {
                string[] fields = csvLines[i].Split(',');
                if (fields.Length >= 4)
                {
                    var obj = new
                    {
                        Name = fields[0].Trim(),
                        Age = int.Parse(fields[1].Trim()),
                        Email = fields[2].Trim(),
                        Score = float.Parse(fields[3].Trim())
                    };
                    jsonArray.Add(obj);
                }
            }
        }

        return jsonArray;
    }
}