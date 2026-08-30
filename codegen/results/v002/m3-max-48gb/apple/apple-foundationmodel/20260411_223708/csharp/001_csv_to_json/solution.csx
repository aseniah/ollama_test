using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;

class Program
{
    static void Main(string[] args)
    {
        var csvData = File.ReadAllLines("input/data.csv");
        var jsonData = JsonConvert.DeserializeObject<dynamic>(string.Join(",", csvData.Skip(1).ToArray()));
        var jsonArray = JsonSerializer.Serialize(jsonData, new JsonSerializerOptions { WriteIndented = true });
        Console.WriteLine(jsonArray);
    }
}