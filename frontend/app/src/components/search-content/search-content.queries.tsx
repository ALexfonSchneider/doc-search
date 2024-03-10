import { useQuery } from '@tanstack/react-query'
import { SearchAPI } from './search-content.services'

const initialData = {
    page: 1,
    size: 0,
    articles: [],
    total_size: 0
}

export const useDocuments = (query: string, selected_keywords: string[], page: number, size: number = 10) => useQuery({
    queryKey: [query, selected_keywords, page, size],
    queryFn: () => SearchAPI.searchDocuments(query, selected_keywords, page, size),
    select: data => data,
    retry: false,
    initialData: initialData
})



