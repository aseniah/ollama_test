using System;
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Nodes;
using System.Globalization;

var inputPath = "input/data.csv";

try {
    using var reader = new StreamReader(inputPath);
    var headerLine = reader.ReadLine();
    if (string.IsNullOrEmpty(headerLine)) {
        Console.WriteLine("Empty input file");
        return;
    }

    // Verify header
    var expectedHeader = "Name,Age,Email,Score";
    if (!headerLine.Equals(expectedHeader, StringComparison.OrdinalIgnoreCase)) {
        Console.WriteLine($"Invalid header. Expected: {expectedHeader}, Got: {headerLine}");
        return;
    }

    var results = new List<JsonNode>();

    while ((headerLine = reader.ReadLine()) != null) {
        if (string.IsNullOrWhiteSpace(headerLine)) {
            continue;
        }

        var columns = headerLine.Split(',');
        if (columns.Length != 4) {
            Console.WriteLine($"Invalid row: {headerLine}");
            continue;
        }

        var name = columns[0].Trim();
        var ageText = columns[1].Trim();
        var email = columns[2].Trim();
        var scoreText = columns[3].Trim();

        if (string.IsNullOrEmpty(name) || string.IsNullOrEmpty(email) || string.IsNullOrEmpty(scoreText)) {
            Console.WriteLine($"Invalid row: {headerLine}");
            continue;
        }

        if (!int.TryParse(ageText, out int age)) {
            Console.WriteLine($"Invalid age: {ageText}");
            continue;
        }

        if (!double.TryParse(scoreText, out float score)) {
            Console.WriteLine($"Invalid score: {scoreText}");
            continue;
        }

        var obj = JsonNode.Parse("{\"Name\":\"" + name + "\",\"Age\":" + age + ",\"Email\":\"" + email + "\",\"Score\":" + score + "}");
        results.Add(obj);
    }

    if (results.Count == 0) {
        Console.WriteLine("[]");
    } else {
        var jsonArray = JsonNode.Parse("[" + string.Join(",", results.Select(r => r.ToString())) + "]");
        Console.Write(jsonArray.ToJsonString());
    }
} catch (Exception ex) {
    Console.WriteLine($"Error: {ex.Message}");
}