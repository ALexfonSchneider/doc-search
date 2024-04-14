import axios from "axios"
import { config } from "@/config"
import { Suggestions } from "./search.interfaces"

const getSuggestions = (query: string) => 
    axios.get<Suggestions>(`${config.api_url}/suggest/queries`, {params: {query: query}}).then(response => response.data)

export const SearchAPI = {
    getQuerySuggestions: getSuggestions
}