import { useQuery } from '@tanstack/react-query'
import { SearchAPI } from './search.services'

export const useQuerySuggestions = (query: string) => useQuery({
    queryKey: [query],
    queryFn: () => query == "" ? {suggestions: []} : SearchAPI.getQuerySuggestions(query),
    select: data => data?.suggestions,
    initialData: {suggestions: []},
    retry: false
})


export const useTags = () => useQuery({
    queryKey: ["query"],
    retry: false,
    queryFn: () =>  SearchAPI.getTags(),
})