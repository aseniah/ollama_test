using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

int main()
{
    string csvPath = "input/input.csv";
    string jsonPath = "input/expected_format.json";
    
    var people = new List<object>();
    
    // Read CSV file
    var csvText = File.ReadAllText(csvPath);
    
    // Parse CSV - assume basic format with first row as headers
    var lines = csvText.Split(new[] { "\r\n", "\n" }, StringSplitOptions.RemoveEmptyEntries);
    
    if (lines.Length > 0)
    {
        var header = lines[0].Split(',');
        
        for (int i = 1; i < lines.Length; i++)
        {
            var values = lines[i].Split(',');
            if (values.Length >= 2) // Assuming at least name and date fields
            {
                var dateStr = values[1]?.Trim();
                if (!string.IsNullOrEmpty(dateStr))
                {
                    try
                    {
                        var birthDate = DateTime.Parse(dateStr);
                        
                        // Calculate age as of July 1, 2025
                        var referenceDate = new DateTime(2025, 7, 1);
                        var ageInYears = referenceDate.Year - birthDate.Year;
                        
                        // Adjust age if birthday hasn't occurred yet in 2025
                        if (referenceDate.Month > birthDate.Month || 
                            (referenceDate.Month == birthDate.Month && 
                             referenceDate.Day < birthDate.Day))
                        {
                            ageInYears -= 1;
                        }
                        else if (referenceDate.Month == birthDate.Month && 
                                 referenceDate.Day == birthDate.Day)
                        {
                            // Born exactly on July 1, 2025
                        }
                        else
                        {
                            // Birthday has occurred in 2025
                        }
                        
                        people.Add(new { Name = values[0]?.Trim(), Age = ageInYears });
                    }
                    catch
                    {
                        // Invalid date, skip
                    }
                }
            }
        }
    }
    
    // Output as JSON array
    var json = JsonSerializer.Serialize(people);
    System.Console.WriteLine(json);
}