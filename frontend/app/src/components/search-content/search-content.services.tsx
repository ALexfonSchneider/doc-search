import { config } from '@/config'
import axios from "axios"
import { SearchResultsPaginate } from './search-content.interfaces'
import { Udk } from '@/lib/reducers/search'

const searchDocuments = (query: string, selected_keywords: string[], year: string | null, udk: Udk | null, page: number, size: number = 10) => {
    const params: any = {
        q: query, page, count: size
    }

    if (selected_keywords.length != 0) {
        params["keywords"] = selected_keywords
    }

    if (year != null) {
        params["year"] = year
    }

    if (udk) {
        params['udk'] = udk.code
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

