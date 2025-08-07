using System.Text.Json.Serialization;

namespace PokeMin.Models
{
    public class PokemonCard
    {
        public string Name { get; set; } = "";
        public string ImageUrl { get; set; } = "";
        public List<string> Types { get; set; } = [];
        public int Height { get; set; }
        public int Weight { get; set; }
    }

    public class Pokemon
    {
        public string Name { get; set; }
        public string Url { get; set; }
    }

    public class PokemonResponse
    {
        public List<Pokemon> Results { get; set; }
    }
}
