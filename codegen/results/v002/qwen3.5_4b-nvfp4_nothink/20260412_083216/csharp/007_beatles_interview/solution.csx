using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

var csvPath = Path.Combine("input", "input.csv");
var expectedPath = Path.Combine("input", "expected_format.json");

var lines = File.ReadAllLines(csvPath);
var firstLine = lines[0];

var headerList = firstLine.Split(',');

var records = new List<Dictionary<string, object>>();

for (var i = 1; i < lines.Length; i++)
{
    var line = lines[i];
    var values = line.Split(',');

    var dict = new Dictionary<string, object>();

    for (var j = 0; j < headerList.Length; j++)
    {
        var header = headerList[j].Trim();
        var val = values[j].Trim();

        if (val == "null" || string.IsNullOrEmpty(val))
        {
            dict[header] = null;
        }
        else
        {
            dict[header] = new DateTime(2025, 7, 1).Subtract(new DateTime(int.Parse(val.Substring(0, 4)), int.Parse(val.Substring(6, 2)), int.Parse(val.Substring(9, 2))))).Days;
            if (header == "Died")
            {
                // Died column is just a date, we don't include it in output
            }
            else if (header == "Birthday")
            {
                // We parsed the date above but stored as days? Let's fix to store actual date for Age calc.
                // Actually the parsing logic above was wrong for Birthday. Let's redo properly per row.
            }
        }
    }

    records.Add(dict);
}