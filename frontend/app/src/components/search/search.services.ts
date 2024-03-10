import axios from "axios"
import { config } from "@/config"
import { MyTag, Suggestions } from "./search.interfaces"

const getSuggestions = (query: string) => 
    axios.get<Suggestions>(`${config.api_url}/suggest/query`, {params: {query: query}}).then(response => response.data)

const getTags = () => 
    axios.get<MyTag[]>(
        `${config.api_url}/metrics`, 
        {
            responseType: "json",
            transformResponse: (data) => {
                const j = JSON.parse(data)
                return j?.word_cloud || []
            }
        }).then(response => {
            return response.data
        })
    
export const SearchAPI = {
    getQuerySuggestions: getSuggestions,
    getTags
}