import { config } from '@/config'
import axios from "axios"
import { SearchResultsPaginate } from './search-content.interfaces'

const searchDocuments = (query: string, selected_keywords: string[], year: string | null, page: number, size: number = 10) => {
    const keywords_query = selected_keywords.length == 0 ? "" : selected_keywords

    const params: any = {
        query, page, keywords_query, size
    }

    if (year != null) {
        params["year"] = year
    }

    return axios.get<SearchResultsPaginate>(
        `${config.api_url}/search`,
        {
            params
        },
    ).then(response => response.data)
    }



export const SearchAPI = {
    searchDocuments
}

