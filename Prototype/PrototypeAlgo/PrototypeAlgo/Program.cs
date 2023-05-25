using System;
using System.Collections.Generic;
using System.Linq;

public class WineTank
{
    public int NumberOfTanks { get; set; }
    public int Capacity { get; set; }
    public string WineType { get; set; }
}

public class Tank
{
    public int Id { get; set; }
    public int Capacity { get; set; }
    public string WineType { get; set; }
}

public class WineBlend
{
    public string WineType { get; set; }
    public double Percentage { get; set; }
}

public class Program
{
    public static void Main()
    {
        var totalVolume = 1000;
        var tankId = 0;

        var formula = new List<WineBlend>
        {
            new WineBlend { WineType = "chardonnay", Percentage = 37 },
            new WineBlend { WineType = "pinot_noir", Percentage = 45 },
            new WineBlend { WineType = "meunier", Percentage = 18 }
        };

        var wineTanks = new List<WineTank>
        {
            new WineTank { NumberOfTanks = 3, Capacity = 50, WineType = "empty" },
            new WineTank { NumberOfTanks = 2, Capacity = 100, WineType = "empty" },
            new WineTank { NumberOfTanks = 1, Capacity = 100, WineType = "pinot_noir" },
            new WineTank { NumberOfTanks = 1, Capacity = 100, WineType = "meunier" },
            new WineTank { NumberOfTanks = 1, Capacity = 50, WineType = "meunier" },
        };

        var tanks = new List<Tank>();

        // Generate individual tanks based on wineTanks
        foreach (var wineTank in wineTanks)
        {
            for (int i = 0; i < wineTank.NumberOfTanks; i++)
            {
                tanks.Add(new Tank { Id = ++tankId, Capacity = wineTank.Capacity, WineType = wineTank.WineType });
            }
        }

        foreach (var blend in formula)
        {
            var requiredQuantity = (int)(totalVolume * blend.Percentage / 100);

            foreach (var tank in tanks)
            {
                if (tank.WineType == blend.WineType && tank.Capacity > 0)
                {
                    if (tank.Capacity >= requiredQuantity)
                    {
                        Console.WriteLine($"Take {requiredQuantity} from tank number {tank.Id} for {blend.WineType}");
                        tank.Capacity -= requiredQuantity;
                        requiredQuantity = 0;
                    }
                    else
                    {
                        Console.WriteLine($"Take all {tank.Capacity} from tank number {tank.Id} for {blend.WineType}");
                        requiredQuantity -= tank.Capacity;
                        tank.Capacity = 0;  // The tank is now empty.
                    }
                }

                if (requiredQuantity == 0)
                    break;
            }

            // If we still need more wine, fill an empty tank
            if (requiredQuantity > 0)
            {
                foreach (var tank in tanks)
                {
                    if (tank.WineType == "empty" && tank.Capacity >= requiredQuantity)
                    {
                        Console.WriteLine($"Fill tank number {tank.Id} with {requiredQuantity} of {blend.WineType}");
                        tank.WineType = blend.WineType;
                        requiredQuantity = 0;
                        break;
                    }
                }
            }

            if (requiredQuantity > 0)
                Console.WriteLine($"Not enough {blend.WineType} available");
        }

    }
}
