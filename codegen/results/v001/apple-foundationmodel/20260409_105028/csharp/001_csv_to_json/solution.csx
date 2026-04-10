using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;

class Program
{
    static void Main()
    {
        if (Args.Length != 1)
        {
            Console.WriteLine("Usage: Program.csx input/data.csv");
            return;
        }

        string csvFilePath = Path.Combine(Environment.GetFolderPath(Environment.SpecialFolder.LocalApplicationData), "input", Args[0]);

        using (StreamReader reader = new StreamReader(csvFilePath))
        {
            using (JsonSerializerOptions options = new JsonSerializerOptions { WriteIndented = true })
            {
                var lines = reader.ReadToEnd().Split('\n');
                if (lines.Length < 2)
                {
                    Console.WriteLine("Error: CSV file does not contain enough rows.");
                    return;
                }

                var jsonArray = JsonSerializer.Deserialize<List<object>>(lines.Skip(1).ToString(), options);
                Console.WriteLine(JsonSerializer.Serialize(jsonArray, options));
            }
        }
    }
}