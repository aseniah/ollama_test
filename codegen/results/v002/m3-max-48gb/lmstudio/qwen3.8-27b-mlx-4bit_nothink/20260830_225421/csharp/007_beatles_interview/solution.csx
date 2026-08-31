int CalculateAge(DateTime birth, DateTime? death, DateTime refDate)
    {
        DateTime end = death ?? refDate;
        if (end < birth) return 0; // Safety
        int age = end.Year - birth.Year;
        // Subtract 1 if birthday hasn't occurred yet in the end year
        if (end.Month < birth.Month || (end.Month == birth.Month && end.Day < birth.Day))
        {
            age--;
        }
        return age;
    }