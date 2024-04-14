import { useQuery } from '@tanstack/react-query'
import { SearchAPI } from './search-content.services'

export const useDocuments = (query: string, selected_keywords: string[], year: string | null, page: number, size: number = 10) => useQuery({
    queryKey: [query, year, selected_keywords, page, size],
    queryFn: () => SearchAPI.searchDocuments(query, selected_keywords, year, page, size),
    select: data => data,
    retry: true
})



