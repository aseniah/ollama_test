using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;

namespace BeatlesAges
{
    class Program
    {
        static void Main()
        {
            string csvPath = "input/input.csv";
            string jsonPath = "output.json";

            // Read CSV file and parse data
            List<object> data = ParseCSV(csvPath);

            // Calculate ages as of July 1, 2025
            List<object> calculatedData = CalculateAges(data);

            // Write calculated data to JSON
            WriteJSON(calculatedData, jsonPath);
        }

        static List<object> ParseCSV(string csvPath)
        {
            List<object> data = new List<object>();
            using (StreamReader reader = new StreamReader(csvPath))
            {
                string line;
                while ((line = reader.ReadLine()) != null)
                {
                    string[] fields = line.Split(',');
                    if (fields.Length == 7)
                    {
                        object entry = new
                        {
                            FirstName = fields[0],
                            LastName = fields[1],
                            Birthday = DateTime.ParseExact(fields[2], "dd/MM/yyyy", CultureInfo.InvariantCulture),
                            Age = CalculateAge(DateTime.ParseExact(fields[2], "dd/MM/yyyy", CultureInfo.InvariantCulture)),
                            Relatives = ParseRelatives(fields)
                        };
                        data.Add(entry);
                    }
                }
            }
            return data;
        }

        static List<object> CalculateAges(List<object> data)
        {
            DateTime currentDate = new DateTime(2025, 7, 1);
            return data.Select(item =>
            {
                object entry = item as object;
                DateTime birthday = (DateTime)entry["Birthday"];
                int age = CalculateAge(birthday);
                return new
                {
                    ...entry,
                    Age = age
                };
            }).ToList();
        }

        static List<object> ParseRelatives(string[] fields)
        {
            List<object> relatives = new List<object>();
            for (int i = 3; i < fields.Length; i += 3)
            {
                string firstName = fields[i];
                string lastName = fields[i + 1];
                string relationship = fields[i + 2];
                relatives.Add(new
                {
                    FirstName = firstName,
                    LastName = lastName,
                    Relationship = relationship
                });
            }
            return relatives;
        }

        static int CalculateAge(DateTime birthday)
        {
            return (int)(currentDate - birthday).Days / 365;
        }

        static void WriteJSON(List<object> data, string jsonPath)
        {
            var json = JsonSerializer.Serialize(data, new JsonSerializerOptions { WriteIndented = true });
            File.WriteAllText(jsonPath, json);
        }
    }
}