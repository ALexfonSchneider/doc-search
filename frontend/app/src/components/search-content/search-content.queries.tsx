import { useQuery } from '@tanstack/react-query'
import { SearchAPI } from './search-content.services'
import { Udk } from '@/lib/reducers/search'

export const useDocuments = (query: string, selected_keywords: string[], year: string | null, udk: Udk | null, page: number, size: number = 10) => useQuery({
    queryKey: [query, year, selected_keywords, page, size, udk],
    queryFn: () => SearchAPI.searchDocuments(query, selected_keywords, year, udk, page, size),
    select: data => data,
    retry: false
})



