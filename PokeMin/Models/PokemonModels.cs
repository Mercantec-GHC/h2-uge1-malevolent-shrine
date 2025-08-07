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

    public class PokemonDetailsData
    {
        public string Name { get; set; }
        public int Height { get; set; }
        public int Weight { get; set; }
        public Sprites Sprites { get; set; }
        public List<PokemonType> Types { get; set; }
    }
    public class Sprites
    {
        [JsonPropertyName("front_default")]
        public string Front_Default { get; set; }
    }
    public class PokemonType
    {
        public TypeInfo Type { get; set; }
    }

    public class TypeInfo
    {
        public string Name { get; set; }
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
