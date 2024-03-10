import { config } from '@/config'
import axios from "axios"
import { SearchResultsPaginate } from './search-content.interfaces'

const searchDocuments = (query: string, selected_keywords: string[], page: number, size: number = 10) => {
    const keywords_query = selected_keywords.length == 0 ? "" : selected_keywords

    return axios.get<SearchResultsPaginate>(
        `${config.api_url}/search`,
        {
            params: {
                query, page, keywords_query, size
            }
        },
    ).then(response => response.data)
    }



export const SearchAPI = {
    searchDocuments
}

