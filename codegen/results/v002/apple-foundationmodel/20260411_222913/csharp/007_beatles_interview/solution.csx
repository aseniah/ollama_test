using System;
using System.IO;
using System.Text.Json;
using System.Text.Json.Linq;

namespace Program
{
    class Program
    {
        static void Main()
        {
            Console.WriteLine("Reading CSV file...");
            JArray input = JArray.Parse(File.ReadAllLines("input/input.csv")).ToJArray();

            Console.WriteLine("Calculating ages...");
            foreach (JObject person in input)
            {
                JArray relatives = JArray.Parse(person["Relatives"]).ToJArray();
                JObject birth = JObject.Parse(person["Birthday"]);
                int age = CalculateAge(birth);
                person["Age"] = age.ToString();

                foreach (JObject relative in relatives)
                {
                    if (relative["Relationship"].Equals("Father"))
                    {
                        person["Father"] = relative["FirstName"] + " " + relative["LastName"];
                    }
                    else if (relative["Relationship"].Equals("Mother"))
                    {
                        person["Mother"] = relative["FirstName"] + " " + relative["LastName"];
                    }
                    else if (relative["Relationship"].Equals("Brother"))
                    {
                        person["Brother"] = relative["FirstName"] + " " + relative["LastName"];
                    }
                    else if (relative["Relationship"].Equals("Sister"))
                    {
                        person["Sister"] = relative["FirstName"] + " " + relative["LastName"];
                    }
                }
            }

            Console.WriteLine("Writing JSON to output...");
            JArray output = JArray.Parse(input.ToString()).ToJArray();
            File.WriteAllText("output.json", output.ToJson());
            Console.WriteLine("JSON array written to output.json");
        }

        static int CalculateAge(JObject birthday)
        {
            JArray birthDate = JArray.Parse(birthday.ToString());
            int year = 2025 - int.Parse(birthDate[0].ToString());
            int month = 1 - int.Parse(birthDate[1].ToString());
            int day = 1 - int.Parse(birthday[2].ToString());
            return year - (month < 3 ? 1 : 0) + (day < 30 ? 1 : 0);
        }
    }
}