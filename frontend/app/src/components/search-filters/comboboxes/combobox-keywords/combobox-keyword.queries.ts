import { useQuery } from "@tanstack/react-query";
import { SuggestionsAPI } from "./combobox-keywords.services";

export const useKeywordsSuggestions = (query: string) => useQuery({
    queryKey: [query],
    queryFn: () => query == "/metrics" ? {suggestions: []} : SuggestionsAPI.getKeywordsSuggestions(query),
    select: data => data?.suggestions,
    initialData: {suggestions: []},
    retry: false
})