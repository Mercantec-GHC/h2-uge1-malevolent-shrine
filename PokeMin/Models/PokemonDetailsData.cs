namespace PokeMin.Models;


    public class PokemonDetailsData
    {
        public string Name { get; set; }
        public int Height { get; set; }
        public int Weight { get; set; }
        public Sprites Sprites { get; set; }
        public List<PokemonType> Types { get; set; }
    }
