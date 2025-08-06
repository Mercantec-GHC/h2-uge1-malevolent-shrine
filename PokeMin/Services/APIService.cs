using PokeMin.Models;
using PokeMin.Pages;
using System.Net.Http.Json;

namespace PokeMin.Services
{
    public class APIService
    {
        private readonly HttpClient _httpClient;
        //       private const string BaseUrl = "";
        public APIService(HttpClient httpClient)
        {
            _httpClient = httpClient;
        }

        public async Task<List<PokemonCard>> GetAllPokemons(int offset = 0, int limit = 10)
        {
            var listResponse = await _httpClient.GetFromJsonAsync<PokemonResponse>($"https://pokeapi.co/api/v2/pokemon?limit={limit}&offset={offset}");
            if (listResponse == null) return [];

            var tasks = listResponse.Results.Select(async entry =>
            {
                var detail = await _httpClient.GetFromJsonAsync<PokemonDetailsData>($"https://pokeapi.co/api/v2/pokemon/{entry.Name}");
                if (detail == null) return null;

                return new PokemonCard
                {
                    Name = detail.Name,
                    ImageUrl = detail.Sprites.Front_Default,
                    Types = detail.Types.Select(t => t.Type.Name).ToList()

                };
            });

            var cards = await Task.WhenAll(tasks);
            return cards.Where(c => c != null).ToList()!;
        }

        public async Task<PokemonCard?> GetPokemonDetailAsync(string nameOrId)
        {
            var detail = await _httpClient.GetFromJsonAsync<PokemonDetailsData>($"https://pokeapi.co/api/v2/pokemon/{nameOrId}");
            if (detail == null) return null;

            return new PokemonCard
            {
                Name = detail.Name,
                ImageUrl = detail.Sprites.Front_Default,
                Types = detail.Types.Select(type => type.Type.Name).ToList(),
                Weight = detail.Weight,
                Height = detail.Height
            };
        }
    }
}
