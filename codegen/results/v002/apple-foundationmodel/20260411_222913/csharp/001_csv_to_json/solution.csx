using System;
using System.Collections.Generic;
using System.Text.Json;

class Program
{
    static void Main(string[] args)
    {
        IList<string> csvLines = new List<string>(File.ReadAllLines("input/data.csv"));
        IList<JsonObject> jsonObjects = new List<JsonObject>();

        for (int i = 1; i < csvLines.Count; i++)
        {
            string[] fields = csvLines[i].Split(',');
            var obj = JsonObject.Create(fields);
            jsonObjects.Add(obj);
        }

        JsonSerializer.Serialize(jsonObjects, Console.Out);
    }
}