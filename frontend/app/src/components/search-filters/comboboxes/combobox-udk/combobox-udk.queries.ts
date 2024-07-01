import { useQuery } from "@tanstack/react-query";
import { SuggestionsAPI } from "./combobox-udk.services";

export const useUdkSuggestions = (query: string) => useQuery({
    queryKey: [query],
    queryFn: () => query == "/suggest/udk" ? {suggestions: []} : SuggestionsAPI.getSuggestUdk(query),
    select: data => data?.suggestions,
    initialData: {suggestions: []},
    retry: false
})