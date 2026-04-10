using System;
using System.IO;
using System.Text.Json;
using System.Text.Json.Linq;

class Program
{
    static void Main(string[] args)
    {
        string inputFilePath = "input/data.json";
        string outputFilePath = "output/filtered_data.json";

        JObject data = JObject.Parse(File.ReadAllText(inputFilePath));

        JArray filteredData = data.Select(item =>
        {
            JObject itemData = JObject.Parse(item);
            if (itemData["active"] != null && itemData["active"].BooleanValue && itemData["age"] != null && itemData["age"].IntegerValue >= 30)
            {
                return JObject.Parse(item);
            }
            return null;
        }).Where(item => item != null);

        JArray.WriteObject(filteredData, outputFilePath);
        Console.WriteLine("Filtered data has been saved to {0}", outputFilePath);
    }
}